package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aelaster/go-taco/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "https://github.com/aelaster/go-taco/")
}

func MenuItemCalculate(w http.ResponseWriter, r *http.Request) {
	var orderItems model.OrderItems
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &orderItems); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()}); err != nil {
			panic(err)
		}
	} else {
		t := RepoCalculateCost(orderItems)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	}
}
