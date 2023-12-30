package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	SELECT, FROM, WHERE, GROUP_BY, ORDER_BY, LIMIT = 0, 1, 2, 3, 4, 5
	SUM, COUNT, AVG, MAX, MIN                      = 1, 2, 3, 4, 5
)

type ViewInfo struct {
	Container     string
	IsTop         bool
	TopSize       int
	SortElems     []string //Elements used to sort in a topK (or map)
	MapKeys       []string //Different views, or "maps of maps".
	CounterFields []CounterField
	StringFields  []string //Could still be numbers, but they are not a sum/avg/count/max/min.

	ViewType string
	View
}

type CounterField struct {
	AggregType int
	FieldName  string
}

//ExternalMap: each entry could be different views
type ExternalMap struct {
	KeyField   string
	ValueField View
}

type Map struct {
	View map[string]View
}

type TopK struct {
	IdField       string
	SortingFields []string //I think this needs to be something else - can be both string or int
	ExtraFields   []string
}

type Counter struct {
	FieldName string
}

type Register struct {
	FieldName string
}

type View interface {
	//addView(view View)
}

//Pre: all fields must be filled up
func (v *ViewInfo) calculateViewType() {
	finalType := ""
	for _, key := range v.MapKeys {
		finalType += "Map[" + key + "]"
	}
	if v.IsTop {
		finalType += "topsum"
	} else if len(v.CounterFields) == 1 && len(v.StringFields) == 0 {
		finalType += "counter"
	} else if len(v.CounterFields) == 0 && len(v.StringFields) == 1 {
		finalType += "register"
	} else if len(v.CounterFields) == 1 && len(v.StringFields) == 1 {
		finalType += "map[" + v.StringFields[0] + "]" + v.CounterFields[0].FieldName //NOTE: THIS WILL BE WRONG IF STRINGFIELDS[0] IS NOT UNIQUE!
	} else {
		finalType += "map[fieldNames]valuesNames" //Note: We have enough information to identify which should be Registers and which should be counter-like types.
	}
	v.ViewType = finalType
}

func (v *ViewInfo) buildView() {
	var innerMostMap *ExternalMap = nil
	if len(v.MapKeys) > 0 {
		currMap := &ExternalMap{KeyField: v.MapKeys[0]}
		v.View = currMap
		for i := 1; i < len(v.MapKeys); i++ {
			nextMap := &ExternalMap{KeyField: v.MapKeys[i]}
			currMap.ValueField = &nextMap
			currMap = nextMap
		}
		innerMostMap = currMap
	}
	var lastView View
	if v.IsTop {
		lastView = &TopK{}
	} else if len(v.CounterFields) == 1 && len(v.StringFields) == 0 {
		lastView = &Counter{FieldName: v.CounterFields[0].FieldName} //We could need to keep info if this is a sum, count, avg, max/min, etc.
	} else if len(v.CounterFields) == 0 && len(v.StringFields) == 1 {
		lastView = &Register{FieldName: v.StringFields[0]}
	} else if len(v.CounterFields) == 1 && len(v.StringFields) == 1 {
		ourMap := &Map{View: make(map[string]View)}
		ourMap.View[v.StringFields[0]] = &Counter{FieldName: v.CounterFields[0].FieldName} //NOTE: THIS WILL BE WRONG IF STRINGFIELDS[0] IS NOT UNIQUE!
		lastView = ourMap
	} else {
		ourMap := &Map{View: make(map[string]View)}
		for _, field := range v.StringFields {
			ourMap.View[field] = &Register{}
			//finalType += "map[fieldNames]valuesNames" //Note: We have enough information to identify which should be Registers and which should be counter-like types.
		}
		for _, field := range v.CounterFields {
			ourMap.View[field.FieldName] = &Counter{} //May need to consider the counter types (avg, count, etc.) at some point
		}
	}
	if innerMostMap == nil {
		v.View = lastView
	} else {
		innerMostMap.ValueField = lastView
	}
}

func (v *ViewInfo) showType() {
	fmt.Println(v.ViewType)
	if v.IsTop {
		fmt.Println("Top size:", v.TopSize)
	}
	fmt.Println("Container:", v.Container)
}

/*

 */

func start() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Number of lines?")
	reader.Scan()
	nLinesS := reader.Text()
	nLines, _ := strconv.Atoi(nLinesS)
	lines := make([][]string, 6)
	origLines := make([]string, 6)
	for i := 0; i < nLines; i++ {
		reader.Scan()
		currLine := reader.Text()
		origLines[i] = currLine
		lines[i] = strings.Split(currLine, " ")
	}
	sortLines(lines, origLines)
	removeCommas(lines)
	v := &ViewInfo{}
	v.processView(lines, origLines)
	v.calculateViewType()
	v.showType()
}

func removeCommas(parts [][]string) {
	for _, slice := range parts {
		//First position never has comma
		for i := 1; i < len(slice); i++ {
			if slice[i][len(slice[i])-1] == ',' {
				slice[i] = slice[i][0 : len(slice[i])-1]
			}
		}
	}
}

