CREATE TABLE IF NOT EXISTS trains (
    id SERIAL PRIMARY KEY,
    ride SERIAL NOT NULL,
    departure SERIAL NOT NULL,
    arrival SERIAL NOT NULL,
    cost DECIMAL NOT NULL,
    departure_time TIME NOT NULL,
    arrival_time TIME NOT NULL
);

COPY trains (ride, departure, arrival, cost, departure_time, arrival_time)
FROM '/etc/postgresql/data/schedule.csv'
DELIMITER ';'
CSV HEADER;