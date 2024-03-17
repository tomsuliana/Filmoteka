package entity

type Actor struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}

type RespID struct {
	ID uint `json:"id"`
}
