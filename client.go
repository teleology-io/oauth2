package oauth2

type Client struct {
	id string
	secret *string
	grants []string
	redirect_uris []string
	scopes []string
}