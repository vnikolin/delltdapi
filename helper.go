package delltdapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func (c *DellTDClient) formatPath(path string) string {
	return fmt.Sprintf("https://%s/%s", c.DellFQDN, path)
}

func (c *DellTDClient) getAPIToken() error {

	authURL, err := url.Parse(c.formatPath("auth/oauth/v2/token"))
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("client_id", c.ClientId)
	data.Set("client_secret", c.ClientSecret)
	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", authURL.String(), strings.NewReader(data.Encode())) //URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	r, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cleanData TokenInfo
	json.Unmarshal(body, &cleanData)
	c.APIToken = cleanData.AccessToken
	//log.Println(string([]byte(body)))

	return err
}

//queryData ... will make REST verbs based on the url
func (c *DellTDClient) QueryData(call string, link string, data []byte) ([]byte, http.Header, int, error) {

	// Create bearer token string
	var bearer = "Bearer " + c.APIToken

	// Create a new request using http
	req, err := http.NewRequest(call, link, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil, 0, err
	}

	// Add authorization header to the req
	req.Header.Add("Authorization", bearer)

	resp, err := c.Client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := errors.New(StatusBadRequest)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 403 {
			err := errors.New(StatusForbidden)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 404 {
			err := errors.New(StatusNotFound)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 409 {
			err := errors.New(StatusConflict)
			return nil, resp.Header, resp.StatusCode, err
		}

	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, resp.StatusCode, err
	}

	return _body, resp.Header, resp.StatusCode, nil
}
