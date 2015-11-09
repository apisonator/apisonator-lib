package apisonator

type ProxiesService struct {
	client *Client
}

type Proxy struct {
	ID        int    `json:"id, omitempty"`
	Endpoint  string `json:"endpoint, omitempty"`
	Subdomain string `json:"subdomain, omitempty"`
	APIKey    string `json:"api_key, omitempty"`
	CreatedAt string `json:"created_at, omitempty"`
	UpdatedAt string `json:"updated_at, omitempty"`
}

func (s *ProxiesService) Create(apikey, endpoint, subdomain string) (*Response, error) {
	u := "api/v1/proxies"
	v := new(Proxy)
	proxy := new(Proxy)
	proxy.APIKey = apikey
	proxy.Endpoint = endpoint
	proxy.Subdomain = subdomain
	Response, err := s.client.Call("POST", u, proxy, v)
	return Response, err
}
