package model

type TaxiData struct {
	VendorID        string  `parquet:"name=vendor_id, type=BYTE_ARRAY"`
	PickupDatetime  string  `parquet:"name=pickup_datetime, type=BYTE_ARRAY"`
	DropoffDatetime string  `parquet:"name=dropoff_datetime, type=BYTE_ARRAY"`
	PassengerCount  int     `parquet:"name=passenger_count, type=INT32"`
	TripDistance    float64 `parquet:"name=trip_distance, type=DOUBLE"`
	RateCode        int     `parquet:"name=rate_code, type=INT32"`
	PULocationID    int     `parquet:"name=pu_location_id, type=INT32"`
	DOLocationID    int     `parquet:"name=do_location_id, type=INT32"`
	PaymentType     string  `parquet:"name=payment_type, type=BYTE_ARRAY"`
	FareAmount      float64 `parquet:"name=fare_amount, type=DOUBLE"`
}

type UserCache struct {
	Username string
	Role     string
	Password string
	Token    string
}
