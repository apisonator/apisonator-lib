package apisonator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client         *http.Client
	baseURL        *url.URL
	Authentication *RegistrationsService
	Proxies        *ProxiesService
}

type Response struct {
	*http.Response
}

func NewClient(instance string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if len(instance) == 0 {
		return nil, fmt.Errorf("No Apisonator instance given.")
	}
	baseURL, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: baseURL,
	}

	c.Authentication = &RegistrationsService{client: c}
	c.Proxies = &ProxiesService{client: c}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.buildURLForRequest(urlStr)
	if err != nil {
		return nil, err
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	return req, nil
}

func (c *Client) Call(method, u string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func (c *Client) buildURLForRequest(urlStr string) (string, error) {
	u := c.baseURL.String()
	if strings.HasSuffix(u, "/") == false {
		u += "/"
	}

	if strings.HasPrefix(urlStr, "/") == true {
		urlStr = urlStr[1:]
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	u += rel.String()

	return u, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return response, err
			}
			err = json.Unmarshal(body, v)
		}
	}
	return response, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("API call to %s failed: %s", r.Request.URL.String(), r.Status)
}
