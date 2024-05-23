package model

import (
	"myapp/dataStore/postgres"
)

type Product struct {
	ContactNumber      int    `json:"cnumber"`
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
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	ContactNumber      int    `json:"contact_number"`
	Email              string `json:"email"`
}

type ProductWithId struct {
	ProductId          int    `json:"productid"`
	ContactNumber      int    `json:"cnumber"`
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

const queryInsertproduct = "INSERT INTO product(contact_number, product_img1, product_img2, product_img3, product_img4, product_title, actual_price, selling_price, category, product_description) Values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) Returning contact_number;"

func (p *Product) Add() error {
	row := postgres.Db.QueryRow(queryInsertproduct, p.ContactNumber, p.Productimg1, p.Productimg2, p.Productimg3, p.Productimg4, p.ProductTitle, p.ActualPrice, p.SellingPrice, p.Category, p.ProductDescription)
	err := row.Scan(&p.ContactNumber)
	return err
}

func GetAllProducts() ([]SellerProduct, error) {
	query := "SELECT p.product_id, p.contact_number, p.product_img1, p.product_img2, p.product_img3, p.product_img4, p.product_title, p.actual_price, p.selling_price, p.category, p.product_description, s.contactnumber, s.firstname, s.lastname, s.email FROM  product p JOIN seller s ON p.contact_number=s.contactnumber"
	rows, err := postgres.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []SellerProduct{}

	for rows.Next() {
		var ps SellerProduct
		err := rows.Scan(&ps.ProductId, &ps.ContactNumber, &ps.Productimg1, &ps.Productimg2, &ps.Productimg3, &ps.Productimg4, &ps.ProductTitle, &ps.ActualPrice, &ps.SellingPrice, &ps.Category, &ps.ProductDescription, &ps.ContactNumber, &ps.FirstName, &ps.LastName, &ps.Email)
		if err != nil {
			return nil, err
		}
		products = append(products, ps)
	}
	return products, nil
}

func (s *SellerProduct) GetSellerProducts() ([]SellerProduct, error) {
	query := "SELECT p.product_id, p.contact_number, p.product_img1, p.product_img2, p.product_img3, p.product_img4, p.product_title, p.actual_price, p.selling_price, p.category, p.product_description, s.contactnumber, s.firstname, s.lastname, s.email FROM  product p JOIN seller s ON p.contact_number=s.contactnumber WHERE s.contactnumber=$1"
	rows, err := postgres.Db.Query(query, s.ContactNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []SellerProduct{}

	for rows.Next() {
		var sp SellerProduct
		err := rows.Scan(&sp.ProductId, &sp.ContactNumber, &sp.Productimg1, &sp.Productimg2, &sp.Productimg3, &sp.Productimg4, &sp.ProductTitle, &sp.ActualPrice, &sp.SellingPrice, &sp.Category, &sp.ProductDescription, &sp.ContactNumber, &sp.FirstName, &sp.LastName, &sp.Email)
		if err != nil {
			return nil, err
		}
		products = append(products, sp)
	}
	return products, nil
}

const queryGetProduct = "SELECT p.product_id, p.contact_number, p.product_img1, p.product_img2, p.product_img3, p.product_img4, p.product_title, p.actual_price, p.selling_price, p.category, p.product_description, s.contactnumber, s.firstname, s.lastname, s.email FROM  product p JOIN seller s ON p.contact_number=s.contactnumber WHERE p.product_id=$1"

func (p *SellerProduct) GetProduct() error {
	return postgres.Db.QueryRow(queryGetProduct, p.ProductId).Scan(&p.ProductId, &p.ContactNumber, &p.Productimg1, &p.Productimg2, &p.Productimg3, &p.Productimg4, &p.ProductTitle, &p.ActualPrice, &p.SellingPrice, &p.Category, &p.ProductDescription, &p.ContactNumber, &p.FirstName, &p.LastName, &p.Email)
}

const queryDeleteProduct = "DELETE FROM product WHERE product_id=$1;"

func (p *ProductWithId) Delete() error {
	_,err:=postgres.Db.Exec(queryDeleteProduct, p.ProductId)
	return err
}
