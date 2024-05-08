package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	expectedResponse := dummyReadAccountRes
	serverResponse := dummyReadAccountServerRes

	mux.HandleFunc("/account/info", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want User
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.Account().Read()
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestAccountService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.Account().Read()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

const dummyReadAccountServerRes = `{
    "user": {
        "id": "32154",
        "type": "Individual",
        "fullname": "Utho Terraform",
        "company": "",
        "email": "terraform@utho.com",
        "address": "Noida",
        "city": "Noida",
        "state": "Uttar Pradesh",
        "country": "IN",
        "postcode": "201301",
        "mobile": "1204840000",
        "mobilecc": "+91",
        "gstnumber": null,
        "supportneed_title": "",
        "supportneed_usecase": "",
        "supportneed_businesstype": "",
        "supportneed_monthlyspend": "",
        "supportneed_employeesize": "",
        "support_fields_required": "No",
        "twofa_settings": "none",
        "currencyprefix": "$",
        "currencyrate": "0.01333",
        "currency": "USD",
        "credit": 0,
        "availablecredit": 11,
        "freecredit": 323.112,
        "currentusages": 59.6,
        "kyc": "0",
        "sms_verified": "1",
        "verify": "1",
        "is_partner": "0",
        "partnerid": "0",
        "twofa": "Completed",
        "kyc_data": [],
        "email_verified": "0",
        "cloudlimit": "25",
        "k8s_limit": "0",
        "is_reseller": "0",
        "singleinvoice": "0",
        "razorpay_customerid": "cust_Ofrr4fd2K3v",
        "razorpay_orderid": null,
        "stripe_customer": null,
        "total_cloudservers": "3",
        "resources": [
            {
                "product": "cloud",
                "count": "3"
            }
        ],
        "rvn": "100.00",
        "c_added": "no",
        "razorpay_sub": "0",
        "affiliate_loginid": "0"
    }
}`

const dummyReadAccountRes = `{
	"id": "32154",
	"type": "Individual",
	"fullname": "Utho Terraform",
	"company": "",
	"email": "terraform@utho.com",
	"address": "Noida",
	"city": "Noida",
	"state": "Uttar Pradesh",
	"country": "IN",
	"postcode": "201301",
	"mobile": "1204840000",
	"mobilecc": "+91",
	"gstnumber": null,
	"supportneed_title": "",
	"supportneed_usecase": "",
	"supportneed_businesstype": "",
	"supportneed_monthlyspend": "",
	"supportneed_employeesize": "",
	"support_fields_required": "No",
	"twofa_settings": "none",
	"currencyprefix": "$",
	"currencyrate": "0.01333",
	"currency": "USD",
	"credit": 0,
	"availablecredit": 11,
	"freecredit": 323.112,
	"currentusages": 59.6,
	"kyc": "0",
	"sms_verified": "1",
	"verify": "1",
	"is_partner": "0",
	"partnerid": "0",
	"twofa": "Completed",
	"kyc_data": [],
	"email_verified": "0",
	"cloudlimit": "25",
	"k8s_limit": "0",
	"is_reseller": "0",
	"singleinvoice": "0",
	"razorpay_customerid": "cust_Ofrr4fd2K3v",
	"razorpay_orderid": null,
	"stripe_customer": null,
	"total_cloudservers": "3",
	"resources": [
		{
			"product": "cloud",
			"count": "3"
		}
	],
	"rvn": "100.00",
	"c_added": "no",
	"razorpay_sub": "0",
	"affiliate_loginid": "0"
}`
