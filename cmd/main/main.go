package main

import (
	"encoding/json"
	"frp-informer/src/utils"
	"net/http"

	log "github.com/sirupsen/logrus"

	FRP "github.com/fatedier/frp/pkg/plugin/server"
)

func FrpHandler(w http.ResponseWriter, req *http.Request) {
	var frpReq FRP.Request
	json.NewDecoder(req.Body).Decode(&frpReq)
	log.Debugf("Frp req: %+v\n", frpReq)
	utils.HandleUserConn(frpReq)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(FRP.Response{
		Reject:   false,
		Unchange: true,
	})
}

func main() {
	http.HandleFunc("/handler", FrpHandler)
	_ = http.ListenAndServe(":9000", nil)
}
