package models

import "database/sql"

type Reading struct {
	TimeStamp                    sql.NullString  `json:"timeStamp"`
	OutsideTemperatureFahrenheit sql.NullFloat64 `json:"outsideTemperatureFahrenheit"`
	OutsideTemperatureCelcius    sql.NullFloat64 `json:"outsideTemperatureCelcius"`
	OutsideHumidity              sql.NullFloat64 `json:"outsideHumidity"`
	CarTemperatureFahrenheit     sql.NullFloat64 `json:"carTemperatureFahrenheit"`
	CarTemperatureCelcius        sql.NullFloat64 `json:"carTemperatureCelcius"`
	CarHumidity                  sql.NullFloat64 `json:"carHumidity"`
}

type DayofWeek struct {
	DayofWeek string `json:"dayofWeek"`
}
