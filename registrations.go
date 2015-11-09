package apisonator

type RegistrationsService struct {
	client *Client
}

type Registration struct {
	ID        int    `json:"id, omitempty"`
	Email     string `json:"email, omitempty"`
	Password  string `json:"password, omitempty"`
	APIKey    string `json:"api_key, omitempty"`
	CreatedAt string `json:"created_at, omitempty"`
	UpdatedAt string `json:"updated_at, omitempty"`
}

func (s *RegistrationsService) Register(email, password string) (string, *Response, error) {
	u := "api/v1/registrations"
	v := new(Registration)
	registration := new(Registration)
	registration.Email = email
	registration.Password = password
	Response, err := s.client.Call("POST", u, registration, v)
	return v.APIKey, Response, err
}

func (s *RegistrationsService) Login(email, password string) (string, *Response, error) {
	u := "api/v1/sessions"
	v := new(Registration)
	registration := new(Registration)
	registration.Email = email
	registration.Password = password
	//TODO: MAKE API RETURN CORRECT JSON
	Response, err := s.client.Call("POST", u, registration, v)
	return v.APIKey, Response, err
}
