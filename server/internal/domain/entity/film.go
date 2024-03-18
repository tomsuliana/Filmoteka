package entity

type Film struct {
	ID          uint    `json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	ReleaseDate string  `json:"ReleaseDate"`
	Rating      float32 `json:"Rating"`
}

type FilmWithActors struct {
	ID          uint    `json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	ReleaseDate string  `json:"ReleaseDate"`
	Rating      float32 `json:"Rating"`
	Actors      []*Actor
}

func ToFilm(filmWithActors *FilmWithActors) *Film {
	return &Film{
		ID:          filmWithActors.ID,
		Name:        filmWithActors.Name,
		Description: filmWithActors.Description,
		ReleaseDate: filmWithActors.ReleaseDate,
		Rating:      filmWithActors.Rating,
	}
}

func ToFilmWithActors(film *Film, actors []*Actor) *FilmWithActors {
	return &FilmWithActors{
		ID:          film.ID,
		Name:        film.Name,
		Description: film.Description,
		ReleaseDate: film.ReleaseDate,
		Rating:      film.Rating,
		Actors:      actors,
	}
}
