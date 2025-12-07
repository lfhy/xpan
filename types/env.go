package types

var (
	ClientId     string
	ClientSecret string
	RedirectUri  = "oob"
	AccessToken  string
	RefreshToken string
)

func getENV(key string) string {
	switch key {
	case "$CLIENT_ID":
		return ClientId
	case "$CLIENT_SECRET":
		return ClientSecret
	case "$REDIRECT_URI":
		return RedirectUri
	default:
		return key
	}
}
