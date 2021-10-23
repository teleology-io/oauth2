
type options struct {
	code_ttle int64
  access_token_ttl int64
  refresh_token_ttl int64
}

type request struct {
	method string
	query map[string]string
	headers map[string]string
	body string
}

type response struct {
  code int
  headers map[string]string
  body string
}

func TokenRequest(o options) func(req request) response {
	return func (r request) response {
		headers := map[string]string{
				"Authorization": "123",
		}
		return response {
			500,
			headers,
			"hello world",
		}
	}
}