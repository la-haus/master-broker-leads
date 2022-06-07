package google_func

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Config() *sheets.Service {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv

}
func ConfigWrite() *sheets.Service {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv

}

func Conn() *sheets.Service {
	// Create a JWT configurations object for the Google service account
	conf := &jwt.Config{
		Email:        "masterbroker@social-selling-accounts.iam.gserviceaccount.com",
		PrivateKey:   []byte("-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQD20c80Kn/mMcHc\ngic85PwyjJpfpGHS7AbD80XMGlsqRqLi4s6mi//Bbi7Md7Q8e9SMyGT4LNZzrLy4\nag+Vr0KdspQlu68jKx+XVlrFp+xfsqeomlcjNhq/r7sVMbXIjtYNUC6mLea0BQOF\nwZNCB+V2sSBr2/3FbkLwaRht1VlSMy+W+5BzZwqpchZb/TblMlyR58RS79BxCOX8\nDeyEOWod2ajiYnDrGyKMFVK9pNa2or6oGZJX8Gki4B/EIzW8n9d1W9v0mc87Pk+m\nyJHucanSr+IITHhG5fNRoH+nBs94ZbKS0Ii/Nx89LcRCzyOfo3kQm+Nte4iQ/6Fi\nEhTOiA47AgMBAAECggEAeX0ziv9iXinevQ35h2FE/BK2R96PYjVYbSo5EGpmtkv3\nkEtIobTmik0UTwOa3WysZjGbdYcIdf6wE+vsyOZwGVRV1K7p2Zs33ROd8Y2i7UXw\nfa6eto/Codi7cwa5bkNUfFAb7iNI3BJdtgBe3hEp8v7jsQOjiYMpAtIQOrvZz/m3\ne9bbTTry8ZgxgVK1SSpwKUMYkPOihBVU+I3R9qimkwW0YwPJXsdWcFgxR5m9EMS8\n4V0f8gfF3Jmqs8K8awEj6hJcJ/Kg+sNfsEMdQh9lkxf7Y1KWPrSl8acXxjbkhGyl\nLQsFeGswDV3ywUo+XbAejrv4kDLdiIRqSRhfEtlSGQKBgQD825DcziGFmSxpvbJo\niWdtUqUq1odUTGdC8fDoO9s2sZAmnU6gJNFGWo7febPoNs3tXBqitvaFQrAD1mX4\nMxzv2ynzgqvyC9J4FrzRG4Yff6qeRtDLPDXzXAMmLeE8Mu8HSNl4lzr0de+CC6I9\n5ynHWui5PfzP0slcUdz/Lf7FUwKBgQD54wi3VciejaVypH0LG9wAh6c+cOlSCJRH\nxAd3d+cvNzGdV2VisPmbT+cavwHBITmiS7TFzjBsJ+aAVIeEVdcc9CJcGa0BKNTn\n3KcUDLlRA+xOcm4OstJ33EZpkZmTi+zQXCYMJieQlFmuHzeIv4nWeljbHm7zgVkb\nHz2cOdzOeQKBgHfEgE82C8a+EL1dLt45BxZPKRCrYJpbPiYFdzho8aJsZH0F0bIk\n/kM2ranIeaCqJMwi7jLYOthQCBMpQE9Z1oD97Fb9M/iIBP8AthrN0K+mxugFXuyE\n7oEAujUhUigzb+ihZvCpoTEdk36jiAuOLMnSk5z6cX72to3V8LmlzaElAoGARlcy\nTOlI39jfibXjvQzKkRueaGVASGdB1jHIKEkIOrI+tRu03mMS4DVNyKmbxNGld5n3\n+PZhixBwQg+JVicTtoLCaUqQ+JKV2+6w2WP7hmKkUjDD6j9MH5FSAMWLhY0NzFkk\nHKdg28HmvKzs5QGFu2oVORkIAmOWwZu5rBBdp4ECgYBaQJS5IIU2qMmajtIA3TR7\nPdfeVipX9Hs1GvNWj34kafylzbY0rWIA7ucCpp76SxaBfMV1xhch/QiTOWK1OBAl\nVvfHb4ZRpd/LLWiMkTWbWqQPqHPMUBuBovgRCK4NK/t/Uaoy1MyNzjYEmONjTrmE\nEdLI+fsMK4F4cAnC563qRg==\n-----END PRIVATE KEY-----\n"),
		PrivateKeyID: "3d95bafea4a0c4684c782cabfdab4b1504384071",
		TokenURL:     "https://oauth2.googleapis.com/token",
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
		},
	}

	client := conf.Client(oauth2.NoContext)

	// Create a service object for Google sheets
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv

}
