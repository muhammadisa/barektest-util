package csvr

import (
	"encoding/csv"
	"errors"
	"net/http"
	"reflect"
	"strings"
)

var (
	ErrorFormatNotSame      = errors.New("csv format is invalid, doesn't match with struct")
	ErrorFormatFieldsLength = errors.New("length of csv header and struct field doesn't same")
	ErrorRetrieveCSVFromURL = errors.New("failed to retrieve csv from url")
	ErrorReadWhileReadCSV   = errors.New("failed to read csv ")
)

type CSVHeader string

func (ch CSVHeader) String() string {
	return string(ch)
}

type CSVResult struct {
	Header CSVHeader
	Values []string
	Length int64
}

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrorRetrieveCSVFromURL
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, ErrorReadWhileReadCSV
	}

	return data, nil
}

func convertHeaderTitle(headerTitle string) string {
	headerTitles := strings.Split(headerTitle, "_")
	headerTitleLength := len(headerTitles)
	converted := make([]string, headerTitleLength)
	for i, ht := range headerTitles {
		converted[i] = strings.Title(strings.ToLower(ht))
	}
	return strings.Join(converted, "")
}

func (ch CSVHeader) BindValues(targetValidation interface{}, value string) error {
	structDetail := reflect.ValueOf(targetValidation).Elem()

	values := strings.Split(value, ",")
	headers := strings.Split(ch.String(), ",")
	fields := structDetail.NumField()

	if fields != len(headers) {
		return ErrorFormatFieldsLength
	}

	for index, header := range headers {
		fieldName := convertHeaderTitle(header)
		varName := structDetail.Type().Field(index).Name
		if varName != fieldName {
			return ErrorFormatNotSame
		}
		structDetail.FieldByName(fieldName).SetString(values[index])
	}
	return nil
}

func DownloadCSV(url string) (res CSVResult, err error) {
	var values []string
	var header string
	data, err := readCSVFromUrl(url)
	if err != nil {
		return res, err
	}
	for idx, row := range data {
		if idx == 0 {
			header = row[0]
			continue
		}
		values = append(values, row[0])
	}
	res.Header = CSVHeader(header)
	res.Length = int64(len(values))
	res.Values = values
	return res, nil
}
