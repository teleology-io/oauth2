package token

import (
	"encoding/json"
	"net/http"
	"oauth2/errors"
	"oauth2/introspection"
	"oauth2/parse"
	"oauth2/types"
)

type DefaultTokenResponse struct {
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
	ExpiresIn    string   `json:"expires_in"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}

type ClientCredentialsResponse struct {
	Scope       []string `json:"scope"`
	TokenType   string   `json:"token_type"`
	AccessToken string   `json:"access_token"`
	ExpiresIn   int64    `json:"expires_in"`
}

func clientCredentials(o types.Options, data map[string]interface{}) *ClientCredentialsResponse {
	accessToken, err := o.CreateAccessToken(data)
	if err != nil {
		return nil
	}

	return &ClientCredentialsResponse{
		Scope:       data["scope"].([]string),
		TokenType:   "Bearer",
		AccessToken: accessToken.Token,
		ExpiresIn:   o.DefaultAccessTokenTTL(),
	}
}

func authorizationCode(o types.Options, data map[string]interface{}) (*DefaultTokenResponse, errors.OAuth2Error) {
	givenCode := data["code"]
	client := data["client"].(types.Client)
	code, err := o.GetCode(givenCode.(string))
	if err != nil || code == nil {
		return nil, *errors.InvalidGrant("Code not found")
	}

	if code.ClientID != client.ID {
		return nil, *errors.InvalidGrant("Code issued elsewhere")
	}

	var refreshToken map[string]string{}
	if o.ValidGrantType(client, "refresh_token") {
		refreshToken, _ = o.CreateRefreshToken(data)
	}

	accessToken, _ := o.CreateAccessToken(data)

	return &DefaultTokenResponse{
		Scope:       data["scope"].([]string),
		TokenType:   "Bearer",
		AccessToken: accessToken.Token,
		RefreshToken: refreshToken.Token,
		ExpiresIn:   o.DefaultAccessTokenTTL(),
	}, nil
}

// // const user = await
//   // TODO: implement this
//   const Code = await options.getCode(code);
//   if (!Code) {
//     throw new InvalidGrant('Code not found');
//   }

//   if (Code.client_id !== client.id) {
//     throw new InvalidGrant('Code issued elsewhere');
//   }

//   const canRefresh = options.validGrantType(client, 'refresh_token');
//   const [accessToken, refreshToken] = await Promise.all([
//     options.createAccessToken({
//       user_id: Code.user_id,
//       client_id: client.id,
//       ttl: options.accessTokenTtl,
//       scope,
//     }),
//     canRefresh
//       ? options.createRefreshToken({
//         user_id: Code.user_id,
//         client_id: client.id,
//         ttl: options.refreshTokenTtl,
//         scope,
//       })
//       : noop(),
//   ]);

//   return {
//     token_type: 'Bearer',
//     expires_in: options.accessTokenTtl,
//     access_token: accessToken.token,
//     // @ts-ignore
//     refresh_token: refreshToken.token,
//     scope,
//   }

func TokenRequest(o types.Options) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("authorization")

		if r.Body == nil {
			errors.ErrorResponse(w, *errors.InvalidRequest("Missing request body"))
			return
		}

		data := parse.CollectData(r)
		client_id := data["client_id"]
		client_secret := data["client_secret"]
		grant_type := data["grant_type"]
		scope := data["scope"]
		redirect_uri := data["redirect_uri"]
		if client_id == nil || client_secret == nil {
			creds, err := introspection.BasicAuthorization(auth)
			if err != nil {
				errors.ErrorResponse(w, *errors.InvalidRequest(err.Error()))
				return
			}

			client_id = creds[0]
			client_secret = creds[1]
			data["client_id"] = client_id
			data["client_secret"] = client_secret
		}

		client, err := o.GetClient(client_id.(string))
		if err != nil || client == nil {
			errors.ErrorResponse(w, *errors.InvalidClient("No client found with the provided credentials"))
			return
		}

		data["client"] = client

		if !o.ValidGrantType(*client, grant_type.(string)) {
			errors.ErrorResponse(w, *errors.UnauthorizedClient("Client unauthorized to make grant_type"))
			return
		}

		if client_secret != nil && !o.ValidSecret(*client, client_secret.(string)) {
			errors.ErrorResponse(w, *errors.InvalidClient("Client secret wrong"))
			return
		}

		if scope != nil && !o.ValidScope(*client, scope.([]string)) {
			errors.ErrorResponse(w, *errors.InvalidScope("Invalid client scope"))
			return
		}

		if redirect_uri != nil && !o.ValidRedirectUri(*client, redirect_uri.(string)) {
			errors.ErrorResponse(w, *errors.InvalidRequest("Redirect not registered with client"))
			return
		}

		if grant_type != "authorization_code" && grant_type != "password" && grant_type != "client_credentials" && grant_type != "refresh_token" {
			errors.ErrorResponse(w, *errors.InvalidRequest("Grant type unsupported"))
			return
		}

		var response interface{}
		if grant_type == "client_credentials" {
			response = clientCredentials(o, data)
		}

		if grant_type == "authorization_code" {
			response = authorizationCode(o, data)
		}

		// success response
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Pragma", "no-cache")

		json, _ := json.Marshal(response)
		w.Write(json)
	}
}