//First pos: keyword. On multi-word keyword, second pos may also have keyword
//Will order it as a normal query
//Select -> From -> Where -> Group by -> Order by -> Limit
func sortLines(toSort [][]string, origLines []string) {
	keywordsToSearch := []string{"SELECT", "FROM", "WHERE", "GROUP", "ORDER", "LIMIT"}
	//Put first words in capital
	for i := 0; i < len(toSort); i++ {
		//Some keywords may not be  present
		if len(toSort[i]) > 0 {
			toSort[i][0] = strings.ToUpper(toSort[i][0])
		}
	}
	for i := 0; i < len(keywordsToSearch); i++ {
		if len(toSort[i]) > 0 && keywordsToSearch[i] == toSort[i][0] {
			continue //Already in the right position
		}
		j := i
		for ; j < len(toSort); j++ {
			if len(toSort[j]) > 0 && keywordsToSearch[i] == toSort[j][0] {
				copy := toSort[i]
				toSort[i] = toSort[j]
				toSort[j] = copy
				copy2 := origLines[i]
				origLines[i] = origLines[j]
				origLines[j] = copy2
			}
		}
		//Keyword not found. Leave empty. But we still need to copy whatever is there to the end
		if j == len(toSort) {
			//Find an empty space at the end to store this
			for k := len(toSort) - 1; k > 0; k-- {
				if len(toSort[k]) == 0 {
					toSort[k] = toSort[i]
					toSort[i] = nil
					origLines[k] = origLines[i]
					origLines[i] = ""
				}
			}
		}
	}
}

func (v *ViewInfo) processView(lines [][]string, origLines []string) {
	v.processFrom(lines[FROM])
	if len(lines[LIMIT]) > 0 {
		v.processLimit(lines[LIMIT])
	}
	if len(lines[ORDER_BY]) > 0 {
		v.processOrderBy(lines[ORDER_BY])
	}
	v.processWhere(lines[WHERE])
	v.processSelect(lines[SELECT], origLines[SELECT])
	if len(lines[GROUP_BY]) > 0 {
		v.processGroupBy(lines[GROUP_BY])
	}
}

func (v *ViewInfo) processFrom(fromLine []string) {
	//Note: Assuming no joins for now
	v.Container = fromLine[1]
}

func (v *ViewInfo) processLimit(limitLine []string) {
	v.IsTop = true
	v.TopSize, _ = strconv.Atoi(limitLine[1])
}

func (v *ViewInfo) processOrderBy(orderByLine []string) {
	v.SortElems = make([]string, len(orderByLine)-2)
	v.IsTop = true
	//First two are "order" and "by".
	//Will copy the elements in case I need to do other processing with the input
	for i := 2; i < len(orderByLine); i++ {
		v.SortElems[i-2] = orderByLine[i]
	}
}

//Pre: all names must be longer than 2 letters
func (v *ViewInfo) processWhere(whereLine []string) {
	//Here we only need to look for the upper case words
	for _, word := range whereLine[1:] {
		if IsAllUpper(word) {
			//Must check if it wasn't already found before
			alreadyFound := false
			for _, found := range v.MapKeys {
				if found == word {
					alreadyFound = true
					break
				}
			}
			if !alreadyFound {
				v.MapKeys = append(v.MapKeys, word)
			}
		}
	}
}

func (v *ViewInfo) processSelect(selectLine []string, origLine string) {
	origLine = origLine[7:] //Skip "select"
	split := strings.Split(origLine, ", ")
	//TODO: This still does not take in consideration:
	//What is inside a sum(), avg(), max(), min(), count()
	//Maths done (e.g. fieldA + fieldB will count as just one field.)
	for _, field := range split {
		//Probably would be cheaper to first verify if there's any "(" in field.
		if strings.HasPrefix(field, "sum(") {
			v.CounterFields = append(v.CounterFields, CounterField{AggregType: SUM, FieldName: field[4 : len(field)-1]}) //-1: take out last ()
		} else if strings.HasPrefix(field, "avg(") {
			v.CounterFields = append(v.CounterFields, CounterField{AggregType: AVG, FieldName: field[4 : len(field)-1]}) //-1: take out last ()
		} else if strings.HasPrefix(field, "max(") {
			v.CounterFields = append(v.CounterFields, CounterField{AggregType: MAX, FieldName: field[4 : len(field)-1]}) //-1: take out last ()
		} else if strings.HasPrefix(field, "min(") {
			v.CounterFields = append(v.CounterFields, CounterField{AggregType: MIN, FieldName: field[4 : len(field)-1]}) //-1: take out last ()
		} else if strings.HasPrefix(field, "count(") {
			v.CounterFields = append(v.CounterFields, CounterField{AggregType: COUNT, FieldName: field[6 : len(field)-1]}) //-1: take out last ()
		} else {
			v.StringFields = append(v.StringFields, field)
		}
	}
}

