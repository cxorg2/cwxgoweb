package api

import (
	"encoding/json"
	"log"
	"net/http"
)

var htmlBody = []byte(`<html>
<head>
<title>Hello</title>
</head>
<body>
Hello, World!
</body>
</html>
`)

// /html 的响应
func htmlPage(w http.ResponseWriter, r *http.Request) {
	addr := r.RemoteAddr
	log.Println("rev mess /htmlPage", addr)
	w.Write(htmlBody)
}

// 响应 根路径的请求
func rootPage(w http.ResponseWriter, r *http.Request) {
	var urls = []string{}
	h1 := "http://" + r.Host

	urls = append(urls, h1+"/html")
	urls = append(urls, h1+"/sleep")
	urls = append(urls, h1+"/getIP")
	urls = append(urls, h1+"/api/readfile")
	urls = append(urls, h1+"/api/env")
	urls = append(urls, h1+"/api/onelog")

	data, err := json.Marshal(urls)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(200)
		w.Write(data)
	}

}
