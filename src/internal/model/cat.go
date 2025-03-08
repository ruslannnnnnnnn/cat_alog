package model

import "time"

type Cat struct {
	Id          string
	Name        string
	DateOfBirth time.Time
	ImageUrl    string
}
