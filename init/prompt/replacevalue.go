package prompt

import "sort"

type ReplaceValue struct {
	ID           string
	OldValue     string
	NewValue     string
	Description  string
	ReplaceOrder int
}

type ReplaceValues []ReplaceValue

func (r ReplaceValues) Sort() ReplaceValues {
	sort.Slice(r, func(i, j int) bool {
		return r[i].ReplaceOrder < r[j].ReplaceOrder
	})

	return r
}

func (rvs ReplaceValues) FetchByID(id string) ReplaceValue {
	for _, rv := range rvs {
		if rv.ID == id {
			return rv
		}
	}
	return ReplaceValue{}
}
