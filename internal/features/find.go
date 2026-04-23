package features

import (
	"errors"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Find(input map[string]struct{}, cols string) ([][]string, error) {
	if len(input) == 0 {
		return nil, errors.New("не указаны данные для поиска")
	}

	f, err := excelize.OpenFile("db.xlsx")
	if err != nil {
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
			// TODO: handle error
		}
	}()

	rows, err := f.Rows(cols)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			// TODO: handle error
			log.Fatal(err)
		}
	}()

	isFirstRow := true

	result := make([][]string, 0)
	for rows.Next() {
		if isFirstRow {
			isFirstRow = false
			continue
		}

		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		length, err := strconv.Atoi(cols)
		if err != nil {
			return nil, err
		}

		row = row[1:]

		if matches(row[1:length+1], input) {
			result = append(result, row)
		}
	}

	return result, nil
}

func matches(row []string, input map[string]struct{}) bool {
	matchCount := 0
	for _, v := range row {
		if _, ok := input[v]; ok {
			matchCount++
		}
	}
	match := matchCount == len(input)
	return match
}
