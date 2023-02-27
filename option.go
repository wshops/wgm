package wgm

type FindPageOption struct {
	fields   []string
	selector interface{}
}

// SetSortField
// @param field []string
// @Description: 初始化SortField, {"age", "-name"}, first sort by age in ascending order, then sort by name in descending order
func (o *FindPageOption) SetSortField(field ...string) *FindPageOption {
	o.fields = field
	return o
}

// SetSelectField
// @param fieldStr bson
// @Description: 初始化Select,
// bson.M{"age": 1} means that only the age field is displayed
// bson.M{"age": 0} means to display other fields except age
func (o *FindPageOption) SetSelectField(bson interface{}) *FindPageOption {
	o.selector = bson
	return o
}

func NewFindPageOption() *FindPageOption {
	return acquireFindPageOption()
}
