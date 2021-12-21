package errors

import (
	"encoding/json"
	"net/http"
)

type OAuth2Error struct {
	Code             int    `json:"code"`
	Message          string `json:"message"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

func AccessDenied(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             403,
		Message:          msg,
		Error:            "access_denied",
		ErrorDescription: "The resource owner or authorization server denied the request.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func InsufficientScope(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             403,
		Message:          msg,
		Error:            "insufficient_scope",
		ErrorDescription: "The request requires higher privileges than provided by the access token.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6750.html#section-3.1",
	}
}

func InvalidClient(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             401,
		Message:          msg,
		Error:            "invalid_client",
		ErrorDescription: "Client authentication failed (e.g., unknown client, no client authentication included, or unsupported authentication method).",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func InvalidGrant(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             400,
		Message:          msg,
		Error:            "invalid_grant",
		ErrorDescription: "The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func InvalidRequest(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             400,
		Message:          msg,
		Error:            "invalid_request",
		ErrorDescription: "The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func InvalidScope(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             400,
		Message:          msg,
		Error:            "invalid_scope",
		ErrorDescription: "The requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func InvalidToken(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             401,
		Message:          msg,
		Error:            "invalid_token",
		ErrorDescription: "The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6750#section-3.1",
	}
}

func ServerError(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             500,
		Message:          msg,
		Error:            "server_error",
		ErrorDescription: "The authorization server encountered an unexpected condition that prevented it from fulfilling the request.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func UnauthorizedClient(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             401,
		Message:          msg,
		Error:            "unauthorized_client",
		ErrorDescription: "The client is not authorized to request an authorization code using this method.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func UnsupportedGrantType(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             400,
		Message:          msg,
		Error:            "unsupported_grant_type",
		ErrorDescription: "The authorization server does not support obtaining an authorization code using this method.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func UnsupportedResponseType(msg string) *OAuth2Error {
	return &OAuth2Error{
		Code:             400,
		Message:          msg,
		Error:            "unsupported_response_type",
		ErrorDescription: "The authorization server does not support obtaining an authorization code using this method.",
		ErrorUri:         "https://tools.ietf.org/html/rfc6749#section-4.1.2.1",
	}
}

func ErrorResponse(w http.ResponseWriter, err OAuth2Error) {
	// Header changes must be before status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(err.Code))
	js, _ := json.Marshal(err)
	w.Write(js)
}
