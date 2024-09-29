package sql

//TODO: Uncomment this file and move it to some other project.
/*
import (
	"fmt"
	"strconv"
	"strings"

	"potionDB/potionDB/components"
	"potionDB/crdt/crdt"
)

//Pre-condition: upd.UpdateArgs is an inbedded Map update
//Pre-condition: The caller already verified that this object is relevant for this view.
func (v *ViewInfo) GenerateViewUpdate(upd antidote.UpdateObjectParams) {
	switch typedUpd := (*upd.UpdateArgs).(type) {
	case crdt.EmbMapUpdateAll:
		v.handleNewObjEmbMap(upd.KeyParams, typedUpd)
	case *crdt.EmbMapUpdateAll:
		v.handleNewObjEmbMap(upd.KeyParams, *typedUpd)
	case crdt.MapRemoveAll:
		v.handleRemObjEmbMap(upd.KeyParams, typedUpd)
	case *crdt.MapRemoveAll:
		v.handleRemObjEmbMap(upd.KeyParams, *typedUpd)
	}
}

//Pre: must have called all the preparation methods...
func (v *ViewInfo) handleNewObjEmbMap(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll) {
	objValues := v.fromCrdtUpdsToValues(upd)
	fmt.Println("Values: ", objValues)
	fmt.Println("Fields needed:", v.FieldsNeeded)
	if !v.ShouldBeUpdated(objValues) {
		fmt.Println("The object above should not generate an update for Q3")
		return
	}
	var viewKeys []string
	if v.Listen.HasInequalityGrouping {
		viewKeys = v.GetKeysWithInequalityForUpdate(objValues)
		fmt.Println(viewKeys, len(viewKeys))
		if len(viewKeys) == 0 { //Can happen when there isn't a single parameter value that respects all inequalities
			fmt.Println("The object above should not generate an update for Q3: no parameter value can satisfact all inequalities.")
			return
		}
	} else {
		viewKeys = []string{v.GetKeyForUpdate(objValues)}
	}
	fmt.Println("The object above should generate an update for Q3")
	fmt.Println("ViewKeys to update:", viewKeys)
	var viewBaseUpd crdt.UpdateArguments
	switch v.ViewType {
	case TOPK:
		viewBaseUpd = v.makeNewTopKViewUpd(objKey, upd, objValues)
	case TOPSUM:
		viewBaseUpd = v.makeNewTopSumViewUpd(objKey, upd, objValues)
	case MAP:
		viewBaseUpd = v.makeNewMapViewUpd(objKey, upd, objValues)
	case COUNTER:
		viewBaseUpd = v.makeNewCounterViewUpd(objKey, upd, objValues)
	case REGISTER:
		viewBaseUpd = v.makeNewRegisterViewUpd(objKey, upd, objValues)
	case AVERAGE:
		viewBaseUpd = v.makeNewAverageViewUpd(objKey, upd, objValues)
	case MAX_CRDT:
		viewBaseUpd = v.makeNewMaxViewUpd(objKey, upd, objValues)
	case MIN_CRDT:
		viewBaseUpd = v.makeNewMinViewUpd(objKey, upd, objValues)
	}
	//viewKeys will probably be used after this
	//TODO: Need to add listen.ID to each of these keys
	ignore(viewKeys)
	ignore(viewBaseUpd)
	/*viewUpds := make([]crdt.UpdateObjectParams, len(viewKeys))
	for i, viewKey := range viewKeys {
		viewUpds[i] = crdt.UpdateObjectParams{KeyParams: antidote.KeyParams{Key: viewKey, Bucket: v.Listen.Bucket, CrdtType: }, UpdateArgs: &viewBaseUpd}
	}*/
