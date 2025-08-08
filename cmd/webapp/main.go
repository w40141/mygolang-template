// Package main
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/w40141/mygolang-template/internal/webapp"
)

// main関数はプログラムのエントリーポイントです。
func main() {
	router := webapp.NewRouter()

	exporsePort, ok := os.LookupEnv("EXPORSE_PORT")
	if !ok {
		log.Fatal("EXPORSE_PORT 環境変数が設定されていません")
	}

	// サーバーが起動することを示すメッセージを出力します。
	fmt.Printf("Server is running on http://localhost:%s", exporsePort)

	// 8080ポートでサーバーを起動します。起動に失敗した場合はエラーを出力し、プログラムを終了します。
	log.Fatal(http.ListenAndServe(":8080", router))
}
