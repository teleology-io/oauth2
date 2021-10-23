package main

import (
  "fmt"
)

type Scope []string

type Token struct {
  token string
  user_id string
  client_id string
  scope Scope
  created_at string
}

type Code struct {
  code string
  user_id string
  client_id string
  scope Scope
  created_at string
}


type AccessToken Token

type RefreshToken Token

type TokenRequest struct {
  user_id string
  client_id string
  scope Scope
  ttl int
}

type User struct {
  id string
  username string
  password string
}

type Client struct {
  id string
  secret string
  grants []string
  redirect_uris []string
  scopes []string
}

type Body struct {
  string
}

type Query struct {}

type Headers struct {}

type Method string

const (
  GET Method = "get"
  POST Method = "post"
  PUT Method = "put"
  DELETE Method = "delete"
  OPTIONS Method = "options"
  PATCH Method = "patch"
)

type Request struct {
  method Method
  headers Headers
  body Body
  query Query
}

type DecisionRequest struct {
  Request
  client_id string
  scope Scope
}

type Response struct {
  code int
  headers Headers
  body Body
}

type IntrospectionResponse struct {
  active bool
  client_id string
  username string
  token_type string
  exp int64
  iat int64
  nbf int64
  sub string
  aud string
  iss string
  jti string
}

type OptionsInterface interface {
  // Decision
  createDecisionPage(req DecisionRequest) string
  createCode(req TokenRequest) Code
  createAccessToken(req TokenRequest) AccessToken
  createRefreshToken(req TokenRequest) AccessToken

  // Token
  getTokenTtl(token Token) int64
  getCode(code string) Code
  getAccessTokenWithIds(user_id string, client_id string) AccessToken
  getAccessToken(token string) AccessToken
  getRefreshToken(token string) RefreshToken
  introspect(token Token) IntrospectionResponse

  // Client
  getClient(id string) Client
  validGrantType(client Client, grant_type string) bool
  validSecret(client Client, client_secret string) bool
  validScope(client Client, scope Scope) bool
  validRedirectUri(client Client, redirect_uri string) bool

  // User
  getUser(id string) User
  getUserByName(username string) User
  validPassword(user User, password string) bool
}

type Options struct {
  codeTtl int64
  accessTokenTtl int64
  refreshTokenTtl int64
}

func (o *Options)  getUser(id string) User {
  user := User{
    "123", "icarus", "",
  }
  return user
}

func main(){
  v := Options{
    30,
    150,
    1000,
  }
  fmt.Println("Its %s", v)
  fmt.Println("get user", v.getUser("123"))
}