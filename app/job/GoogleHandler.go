package job

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
)

type GoogleHandler struct {
}

func (g GoogleHandler) Config() *oauth2.Config {
	secret, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Println(err)
	}
	// Creates a oauth2.Config using the secret
	// The second parameter is the scope, in this case we only want to send email
	conf, err := google.ConfigFromJSON(secret, gmail.MailGoogleComScope)
	if err != nil {
		log.Print(err)
	}
	return conf

}
func (g GoogleHandler) GetUrl() string {
	url := g.Config().AuthCodeURL("", oauth2.AccessTypeOnline)
	return url
}
func (g GoogleHandler) GetTokenGmail(code string) *oauth2.Token {
	fmt.Println(code)
	tok, err := g.Config().Exchange(nil, code)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return tok
}

func (g GoogleHandler) GetClient(tok *oauth2.Token) *gmail.Service {
	client := g.Config().Client(nil, tok)

	gmailService, err := gmail.New(client)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return gmailService
}
