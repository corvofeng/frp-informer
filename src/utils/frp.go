package utils

import (
	"encoding/json"
	"fmt"

	FRP "github.com/fatedier/frp/pkg/plugin/server"
	log "github.com/sirupsen/logrus"
)

func HandleUserConn(req FRP.Request) {
	var userConn FRP.NewUserConnContent
	content, err := json.Marshal(req.Content)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(content, &userConn)

	log.Infof("Get user conn: %+v\n", userConn)
}
