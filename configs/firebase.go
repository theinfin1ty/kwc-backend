package configs

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func InitializeFirebase() *firebase.App {
	config := FirebaseConfig{
		Type:                    GetEnvVariable("FIREBASE_ACCOUNT_TYPE"),
		ProjectId:               GetEnvVariable("FIREBASE_PROJECT_ID"),
		PrivateKeyId:            GetEnvVariable("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              GetEnvVariable("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             GetEnvVariable("FIREBASE_CLIENT_EMAIL"),
		ClientId:                GetEnvVariable("FIREBASE_CLIENT_ID"),
		AuthUri:                 GetEnvVariable("FIREBASE_AUTH_URI"),
		TokenUri:                GetEnvVariable("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertUrl: GetEnvVariable("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl:       GetEnvVariable("FIREBASE_CLIENT_X509_CERT_URL"),
	}

	firebaseConfigJson, err := json.Marshal(config)

	if err != nil {
		panic("Unable to convert firebase config json")
	}

	opt := option.WithCredentialsJSON(firebaseConfigJson)

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		panic("Unable to initialize firebase")
	}

	return app
}

var FirebaseAdmin = InitializeFirebase()
