package database

import (
	"database/sql"
	"fmt"
	"log"
	"trains/config"
	"trains/models"
	"trains/utils"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres() *Postgres {
	fmt.Println(config.PostgresPassword)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.PostgresHost, config.PostgresPort, config.PostgresUser,
		config.PostgresPassword, config.PostgresDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &Postgres{db}
}

func (p *Postgres) FetchAllRoutes() []models.Route {
	rows, err := p.Db.Query("SELECT departure, arrival FROM trains")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	route := new(models.Route)
	var routes []models.Route
	for rows.Next() {
		if err := rows.Scan(&route.Departure, &route.Arrival); err != nil {
			log.Fatal(err)
		}
		if !utils.Contains(routes, *route) {
			routes = append(routes, *route)
		}
	}
	return routes
}

func (p *Postgres) FindMinCostForTheRoute(route models.Route) float64 {
	req := fmt.Sprintf(`SELECT cost FROM trains WHERE departure = %s AND arrival = %s AND cost = (
		SELECT MIN(cost) FROM trains WHERE departure = %s AND arrival = %s);`,
		route.Departure, route.Arrival, route.Departure, route.Arrival)
	rows, err := p.Db.Query(req)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var cost float64
	if rows.Next() {
		if err = rows.Scan(&cost); err != nil {
			log.Fatal(err)
		}
	}
	return cost
}

func (p *Postgres) FindMinCostRidesForTheRoute(route models.Route) []int {
	req := fmt.Sprintf(`SELECT ride FROM trains WHERE departure = %s AND arrival = %s AND cost = (
		SELECT MIN(cost) FROM trains WHERE departure = %s AND arrival = %s);`,
		route.Departure, route.Arrival, route.Departure, route.Arrival)
	rows, err := p.Db.Query(req)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var rides []int
	var ride int
	for rows.Next() {
		if err = rows.Scan(&ride); err != nil {
			log.Fatal(err)
		}
		rides = append(rides, ride)
	}
	return rides
}
