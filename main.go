package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"oauth2/token"
	"oauth2/types"
)

type Impl struct {
}

func some(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// var options = types.Options{
// 	CodeTTL:         300,
// 	AccessTokenTTL:  3600,
// 	RefreshTokenTTL: 1209600,
// }
func (i Impl) DefaultCodeTTL() int64 {
	return 300
}

func (i Impl) DefaultAccessTokenTTL() int64 {
	return 3600
}

func (i Impl) DefaultRefreshTokenTTL() int64 {
	return 1209600
}

func (i Impl) CreateDecisionPage(req types.Request) (string, error) {
	return "", errors.New("CreateDecisionPage not implemented")
}

func (i Impl) CreateCode(req types.Request) (*types.Code, error) {
	return nil, errors.New("CreateCode not implemented")
}

func (i Impl) CreateAccessToken(req types.Request) (*types.Token, error) {
	return nil, errors.New("CreateAccessToken not implemented")
}

func (i Impl) CreateRefreshToken(req types.Request) (*types.Token, error) {
	return nil, errors.New("CreateRefreshToken not implemented")
}

// Token
func (i Impl) GetTokenTTL(token types.Token) (int64, error) {
	return -1, errors.New("GetTokenTTL not implemented")
}

func (i Impl) GetCode(code string) (*types.Code, error) {
	return nil, errors.New("GetCode not implemented")
}

func (i Impl) GetAccessTokenWithIds(user_id string, client_id string) (*types.Token, error) {
	return nil, errors.New("GetAccessTokenWithIds not implemented")
}

func (i Impl) GetAccessToken(user_id string, client_id string) (*types.Token, error) {
	return nil, errors.New("GetAccessToken not implemented")
}

func (i Impl) GetRefreshToken(user_id string, client_id string) (*types.Token, error) {
	return nil, errors.New("GetRefreshToken not implemented")
}

func (i Impl) Introspect(token types.Token) (*types.IntrospectionResponse, error) {
	return nil, errors.New("Introspect not implemented")
}

// Client
func (i Impl) GetClient(id string) (*types.Client, error) {
	secret := "fake_client_secret"
	return &types.Client{
		Grants:       []string{"password"},
		RedirectURIS: make([]string, 0),
		Scopes:       []string{"user:read", "user:write"},
		Secret:       &secret,
		ID:           "fake_client_key",
	}, nil
}

func (i Impl) ValidGrantType(client types.Client, grant_type string) bool {
	return some(client.Grants, grant_type)
}

func (i Impl) ValidSecret(client types.Client, client_secret string) bool {
	return *client.Secret == client_secret
}

func (i Impl) ValidScope(client types.Client, scopes []string) bool {
	fmt.Println("valid scope", client.Scopes, scopes)
	for _, sc := range scopes {
		if !some(client.Scopes, sc) {
			return false
		}
	}
	return true
}

func (i Impl) ValidRedirectUri(client types.Client, redirect_uri string) bool {
	return some(client.RedirectURIS, redirect_uri)
}

// User
func (i Impl) GetUser(id string) (*types.User, error) {
	return nil, errors.New("GetUser not implemented")
}

func (i Impl) GetUserByName(username string) (*types.User, error) {
	return nil, errors.New("GetUserByName not implemented")
}

func (i Impl) ValidPassword(user types.User, password string) bool {
	return false
}

var options = Impl{}

func main() {
	http.HandleFunc("/token", token.TokenRequest(options))
	fmt.Println("Server started at port 9096")
	log.Fatal(http.ListenAndServe(":9096", nil))
}

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	options := types.Options{}

// 	token.TokenRequest(options)(w, r)
// }
