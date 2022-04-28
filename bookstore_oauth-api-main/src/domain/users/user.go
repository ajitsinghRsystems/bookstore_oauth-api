package users

type User struct{
Id int64 `json:"id"`
FirstName string `json:"first_name"`
LastName string	`json:"Last_Name"`
Email string `json:"email"`

}

type UserLoginRequest struct{
	Email string	`json:"email"`
	Password string `json:"Password"`
}