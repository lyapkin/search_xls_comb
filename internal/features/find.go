package features

import (
	"errors"
	"log"
	"strconv"

	"github.com/lyapkin/search_comg_xls/internal/state"
	"github.com/xuri/excelize/v2"
)

func Find(input map[string]struct{}, cols string) ([]state.Row, error) {
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

	rowIdx := 0
	result := make([]state.Row, 0)
	for rows.Next() {
		rowIdx++

		if rowIdx == 1 {
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

		if len(row) < length+2 {
			continue
		}

		if matches(row[2:length+2], input) {
			n, err := strconv.Atoi(row[0])
			if err != nil {
				n = 0
			}

			date := row[1]

			data := make([]int, 0)
			for _, val := range row[2 : length+2] {
				v, err := strconv.Atoi(val)
				if err != nil {
					return nil, err
				}
				data = append(data, v)
			}

			sum, err := strconv.Atoi(row[2+length])
			if err != nil {
				return nil, err
			}

			even, err := strconv.Atoi(row[2+length+1])
			if err != nil {
				return nil, err
			}

			odd, err := strconv.Atoi(row[2+length+2])
			if err != nil {
				return nil, err
			}

			r := state.Row{
				Number: n,
				Date:   date,
				Data:   data,
				Sum:    sum,
				Even:   even,
				Odd:    odd,
			}
			result = append(result, r)
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
