package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type AppEnv struct {
	SrcIp    string
	HostName string
	Req      ReqInfo
}

type ReqInfo struct {
	Method     string
	Path       string
	Proto      string
	Host       string
	RequestURI string
	XFF_IP     string
}

func getEnv(w http.ResponseWriter, r *http.Request) {
	var xff string
	if len(r.Header["X-Forwarded-For"]) != 0 {
		xff = r.Header["X-Forwarded-For"][0]
	} else {
		xff = ""
	}

	s1 := AppEnv{
		SrcIp:    r.RemoteAddr,
		HostName: os.Getenv("HOSTNAME"),
		Req: ReqInfo{
			Method:     r.Method,
			Path:       r.URL.Path,
			Proto:      r.Proto,
			Host:       r.Host,
			RequestURI: r.RequestURI,
			XFF_IP:     xff,
		},
	}

	w.WriteHeader(200)
	data, err := json.Marshal(s1)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(data)

}
