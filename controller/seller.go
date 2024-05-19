package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var seller model.Seller
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&seller); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	// convert contact number to int from string
	saveErr := seller.Create()
	if saveErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpresp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Seller registered"})
}

func AddProfile(w http.ResponseWriter, r *http.Request) {
	var profile model.SellerProfile
	err := r.ParseMultipartForm(10 << 20) //10MB max file size
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "Error parsing multipart form")
		return
	}

	file, _, err := r.FormFile("profile_picture")
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "Error retrieving profile picture")
		return
	}
	defer file.Close()

	// Read the file data into a []byte slice
	profilePictureBytes, err := io.ReadAll(file)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusInternalServerError, "Error reading profile picture data")
		return
	}

	uid := r.FormValue("userid")
	uidInt, _ := strconv.Atoi(uid)

	profile.UserId = uidInt
	profile.ProfilePicture = profilePictureBytes

	saveErr := profile.Add()
	if saveErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpresp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Profile uploaded"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var seller model.Seller

	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := seller.Get()
	if getErr != nil {
		httpresp.RespondWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}
	authCookie := http.Cookie{
		Name:    "session-cookie",
		Value:   "hello-world",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	http.SetCookie(w, &authCookie)

	contactNumber := http.Cookie{
		Name:  "contactnumber",
		Value: strconv.Itoa(seller.ContactNumber),
	}
	http.SetCookie(w, &contactNumber)

	httpresp.RespondWithJson(w, http.StatusOK, map[string]string{"message": "success"})
}

func GetSeller(w http.ResponseWriter, r *http.Request) {
	pnumber:=mux.Vars(r)["phonenumber"]
	phonenumber, numErr:=getPnumber(pnumber)
	if numErr!=nil{
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	p:=model.SellerWithProfile{ContactNumber: phonenumber}
	getErr:=p.Read()
	if getErr!=nil{
		switch getErr{
		case sql.ErrNoRows:
			httpresp.RespondWithError(w, http.StatusNotFound, "user not found")
		default:
			httpresp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpresp.RespondWithJson(w, http.StatusOK, p)
}

func getPnumber(pnumberParam string) (int, error){
	phonenumber, err:=strconv.Atoi(pnumberParam)
	if err != nil{
		return 0, err
	}
	return phonenumber, nil
}
