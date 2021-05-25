package magistergo

import "net/http"

type Magister_I interface {

}

type Magister struct {
	Tenant string
	AccessToken string
	RefreshToken string
	AccessTokenExpiresAt int64
	Authority string
	Endpoints Endpoints
	//ClientID string
	UserID string
	HTTPClient http.Client
}

type Persoon struct {
	Id int64 `json:"Id"`
}

type AccountData struct {
	Persoon Persoon `json:"Persoon"`
}

//type Magister struct {
//	School string
//	Username string
//	Password string
//	Tenant string
//	Authority string
//	ClientID string
//	RedirectURI string
//	Scope string
//	ResponseType string
//	ACRValues string // Idk what this is
//	DefaultState string
//	DefaultNonce string
//	AuthCode string
//	Endpoints Endpoints
//	HTTPClient http.Client
//	//CookieJar *cookiejar.Jar
//}

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

type LoginOptions struct {
	Username string
	Password string
}

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	expiresIn int64 `json:"expires_in"`
	ExpiresAt int64
}