package handlers

import (
	"encoding/json"
	"fmt"
	charitydto "holyways/dto/charity"
	dto "holyways/dto/result"
	"holyways/models"
	"holyways/repositories"
	"net/http"
	"os"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	//"github.com/go-playground/validator/v10"
	//"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCharity struct {
	CharityRepository repositories.CharityRepository
}

func HandlerCharity(CharityRepository repositories.CharityRepository) *handlerCharity {
	return &handlerCharity{CharityRepository}
}

func (h *handlerCharity) FindCharities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	charity, err := h.CharityRepository.FindCharities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: charity}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCharity) GetCharity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var charity models.Charity
	charity, err := h.CharityRepository.GetCharity(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: charity}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCharity) GetCharityByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["user_id"].(float64))

	var charity []models.Charity
	charity, err := h.CharityRepository.GetCharityByUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: charity}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCharity) CreateCharity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["user_id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)



	// request := charitydto.CreateCharityRequest{
	// 	Title:       ,
	// 	Description: ,
	// 	Goal:        goal,
	// 	UserID:      userId,
	// }

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filename, uploader.UploadParams{Folder: "holyways"})

	if err != nil {
		fmt.Println(err.Error())
	}

	goal, _ := strconv.Atoi(r.FormValue("goal"))
	userID, _ := strconv.Atoi(r.FormValue("user_id"))

	charity := models.Charity{
		Title:       r.FormValue("title"),
		Image:       resp.SecureURL,
		Goal:        goal,
		Description: r.FormValue("description"),
		Status:      "Running",
		UserID:      userID,
		Donation:    0,
	}

	charity, err = h.CharityRepository.CreateCharity(charity)

	if err !=nil{
		w. WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	charity, _ = h.CharityRepository.GetCharity(charity.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: charity}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCharity) UpdateCharity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(charitydto.UpdateCharityRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "err.Error()"}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	charity, err := h.CharityRepository.GetCharity(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "aa.Error()"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Status != "" {
		charity.Status = request.Status
	}

	data, err := h.CharityRepository.UpdateCharity(charity, id)
	fmt.Println("charity Status : ", charity.Status)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err.sError()"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCharity) DeleteCharity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	charity, err := h.CharityRepository.GetCharity(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CharityRepository.DeleteCharity(id, charity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
