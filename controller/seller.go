package controller

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	pnumberint "myapp/utils/pnumberInt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var seller model.Seller
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&seller); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	hashedPassword, err := httpresp.HashPassword(seller.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	seller.Password = hashedPassword

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

	cnumber := r.FormValue("cnumber")
	cnumberInt, _ := strconv.Atoi(cnumber)

	profile.ContactNumber = cnumberInt
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

	storedHashedPassword, pass_err := model.GetUserHashedPassword(seller.ContactNumber)

	if pass_err != nil {
		if pass_err == sql.ErrNoRows {
			httpresp.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		} else {
			httpresp.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	by_err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(seller.Password))

	if by_err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
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
	if !VerifyCookie(w, r) {
		return
	}
	pnumber := mux.Vars(r)["phonenumber"]
	phonenumber, numErr := pnumberint.GetPnumber(pnumber)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	p := model.SellerWithProfile{ContactNumber: phonenumber}
	getErr := p.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpresp.RespondWithError(w, http.StatusNotFound, "user not found")
		default:
			httpresp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpresp.RespondWithJson(w, http.StatusOK, p)
}

func GetSellerDetails(w http.ResponseWriter, r *http.Request) {
	pnumber := mux.Vars(r)["phonenumber"]
	phonenumber, numErr := pnumberint.GetPnumber(pnumber)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	p := model.Seller{ContactNumber: phonenumber}
	getErr := p.GetDetails()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpresp.RespondWithError(w, http.StatusNotFound, "user not found")
		default:
			httpresp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpresp.RespondWithJson(w, http.StatusOK, p)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	pnumber := mux.Vars(r)["phonenumber"]
	phonenumber, numErr := pnumberint.GetPnumber(pnumber)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}

	var profileData struct {
		ProfilePicture string `json:"profilepicture"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&profileData); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	// Decode base64 image data
	imageData, err := base64.StdEncoding.DecodeString(profileData.ProfilePicture[len("data:image/jpeg;base64,"):])
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid base64 data")
		return
	}

	sellerProfile := &model.SellerProfile{
		ContactNumber:  phonenumber,
		ProfilePicture: imageData,
	}

	updateErr := sellerProfile.UpdatePic()
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpresp.RespondWithError(w, http.StatusNotFound, "user not found")
			return
		default:
			httpresp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpresp.RespondWithJson(w, http.StatusOK, sellerProfile)
	}
}

func UpdateSellerDetails(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	pnumber := mux.Vars(r)["phonenumber"]
	oldpnumber, numErr := pnumberint.GetPnumber(pnumber)
	if numErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, numErr.Error())
		return
	}
	var seller model.Seller
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&seller); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	updateErr := seller.UpdateDetails(oldpnumber)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpresp.RespondWithError(w, http.StatusNotFound, "user not found")
		default:
			httpresp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpresp.RespondWithJson(w, http.StatusOK, seller)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// removing the authorization cookie on logout 
	http.SetCookie(w, &http.Cookie{
		Name:    "session-cookie",
		Expires: time.Now(),
	})

	// removing the contactnumber cookie
	http.SetCookie(w, &http.Cookie{
		Name: "contactnumber",
		Expires: time.Now(),
	})
	httpresp.RespondWithJson(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			httpresp.RespondWithError(w, http.StatusSeeOther, "cookie not found")
			return false
		}
		httpresp.RespondWithError(w, http.StatusInternalServerError, "internal server error")
		return false
	}

	if cookie.Value != "hello-world" {
		httpresp.RespondWithError(w, http.StatusSeeOther, "cookie does not match")
		return false
	}
	return true
}
