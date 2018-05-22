package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PacktPublishing/Echo-Essentials/chapter8/middlewares"
	"github.com/PacktPublishing/Echo-Essentials/chapter8/renderings"
	"github.com/labstack/echo"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	e.Pre(middlewares.RequestIDMiddleware)
	e.GET("/health-check", HealthCheck)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/health-check", nil)

	e.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Error("unexpected status code: ", resp.Status)
	}
	healthCheckResponse := new(renderings.HealthCheckResponse)

	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(healthCheckResponse)
	if err != nil {
		t.Error("error decoding", err)
	}

	if healthCheckResponse.Message != "Everything is good!" {
		t.Error("invalid response message: ", healthCheckResponse.Message)
	}
}

func BenchmarkHealthCheck(b *testing.B) {
	e := echo.New()
	e.Pre(middlewares.RequestIDMiddleware)
	e.GET("/health-check", HealthCheck)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/health-check", nil)

	for i := 0; i < b.N; i++ {
		e.ServeHTTP(w, r)
	}
}
