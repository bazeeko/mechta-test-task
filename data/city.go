package data

import (
	"database/sql"
	"fmt"
)

type City struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	CountryCode string `json:"country_code"`
}

type CityRepository struct {
	*sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{db}
}

func (db *CityRepository) GetCities() ([]City, error) {
	cities := []City{}

	rows, err := db.Query(`SELECT id, name, code, country_code FROM cities`)
	if err != nil {
		return nil, fmt.Errorf("GetCities: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		city := City{}

		err := rows.Scan(&city.Id, &city.Name, &city.Code, &city.CountryCode)
		if err != nil {
			return nil, fmt.Errorf("GetCities: %w", err)
		}

		cities = append(cities, city)
	}

	return cities, nil
}

func (db *CityRepository) AddCity(city City) error {
	_, err := db.Exec(`INSERT INTO cities (name, code, country_code) VALUES ($1, $2, $3)`,
		city.Name, city.Code, city.CountryCode)

	if err != nil {
		return fmt.Errorf("GetCities: %w", err)
	}

	return nil
}

func (db *CityRepository) GetCityById(id int) (City, error) {
	city := City{}

	err := db.QueryRow(`SELECT id, name, code, country_code FROM cities WHERE id=$1`, id).
		Scan(&city.Id, &city.Name, &city.Code, &city.CountryCode)

	if err != nil {
		return city, fmt.Errorf("GetCityById: %w", err)
	}

	return city, nil
}

func (db *CityRepository) DeleteCityById(id int) error {
	_, err := db.Exec(`DELETE FROM cities WHERE id=$1`, id)

	if err != nil {
		return fmt.Errorf("DeleteCityById: %w", err)
	}

	return nil
}

func (db *CityRepository) UpdateCityById(id int, city City) error {
	_, err := db.Exec(`UPDATE cities SET name=$1, code=$2, country_code=$3 WHERE id=$4`,
		city.Name, city.Code, city.CountryCode, id)

	if err != nil {
		return fmt.Errorf("UpdateCityById: %w", err)
	}

	return nil
}
