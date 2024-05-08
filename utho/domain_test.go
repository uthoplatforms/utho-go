package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainService_CreateDomain_happyPath(t *testing.T) {
	token := "token"

	payload := CreateDomainParams{
		Domain: "example.com",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/dns/adddomain", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	got, err := client.Domain().CreateDomain(payload)

	var want BasicResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestDomainService_CreateDomain_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Domain().CreateDomain(CreateDomainParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestDomainService_ReadDomain_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	domainName := "example.com"
	expectedResponse := dummyReadDomainRes
	serverResponse := dummyReadDomainServerRes

	mux.HandleFunc("/dns/"+domainName, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Domain
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Domain().ReadDomain(domainName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestDomainService_ReadDomain_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Domain().ReadDomain("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestDomainService_ListDomain_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyReadDomainRes
	serverResponse := dummyListDomainServerRes

	mux.HandleFunc("/dns", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Domain
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Domain().ListDomains()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d domain to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestDomainService_ListDomain_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	domain, err := client.Domain().ListDomains()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if domain != nil {
		t.Errorf("Was not expecting any domain to be returned, instead got %v", domain)
	}
}

func TestDomainService_DeleteDomain_happyPath(t *testing.T) {
	token := "token"
	domainName := "someDomainName"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/dns/"+domainName+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Domain().DeleteDomain(domainName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestDomainService_DeleteDomain_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Domain().DeleteDomain("someDomainName")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// Dns Record tests
func TestDomainService_CreateDnsRecord_happyPath(t *testing.T) {
	token := "token"

	payload := CreateDnsRecordParams{
		Domain:   "example.com",
		Type:     "A",
		Hostname: "example22.com",
		Value:    "1.1.12.1",
		TTL:      "65444",
		Porttype: "TCP",
		Port:     "5060",
		Priority: "10",
		Wight:    "100",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/dns/"+payload.Domain+"/record/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateBasicResponseJson)
	})

	got, err := client.Domain().CreateDnsRecord(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateBasicResponseJson), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestDomainService_CreateDnsRecord_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Domain().CreateDnsRecord(CreateDnsRecordParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestDomainService_ReadDnsRecord_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	domainName := "example.com"
	dnsRecordID := "12331"
	expectedResponse := dummyReadDomainRes
	serverResponse := dummyReadDomainServerRes

	mux.HandleFunc("/dns/"+domainName, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want DnsRecord
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Domain().ReadDnsRecord(domainName, dnsRecordID)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestDomainService_ReadDnsRecord_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Domain().ReadDnsRecord("example.com", "122134")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestDomainService_ListDnsRecord_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	domainName := "example.com"
	expectedResponse := dummyReadDomainRes
	serverResponse := dummyListDomainServerRes

	mux.HandleFunc("/dns", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []DnsRecord
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Domain().ListDnsRecords(domainName)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d dnsrecord to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestDomainService_ListDnsRecord_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	dnsrecord, err := client.Domain().ListDnsRecords("example.com")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if dnsrecord != nil {
		t.Errorf("Was not expecting any dnsrecord to be returned, instead got %v", dnsrecord)
	}
}

func TestDomainService_DeleteDnsRecord_happyPath(t *testing.T) {
	token := "token"
	domainName := "someDomainName"
	recordId := "53211"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/dns/"+domainName+"/record/"+recordId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.Domain().DeleteDnsRecord(domainName, recordId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestDomainService_DeleteDnsRecord_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.Domain().DeleteDnsRecord("someDnsRecordName", "123543")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

const dummyReadDomainRes = `{
	"domain": "examqweple.com",
	"nspoint": "NO",
	"created_at": "2024-05-03 21:29:23",
	"dnsrecord_count": "2",
	"records": [
		{
			"id": "25244",
			"hostname": "example22.com.examqweple.com",
			"type": "A",
			"value": "1.1.1.1",
			"ttl": "65444",
			"priority": "10"
		},
		{
			"id": "25245",
			"hostname": "example22.com.examqweple.com",
			"type": "A",
			"value": "1.1.12.1",
			"ttl": "65444",
			"priority": "10"
		}
	],
	"record": {
		"id": "25245",
		"hostname": "example22.com.examqweple.com",
		"type": "A",
		"value": "1.1.12.1",
		"ttl": "65444",
		"priority": "10"
	}
}
`

const dummyReadDomainServerRes = `{
    "domains": [
        {
            "domain": "examqweple.com",
            "nspoint": "NO",
            "created_at": "2024-05-03 21:29:23",
            "dnsrecord_count": "2",
            "records": [
                {
                    "id": "25244",
                    "hostname": "example22.com.examqweple.com",
                    "type": "A",
                    "value": "1.1.1.1",
                    "ttl": "65444",
                    "priority": "10"
                },
                {
                    "id": "25245",
                    "hostname": "example22.com.examqweple.com",
                    "type": "A",
                    "value": "1.1.12.1",
                    "ttl": "65444",
                    "priority": "10"
                }
            ],
            "record": {
                "id": "25245",
                "hostname": "example22.com.examqweple.com",
                "type": "A",
                "value": "1.1.12.1",
                "ttl": "65444",
                "priority": "10"
            }
        }
    ]
}
`

const dummyListDomainServerRes = `[
	{
		"domain": "examqweple.com",
		"nspoint": "NO",
		"created_at": "2024-05-03 21:29:23",
		"dnsrecord_count": "2",
		"records": [
			{
				"id": "25244",
				"hostname": "example22.com.examqweple.com",
				"type": "A",
				"value": "1.1.1.1",
				"ttl": "65444",
				"priority": "10"
			},
			{
				"id": "25245",
				"hostname": "example.examqweple.com",
				"type": "A",
				"value": "1.1.12.1",
				"ttl": "65444",
				"priority": "10"
			}
		],
		"record": {
			"id": "25245",
			"hostname": "example.examqweple.com",
			"type": "A",
			"value": "1.1.12.1",
			"ttl": "65444",
			"priority": "10"
		}
	},{
		"domain": "example2.com",
		"nspoint": "NO",
		"created_at": "2024-05-03 21:29:23",
		"dnsrecord_count": "2",
		"records": [
			{
				"id": "25244",
				"hostname": "example.example2.com",
				"type": "A",
				"value": "1.1.1.1",
				"ttl": "65444",
				"priority": "10"
			},
			{
				"id": "25245",
				"hostname": "example.example2.com",
				"type": "A",
				"value": "1.1.12.1",
				"ttl": "65444",
				"priority": "10"
			}
		],
		"record": {
			"id": "25245",
			"hostname": "example22.com.examqweple.com",
			"type": "A",
			"value": "1.1.12.1",
			"ttl": "65444",
			"priority": "10"
		}
	}
]`
