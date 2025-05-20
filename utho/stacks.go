package utho

import (
	"errors"
)

type StacksService service

type Stacks struct {
	Stacks  []Stack `json:"stacks"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}
type Stack struct {
	ID            string   `json:"id"`
	IsOwner       string   `json:"is_owner"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Distro        []string `json:"distro"`
	LogoURL       string   `json:"logo_url"`
	IsPublic      string   `json:"is_public"`
	IsMarketplace string   `json:"is_marketplace"`
	Status        string   `json:"status"`
	Script        string   `json:"script"`
	Fields        []any    `json:"fields"`
}

type CreateStacksParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Status      string `json:"status"`
	IsPublic    string `json:"is_public"`
	Script      string `json:"script"`
}

func (s *StacksService) Create(params CreateStacksParams) (*CreateResponse, error) {
	reqUrl := "stacks"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var stacks CreateResponse
	_, err := s.client.Do(req, &stacks)
	if err != nil {
		return nil, err
	}
	if stacks.Status != "success" && stacks.Status != "" {
		return nil, errors.New(stacks.Message)
	}

	return &stacks, nil
}

func (s *StacksService) Read(stackId string) (*Stack, error) {
	reqUrl := "stacks"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var stacks Stacks
	_, err := s.client.Do(req, &stacks)
	if err != nil {
		return nil, err
	}
	if stacks.Status != "success" && stacks.Status != "" {
		return nil, errors.New(stacks.Message)
	}

	var stack Stack
	for _, r := range stacks.Stacks {
		if r.ID == stackId {
			stack = r
			break
		}
	}
	if len(stack.ID) == 0 {
		return nil, errors.New("stack not found")
	}

	return &stack, nil
}

func (s *StacksService) List() ([]Stack, error) {
	reqUrl := "stacks"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var stacks Stacks
	_, err := s.client.Do(req, &stacks)
	if err != nil {
		return nil, err
	}
	if stacks.Status != "success" && stacks.Status != "" {
		return nil, errors.New(stacks.Message)
	}

	return stacks.Stacks, nil
}

type UpdateStacksParams struct {
	StackId     string
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Status      string `json:"status"`
	IsPublic    string `json:"is_public"`
	Script      string `json:"script"`
}

func (s *StacksService) Update(params UpdateStacksParams) (*UpdateResponse, error) {
	reqUrl := "stacks/" + params.StackId
	req, _ := s.client.NewRequest("PUT", reqUrl, &params)

	var stacks UpdateResponse
	_, err := s.client.Do(req, &stacks)
	if err != nil {
		return nil, err
	}
	if stacks.Status != "success" && stacks.Status != "" {
		return nil, errors.New(stacks.Message)
	}

	return &stacks, nil
}

func (s *StacksService) Delete(stackId string) (*DeleteResponse, error) {
	reqUrl := "stacks/" + stackId
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}
	if delResponse.Status != "success" && delResponse.Status != "" {
		return nil, errors.New(delResponse.Message)
	}

	return &delResponse, nil
}
