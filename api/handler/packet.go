package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/loggibox/loggibox-api/pkg/entity"
	"github.com/loggibox/loggibox-api/pkg/packet"
)

func packetIndex(service packet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading packets"
		var data []entity.Packet
		var err error
		query := r.URL.Query().Encode()
		switch {
		case query == "":
			data, err = service.FindAll()
		default:
			data, err = service.Search(query)
		}

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			resp := entity.HTTPResp{
				Messages: []string{"Is empty"},
				Result:   []string{},
			}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
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

func packetAdd(service packet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding packet"
		var b *entity.Packet
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

func packetFind(service packet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading packet"
		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.Find(id)
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

func packetDelete(service packet.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing packet"
		vars := mux.Vars(r)
		id := vars["id"]
		err := service.Delete(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakePacketHandlers make url handlers
func MakePacketHandlers(r *mux.Router, n negroni.Negroni, service packet.UseCase) {
	r.Handle("/packets", n.With(
		negroni.Wrap(packetIndex(service)),
	)).Methods("GET", "OPTIONS").Name("packetIndex")

	r.Handle("/packets", n.With(
		negroni.Wrap(packetAdd(service)),
	)).Methods("POST", "OPTIONS").Name("packetAdd")

	r.Handle("/packets/{id}", n.With(
		negroni.Wrap(packetFind(service)),
	)).Methods("GET", "OPTIONS").Name("packetFind")

	r.Handle("/packets/{id}", n.With(
		negroni.Wrap(packetDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("packetDelete")
}
