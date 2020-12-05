package models

type Submit struct {
	Email  string `json:"email_addr"`
	WmlUrl string `json:"wml_url"`
	Token  string `json:"iam_token"`
	Submit bool   `json:"submit_confirmation"`
}

