package main

import (
	"github.com/mstrangfeld/oauth2_security_examples/aidp"
	"github.com/mstrangfeld/oauth2_security_examples/client"
	"github.com/mstrangfeld/oauth2_security_examples/hidp"
	"golang.org/x/oauth2"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"
)

const (
	localhost = "http://localhost"

	hidpPort = "3333"
	hidpURL  = localhost + ":" + hidpPort

	aidpPort = "6363"
	aidpURL  = localhost + ":" + "6363"

	clientPort = "4242"
	clientURL  = localhost + ":" + clientPort
)

var hidpConf = oauth2.Config{
	ClientID:     "clientID",
	ClientSecret: "clientSecret",
	Scopes:       []string{"all"},
	RedirectURL:  clientURL + "/oauth2",
	Endpoint: oauth2.Endpoint{
		AuthURL:  hidpURL + "/authorize",
		TokenURL: hidpURL + "/token",
	},
}

var aidpConf = oauth2.Config{
	ClientID:     "clientID",
	ClientSecret: "clientSecret",
	Scopes:       []string{"all"},
	RedirectURL:  clientURL + "/oauth2",
	Endpoint: oauth2.Endpoint{
		AuthURL:  aidpURL + "/authorize",
		TokenURL: aidpURL + "/token",
	},
}

var clientStore = store.NewClientStore()

var aidpServerConf = aidp.Config{
	Port:   aidpPort,
	HIdP:   hidpURL,
	Client: clientURL,
}

func main() {

	clientStore.Set("clientID", &models.Client{
		ID:     "clientID",
		Secret: "clientSecret",
		Domain: clientURL,
	})

	clientStore.Set("attackerID", &models.Client{
		ID:     "attackerID",
		Secret: "attackerSecret",
		Domain: aidpURL,
	})

	go hidp.StartHIdP(hidpPort, clientStore)
	go aidp.StartAIdP(aidpServerConf)

	client.StartClient(clientPort, hidpURL, map[string]oauth2.Config{
		"aidp": aidpConf,
		"hidp": hidpConf,
	})
}
