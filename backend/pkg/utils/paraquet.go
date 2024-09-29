package utils

import (
	"fmt"
	"io"
	"net/http"
	"nyctaxi_mapup/pkg/config"
	"nyctaxi_mapup/pkg/model"
	"nyctaxi_mapup/pkg/repository"
	"os"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func DownloadParquetFile() error {
	url, file := config.GetParaquetFileURLAndPath()
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	outFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	return err
}

func ProcessParquetFile() error {
	fr, err := local.NewLocalFileReader("yellow_tripdata_2023-12.parquet")
	if err != nil {
		return fmt.Errorf("can't open file %v", err)
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(model.TaxiData), 4)
	if err != nil {
		return fmt.Errorf("can't create parquet reader %v", err)
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	taxiData := make([]model.TaxiData, num)
	if err = pr.Read(&taxiData); err != nil {
		return fmt.Errorf("can't read from parquet file %v", err)
	}

	for _, data := range taxiData {
		err := repository.InsertData(data)
		if err != nil {
			return fmt.Errorf("error inserting data into db: %v", err)
		}
	}
	return nil
}
