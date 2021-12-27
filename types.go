package delltdapi

// TokenInfo from API
type TokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"oob"`
}

// Declaring the Constant Values
const (
	StatusUnauthorized        = "Unauthorized"
	StatusInternalServerError = "Server Error"
	StatusBadRequest          = "Bad Request"
	StatusConflict            = "Conflict Already Exists"
	StatusNotFound            = "Not Found"
	StatusForbidden           = "Forbidden"
)

type AssetInfo []struct {
	HostName   string
	ServiceTag string `json:"serviceTag"`
	//ProductId    string `json:"productId"`
	Entitlements []struct {
		EndDate          string `json:"endDate"`
		ServiceLevelCode string `json:"serviceLevelCode"`
	} `json:"entitlements"`
}
