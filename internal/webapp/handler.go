// Package webapp provides HTTP handlers for the web application.
package webapp

import (
	"fmt"
	"net/http"
)

// helloHandlerは /hell エンドポイントへのリクエストを処理します。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// GETメソッド以外は許可しない
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}
	// レスポンスとして "Hello, World!" を書き込みます。
	fmt.Fprintf(w, "Hello, World!")
}

// healthcheckHandlerは /healthcheck エンドポイントへのリクエストを処理します。
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// GETメソッド以外は許可しない
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}
	// HTTPステータスコードを200 OKに設定します。
	w.WriteHeader(http.StatusOK)
	// レスポンスボディとして "OK" を書き込みます。
	w.Write([]byte("OK"))
}

// pingHandlerは /ping エンドポイントへのリクエストを処理します。
func pingHandler(w http.ResponseWriter, r *http.Request) {
	// GETメソッド以外は許可しない
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}
	// HTTPステータスコードを200 OKに設定します。
	w.WriteHeader(http.StatusOK)
	// レスポンスボディとして "pong" を書き込みます。
	w.Write([]byte("pong"))
}

// NewRouter は新しいHTTPルーターを作成し、エンドポイントとハンドラーを設定します。
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// /hell のリクエストを hellHandler で処理するように設定します。
	mux.HandleFunc("/hello", helloHandler)
	// /healthcheck のリクエストを healthcheckHandler で処理するように設定します。
	mux.HandleFunc("/healthcheck", healthcheckHandler)
	// /ping のリクエストを pingHandler で処理するように設定します。
	mux.HandleFunc("/ping", pingHandler)

	return mux
}
