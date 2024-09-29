package model

type TaxiData struct {
	VendorID        string  `json:"vendor_id"`
	PickupDatetime  string  `json:"pickup_datetime"`
	DropoffDatetime string  `json:"dropoff_datetime"`
	PassengerCount  int     `json:"passenger_count"`
	TripDistance    float64 `json:"trip_distance"`
	RateCode        int     `json:"rate_code"`
	PULocationID    int     `json:"pu_location_id"`
	DOLocationID    int     `json:"do_location_id"`
	PaymentType     string  `json:"payment_type"`
	FareAmount      float64 `json:"fare_amount"`
}
