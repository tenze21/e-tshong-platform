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
	ContactNumber         int    `json:"cnumber"`
	ProfilePicture []byte `json:"profilepicture"`
}

type SellerWithProfile struct {
	ProfilePicture []byte `json:"profilepicture"`
	FirstName      string `json:"fname"`
	LastName       string `josn:"lname"`
	ContactNumber  int    `json:"cnumber"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
}

const queryInsertSellerProfile = "INSERT INTO seller_profile(contact_number, profile_picture) VALUES($1, $2) Returning profile_picture;"

func (profile *SellerProfile) Add() error {
	row := postgres.Db.QueryRow(queryInsertSellerProfile, profile.ContactNumber, profile.ProfilePicture)
	err := row.Scan(&profile.ProfilePicture)
	return err
}

const queryGetSeller = "SELECT contactnumber, password FROM seller WHERE contactnumber=$1 and password=$2;"

func (s *Seller) Get() error {
	return postgres.Db.QueryRow(queryGetSeller, s.ContactNumber, s.Password).Scan(&s.ContactNumber, &s.Password)
}

const queryGetSellerWithProfile = "SELECT p.profile_picture, s.firstname, s.lastname, s.contactnumber, s.email, s.gender From seller s JOIN seller_profile p ON s.contactnumber=p.contact_number WHERE s.contactnumber=$1 "

func (p *SellerWithProfile) Read() error {
	return postgres.Db.QueryRow(queryGetSellerWithProfile, p.ContactNumber).Scan(&p.ProfilePicture, &p.FirstName, &p.LastName, &p.ContactNumber, &p.Email, &p.Gender)
}

const queryGetSellerDetails= "SELECT firstname, lastname, contactnumber, email, gender FROM seller WHERE contactnumber=$1;"
func (s *Seller) GetDetails() error{
	return postgres.Db.QueryRow(queryGetSellerDetails, s.ContactNumber).Scan(&s.FirstName, &s.LastName, &s.ContactNumber, &s.Email, &s.Gender)
}

const queryUpdateProfile = "UPDATE seller_profile SET profile_picture=$1 WHERE contact_number=$2"
func (p *SellerProfile) UpdatePic() error {
    _, err := postgres.Db.Exec(queryUpdateProfile, p.ProfilePicture, p.ContactNumber)
    return err
}


const queryUpdateDetails= "UPDATE seller SET firstname=$1, lastname=$2, contactnumber=$3, email=$4, gender=$5 WHERE contactnumber=$6;"
func (s *Seller) UpdateDetails(oldpnumber int) error{
	_,err:= postgres.Db.Exec(queryUpdateDetails, s.FirstName, s.LastName, s.ContactNumber, s.Email, s.Gender, oldpnumber)
    return err
}