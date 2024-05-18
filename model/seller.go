package model

import "myapp/dataStore/postgres"

type Seller struct {
	FirstName     string `json:"fname"`
	LastName      string `json:"lname"`
	ContactNumber int    `json:"cnumber"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Password      string `json:"password"`
}

const queryInsertSeller = "INSERT INTO seller(firstname, lastname, contactnumber, email, gender, password) VALUES($1, $2, $3, $4, $5, $6) Returning contactnumber;"

func (s *Seller) Create() error {
	row := postgres.Db.QueryRow(queryInsertSeller, s.FirstName, s.LastName, s.ContactNumber, s.Email, s.Gender, s.Password)
	err := row.Scan(&s.ContactNumber)
	return err
}

type SellerProfile struct {
	UserId         int    `json:"userid"`
	ProfilePicture []byte `json:"profilepicture"`
}

const queryInsertSellerProfile = "INSERT INTO seller_profile(user_id, profile_picture) VALUES($1, $2) Returning profile_picture;"

func (profile *SellerProfile) Add() error {
	row := postgres.Db.QueryRow(queryInsertSellerProfile, profile.UserId, profile.ProfilePicture)
	err := row.Scan(&profile.ProfilePicture)
	return err
}

const queryGetSeller="SELECT contactnumber, password FROM seller WHERE contactnumber=$1 and password=$2;"
func (s *Seller) Get() error{
	return postgres.Db.QueryRow(queryGetSeller, s.ContactNumber, s.Password).Scan(&s.ContactNumber, &s.Password)
}