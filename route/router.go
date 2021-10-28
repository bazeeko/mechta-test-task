package route

import (
	"net/http"

	"github.com/bazeeko/mechta-test-task/data"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router     *mux.Router
	Repository *data.CityRepository
}

type route struct {
	Endpoint string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

func NewHandler(db *data.CityRepository) *Handler {
	return &Handler{
		mux.NewRouter(),
		db,
	}
}

func (h *Handler) Init() {
	routes := []route{
		{
			Endpoint: "/cities",
			Method:   http.MethodGet,
			Function: h.GetAllCities,
		},
		{
			Endpoint: "/cities",
			Method:   http.MethodPost,
			Function: h.AddSingleCity,
		},
		{
			Endpoint: "/cities/{id:[0-9]+}",
			Method:   http.MethodGet,
			Function: h.GetSingleCity,
		},
		{
			Endpoint: "/cities/{id:[0-9]+}",
			Method:   http.MethodPut,
			Function: h.UpdateSingleCity,
		},
		{
			Endpoint: "/cities/{id:[0-9]+}",
			Method:   http.MethodDelete,
			Function: h.DeleteSingleCity,
		},
	}

	h.Router.Use(SetHeaders)

	for _, r := range routes {
		h.Router.HandleFunc(r.Endpoint, r.Function).Methods(r.Method)
	}

	http.Handle("/", h.Router)
}
