package controller

import (
	"database/sql"
	"io"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	pnumberint "myapp/utils/pnumberInt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	pnumber := mux.Vars(r)["phonenumber"]
	phonenumber, numErr := pnumberint.GetPnumber(pnumber)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	var product model.Product
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "Error parsing multipart form")
		return
	}

	product.ContactNumber = phonenumber
	product.ProductTitle = r.FormValue("ptitle")
	product.ActualPrice, _ = strconv.Atoi(r.FormValue("aprice"))
	product.SellingPrice, _ = strconv.Atoi(r.FormValue("sprice"))
	product.Category = r.FormValue("category")
	product.ProductDescription = r.FormValue("pdescription")

	// Process product images
	imageFields := []string{"productimg1", "productimg2", "productimg3", "productimg4"}
	imageBytes := []*[]byte{&product.Productimg1, &product.Productimg2, &product.Productimg3, &product.Productimg4}

	for i, field := range imageFields {
		file, _, err := r.FormFile(field)
		if err != nil {
			if err == http.ErrMissingFile {
				continue //skip is the file is not provided and process the next image
			}
			httpresp.RespondWithError(w, http.StatusBadRequest, "error retrieving "+field)
			return
		}
		defer file.Close()

		imgBytes, err := io.ReadAll(file)
		if err != nil {
			httpresp.RespondWithError(w, http.StatusInternalServerError, "error reading "+field+" data")
			return
		}

		// Assign the read bytes to the corresponding Product field
		*imageBytes[i] = imgBytes
	}

	saveErr := product.Add()
	if saveErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpresp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Product uploaded"})
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, geterr := model.GetAllProducts()
	if geterr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, geterr.Error())
		return
	}
	httpresp.RespondWithJson(w, http.StatusOK, products)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	pnumber := mux.Vars(r)["phonenumber"]
	phonenumber, numErr := pnumberint.GetPnumber(pnumber)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	p := model.SellerProduct{ContactNumber: phonenumber}
	products, getErr := p.GetSellerProducts()
	if getErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
	}
	httpresp.RespondWithJson(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	pid := mux.Vars(r)["pid"]
	productid, numErr := pnumberint.GetPnumber(pid)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	p := model.SellerProduct{ProductId: productid}
	getErr := p.GetProduct()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpresp.RespondWithError(w, http.StatusNotFound, "product not found")
		default:
			httpresp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	// Set Product id cookie
	productId := http.Cookie{
		Name:    "product_id",
		Value:   strconv.Itoa(p.ProductId),
		Expires: time.Now().Add(10 * time.Minute),
	}
	http.SetCookie(w, &productId)
	httpresp.RespondWithJson(w, http.StatusOK, p)
}

