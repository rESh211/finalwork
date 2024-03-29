package step9

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"работы/Скилбокс/finalwork/pkg/config"

	"github.com/gorilla/mux"
)

func simulatorListenAndServeHTTP() {
	router := mux.NewRouter()

	router.HandleFunc("/mms", handleMMS).Methods("GET")
	router.HandleFunc("/support", handleSupport).Methods("GET")
	router.HandleFunc("/incident", handleIncident).Methods("GET")

	log.Fatal(http.ListenAndServe(config.GlobalConfig.SimulatorAddr, router))
}

func handleMMS(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(config.GlobalConfig.MMSFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	if random%5 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func handleSupport(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(config.GlobalConfig.SupportFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	if random%5 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleIncident(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(config.GlobalConfig.IncidentFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	if random%5 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func StartSimulatorServer() {
	simulatorListenAndServeHTTP()
}
