package utho

import (
	"errors"
)

type ActionService service

type Actions struct {
	Actions []Action `json:"actions"`
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
}

type Action struct {
	Userid       string `json:"userid" faker:"oneof:197456,11111,22222"`
	ID           string `json:"id" faker:"oneof:97267,94188,124214"`
	Action       string `json:"action" faker:"oneof:start,stop,restart"`
	ResourceType string `json:"resource_type" faker:"oneof:cloud,vm,db"`
	ResourceID   string `json:"resource_id" faker:"oneof:1277087,1627803,1277721"`
	StartedAt    string `json:"started_at" faker:"date"`
	CompletedAt  string `json:"completed_at" faker:"oneof:0000-00-00 00:00:00,2025-05-12 15:00:00"`
	Process      string `json:"process" faker:"oneof:95,96,97"`
	Status       string `json:"status" faker:"oneof:Pending,Support,Success"`
	Message      string `json:"message,omitempty"`
}

func (s *ActionService) List() ([]Action, error) {
	actionUrl := "actions"
	req, _ := s.client.NewRequest("GET", actionUrl)

	var actions Actions
	_, err := s.client.Do(req, &actions)
	if err != nil {
		return nil, err
	}
	if actions.Status != "success" && actions.Status != "" {
		return nil, errors.New(actions.Message)
	}

	return actions.Actions, nil
}

func (s *ActionService) Read(id string) (*Action, error) {
	actionUrl := "actions/" + id
	req, _ := s.client.NewRequest("GET", actionUrl)

	var action Action
	if _, err := s.client.Do(req, &action); err != nil {
		return nil, errors.New("failed to fetch action information: " + err.Error())
	}

	if action.Status == "error" {
		return nil, errors.New("action service error: " + action.Message)
	}

	if action.ID == "" {
		return nil, errors.New("action not found in response")
	}

	return &action, nil
}
