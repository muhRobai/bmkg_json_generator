package route

import (
	"github.com/bmkg_json_generator/service"
	"github.com/gorilla/mux"
	"net/http"
	// model "bmkg/bmkg_api/model"
	"encoding/json"
)

type RouteAPI struct {
	service.NewApi
}

func initApi() RouteAPI {
	api := service.NewApiService()
	return RouteAPI{api}
}

func (c *RouteAPI) recentEarthQuakesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := c.RecentEarthQuakes()
	if err != nil {
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// func earthQuakesHandler(w http.ResponseWriter, r *http.Request) {

// }

// func earthQuakesFeelHandler(w http.ResponseWriter, r *http.Request) {

// }

// func earthQuakesTsunamiHandler(w http.ResponseWriter, r *http.Request) {

// }

// func earthQuakesImageHandler(w http.ResponseWriter, r *http.Request) {

// }

func StartHttp() http.Handler {
	api := initApi()
	r := mux.NewRouter()
	r.HandleFunc("/api/bmkl/recent-earth-quakes", api.recentEarthQuakesHandler).Methods("GET")
	// r.HandleFunc("/api/bmkl/earth-quakes", earthQuakesHandler).Methods("GET")
	// r.HandleFunc("/api/bmkl/earth-quakes-feel", earthQuakesFeelHandler).Methods("GET")
	// r.HandleFunc("/api/bmkl/earth-quakes-tsunami", earthQuakesTsunamiHandler).Methods("GET")
	// r.HandleFunc("/api/bmkl/earth-quakes-image", earthQuakesImageHandler).Methods("GET")

	return r
}
