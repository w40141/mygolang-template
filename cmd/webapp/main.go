// Package main
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/w40141/mygolang-template/internal/webapp"
)

const serverPort = "8080"

// main関数はプログラムのエントリーポイントです。
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	router := webapp.NewRouter()

	exporsePort, ok := os.LookupEnv("EXPORSE_PORT")
	if !ok {
		slog.Error("EXPORSE_PORT 環境変数が設定されていません")
		os.Exit(1)
	}

	slog.Info("Server is running", "port", exporsePort)

	if e := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), router); e != nil {
		slog.Error("Failed to start server", "error", e)
		os.Exit(1)
	}
}
