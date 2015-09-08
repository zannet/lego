package models

type User struct {
    First 		string `json:"first"`
    Last 		string `json:"last"`
    Username 	string `json:"username"`
    Password 	string `json:"password"`
    Image		string `json:"image"`
}