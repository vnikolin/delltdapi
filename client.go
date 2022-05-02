package delltdapi

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
        "net/url"
)

//DellTDClient ... Contstructor required Variables
type DellTDClient struct {
	DellFQDN     string
	ClientId     string
	ClientSecret string
	APIToken     string

	Client *http.Client
}

//NewDellTDClient ... Initializes the Constructor with the above variables
func NewDellTDClient(target string, username string, password string, apiToken string) (*DellTDClient, error) {

	// Create a new Client instance
	var err error
        proxyUrl, err := url.Parse("http://sub.proxy.att.com:8080")
	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
                Proxy: http.ProxyURL(proxyUrl),
	}
	c := &DellTDClient{DellFQDN: target, ClientId: username, ClientSecret: password, APIToken: apiToken}
	c.Client = &http.Client{Transport: tr, Jar: cookieJar}

	// Get an API Token if not provided
	if apiToken == "" {
		c.getAPIToken()
	}

	return c, err
}
