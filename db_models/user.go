package models

type CCUser struct {
	Id 				string `json:"id"`
    Firstname 		string `json:"firstname"`
    Lastname 		string `json:"lastname"`
    Username 		string `json:"username"`
    Password 		string `json:"password"`
    Email_id		string `json:"email_id"`
    Image			string `json:"image"`
}