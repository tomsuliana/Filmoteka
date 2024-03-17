package entity

type Actor struct {
	ID       uint   `json:"Id"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Birthday string `json:"Birthday"`
	Gender   string `json:"Gender"`
}

type RespID struct {
	ID uint `json:"id"`
}
