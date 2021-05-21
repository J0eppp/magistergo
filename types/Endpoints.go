package types

// Endpoints contains all the information about the Magister endpoints (I guess)
type Endpoints struct {
	Issuer                             string     `json:"issuer"`
	JWKSUri                            string     `json:"jwks_uri"`
	AuthorizationEndpoint              string     `json:"authorization_endpoint"`
	TokenEndpoint                      string     `json:"token_endpoint"`
	UserInfoEndpoint                   string     `json:"userinfo_endpoint"`
	EndSessionEndpoint                 string     `json:"end_session_endpoint"`
	CheckSessionIframe                 string     `json:"check_session_endpoint"`
	RevocationEndpoint                 string     `json:"revocation_endpoint"`
	FrontChannelLogoutSupported        bool       `json:"frontchannel_logout_supported"`
	FrontChannelLogoutSessionSupported bool       `json:"frontchannel_logout_session_supported"`
	BackChannelLogoutSupported         bool       `json:"backchannel_logout_supported"`
	BackChannelLogoutSessionSupported  bool       `json:"backchannel_logout_session_supported"`
	ScopesSupported                    [4]string  `json:"scopes_supported"`
	ClaimsSupported                    [19]string `json:"claims_supported"`
	GrantTypesSupported                [6]string  `json:"grant_types_supported"`
	ResponseTypesSupported             [7]string  `json:"response_types_supported"`
	ResponseModesSupported             [4]string  `json:"response_modes_supported"`
	TokenEndpointAuthMethodsSupported  [2]string  `json:"token_endpoint_auth_methods_supported"`
	IDTokenSigningAlgValuesSupported   [1]string  `json:"id_token_signing_alg_values_supported"`
	SubjectTypesSupported              [1]string  `json:"subject_types_supported"`
	CodeChallengeMethodsSupported      [2]string  `json:"code_challenge_methods_supported"`
	RequestParameterSupported          bool       `json:"request_parameter_supported"`
	TenantsEndpoint                    string     `json:"tenants_endpoint"`
}
