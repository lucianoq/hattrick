package raw

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mrjones/oauth"
)

// HattrickCHPPOauth ...
const HattrickCHPPOauth = "https://chpp.hattrick.org/oauth/"

// Raw ...
type Raw struct {
	client *http.Client
}

// NewRaw ...
func NewRaw(consumerKey, consumerSecret, accessToken, accessSecret string, accessAdditionalData map[string]string) (*Raw, error) {
	client, err := getConsumer(consumerKey, consumerSecret).MakeHttpClient(&oauth.AccessToken{
		Token:          accessToken,
		Secret:         accessSecret,
		AdditionalData: accessAdditionalData,
	})
	if err != nil {
		return nil, err
	}

	return &Raw{
		client: client,
	}, nil
}

func (raw *Raw) callGet(u string) ([]byte, error) {
	response, err := raw.client.Get(u)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bits, nil
}

// GetRawXML ...
func (raw *Raw) GetRawXML(file, version string, values map[string]string) ([]byte, error) {
	u := buildURL(file, version, values)
	buf, err := raw.callGet(u)
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
			RequestTokenUrl:   HattrickCHPPOauth + "request_token.ashx",
			AuthorizeTokenUrl: HattrickCHPPOauth + "authorize.aspx",
			AccessTokenUrl:    HattrickCHPPOauth + "access_token.ashx",
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
