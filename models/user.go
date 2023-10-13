package models

type Gender string

const (
	Male   Gender = "male"
	Female        = "female"
	Other         = "other"
)

type User struct {
	Name        string
	Gender      Gender
	PhoneNumber string
	Pincode     Pincode
}
