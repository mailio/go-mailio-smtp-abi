# Generic interface around receiving and sending email

The Email service can be AWS SES, Mailgun, Postal or any other. 

You can find example implementation using Mailgun's service [here](https://github.com/mailio/go-mailio-mailgun-smtp-handler)

The code inherits the same principle as datastore/sql pluggable interfaces. In this case this interface must be implemented by your custom SMTP handler: 

`go get github.com/mailio/go-mailio-mailgun-smtp-handler`

```go
type SmtpHandler interface {
	// ReceiveMail is a method called on the specific ESP handler webhook implementation
	ReceiveMail(request http.Request) (*types.Mail, error)
	// SendMimeMail returns generated message id or error
	SendMimeMail(mime []byte, to []mail.Address) (string, error)
}
```
