package domain

type RegisterRequest struct {
	UserPort   string
	ServerHost string
	Proxy      string

	From string
	To   string

	Realm    string
	Username string
	Password string

	Expires int

	UserAgent string
}
