package timeseries

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

type CsvConfig struct {
	Comma     rune
	TimeStamp string
	Maps      map[string]func(string) (interface{}, error)
}

func NewCsvConfig(timestamp string) *CsvConfig {
	return &CsvConfig{
		Comma:     ',',
		TimeStamp: timestamp,
		Maps:      make(map[string]func(string) (interface{}, error), 0),
	}
}

func FromCsv(in string, config *CsvConfig) (res *TimeArray, err error) {
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = config.Comma
	r.LazyQuotes = true

	raw, err := r.ReadAll()
	if err != nil {
		return
	}

	fmt.Println(raw)

	values := make(map[string][]interface{}, 0)
	var timestamp []interface{}
	rowsRaw := raw[1:len(raw)]
	for ci, k := range raw[0] {
		value := make([]interface{}, len(rowsRaw))
		if mapping, ok := config.Maps[k]; ok {
			for i, v := range rowsRaw {
				value[i], err = mapping(v[ci])
				if err != nil {
					return
				}
			}
		} else {
			for i, v := range rowsRaw {
				value[i] = v[ci]
			}
		}
		if k != config.TimeStamp {
			values[k] = value
		} else {
			timestamp = value
		}
	}

	if len(timestamp) == 0 {
		err = fmt.Errorf("Missing Timestamp column \"%v\"", config.TimeStamp)
		return
	}

	res = FromData(timestamp, values)

	return
}

func FromCsvFile(filename string, config *CsvConfig) (res *TimeArray, err error) {
	csv, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	return FromCsv(string(csv), config)
}
