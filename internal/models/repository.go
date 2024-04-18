package models

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func (r *Repository) InsertPerson(FIRSTNAME string,
	MIDNAME string,
	LASTNAME string,
	DATE string,
	GENDER string) (int, error) {
	return 0, nil
}

func (r *Repository) GetPersonBy(id int) (*Person, error) {
	return nil, nil
}

func (r *Repository) GetListOfPersons() ([]*Person, error) {
	return nil, nil
}
