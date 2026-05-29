package config

type SIPConfig struct {
	UserPort   string
	ServerHost string
	Proxy      string

	From string
	To   string

	Username string
	Password string
	Realm    string

	Expire int

	UserAgent string
}
