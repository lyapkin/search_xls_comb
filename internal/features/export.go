package features

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/lyapkin/search_comg_xls/internal/state"
	"github.com/xuri/excelize/v2"
)

func Export(data []state.Row) error {
	if len(data) == 0 {
		return errors.New("нет данных для экспорта")
	}

	folder, err := getFolderOrCreate()
	if err != nil {
		return err
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
			// TODO: handle error
		}
	}()

	sheetName := f.GetSheetName(0)
	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}
	headerCell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return err
	}
	header := make([]any, 0)
	header = append(header, "№", "Дата")
	for i := 1; i <= len(data[0].Data); i++ {
		header = append(header, fmt.Sprint(i))
	}
	header = append(header, "Сумма", "Чет", "Нечет", "Кол-во")
	sw.SetRow(headerCell, header)

	for i, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, i+2)
		if err != nil {
			return err
		}
		n := row.Number
		date := row.Date
		sum := row.Sum
		even := row.Even
		odd := row.Odd
		count := row.Count
		r := []any{n, date}
		for _, val := range row.Data {
			r = append(r, val)
		}
		r = append(r, sum, even, odd)

		if count != 0 {
			r = append(r, count)
		} else {
			r = append(r, nil)
		}

		sw.SetRow(cell, r)
	}

	if err := sw.Flush(); err != nil {
		return err
	}

	inputLength := fmt.Sprint(len(data[0].Data))
	fileName := time.Now().Format("02-01-06_15-04-05") + "_" + inputLength + ".xlsx"
	if err := f.SaveAs(filepath.Join(folder, fileName)); err != nil {
		return err
	}

	return nil
}

func getFolderOrCreate() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	currentDir := filepath.Dir(exePath)

	newFolderPath := filepath.Join(currentDir, "результаты")

	err = os.MkdirAll(newFolderPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Ошибка при создании папки: %v\n", err)
		return "", err
	}

	return newFolderPath, nil
}
