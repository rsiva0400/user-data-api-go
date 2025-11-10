package models

type User struct {
	Userid int32  `json:"userid"`
	Name   string `json:"name"`
	Dob    string `json:"dob"`
	Age    int    `json:"age"`
}

type UserRequest struct {
	Name string `json:"name" validate:"required,min=2,alpha"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
