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
	Userid       string `json:"userid"`
	ID           string `json:"id"`
	Action       string `json:"action"`
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
	StartedAt    string `json:"started_at"`
	CompletedAt  string `json:"completed_at"`
	Process      string `json:"process"`
	Status       string `json:"status"`
}

func (s *ActionService) ListAction() (*[]Action, error) {
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

	return &actions.Actions, nil
}
