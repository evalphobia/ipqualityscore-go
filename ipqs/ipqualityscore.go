package ipqs

import (
	"github.com/evalphobia/ipqualityscore-go/client"
	"github.com/evalphobia/ipqualityscore-go/config"
	"github.com/evalphobia/ipqualityscore-go/log"
)

// IPQualityScore is service struct for IPQualityScore API.
type IPQualityScore struct {
	client *client.Client
	logger log.Logger
}

// New creates IPQualityScore from Config data.
func New(conf config.Config) (*IPQualityScore, error) {
	cli, err := conf.Client()
	if err != nil {
		return nil, err
	}

	return &IPQualityScore{
		client: cli,
		logger: log.DefaultLogger,
	}, nil
}

// SetLogger sets internal API logger.
func (s *IPQualityScore) SetLogger(logger log.Logger) {
	s.logger = logger
}

// IPReputation executes Proxy & VPN Detection API.
// ref: https://www.ipqualityscore.com/documentation/proxy-detection/overview
func (s *IPQualityScore) IPReputation(ipaddr string, opt ...IPReputationOption) (IPResponse, error) {
	params := make(map[string]interface{})
	params["ip"] = ipaddr
	var o IPReputationOption
	if len(opt) != 0 {
		o = opt[0]
	}

	if o.Strictness > 0 {
		params["strictness"] = o.Strictness
	}
	if o.TransactionStrictness > 0 {
		params["transaction_strictness"] = o.TransactionStrictness
	}
	if o.Fast {
		params["fast"] = o.Fast
	}
	if o.Mobile {
		params["mobile"] = o.Mobile
	}
	if o.AllowPublicAccessPoints {
		params["allow_public_access_points"] = o.AllowPublicAccessPoints
	}
	if o.LighterPenalties {
		params["lighter_penalties"] = o.LighterPenalties
	}
	if o.UserAgent != "" {
		params["user_agent"] = o.UserAgent
	}
	if o.UserLanguage != "" {
		params["user_language"] = o.UserLanguage
	}

	resp := IPResponse{}
	err := s.client.CallGET("/api/json/ip", params, &resp)
	return resp, err
}

type IPReputationOption struct {
	Strictness              int
	Fast                    bool
	UserAgent               string
	UserLanguage            string
	Mobile                  bool
	AllowPublicAccessPoints bool
	LighterPenalties        bool
	TransactionStrictness   int
}

// EmailValidation executes Email Validation API.
// ref: https://www.ipqualityscore.com/documentation/email-validation/overview
func (s *IPQualityScore) EmailValidation(email string, opt ...EmailValidationOption) (EmailResponse, error) {
	params := make(map[string]interface{})
	params["email"] = email
	var o EmailValidationOption
	if len(opt) != 0 {
		o = opt[0]
	}

	if o.Strictness > 0 {
		params["strictness"] = o.Strictness
	}
	if o.Fast {
		params["fast"] = o.Fast
	}
	if o.Timeout > 0 {
		params["timeout"] = o.Timeout
	}
	if o.SuggesDomain {
		params["suggest_domain"] = o.SuggesDomain
	}
	if o.AbuseStrictness > 0 {
		params["abuse_strictness"] = o.AbuseStrictness
	}

	resp := EmailResponse{}
	err := s.client.CallGET("/api/json/email", params, &resp)
	return resp, err
}

type EmailValidationOption struct {
	Strictness      int
	Fast            bool
	Timeout         int
	SuggesDomain    bool
	AbuseStrictness int
}
