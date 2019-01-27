package pr0gramm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ajg/form"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Session struct {
	client http.Client
}

func NewSession(client http.Client) *Session {
	client.Jar, _ = cookiejar.New(nil)
	return &Session{client: client}
}

func (sess *Session) Login(username, password string) (*LoginResponse, error) {
	body := make(url.Values)
	body.Set("name", username)
	body.Set("password", password)

	uri := "https://pr0gramm.com/api/user/login"
	resp, err := sess.client.PostForm(uri, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return &response, err
}

func (sess *Session) apiGET(path string, query url.Values, target interface{}) error {
	uri := "https://pr0gramm.com/api" + path

	if query != nil {
		uri += "?" + query.Encode()
	}

	response, err := http.DefaultClient.Get(uri)
	if err != nil {
		return err
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode/100 != 2 {
		return fmt.Errorf("error %d", response.StatusCode)
	}

	return json.NewDecoder(response.Body).Decode(target)
}

func (sess *Session) apiPOST(path string, query url.Values, body interface{}, target interface{}) error {
	uri := "https://pr0gramm.com/api" + path

	if query != nil {
		uri += "?" + query.Encode()
	}

	bodyValues, err := encodeToValues(body)
	if err != nil {
		return fmt.Errorf("encode body: %s", err)
	}

	var nonce string

	// extract nonce from 'id' in 'me' cookie.
	cookies := sess.client.Jar.Cookies(&url.URL{Scheme: "https", Host: "pr0gramm.com", Path: "/api/"})
	for _, cookie := range cookies {
		if cookie.Name == "me" {
			var decoded struct{ Id string }
			_ = json.Unmarshal([]byte(cookie.Value), &decoded)

			if len(decoded.Id) > 16 {
				nonce = decoded.Id[:16]
				break
			}
		}
	}

	if nonce == "" {
		return errors.New("not authorized")
	}

	bodyContent := bodyValues.Encode() + "&_nonce=" + nonce

	response, err := http.DefaultClient.Post(
		uri, "application/x-www-form-urlencoded",
		strings.NewReader(bodyContent))

	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}

	defer func() {
		// discard and close body
		_, _ = io.Copy(ioutil.Discard, response.Body)
		_ = response.Body.Close()
	}()

	if response.StatusCode/100 != 2 {
		return fmt.Errorf("error %d", response.StatusCode)
	}

	if target == nil {
		return nil
	}

	return json.NewDecoder(response.Body).Decode(target)
}

func encodeToValues(body interface{}) (url.Values, error) {
	if values, ok := body.(url.Values); ok {
		return values, nil
	} else {
		return form.EncodeToValues(body)
	}
}
