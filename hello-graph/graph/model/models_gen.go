// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PushNotification struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Image       string `json:"image"`
	DeviceToken string `json:"deviceToken"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
