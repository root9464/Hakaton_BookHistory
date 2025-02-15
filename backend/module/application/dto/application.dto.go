package application_dto

type CreateApplicationDto struct {
	SenderID string `form:"sender_id"`
	// сведения о жителе Оренбургской области
	FIO                         string `form:"fio"`
	DateOfBirth                 string `form:"date_of_birth"`
	PlaceOfBirth                string `form:"place_of_birth"`
	NameOfMillitaryCommissariat string `form:"name_of_millitary_commissariat"`
	MillitaryRank               string `form:"millitary_rank"`
	DateOfDeath                 string `form:"date_of_death"`
	BurialPlace                 string `form:"burial_place"`
	BiographicalFacts           string `form:"biographical_facts"`
}
