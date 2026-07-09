package parsed

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/mrjones/oauth"
)

// hattrickCHPPOauth is the base URL for Hattrick's CHPP OAuth1 endpoints
// (request_token, authorize, access_token).
const hattrickCHPPOauth = "https://chpp.hattrick.org/oauth/"

// rawClient is the lowest-level CHPP client: it issues OAuth1-signed HTTP
// requests to chpp.hattrick.org and returns the raw response bytes,
// without any XML decoding.
type rawClient struct {
	client *http.Client
}

// newRawClient builds a rawClient using the application's OAuth1 consumer
// key/secret and a user's access token/secret (plus any additional data
// returned alongside the access token during authorization, e.g. a session
// handle needed to refresh it).
func newRawClient(consumerKey, consumerSecret, accessToken, accessSecret string, accessAdditionalData map[string]string) (*rawClient, error) {
	client, err := getConsumer(consumerKey, consumerSecret).MakeHttpClient(&oauth.AccessToken{
		Token:          accessToken,
		Secret:         accessSecret,
		AdditionalData: accessAdditionalData,
	})
	if err != nil {
		return nil, err
	}

	return &rawClient{
		client: client,
	}, nil
}

func (c *rawClient) callGet(u string) ([]byte, error) {
	response, err := c.client.Get(u)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bits, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bits, nil
}

func (c *rawClient) callPost(u, contentType, body string) ([]byte, error) {
	response, err := c.client.Post(u, contentType, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bits, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bits, nil
}

// getRawXML issues a signed GET request for the given CHPP file/version
// with the given query parameters, and returns the raw (undecoded)
// response body.
func (c *rawClient) getRawXML(file, version string, values map[string]string) ([]byte, error) {
	u := buildURL(file, version, values)
	buf, err := c.callGet(u)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// postRawXML sends a POST request whose body is not part of the query
// string (e.g. the "lineup" JSON payload for matchorders'
// actionType=setmatchorder), and returns the raw response body.
//
// The request body is sent with a non-form Content-Type so that the OAuth1
// signature (computed by the underlying client) is derived from the query
// string alone, matching the CHPP contract: "All parameters except lineup
// should still be sent in querystring."
func (c *rawClient) postRawXML(file, version string, values map[string]string, body string) ([]byte, error) {
	u := buildURL(file, version, values)
	buf, err := c.callPost(u, "application/json", body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func getConsumer(consumerKey, consumerSecret string) *oauth.Consumer {
	return oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   hattrickCHPPOauth + "request_token.ashx",
			AuthorizeTokenUrl: hattrickCHPPOauth + "authorize.aspx",
			AccessTokenUrl:    hattrickCHPPOauth + "access_token.ashx",
		})
}

func buildURL(file, version string, values map[string]string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "chpp.hattrick.org",
		Path:   "chppxml.ashx",
	}

	q := u.Query()
	q.Set("file", file)
	q.Set("version", version)

	for k, v := range values {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	return u.String()
}
