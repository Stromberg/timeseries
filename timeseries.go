package timeseries

import (
	"sort"
)

type TimeArray struct {
	timestamp []interface{}
	values    map[string][]interface{}
}

func FromData(timestamp []interface{}, values map[string][]interface{}) *TimeArray {
	// Check length of all inputs
	return &TimeArray{
		timestamp: timestamp,
		values:    values,
	}
}

func (ta *TimeArray) TimeStamp() []interface{} {
	return ta.timestamp
}

func (ta *TimeArray) ColNames() []string {
	keys := []string{}
	for k := range ta.values {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	return keys
}

func (ta *TimeArray) Values(colnames ...string) *TimeArray {
	values := make(map[string][]interface{}, 0)
	for _, name := range colnames {
		values[name] = ta.values[name]
	}
	return FromData(ta.timestamp, values)
}

func (ta *TimeArray) AsSlice(name string) []interface{} {
	return ta.values[name]
}

func (ta *TimeArray) Rename(oldName, newName string) bool {
	oldCol, ok := ta.values[oldName]
	if !ok {
		return false
	}
	ta.values[newName] = oldCol
	delete(ta.values, oldName)
	return true
}

func (ta *TimeArray) RemoveIndices(indices []int) *TimeArray {
	ts := ta.removeIndices(indices, ta.timestamp)
	values := map[string][]interface{}{}

	for k, v := range ta.values {
		values[k] = ta.removeIndices(indices, v)
	}

	return FromData(ts, values)
}

func (ta *TimeArray) ChangeTimeStamp(oldTs, newTs interface{}) {
	for i, v := range ta.timestamp {
		if v == oldTs {
			ta.timestamp[i] = newTs
			return
		}
	}
}
func (ta *TimeArray) removeIndices(indices []int, s []interface{}) []interface{} {
	res := []interface{}{}

	pi := 0
	for i, v := range s {
		if pi >= len(indices) || i < indices[pi] {
			res = append(res, v)
		} else if i == indices[pi] {
			pi++
		}
	}

	return res
}
