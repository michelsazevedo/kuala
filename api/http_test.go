package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/michelsazevedo/kuala/domain"
)

func TestGetAllHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/jobs", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).GetAll)

	t.Run("Returns Reponse Code 200", func(t *testing.T) {
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusOK, status)
		}
	})

	t.Run("Returns The Job List", func(t *testing.T) {
		var jobs []domain.Job

		if err := json.NewDecoder(rr.Body).Decode(&jobs); err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}

		resultTotal := len(jobs)
		expectedTotal := 1

		if resultTotal != expectedTotal {
			t.Errorf("Expected: %d. Got: %d.", expectedTotal, resultTotal)
		}
	})
}

func TestGetHandler(t *testing.T) {
	t.Run("Returns Reponse Code 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/jobs/", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Get)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusOK, status)
		}
	})

	t.Run("Returns Reponse Code 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/jobs/", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "2")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Get)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusNotFound, status)
		}
	})
}

func TestCreateHandler(t *testing.T) {
	t.Run("Return Response code 200", func(t *testing.T) {
		now := time.Now()
		job := domain.Job{
			Title:       "Developer",
			Description: "Developer",
			CompanyId:   1,
			Tags:        []string{"Backend"},
			Featured:    false,
			PublishedAt: now,
			CreatedAt:   now,
			ExpiresAt:   now.AddDate(0, 0, 30),
		}
		reqBody, _ := json.Marshal(job)
		req, err := http.NewRequest(http.MethodPost, "/jobs/", bytes.NewBuffer(reqBody))

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Post)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusOK, status)
		}
	})

	t.Run("Return Response code 423", func(t *testing.T) {
		now := time.Now()
		jobWithBlankTitle := domain.Job{
			Title:       "",
			Description: "Developer",
			CompanyId:   1,
			Tags:        []string{"Backend"},
			Featured:    false,
			PublishedAt: now,
			CreatedAt:   now,
			ExpiresAt:   now.AddDate(0, 0, 30),
		}
		reqBody, _ := json.Marshal(jobWithBlankTitle)
		req, err := http.NewRequest(http.MethodPost, "/jobs/", bytes.NewBuffer(reqBody))

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Post)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusUnprocessableEntity, status)
		}
	})
}

func TestUpdatePutHandler(t *testing.T) {
	t.Run("Return Response code 200", func(t *testing.T) {
		now := time.Now()
		job := domain.Job{
			Title:       "Developer",
			Description: "Developer",
			CompanyId:   1,
			Tags:        []string{"Backend"},
			Featured:    false,
			PublishedAt: now,
			CreatedAt:   now,
			ExpiresAt:   now.AddDate(0, 0, 30),
		}
		reqBody, _ := json.Marshal(job)
		req, err := http.NewRequest(http.MethodPut, "/jobs/", bytes.NewBuffer(reqBody))

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Post)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusOK, status)
		}
	})

	t.Run("Returns Reponse Code 404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/jobs/", nil)
		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "2")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Delete)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusUnprocessableEntity, status)
		}
	})

	t.Run("Return Response code 423", func(t *testing.T) {
		now := time.Now()
		jobWithBlankTitle := domain.Job{
			Title:       "",
			Description: "Developer",
			CompanyId:   1,
			Tags:        []string{"Backend"},
			Featured:    false,
			PublishedAt: now,
			CreatedAt:   now,
			ExpiresAt:   now.AddDate(0, 0, 30),
		}
		reqBody, _ := json.Marshal(jobWithBlankTitle)
		req, err := http.NewRequest(http.MethodPut, "/jobs/", bytes.NewBuffer(reqBody))

		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Post)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusUnprocessableEntity, status)
		}
	})
}

func TestDeleteHandler(t *testing.T) {
	t.Run("Returns Response code 400", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/jobs/", nil)
		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Delete)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusBadRequest, status)
		}
	})

	t.Run("Returns Reponse Code 204", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/jobs/", nil)
		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Delete)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusNoContent, status)
		}
	})

	t.Run("Returns Reponse Code 423", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/jobs/", nil)
		if err != nil {
			t.Errorf("Error creating a new request: %v", err)
		}

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "2")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(NewHandler(domain.NewServiceMock()).Delete)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.",
				http.StatusUnprocessableEntity, status)
		}
	})
}
