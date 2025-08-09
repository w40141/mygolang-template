// Package webapp provides HTTP handlers for the web application.
package webapp

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// helloHandlerは /hello エンドポイントへのリクエストを処理します。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	message := map[string]string{
		"message": "Hello, World!",
	}

	writeJSON(w, http.StatusOK, message)
}

// healthcheckHandlerは /healthcheck エンドポイントへのリクエストを処理します。
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	message := map[string]string{
		"message": "Service is healthy",
		"now":     time.Now().Format(time.RFC3339),
	}

	writeJSON(w, http.StatusOK, message)
}

type errorResponse struct {
	Title string `json:"title"`
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(v)
	if err != nil {
		slog.Error("Error marshalling JSON", "error", err)

		code = http.StatusInternalServerError
		er := errorResponse{Title: "Internal Server Error"}

		res, err = json.Marshal(er)
		if err != nil {
			slog.Error("Error marshalling error response", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			return
		}
	}

	w.WriteHeader(code)

	if _, err := w.Write(res); err != nil {
		slog.Error("Error writing response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}
}

// NewRouter は新しいHTTPルーターを作成し、エンドポイントとハンドラーを設定します。
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/healthcheck", healthcheckHandler)

	return mux
}
