CREATE TABLE nyc_taxi_data (
    vendor_id VARCHAR(50),
    pickup_datetime TIMESTAMP,
    dropoff_datetime TIMESTAMP,
    passenger_count INT,
    trip_distance DOUBLE PRECISION,
    rate_code INT,
    pu_location_id INT,
    do_location_id INT,
    payment_type VARCHAR(50),
    fare_amount DOUBLE PRECISION
);