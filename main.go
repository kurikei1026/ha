package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// *HelloHandler がインターフェース http.Handler を実装
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	// HelloHandler 型の変数を宣言
	handler := HelloHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
