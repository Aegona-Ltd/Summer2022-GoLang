package models

type Form struct {
	Id            int    `json:"id"`
	FullName      string `json:"fullname"`
	CompanyName   string `json:"companyname"`
	BusinessPhone string `json:"businessphone"`
	Email         string `json:"email"`
	Message       string `json:"message"`
}

func (f *Form) TableName() string {
	return "form"
}
