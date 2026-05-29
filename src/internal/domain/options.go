package domain

type OptionsRequest struct {
	UserPort   string
	ServerHost string
	Proxy      string

	From string
	To   string

	UserAgent string
}
