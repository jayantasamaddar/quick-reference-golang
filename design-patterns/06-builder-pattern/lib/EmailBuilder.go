package lib

import "strings"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

// Builder methods
func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain '@'")
	}
	eb.email.from = from
	return eb
}
func (eb *EmailBuilder) To(to string) *EmailBuilder {
	eb.email.to = to
	return eb
}
func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}
func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

// Send the email
func sendMailImpl(email *email) {
	// Implementation
}

// Builder Parameter: Takes a builder and does something with it, typically calls something on the builder
type build func(*EmailBuilder)

// Publicly exposed function that people are meant to use
func SendEmail(action build) {
	b := EmailBuilder{}
	action(&b)
	sendMailImpl(&b.email)
}
