package model

import (
	"myapp/dataStore/postgres"
)

type Product struct {
	UserId             int    `json:"userid"`
	Productimg1        []byte `json:"productimg1"`
	Productimg2        []byte `json:"productimg2"`
	Productimg3        []byte `json:"productimg3"`
	Productimg4        []byte `json:"productimg4"`
	ProductTitle       string `json:"ptitle"`
	ActualPrice        int    `json:"aprice"`
	SellingPrice       int    `json:"sprice"`
	Category           string `json:"category"`
	ProductDescription string `json:"pdescription"`
}

type SellerProduct struct {
	ProductId          int    `json:"productid"`
	Productimg1        []byte `json:"productimg1"`
	Productimg2        []byte `json:"productimg2"`
	Productimg3        []byte `json:"productimg3"`
	Productimg4        []byte `json:"productimg4"`
	ProductTitle       string `json:"ptitle"`
	ActualPrice        int    `json:"aprice"`
	SellingPrice       int    `json:"sprice"`
	Category           string `json:"category"`
	ProductDescription string `json:"pdescription"`
	UserID             int    `json:"user_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	ContactNumber      int    `json:"contact_number"`
	Email              string `json:"email"`
}

const queryInsertproduct = "INSERT INTO product(user_id, product_img1, product_img2, product_img3, product_img4, product_title, actual_price, selling_price, category, product_description) Values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) Returning user_id;"

func (p *Product) Add() error {
	row := postgres.Db.QueryRow(queryInsertproduct, p.UserId, p.Productimg1, p.Productimg2, p.Productimg3, p.Productimg4, p.ProductTitle, p.ActualPrice, p.SellingPrice, p.Category, p.ProductDescription)
	err := row.Scan(&p.UserId)
	return err
}

func GetAllProducts() ([]SellerProduct, error) {
	query := "SELECT p.product_id,p.user_id,p.product_img1, p.product_img2, p.product_img3, p.product_img4, p.product_title, p.actual_price, p.selling_price, p.category, p.product_description, s.userid, s.firstname, s.lastname, s.contactnumber, s.email FROM  product p JOIN seller s ON p.user_id=s.userid"
	rows, err := postgres.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []SellerProduct{}

	for rows.Next() {
		var ps SellerProduct
		err := rows.Scan(&ps.ProductId, &ps.UserID, &ps.Productimg1, &ps.Productimg2, &ps.Productimg3, &ps.Productimg4, &ps.ProductTitle, &ps.ActualPrice, &ps.SellingPrice, &ps.Category, &ps.ProductDescription, &ps.UserID, &ps.FirstName, &ps.LastName, &ps.ContactNumber, &ps.Email)
		if err != nil {
			return nil, err
		}
		products = append(products, ps)
	}
	return products, nil
}
