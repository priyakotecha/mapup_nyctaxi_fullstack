package repository

import (
	"database/sql"
	"fmt"
	"nyctaxi_mapup/pkg/config"
	"nyctaxi_mapup/pkg/model"
)

var db *sql.DB

func InitializeDB() error {
	connStr := config.GetPostgresConnectionString()

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database:", err)
	}

	fmt.Println("Successfully connected to database!")
	return nil
}

func GetTaxiData() ([]model.TaxiData, error) {
	var results []model.TaxiData
	rows, err := db.Query("SELECT vendor_id, pickup_datetime, dropoff_datetime, passenger_count, trip_distance, rate_code, pu_location_id, do_location_id, payment_type, fare_amount FROM nyc_taxi_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TaxiData
		err := rows.Scan(&data.VendorID, &data.PickupDatetime, &data.DropoffDatetime, &data.PassengerCount, &data.TripDistance, &data.RateCode, &data.PULocationID, &data.DOLocationID, &data.PaymentType, &data.FareAmount)
		if err != nil {
			return nil, err
		}
		results = append(results, data)
	}
	return results, nil
}
