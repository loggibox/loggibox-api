package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/loggibox/loggibox-api/pkg/entity"
	"github.com/loggibox/loggibox-api/pkg/user"
)

func userIndex(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*entity.User
		var err error
		data, err = service.FindAll()

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		resp := entity.HTTPResp{
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func userAdd(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var b *entity.User
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		b.ID, err = service.Store(b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		resp := entity.HTTPResp{
			Code:   http.StatusCreated,
			Result: b,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func userFind(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.Find(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		resp := entity.HTTPResp{
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func userDelete(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
		vars := mux.Vars(r)
		id := vars["id"]
		err := service.Delete(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func userAuth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error auth user"

		type auth struct {
			ID    int    `json:"id"`
			Token string `json:"token"`
		}

		data := auth{
			ID:    1,
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjo2MCwiY3JlYXRlZEF0IjoiMjAxOC0wOC0wMlQxNjozMzoxMS4xNDQ2NDdaIiwidXBkYXRlZEF0IjoiMjAxOC0wOC0wMlQxNjozMzoxMS4xNDQ2NDdaIiwiZGVsZXRlZEF0IjpudWxsLCJmYWNlYm9va0lEIjoiMTIzNDEyMzQifSwiZXhwIjoxNTMzMzE0NzE5LCJpc3MiOiJtdSJ9.on2uZ0WIpdlAitBrGzISZ4tWoSRD5__Vswgl84Yaql8",
		}

		w.WriteHeader(http.StatusOK)
		resp := entity.HTTPResp{
			Code:   http.StatusOK,
			Result: data,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/users", n.With(
		negroni.Wrap(userIndex(service)),
	)).Methods("GET", "OPTIONS").Name("userIndex")

	r.Handle("/users", n.With(
		negroni.Wrap(userAdd(service)),
	)).Methods("POST", "OPTIONS").Name("userAdd")

	r.Handle("/users/{id}", n.With(
		negroni.Wrap(userFind(service)),
	)).Methods("GET", "OPTIONS").Name("userFind")

	r.Handle("/users/{id}", n.With(
		negroni.Wrap(userDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("userDelete")

	r.Handle("/auth", n.With(
		negroni.Wrap(userAuth()),
	)).Methods("POST", "OPTIONS").Name("userAuth")
}
