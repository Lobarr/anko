package packages

import "github.com/jordan-wright/email"

func init() {
	Packages["email"] = map[string]interface{}{
		"NewEmail":           email.NewEmail,
		"NewEmailFromReader": email.NewEmailFromReader,
		"Pool":               email.NewPool,
	}
	PackageTypes["email"] = map[string]interface{}{
		"Email":      email.Email{},
		"Attachment": email.Attachment{},
		"Pool":       email.Pool{},
	}
}
