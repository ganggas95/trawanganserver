package job

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/plus/v1"
	"net/http"
)

type GplusHandler struct {
}

const (
	clientID        = "75703245588-v7co3b5ek6nc0qdl6p5q7fve73f7cb8j.apps.googleusercontent.com"
	clientSecret    = "MpJ_eMCSnEhydLkSksfzzUw1"
	applicationName = "Google+ Go Quickstart"
)

// config is the configuration specification supplied to the OAuth package.
var config = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	// Scope determines which API calls you are authorized to make
	Scopes:   []string{plus.UserinfoEmailScope},
	Endpoint: google.Endpoint,
	// Use "postmessage" for the code-flow for server side apps
	RedirectURL: "https://trawanganserver.herokuapp.com/verifygplus",
}

func (g GplusHandler) GetUrlPlus() string {
	url := config.AuthCodeURL("", oauth2.AccessTypeOnline)
	return url
}
func (g GplusHandler) GetTokenPlus(code string) *oauth2.Token {
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}
	return token
}

func (g GplusHandler) GetClientPlus(token *oauth2.Token) *http.Client {
	client := config.Client(oauth2.NoContext, token)
	return client
}

func (g GplusHandler) GetServicePlus(client *http.Client) *plus.Service {
	service, err := plus.New(client)
	if err != nil {
		panic(err)
	}
	return service
}

func (g GplusHandler) GetPeoplePlus(service *plus.Service) *plus.Person {
	people := service.People.Get("me")
	peoplefeed, err := people.Do()
	if err != nil {
		panic(err)
	}
	return peoplefeed

}
