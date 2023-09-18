package service

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/jszwec/csvutil"
	"github.com/shplume/csv-import/model"
)

func Load() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("!!!!!!!!!!!!!!!! Service init err: %s\n", err.Error())
		return
	}

	csvDir := path.Join(dir, model.CSVDir)
	if _, err := os.Stat(csvDir); err != nil {
		fmt.Printf("!!!!!!!!!!!!!!!! Service init err: %s\n", err.Error())
		return
	}

	err = loadCSV(csvDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	storeData(data)
}

func loadCSV(dir string) error {
	file, err := os.Open(path.Join(dir, model.DataFile))
	if err != nil {
		return fmt.Errorf("Load with error: %s", err.Error())
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Load with error: %s", err.Error())
	}

	if err = csvutil.Unmarshal(content, &data); err != nil {
		return fmt.Errorf("Load with error: %s", err.Error())
	}

	return nil
}

func storeData(records []model.VideoData) {
	if err := dbConn.Model(&model.VideoData{}).Create(records).Error; err != nil {
		fmt.Printf("Store with error: %s\n", err.Error())
	}
}
