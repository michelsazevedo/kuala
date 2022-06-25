package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/michelsazevedo/kuala/domain"
)

type JobHandler interface {
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type handler struct {
	jobService domain.Service
}

//NewHandler ...
func NewHandler(jobService domain.Service) JobHandler {
	return &handler{jobService: jobService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(requestId, 10, 64)
	if err != nil {
		log.Printf("Error to converter string id %s to int64", requestId)
	}

	job, err := h.jobService.Find(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&job)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jobs, err := h.jobService.FindAll()

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&jobs)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	job := &domain.Job{}
	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.jobService.Create(job)
	if err != nil {
		log.Default().Print(err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(&job)
}

func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	job := &domain.Job{}
	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.jobService.Update(job)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&job)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestId := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(requestId, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.jobService.Delete(id); err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
