package packages

import sendgrid "github.com/sendgrid/sendgrid-go"

func init() {
	Packages["sendgrid"] = map[string]interface{}{
		"API":           sendgrid.API,
		"GetRequest":    sendgrid.GetRequest,
		"Version":       sendgrid.Version,
		"NewSendClient": sendgrid.NewSendClient,
		"DefaultClient": sendgrid.DefaultClient,
	}
	PackageTypes["sendgrid"] = map[string]interface{}{
		"Client": sendgrid.Client{},
	}
}
