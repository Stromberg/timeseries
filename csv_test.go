package timeseries

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCsvFromFile(t *testing.T) {
	config := NewCsvConfig("Date")
	config.Maps["Date"] = func(s string) (interface{}, error) {
		ds, err := time.Parse("Jan 06", s)
		if err != nil {
			return ds, err
		}
		return ds.Format("2006-01-02"), nil
	}
	config.Maps["Price"] = func(s string) (interface{}, error) {
		return strconv.ParseFloat(s, 64)
	}

	ta, err := FromCsvFile("test.csv", config)
	assert.NoError(t, err)
	assert.EqualValues(t, []string{"Change", "High", "Low", "Open", "Price", "Volume"}, ta.ColNames())
	assert.EqualValues(t, []string{"Price", "Volume"}, ta.Values("Volume", "Price").ColNames())

	assert.EqualValues(t, []interface{}{"2005-06-01", "2005-07-01", "2005-08-01", "2005-09-01", "2005-10-01"}, ta.TimeStamp())
	assert.EqualValues(t, []interface{}{96.70, 93.11, 95.99, 92.20, 89.78}, ta.AsSlice("Price"))

	ta.Rename("Price", "Value")
	assert.EqualValues(t, []string{"Change", "High", "Low", "Open", "Value", "Volume"}, ta.ColNames())

	fmt.Println(ta)
}
