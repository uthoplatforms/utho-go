package utho

import (
	"errors"
)

type AccountService service

type Account struct {
	User    User   `json:"user"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type User struct {
	// Basic Information
	ID       string `json:"id,omitempty" faker:"boundary_start=100000,boundary_end=999999"`
	Type     string `json:"type,omitempty" faker:"oneof: Individual, Business"`
	Fullname string `json:"fullname,omitempty" faker:"name"`
	Company  string `json:"company,omitempty" faker:"company"`
	Email    string `json:"email,omitempty" faker:"email"`

	// Address Information
	Address  string `json:"address,omitempty" faker:"street_address"`
	City     string `json:"city,omitempty" faker:"city"`
	State    string `json:"state,omitempty" faker:"state"`
	Country  string `json:"country,omitempty" faker:"country_code"`
	Postcode string `json:"postcode,omitempty" faker:"postcode"`

	// Contact Information
	Mobile   string `json:"mobile,omitempty" faker:"phone_number"`
	Mobilecc string `json:"mobilecc,omitempty" faker:"oneof: +1, +44, +91, +81"`

	// GST Information
	Gstnumber string `json:"gstnumber,omitempty" faker:"len=15"`

	// Support Information
	SupportneedTitle        string `json:"supportneed_title,omitempty" faker:"word"`
	SupportneedUsecase      string `json:"supportneed_usecase,omitempty" faker:"sentence"`
	SupportneedBusinesstype string `json:"supportneed_businesstype,omitempty" faker:"word"`
	SupportneedMonthlyspend string `json:"supportneed_monthlyspend,omitempty" faker:"word"`
	SupportneedEmployeesize string `json:"supportneed_employeesize,omitempty" faker:"word"`
	SupportFieldsRequired   string `json:"support_fields_required,omitempty" faker:"word"`

	// Financial Information
	Currencyprefix  string  `json:"currencyprefix,omitempty" faker:"oneof: $, €, £"`
	Currencyrate    string  `json:"currencyrate,omitempty" faker:"amount_with_currency"`
	Currency        string  `json:"currency,omitempty" faker:"oneof: USD, EUR, GBP"`
	Credit          float64 `json:"credit,omitempty" faker:"amount"`
	Availablecredit float64 `json:"availablecredit,omitempty" faker:"amount"`
	Freecredit      float64 `json:"freecredit,omitempty" faker:"amount"`
	Currentusages   float64 `json:"currentusages,omitempty" faker:"amount"`

	// Verification and Security
	Kyc           string   `json:"kyc,omitempty" faker:"oneof: 0, 1"`
	KycData       []string `json:"kyc_data,omitempty"`
	SmsVerified   string   `json:"sms_verified,omitempty" faker:"oneof: 0, 1"`
	Verify        string   `json:"verify,omitempty" faker:"oneof: 0, 1"`
	Twofa         string   `json:"twofa,omitempty" faker:"oneof: Completed, Pending"`
	TwofaSettings string   `json:"twofa_settings,omitempty" faker:"oneof: enabled, none"`
	EmailVerified string   `json:"email_verified,omitempty" faker:"oneof: 0, 1"`

	// Partner and Reseller Information
	IsPartner  string `json:"is_partner,omitempty" faker:"oneof: 0, 1"`
	Partnerid  string `json:"partnerid,omitempty" faker:"uuid_digit"`
	IsReseller string `json:"is_reseller,omitempty" faker:"oneof: 0, 1"`

	// Cloud and Resource Information
	Cloudlimit        string      `json:"cloudlimit,omitempty" faker:"oneof: 10, 25, 50, 100, 200"`
	K8SLimit          string      `json:"k8s_limit,omitempty" faker:"oneof: 0, 5, 10"`
	TotalCloudservers string      `json:"total_cloudservers,omitempty" faker:"oneof: 1, 3, 5, 10, 15"`
	Resources         []Resources `json:"resources,omitempty"`

	// Billing Information
	Singleinvoice      string `json:"singleinvoice,omitempty" faker:"oneof: 0, 1"`
	RazorpayCustomerid string `json:"razorpay_customerid,omitempty" faker:"uuid_digit"`
	RazorpayOrderid    string `json:"razorpay_orderid,omitempty" faker:"uuid_digit"`
	RazorpaySub        string `json:"razorpay_sub,omitempty" faker:"oneof: 0, 1"`
	StripeCustomer     string `json:"stripe_customer,omitempty" faker:"uuid_digit"`

	// Miscellaneous
	Permissions      string `json:"permissions,omitempty" faker:"oneof: admin, user, full"`
	Rvn              string `json:"rvn,omitempty" faker:"amount"`
	CAdded           string `json:"c_added,omitempty" faker:"oneof: yes, no"`
	AffiliateLoginid string `json:"affiliate_loginid,omitempty" faker:"uuid_digit"`
}

type Resources struct {
	Product string `json:"product,omitempty" faker:"oneof: cloud, kubernetes_master, kubernetes_worker, loadbalancer"`
	Count   string `json:"count,omitempty" faker:"oneof: 1, 2, 3, 5, 7, 10"`
}

func (s *AccountService) Read() (*User, error) {
	userUrl := "account/info"
	req, _ := s.client.NewRequest("GET", userUrl)

	var account Account
	if _, err := s.client.Do(req, &account); err != nil {
		return nil, errors.New("failed to fetch account information: " + err.Error())
	}

	if account.Status != "success" && account.Status != "" {
		return nil, errors.New("account service error: " + account.Message)
	}

	if account.User.ID == "" {
		return nil, errors.New("user not found in account information")
	}

	return &account.User, nil
}
