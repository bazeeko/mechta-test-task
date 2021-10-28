package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bazeeko/mechta-test-task/data"
	"github.com/gorilla/mux"
)

func (h *Handler) GetAllCities(w http.ResponseWriter, r *http.Request) {
	cities, err := h.Repository.GetCities()
	if err != nil {
		log.Printf("GetAllCities: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(cities)
	if err != nil {
		log.Printf("GetAllCities: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (h *Handler) AddSingleCity(w http.ResponseWriter, r *http.Request) {
	city := data.City{}

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		log.Printf("AddSingleCity: %s", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = h.Repository.AddCity(city)
	if err != nil {
		log.Printf("AddSingleCity: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetSingleCity(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	city, err := h.Repository.GetCityById(id)
	if err != nil {
		resp := make(map[string]string)
		resp["message"] = "City Not Found"

		json.NewEncoder(w).Encode(resp)
		return
	}

	err = json.NewEncoder(w).Encode(city)
	if err != nil {
		log.Printf("GetSingleCity: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateSingleCity(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	city := data.City{}

	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		log.Printf("UpdateSingleCity: %s", err)
		http.Error(w, "", http.StatusBadRequest)
	}

	err = h.Repository.UpdateCityById(id, city)
	if err != nil {
		log.Printf("UpdateSingleCity: %s", err)

		resp := make(map[string]string)
		resp["message"] = "City Not Found"

		json.NewEncoder(w).Encode(resp)
	}
}

func (h *Handler) DeleteSingleCity(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := h.Repository.DeleteCityById(id)
	if err != nil {
		log.Printf("DeleteSingleCity: %s", err)

		resp := make(map[string]string)
		resp["message"] = "City Not Found"

		json.NewEncoder(w).Encode(resp)
	}
}
