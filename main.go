package main

import (
	"fmt"
	"net/http"

	"ginexample/common"
	"ginexample/common/config"
	"ginexample/handlers"
	"ginexample/utils/log"
)

func init() {
	err := common.LoadNormalEnv()
	if err != nil {
		log.Warnf("can't load env val from .env file.")
	}
	common.Setup()
}

func main() {
	r := handlers.SetupRouter()
	port := config.Conf.Port
	if port == "" {
		port = "9000"
	}
	port = fmt.Sprintf(":%s", port)
	log.Printf("listen on %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
