package models

import "database/sql"

type Reading struct {
	TimeStamp          sql.NullString  `json:"timeStamp"`
	OutsideTemperature sql.NullFloat64 `json:"outsideTemperature"`
	OutsideHumidity    sql.NullFloat64 `json:"outsideHumidity"`
	CarTemperature     sql.NullFloat64 `json:"carTemperature"`
	CarHumidity        sql.NullFloat64 `json:"carHumidity"`
}

type DayofWeek struct {
	DayofWeek string `json:"dayofWeek"`
}
