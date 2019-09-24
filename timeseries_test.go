package timeseries_test

import (
	"fmt"
	"testing"

	"github.com/Stromberg/timeseries"
	"github.com/stretchr/testify/assert"
)

func genTimeStamps(n int) []interface{} {
	res := []interface{}{}
	for i := 0; i < n; i++ {
		res = append(res, fmt.Sprintf("%v", i))
	}
	return res
}

func genValues(n int, m float64) []interface{} {
	res := []interface{}{}
	for i := 0; i < n; i++ {
		res = append(res, float64(i)*m)
	}
	return res
}

func TestRemoveIndices(t *testing.T) {
	ts := genTimeStamps(10)
	values := map[string][]interface{}{
		"a": genValues(10, 2),
		"b": genValues(10, 3),
	}
	ta := timeseries.FromData(ts, values)

	res := ta.RemoveIndices([]int{3, 7, 8})

	assert.Equal(t, 7, len(res.TimeStamp()))
	assert.EqualValues(t, []interface{}{"0", "1", "2", "4", "5", "6", "9"}, res.TimeStamp())
	assert.EqualValues(t, []string{"a", "b"}, res.ColNames())
	assert.EqualValues(t, []interface{}{0.0, 2.0, 4.0, 8.0, 10.0, 12.0, 18.0}, res.AsSlice("a"))
	assert.EqualValues(t, []interface{}{0.0, 3.0, 6.0, 12.0, 15.0, 18.0, 27.0}, res.AsSlice("b"))
}

func TestChangeTimeStamp(t *testing.T) {
	ts := genTimeStamps(5)
	values := map[string][]interface{}{
		"a": genValues(5, 2),
		"b": genValues(5, 3),
	}
	ta := timeseries.FromData(ts, values)

	ta.ChangeTimeStamp("3", "10")
	assert.Equal(t, 5, len(ta.TimeStamp()))
	assert.EqualValues(t, []interface{}{"0", "1", "2", "10", "4"}, ta.TimeStamp())
}
