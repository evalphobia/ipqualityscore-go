package ipqs

type ErrData struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func (r *ErrData) SetStatusCode(code int) {
	r.StatusCode = code
}

func (r ErrData) HasError() bool {
	return r.StatusCode >= 400
}

// IPResponse is a response data of Proxy & VPN Detection API.
// ref: https://www.ipqualityscore.com/documentation/proxy-detection/overview
type IPResponse struct {
	ErrData
	RequestID string `json:"request_id"`

	FraudScore   float64 `json:"fraud_score"`
	Host         string  `json:"host"`
	ISP          string  `json:"isp"`
	ASN          int64   `json:"asn"`
	Organization string  `json:"organization"`
	CountryCode  string  `json:"country_code"`
	Region       string  `json:"region"`
	City         string  `json:"city"`
	ZipCode      string  `json:"zip_code"`
	Timezone     string  `json:"timezone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`

	Mobile         bool   `json:"mobile"`
	IsCrawler      bool   `json:"is_crawler"`
	Proxy          bool   `json:"proxy"`
	VPN            bool   `json:"vpn"`
	Tor            bool   `json:"tor"`
	ActiveVPN      bool   `json:"active_vpn"`
	ActiveTor      bool   `json:"active_tor"`
	RecentAbuse    bool   `json:"recent_abuse"`
	BotStatus      bool   `json:"bot_status"`
	ConnectionType string `json:"connection_type"` // premium only
	AbuseVelocity  string `json:"abuse_velocity"`  // premium only

	// with user-agent parameter
	OperatingSystem   string            `json:"operating_system"`
	Browser           string            `json:"browser"`
	DeviceModel       string            `json:"device_model"`
	DeviceBrand       string            `json:"device_brand"`
	TransactionDetail TransactionDetail `json:"transaction_details"`
}

type TransactionDetail struct {
	ValidBillingAddress       bool     `json:"valid_billing_address"`
	ValidShippingAddress      bool     `json:"valid_shipping_address"`
	ValidBillingEmail         bool     `json:"valid_billing_email"`
	ValidShippingEmail        bool     `json:"valid_shipping_email"`
	RiskyBillingPhone         bool     `json:"risky_billing_phone"`
	RiskyShippingPhone        bool     `json:"risky_shipping_phone"`
	BillingPhoneCarrier       string   `json:"billing_phone_carrier"`
	ShippingPhoneCarrier      string   `json:"shipping_phone_carrier"`
	BillingPhoneLineType      string   `json:"billing_phone_line_type"`
	ShippingPhoneLineType     string   `json:"shipping_phone_line_type"`
	BillingPhoneCountryCode   string   `json:"billing_phone_country_code"`
	ShippingPhoneCountryCode  string   `json:"shipping_phone_country_code"`
	FraudulentBehavior        bool     `json:"fraudulent_behavior"`
	BinCountry                string   `json:"bin_country"`
	RiskScore                 float64  `json:"risk_score"`
	RiskFactors               []string `json:"risk_factors"`
	IsPrepaidCard             bool     `json:"is_prepaid_card"`
	RiskyUserName             bool     `json:"risky_user_name"`
	ValidBillingPhone         bool     `json:"valid_billing_phone"`
	ValidShippingPhone        bool     `json:"valid_shipping_phone"`
	LeakedBillingEmail        bool     `json:"leaked_billing_email"`
	LeakedShippingEmail       bool     `json:"leaked_shipping_email"`
	LeakedUserData            bool     `json:"leaked_user_data"`
	PhoneNameIdentityMatch    bool     `json:"phone_name_identity_match"`
	PhoneEmailIdentityMatch   bool     `json:"phone_email_identity_match"`
	PhoneAddressIdentityMatch bool     `json:"phone_address_identity_match"`
	EmailNameIdentityMatch    bool     `json:"email_name_identity_match"`
	NameAddressIdentityMatch  bool     `json:"name_address_identity_match"`
	AddressEmailIdentityMatch bool     `json:"address_email_identity_match"`
}

// EmailResponse is a response data of Email Validation API.
// ref: https://www.ipqualityscore.com/documentation/email-validation/overview
type EmailResponse struct {
	ErrData
	RequestID string `json:"request_id"`

	Valid              bool    `json:"valid"`
	Disposable         bool    `json:"disposable"`
	SMTPScore          float64 `json:"smtp_score"`
	OverallScore       float64 `json:"overall_score"`
	FirstName          string  `json:"first_name"`
	Generic            bool    `json:"generic"`
	Common             bool    `json:"common"`
	DNSValid           bool    `json:"dns_valid"`
	Honeypot           bool    `json:"honeypot"`
	Deliverability     string  `json:"deliverability"`
	FrequentComplainer bool    `json:"frequent_complainer"`
	SpamTrapScore      string  `json:"spam_trap_score"`
	CatchAll           bool    `json:"catch_all"`
	TimedOut           bool    `json:"timed_out"`
	Suspect            bool    `json:"suspect"`
	RecentAbuse        bool    `json:"recent_abuse"`
	FraudScore         float64 `json:"fraud_score"`
	SuggestedDomain    string  `json:"suggested_domain"`
	Leaked             bool    `json:"leaked"`
	DomainAge          AgeData `json:"domain_age"`
	FirstSeen          AgeData `json:"first_seen"`
	SanitizedEmail     string  `json:"sanitized_email"`
}

type AgeData struct {
	Human     string `json:"human"`
	Timestamp int64  `json:"timestamp"`
	ISO       string `json:"iso"`
}
