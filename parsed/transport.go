package parsed

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

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

	// Kept for the manually-signed matchorders POST (see postRawXML).
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
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
		client:         client,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		accessToken:    accessToken,
		accessSecret:   accessSecret,
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

// postRawXML sends the matchorders POST and returns the raw response body.
//
// CHPP expects the "lineup" JSON as an application/x-www-form-urlencoded
// body field, with every other parameter in the query string, and the OAuth1
// signature computed over the query-string parameters only (lineup excluded).
// The mrjones client can't express that: for a form body it signs the body
// and drops the query string, and for a JSON body CHPP never reads the lineup
// field ("JSON data for lineup was not POSTed"). So the request is signed by
// hand here over the query string alone.
func (c *rawClient) postRawXML(file, version string, values map[string]string, lineup string) ([]byte, error) {
	u := buildURL(file, version, values)

	header, err := c.oauthHeader(u, map[string]string{"lineup": lineup})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader("lineup="+oauthEscape(lineup)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", header)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// oauthHeader builds an OAuth1 (HMAC-SHA1) Authorization header signed over
// the request's query-string parameters plus the given form body parameters,
// as OAuth1 requires for an application/x-www-form-urlencoded POST.
func (c *rawClient) oauthHeader(rawURL string, bodyParams map[string]string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	nonceBytes := make([]byte, 16)
	if _, err := rand.Read(nonceBytes); err != nil {
		return "", err
	}
	oauthParams := map[string]string{
		"oauth_consumer_key":     c.consumerKey,
		"oauth_token":            c.accessToken,
		"oauth_nonce":            hex.EncodeToString(nonceBytes),
		"oauth_timestamp":        fmt.Sprintf("%d", time.Now().Unix()),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_version":          "1.0",
	}

	// Base-string parameters: query-string params plus the oauth_* params.
	var pairs []string
	for k, vs := range u.Query() {
		for _, v := range vs {
			pairs = append(pairs, oauthEscape(k)+"="+oauthEscape(v))
		}
	}
	for k, v := range oauthParams {
		pairs = append(pairs, oauthEscape(k)+"="+oauthEscape(v))
	}
	for k, v := range bodyParams {
		pairs = append(pairs, oauthEscape(k)+"="+oauthEscape(v))
	}
	sort.Strings(pairs)

	baseURL := u.Scheme + "://" + u.Host + u.Path
	base := http.MethodPost + "&" + oauthEscape(baseURL) + "&" + oauthEscape(strings.Join(pairs, "&"))
	key := oauthEscape(c.consumerSecret) + "&" + oauthEscape(c.accessSecret)

	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(base))
	oauthParams["oauth_signature"] = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	var parts []string
	for k, v := range oauthParams {
		parts = append(parts, oauthEscape(k)+`="`+oauthEscape(v)+`"`)
	}
	sort.Strings(parts)
	return "OAuth " + strings.Join(parts, ", "), nil
}

// oauthEscape percent-encodes per RFC 5849 (OAuth): unreserved characters
// A-Z a-z 0-9 - . _ ~ are left as-is, everything else is %XX uppercase.
func oauthEscape(s string) string {
	const upperhex = "0123456789ABCDEF"
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '~' {
			b.WriteByte(ch)
			continue
		}
		b.WriteByte('%')
		b.WriteByte(upperhex[ch>>4])
		b.WriteByte(upperhex[ch&0x0f])
	}
	return b.String()
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