func (v *ViewInfo) processGroupBy(groupByLine []string) {

}

//Which order do I want to process by?
//1st - FROM: identifies "group"
//2nd - LIMIT: forces TopK/TopSum
//3rd - ORDER BY: defines the ID of TopK or Map. If TopK, the view will be sorted, if Map, the client sorts. Either way, this will be the key(s)
//4th - WHERE: defines keys and number of map levels
//5th - SELECT: can also find if it's a counter or top. Obtain set of fields of the view.
//6th - GROUP BY: useless for now. Will be useful for aggregations later I suppose. (like ID to search for?)

/*
Rules:
Order by -> topk or topsum. Defines which components are part of topk/topsum.
Limit -> topk or topsum. Defines size of topk/topsum.
From -> "container" of objects to consider (in which kind of objects should an update be triggered)
Where -> which values to be considered.
	any parameter (e.g. DATE) -> key of map or different view. (equivalent solutions)
		more than one parameter -> can do compound key or "maps inside of maps". We do not, however, have any way to evaluate which parameter is more decisive. (e.g.: year is less selective than day)
	other fields -> for the included keys, the values that certain properties must have.
		maybe assume we only know how to process ANDs?
Select:
	sum -> counter or top.
	count -> counter or top.
	other fields -> if topK, extra data on topk. Otherwise, map with registers?
		- There's a special case in which, if it's only one extra element, could be a key of a map.
Group by -> I don't think this is really helpful. We need, however, one way to know
how to apply the sum/other aggregations. How do we define this? Which restrictions should we make?
Limit -> topk/topsum
*/

/*
Q3:
	Must use lineitem due to shipdate. Each lineitem in an order has a different shipdate.
	Add o.orderdate and c.mktsegment back to lineitem.
	Alternatively, sum of sale price (revenue) and c.mktsegment to orders.
	Will assume the later.

Select orderkey, sum(extendedprice*(1-discount) as revenue, orderdate, shippriority
From lineitem
Where mktsegment = SEGMENT and orderdate < DATE and shipdate > DATE and DATE >= 1995-03-01 and DATE <= 1995-03-31
Group by orderkey, orderdate, shippriority


Q5:
	Must still use lineitem. This because we need to ensure nationkey of the item's supplier matches the nationkey of the customer's
		Different lineitems in the same order can belong to different countries
	Can either add customer's nation or supplier's nation to lineitem. Will add customer's nation. Also need o.orderdate. Also must add Region's name. Must add both nation keys.

Select n_name, sum(extendedprice * (1-discount)) as revenue
From lineitem
Where c_nationkey = s_nationkey and r_name = REGION and orderdate.year == YEAR and YEAR >= 1993 and YEAR <= 1997
Group by n_name
Order by revenue desc


Q11:
	Must use partsupp (need supplycost and availqty).
	Must add supplier's nationname to partsupp (for filtering)



View 1 (actual top):
Select partkey, sum(supplycost + availqty) as value
From partsupp
Where n_name = NATION
Group by partkey
Order by value desc
//Limiting is then done aside from the view with the sum view's value.

View 2 (sum):
Select sum(supplycost * availqty) as value
From partsupp
Where n_name = NATION


Q14:
	Must be lineitem, as the same product can have different extendedprice and discounts (and tax too).
	Must add p.type to lineitem.

First view: all

Select sum(extendedprice * (1-discount))
From lineitem
Where shipdate.year = YEAR and shipdate.month = MONTH and YEAR >= 1993 and YEAR <= 1997

Second view: promo
Select sum(extendedprice * (1-discount))
From lineitem
Where p_type startsWith "PROMO" and shipdate.year = YEAR and shipdate.month = MONTH and YEAR >= 1993 and YEAR <= 1997


Q15:
	Sadly must still be lineitem because of the shipdate...
	So we must add supplier's name, address and phone to lineitem.

Select suppkey, suppname, suppaddress, suppphone, sum(extendedprice * (1-discount)) as revenue
From lineitem
Where shipdate >= QUARTER and shipdate < QUARTER + 1 and QUARTER >= 1993.1 and QUARTER < 1998.1
Group by suppkey, suppname, suppaddress, suppphone
Order by revenue desc, suppkey asc
Limit 1


Q18:
	Could do this from order if we add a totalquantity and customer's name field to order.
	If using lineitem, either need to add o.totalprice or do a sum(). Also must add customer's name field.
		- Given we would need to add customer's name field, better do directly on orders.
	totalprice is extendedprice with tax and after applying discount

Select c.name, custkey, orderkey, orderdate, totalprice, totalquantity
From orders
Where totalquantity > QUANTITY and QUANTITY >= 312 and QUANTITY <= 315
Order by totalprice desc, orderdate asc
Limit 100
*/

func IsAllUpper(test string) bool {
	for _, char := range test {
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}
