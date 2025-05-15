package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestDomainService_CreateDomain_happyPath(t *testing.T) {
	var payload CreateDomainParams
	_ = faker.FakeData(&payload)

	token := faker.UUIDDigit()

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResponse BasicResponse
	_ = faker.FakeData(&fakeResponse)
	fakeResponse.Status = "success"
	fakeResponseJson, _ := json.Marshal(fakeResponse)

	mux.HandleFunc("/dns/adddomain", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(fakeResponseJson))
	})

	got, err := client.Domain().CreateDomain(payload)

	assert.Nil(t, err)
	assert.Equal(t, fakeResponse, *got)
}

func TestDomainService_CreateDomain_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Domain().CreateDomain(CreateDomainParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestDomainService_ReadDomain_happyPath(t *testing.T) {
	var fakeDomain DnsDomains
	_ = faker.FakeData(&fakeDomain)
	if len(fakeDomain.Domains) == 0 {
		var d Domain
		_ = faker.FakeData(&d)
		fakeDomain.Domains = append(fakeDomain.Domains, d)
	} else if len(fakeDomain.Domains) > 1 {
		fakeDomain.Domains = fakeDomain.Domains[:1]
	}
	fakeDomain.Status = "success"

	client, mux, _, teardown := setup(faker.UUIDDigit())
	defer teardown()

	domainName := fakeDomain.Domains[0].Domain
	serverResponseJson, _ := json.Marshal(fakeDomain)

	mux.HandleFunc("/dns/"+domainName, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponseJson))
	})

	want := fakeDomain.Domains[0]

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

	var fakeDomains DnsDomains
	_ = faker.FakeData(&fakeDomains)
	if len(fakeDomains.Domains) < 2 {
		for i := len(fakeDomains.Domains); i < 2; i++ {
			var d Domain
			_ = faker.FakeData(&d)
			fakeDomains.Domains = append(fakeDomains.Domains, d)
		}
	}
	fakeDomains.Status = "success"
	serverResponseJson, _ := json.Marshal(fakeDomains)

	mux.HandleFunc("/dns", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponseJson))
	})

	want := fakeDomains.Domains

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
	var domain Domain
	_ = faker.FakeData(&domain)

	token := faker.UUIDDigit()
	domainName := domain.Domain

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResponse DeleteResponse
	_ = faker.FakeData(&fakeResponse)
	fakeResponse.Status = "success"
	fakeResponse.Message = "success"
	fakeResponseJson, _ := json.Marshal(fakeResponse)

	mux.HandleFunc("/dns/"+domainName+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(fakeResponseJson))
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

func TestDomainService_CreateDnsRecord_happyPath(t *testing.T) {
	var payload CreateDnsRecordParams
	_ = faker.FakeData(&payload)
	payload.Domain = faker.DomainName()

	token := faker.UUIDDigit()

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResponse CreateResponse
	_ = faker.FakeData(&fakeResponse)
	fakeResponse.Status = "success"
	fakeResponseJson, _ := json.Marshal(fakeResponse)

	mux.HandleFunc("/dns/"+payload.Domain+"/record/add", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(fakeResponseJson))
	})

	got, err := client.Domain().CreateDnsRecord(payload)

	assert.Nil(t, err)
	assert.Equal(t, fakeResponse, *got)
}

func TestDomainService_CreateDnsRecord_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.Domain().CreateDnsRecord(CreateDnsRecordParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestDomainService_ReadDnsRecord_happyPath(t *testing.T) {
	var fakeDomainResponse DnsDomains
	_ = faker.FakeData(&fakeDomainResponse)
	if len(fakeDomainResponse.Domains) == 0 {
		var d Domain
		_ = faker.FakeData(&d)
		if len(d.Records) == 0 {
			var r DnsRecord
			_ = faker.FakeData(&r)
			d.Records = append(d.Records, r)
		}
		fakeDomainResponse.Domains = append(fakeDomainResponse.Domains, d)
	} else if len(fakeDomainResponse.Domains) > 1 {
		fakeDomainResponse.Domains = fakeDomainResponse.Domains[:1]
		if len(fakeDomainResponse.Domains[0].Records) == 0 {
			var r DnsRecord
			_ = faker.FakeData(&r)
			fakeDomainResponse.Domains[0].Records = append(fakeDomainResponse.Domains[0].Records, r)
		}
	} else {
		if len(fakeDomainResponse.Domains[0].Records) == 0 {
			var r DnsRecord
			_ = faker.FakeData(&r)
			fakeDomainResponse.Domains[0].Records = append(fakeDomainResponse.Domains[0].Records, r)
		}
	}
	fakeDomainResponse.Status = "success"

	client, mux, _, teardown := setup(faker.UUIDDigit())
	defer teardown()

	domainName := fakeDomainResponse.Domains[0].Domain
	dnsRecordID := fakeDomainResponse.Domains[0].Records[0].ID
	serverResponseJson, _ := json.Marshal(fakeDomainResponse)

	mux.HandleFunc("/dns/"+domainName, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponseJson))
	})

	want := fakeDomainResponse.Domains[0].Records[0]

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

	var fakeDomainResponse DnsDomains
	_ = faker.FakeData(&fakeDomainResponse)
	if len(fakeDomainResponse.Domains) == 0 {
		var d Domain
		_ = faker.FakeData(&d)
		if len(d.Records) < 2 {
			for i := len(d.Records); i < 2; i++ {
				var r DnsRecord
				_ = faker.FakeData(&r)
				d.Records = append(d.Records, r)
			}
		}
		fakeDomainResponse.Domains = append(fakeDomainResponse.Domains, d)
	} else if len(fakeDomainResponse.Domains) > 1 {
		fakeDomainResponse.Domains = fakeDomainResponse.Domains[:1]
		if len(fakeDomainResponse.Domains[0].Records) < 2 {
			for i := len(fakeDomainResponse.Domains[0].Records); i < 2; i++ {
				var r DnsRecord
				_ = faker.FakeData(&r)
				fakeDomainResponse.Domains[0].Records = append(fakeDomainResponse.Domains[0].Records, r)
			}
		}
	} else {
		if len(fakeDomainResponse.Domains[0].Records) < 2 {
			for i := len(fakeDomainResponse.Domains[0].Records); i < 2; i++ {
				var r DnsRecord
				_ = faker.FakeData(&r)
				fakeDomainResponse.Domains[0].Records = append(fakeDomainResponse.Domains[0].Records, r)
			}
		}
	}
	fakeDomainResponse.Status = "success"

	domainName := fakeDomainResponse.Domains[0].Domain
	serverResponseJson, _ := json.Marshal(fakeDomainResponse)

	mux.HandleFunc("/dns/"+domainName, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponseJson))
	})

	want := fakeDomainResponse.Domains[0].Records

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
	var fakeDomain Domain
	_ = faker.FakeData(&fakeDomain)
	if len(fakeDomain.Records) == 0 {
		var r DnsRecord
		_ = faker.FakeData(&r)
		fakeDomain.Records = append(fakeDomain.Records, r)
	}

	token := faker.UUIDDigit()
	domainName := fakeDomain.Domain
	recordId := fakeDomain.Records[0].ID

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResponse DeleteResponse
	_ = faker.FakeData(&fakeResponse)
	fakeResponse.Status = "success"
	fakeResponse.Message = "success"
	fakeResponseJson, _ := json.Marshal(fakeResponse)

	mux.HandleFunc("/dns/"+domainName+"/record/"+recordId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, string(fakeResponseJson))
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
