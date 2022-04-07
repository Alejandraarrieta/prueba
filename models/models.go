package models

type UserData struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Lastname string `json:"lastname" bson:"lastname,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Photo    string `json:"photo,omitempty" bson:"photo,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Workunit string `json:"workunit,omitempty" bson:"workunit,omitempty"`
	Area     string `json:"area,omitempty" bson:"area,omitempty"`
	Role     string `json:"role,omitempty" bson:"role,omitempty"`
}

type UserLog struct {
	Username string `json:"username,omitempty"`
}

//struct to return token 
type ResponseToken struct{
	Token string `json:"token"`
}
