package utho

import (
	"errors"
)

type MonitoringService service

type Alerts struct {
	Alerts  []Alert `json:"alerts"`
	Status  string  `json:"status,omitempty"`
	Message string  `json:"message,omitempty"`
}
type Alert struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	RefIds   string `json:"ref_ids"`
	RefType  string `json:"ref_type"`
	Compare  string `json:"compare"`
	Value    string `json:"value"`
	For      string `json:"for"`
	Contacts string `json:"contacts"`
	Status   string `json:"status"`
}

type Contacts struct {
	Contacts []Contact `json:"contacts"`
	Status   string    `json:"status,omitempty"`
	Message  string    `json:"message,omitempty"`
}
type Contact struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Slack        string `json:"slack"`
	Mobilenumber string `json:"mobilenumber"`
}

type CreateAlertParams struct {
	Name     string `json:"name"`
	RefType  string `json:"ref_type"`
	Type     string `json:"type"`
	Compare  string `json:"compare"`
	Value    string `json:"value"`
	For      string `json:"for"`
	Contacts string `json:"contacts"`
	Status   string `json:"status"`
	RefIds   string `json:"ref_ids"`
}

func (s *MonitoringService) CreateAlert(params CreateAlertParams) (*BasicResponse, error) {
	reqUrl := "alert"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var alert BasicResponse
	_, err := s.client.Do(req, &alert)
	if err != nil {
		return nil, err
	}
	if alert.Status != "success" && alert.Status != "" {
		return nil, errors.New(alert.Message)
	}

	return &alert, nil
}

func (s *MonitoringService) ReadAlert(alertId string) (*Alert, error) {
	reqUrl := "alert"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var alerts Alerts
	_, err := s.client.Do(req, &alerts)
	if err != nil {
		return nil, err
	}
	if alerts.Status != "success" && alerts.Status != "" {
		return nil, errors.New(alerts.Message)
	}

	var alert Alert
	for _, v := range alerts.Alerts {
		if v.ID == alertId {
			alert = v
		}
	}

	return &alert, nil
}

func (s *MonitoringService) ListAlerts() (*[]Alert, error) {
	reqUrl := "alert"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var alert Alerts
	_, err := s.client.Do(req, &alert)
	if err != nil {
		return nil, err
	}
	if alert.Status != "success" && alert.Status != "" {
		return nil, errors.New(alert.Message)
	}

	return &alert.Alerts, nil
}

type UpdateAlertParams struct {
	AlertId  string
	Name     string `json:"name"`
	RefType  string `json:"ref_type"`
	Type     string `json:"type"`
	Compare  string `json:"compare"`
	Value    string `json:"value"`
	For      string `json:"for"`
	Contacts string `json:"contacts"`
	Status   string `json:"status"`
	RefIds   string `json:"ref_ids"`
}

func (s *MonitoringService) UpdateAlert(params UpdateAlertParams) (*BasicResponse, error) {
	reqUrl := "alert/" + params.AlertId + "/update"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var alert BasicResponse
	_, err := s.client.Do(req, &alert)
	if err != nil {
		return nil, err
	}
	if alert.Status != "success" && alert.Status != "" {
		return nil, errors.New(alert.Message)
	}

	return &alert, nil
}

// func (s *MonitoringService) DeleteAlert(alertId string) (*DeleteResponse, error) {
// 	reqUrl := "alert/" + alertId
// 	req, _ := s.client.NewRequest("DELETE", reqUrl)

// 	var delResponse DeleteResponse
// 	if _, err := s.client.Do(req, &delResponse); err != nil {
// 		return nil, err
// 	}

//		return &delResponse, nil
//	}
//

// /////////////////////////////////////////////////////////////////

type CreateContactParams struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobilenumber string `json:"mobilenumber"`
	Status       string `json:"status"`
}

func (s *MonitoringService) CreateContact(params CreateContactParams) (*BasicResponse, error) {
	reqUrl := "alert/contact/add"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var contact BasicResponse
	_, err := s.client.Do(req, &contact)
	if err != nil {
		return nil, err
	}
	if contact.Status != "success" && contact.Status != "" {
		return nil, errors.New(contact.Message)
	}

	return &contact, nil
}

func (s *MonitoringService) ReadContact(contactId string) (*Contact, error) {
	reqUrl := "alert/contact/list"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var contacts Contacts
	_, err := s.client.Do(req, &contacts)
	if err != nil {
		return nil, err
	}
	if contacts.Status != "success" && contacts.Status != "" {
		return nil, errors.New(contacts.Message)
	}

	var contact Contact
	for _, v := range contacts.Contacts {
		if v.ID == contactId {
			contact = v
		}
	}

	return &contact, nil
}

func (s *MonitoringService) ListContacts() (*[]Contact, error) {
	reqUrl := "alert/contact/list"
	req, _ := s.client.NewRequest("GET", reqUrl)

	var contact Contacts
	_, err := s.client.Do(req, &contact)
	if err != nil {
		return nil, err
	}
	if contact.Status != "success" && contact.Status != "" {
		return nil, errors.New(contact.Message)
	}

	return &contact.Contacts, nil
}

type UpdateContactParams struct {
	ContactId    string
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobilenumber string `json:"mobilenumber"`
	Status       string `json:"status"`
}

func (s *MonitoringService) UpdateContact(params UpdateContactParams) (*BasicResponse, error) {
	reqUrl := "alert/contact/" + params.ContactId + "/update"
	req, _ := s.client.NewRequest("POST", reqUrl, &params)

	var contact BasicResponse
	_, err := s.client.Do(req, &contact)
	if err != nil {
		return nil, err
	}
	if contact.Status != "success" && contact.Status != "" {
		return nil, errors.New(contact.Message)
	}

	return &contact, nil
}

func (s *MonitoringService) DeleteContact(contactId string) (*DeleteResponse, error) {
	reqUrl := "alert/contact/" + contactId + "/delete"
	req, _ := s.client.NewRequest("DELETE", reqUrl)

	var delResponse DeleteResponse
	if _, err := s.client.Do(req, &delResponse); err != nil {
		return nil, err
	}

	return &delResponse, nil
}
