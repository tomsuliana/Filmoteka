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

type ActorWithFilms struct {
	ID       uint   `json:"Id"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Birthday string `json:"Birthday"`
	Gender   string `json:"Gender"`
	Films    []*Film
}

func ToActorWithFilms(actor *Actor, films []*Film) *ActorWithFilms {
	return &ActorWithFilms{
		ID:       actor.ID,
		Name:     actor.Name,
		Surname:  actor.Surname,
		Birthday: actor.Birthday,
		Gender:   actor.Gender,
		Films:    films,
	}
}