//TODO: Finish
//Note: Parameters are not present in select (even though they are in theory supported)
/*}

//Assuming full object delete
func (v *ViewInfo) handleRemObjEmbMap(objKey antidote.KeyParams, upd crdt.MapRemoveAll) {

}

func (v *ViewInfo) fromCrdtUpdsToValues(upd crdt.EmbMapUpdateAll) map[string]interface{} {
	updValues := make(map[string]interface{})
	for key, dataType := range v.FieldsNeeded {
		switch typedUpd := upd.Upds[key].(type) {
		//Counter
		case crdt.Increment:
			updValues[key] = int(typedUpd.Change)
		case crdt.Decrement:
			updValues[key] = int(-typedUpd.Change)
		//Register
		case crdt.SetValue:
			stringValue := typedUpd.NewValue.(string)
			switch dataType {
			case INT_TYPE:
				updValues[key], _ = strconv.Atoi(stringValue)
			case FLOAT_TYPE:
				updValues[key], _ = strconv.ParseFloat(stringValue, 64)
			case NUMBER_TYPE:
				number, err := strconv.Atoi(stringValue)
				if err != nil {
					updValues[key], _ = strconv.ParseFloat(stringValue, 64)
				} else {
					updValues[key] = number
				}
			case DATE_TYPE:
				updValues[key] = Date{}.FromString(stringValue)
			default: //For now for all other types we want as string
				updValues[key] = stringValue
			}
		default:
			fmt.Printf("Unknown type: %T (key: %s)\n", typedUpd, key)
		}
	}
	fmt.Printf("FieldTypes: %v\n", v.FieldsNeeded)
	return updValues
}

func (v *ViewInfo) makeNewTopKViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	id, value, extra := v.makeNewTopUpdHelper(objValues)
	fmt.Printf("TOPK update: (%d, %d, %s)\n", id, value, string(extra))
	return crdt.TopKAdd{TopKScore: crdt.TopKScore{Id: id, Score: value, Data: &extra}}
}

func (v *ViewInfo) makeNewTopSumViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	id, value, extra := v.makeNewTopUpdHelper(objValues)
	fmt.Printf("TOPSUM update: (%d, %d, %s)\n", id, value, string(extra))
	return crdt.TopSAdd{TopKScore: crdt.TopKScore{Id: id, Score: value, Data: &extra}}
}

func (v *ViewInfo) makeNewTopUpdHelper(objValues map[string]interface{}) (id, value int32, extra []byte) {
	//Current limitation: both id and value must be ints.
	var idS, valueS, extraS string
	var extraSb strings.Builder
	for _, key := range v.TopKType.ExtraData {
		//Now can use AllSelect
		extraSb.WriteString(ValueToString(v.Listen.AllSelect[key]))
		extraSb.WriteRune('_')
	}
	idS, valueS, extraS = ValueToString(v.TopKType.Id), ValueToString(v.TopKType.Value), extraSb.String()
	if len(extraS) > 0 {
		extraS = extraS[:len(extraS)-1]
	}
	idInt, _ := strconv.Atoi(idS)
	valueInt, _ := strconv.Atoi(valueS)
	return int32(idInt), int32(valueInt), []byte(extraS) //[]byte(extraS[:len(extraS)-1])
}

func (v *ViewInfo) makeNewMapViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	//v.MapType
	upds := make(map[string]crdt.UpdateArguments)
	for id, nameable := range v.Listen.SelectNameable { //Always register
		upds[id] = crdt.SetValue{NewValue: ValueToString(nameable.CalculateValue(objValues))}
	}
	for id, asClause := range v.Listen.SelectAsClause { //Always register
		upds[id] = crdt.SetValue{NewValue: ValueToString(asClause.Math.GetCalculable().CalculateValue(objValues))}
	}
	for id, asAggr := range v.Listen.SelectAsAggregation {
		value := asAggr.Math.GetCalculable().CalculateValue(objValues)
		intV := int64(0)
		var currUpd crdt.UpdateArguments
		switch typedV := value.(type) {
		case int:
			intV = int64(typedV)
		case float64:
			intV = int64(typedV)
		}
		switch asAggr.AggrFunc {
		case SUM:
			currUpd = crdt.Increment{Change: int32(intV)}
		case AVG:
			currUpd = crdt.AddValue{Value: intV}
		case MAX:
			currUpd = crdt.MaxAddValue{Value: intV}
		case MIN:
			currUpd = crdt.MinAddValue{Value: intV}
		}
		upds[id] = currUpd
	}
	for id, _ := range v.Listen.SelectAsCount {
		upds[id] = crdt.Increment{Change: 1}
	}
	return crdt.EmbMapUpdateAll{Upds: upds}
}

func (v *ViewInfo) makeNewCounterViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	//We only support basic counting for now. (COUNT(*))
	return crdt.Increment{Change: 1}
}

func (v *ViewInfo) makeNewRegisterViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	//It's always in "SelectAsClause" or "SelectNameable"
	if len(v.Listen.SelectNameable) > 0 {
		for _, nameable := range v.Listen.SelectNameable { //Fake for... we're only iterating one position
			return crdt.SetValue{NewValue: ValueToString(nameable.CalculateValue(objValues))}
		}
	}
	for _, asClause := range v.Listen.SelectAsClause { //Fake for...
		return crdt.SetValue{NewValue: ValueToString(asClause.Math.GetCalculable().CalculateValue(objValues))}
	}
	return nil
}

func (v *ViewInfo) makeNewAverageViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	return crdt.AddValue{Value: v.getIntSingleAggr(objValues)}
}

func (v *ViewInfo) makeNewMaxViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	return crdt.MaxAddValue{Value: v.getIntSingleAggr(objValues)}
}

func (v *ViewInfo) makeNewMinViewUpd(objKey antidote.KeyParams, upd crdt.EmbMapUpdateAll, objValues map[string]interface{}) crdt.UpdateArguments {
	return crdt.MinAddValue{Value: v.getIntSingleAggr(objValues)}
}

func (v *ViewInfo) getIntSingleAggr(objValues map[string]interface{}) int64 {
	for _, aggr := range v.Listen.SelectAsAggregation { //Fake for.
		value := aggr.Math.GetCalculable().CalculateValue(objValues)
		switch typedV := value.(type) {
		case int:
			return int64(typedV)
		case float64:
			return int64(typedV)
		}
	}
	return 0
}
*/
