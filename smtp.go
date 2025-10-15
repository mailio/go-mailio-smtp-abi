package mailiosmtpabi

import (
	"net/http"
	"net/mail"
)

type SmtpHandler interface {
	// ReceiveMail is a method called on the specific ESP handler webhook implementation
	ReceiveMail(request http.Request) (*Mail, error)
	// SendMimeMail returns generated message id or error
	SendMimeMail(from mail.Address, mime []byte, to []mail.Address) (string, error)
	// return list of supported domains by the smtp
	ListDomains() ([]string, error)
}
