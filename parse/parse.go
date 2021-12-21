package parse

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

func ParseBody(in io.ReadCloser) interface{} {
	body, ioErr := ioutil.ReadAll(in)
	if ioErr != nil {
		return map[string]interface{}{}
	}

	var response map[string]interface{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr == nil {
		return response
	}

	return string(body)
}

func ParseUrlValues(values url.Values) map[string]interface{} {
	response := map[string]interface{}{}

	for k, v := range values {
		if k == "scope" {
			response[k] = v
		} else if len(v) == 1 {
			response[k] = v[0]
		} else {
			response[k] = v
		}
	}

	return response
}

func merge(ms ...map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

func CollectData(req *http.Request) map[string]interface{} {
	// parse any form fields
	req.ParseForm()

	body := ParseBody(req.Body)
	params := ParseUrlValues(req.URL.Query())
	form := ParseUrlValues(req.Form)
	kind := reflect.TypeOf(body).Kind()
	if kind == reflect.String {
		return merge(form, params)
	}

	return merge(body.(map[string]interface{}), form, params)
}
