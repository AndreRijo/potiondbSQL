package sql

type CalculableStack []Calculable

func (st *CalculableStack) Push(value Calculable) {
	*st = append(*st, value)
}

func (st *CalculableStack) Pool() Calculable {
	if st.IsEmpty() {
		return nil
	}
	return (*st)[len(*st)-1]
}

func (st *CalculableStack) Pop() Calculable {
	if st.IsEmpty() {
		return nil
	}
	toReturn := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]

	return toReturn
}

func (st *CalculableStack) ReplaceTop(value Calculable) {
	if st.IsEmpty() {
		st.Push(value)
	} else {
		(*st)[len(*st)-1] = value
	}
}

func (st *CalculableStack) IsEmpty() bool {
	return len(*st) == 0
}
