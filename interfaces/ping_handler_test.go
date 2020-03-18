package interfaces

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func sendRequest(t *testing.T, method, url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	return res
}

func TestHandlerPong_GetPong(t *testing.T) {
	t.Log("TestHandlerPong_GetPong should return status OK")
	router := mux.NewRouter()
	w := httptest.NewRecorder()

	handler := NewHandlerPong()
	router.HandleFunc("/ping", handler.GetPong).Methods("GET")

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/ping", "http://localhost:8000"), nil)
	defer req.Body.Close()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,"pong", w.Body.String())
}
