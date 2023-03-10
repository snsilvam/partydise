package models

type OurServices struct {
	ID                    string `json: id`
	Title                 string `json: title`
	DescriptionOfServices string `json: descriptionOfServices`
	mgOfServices          string `json: ImgOfServices`
	status                string `json: status`
}
