package cookies

type AuthResponse struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
