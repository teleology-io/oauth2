package types

type Token struct {
	Scope     []string
	Token     string
	UserID    string
	ClientID  string
	CreatedAt string
}

type Code struct {
	Scope     []string
	Code      string
	UserID    string
	ClientID  string
	CreatedAt string
}

type Request struct {
	Data    map[string]interface{}
	Query   map[string]string
	Headers map[string]string
	Method  string
}

type Response struct {
	Headers map[string]string
	Body    string
	Code    int
}

type IntrospectionResponse struct {
	ClientId  string
	Username  string
	TokenType string
	Sub       string
	Aud       string
	Iss       string
	Jti       string
	Exp       int64
	Iat       int64
	Nbf       int64
	Active    bool
}

type User struct {
	Password *string
	ID       string
	Username string
}

type Client struct {
	Grants       []string
	RedirectURIS []string
	Scopes       []string
	Secret       *string
	ID           string
}

type Options interface {
	DefaultCodeTTL() int64
	DefaultAccessTokenTTL() int64
	DefaultRefreshTokenTTL() int64

	CreateDecisionPage(data map[string]interface{}) (string, error)
	CreateCode(data map[string]interface{}) (*Code, error)
	CreateAccessToken(data map[string]interface{}) (*Token, error)
	CreateRefreshToken(data map[string]interface{}) (*Token, error)

	GetTokenTTL(token Token) (int64, error)
	GetCode(code string) (*Code, error)
	GetAccessTokenWithIds(user_id string, client_id string) (*Token, error)
	GetAccessToken(user_id string, client_id string) (*Token, error)
	GetRefreshToken(user_id string, client_id string) (*Token, error)

	Introspect(token Token) (*IntrospectionResponse, error)

	GetClient(id string) (*Client, error)
	ValidGrantType(client Client, grant_type string) bool
	ValidSecret(client Client, client_secret string) bool
	ValidScope(client Client, scopes []string) bool
	ValidRedirectUri(client Client, redirect_uri string) bool
	GetUser(id string) (*User, error)
	GetUserByName(username string) (*User, error)
	ValidPassword(user User, password string) bool
}
