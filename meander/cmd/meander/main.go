package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"strconv"
	"strings"

	"github.com/RomanosTrechlis/GoProgBlueprints/meander"
)

// http://localhost:8080/recomendations?lat=51.520707&lng=-0.153809&radius=5000&journey=cafe|bar|casino|restaurant&cost=$...$$$
func main() {
	meander.APIKey = os.Getenv("GOOGLE_API_KEY_1")
	if meander.APIKey == "" {
		log.Fatalln("Couldn't get Google's API key from environment")
		return
	}
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.HandleFunc("/recomendations", cors(func(w http.ResponseWriter, r *http.Request) {
		var queryValues = r.URL.Query()
		q := &meander.Query{
			Journey: strings.Split(queryValues.Get("journey"), "|"),
		}
		var err error
		q.Lat, err = strconv.ParseFloat(queryValues.Get("lat"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Lng, err = strconv.ParseFloat(queryValues.Get("lng"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Radius, err = strconv.Atoi(queryValues.Get("radius"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.CostRangeStr = queryValues.Get("cost")
		places := q.Run()
		respond(w, r, places)
	}))
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
