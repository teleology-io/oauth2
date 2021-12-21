package models

type Client struct {
	Grants        []string
	Redirect_uris []string
	Scopes        []string
	Secret        *string
	Id            string
}
