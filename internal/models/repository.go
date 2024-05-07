package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

type ID struct {
	Value int
}

func (r *Repository) getGenderId(GENDER string) (int, error) {

	id := &ID{}

	genIdReq := `SELECT ID FROM GENDER WHERE GENDER = ?`

	err := r.DB.QueryRow(genIdReq, GENDER).Scan(&id.Value)

	if err != nil {
		return 0, err
	}

	return id.Value, nil

}

func (r *Repository) InsertPerson(FIRSTNAME string,
	MIDNAME string,
	LASTNAME string,
	DATE string,
	GENDER string) (int, error) {

	genderId, errGender := r.getGenderId(GENDER)

	if errGender != nil {
		return 0, errGender
	}

	// genderId := fmt.Sprintf("(SELECT Id FROM GENDER WHERE Gender = '%s')", GENDER)

	stmt := `INSERT INTO persons (FIRSTNAME,
			MIDNAME,
			LASTNAME,
			DATE,
			GENDER_ID)
			VALUES(?,?,?,?,?)`

	result, err := r.DB.Exec(stmt, FIRSTNAME, MIDNAME, LASTNAME, DATE, genderId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *Repository) GetPersonBy(id int) (*Person, error) {

	p := &Person{}

	request := `SELECT p.ID, p.FIRSTNAME, p.MIDNAME, p.LASTNAME, p.DATE, 
	g.Gender FROM persons as p 	join Gender as g ON p.GENDER_ID = g.Id
	where p.ID = ?`

	err := r.DB.QueryRow(request, id).Scan(&p.ID, &p.FIRSTNAME, &p.MIDNAME, &p.LASTNAME, &p.DATE, &p.GENDER)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return p, nil
}

func (r *Repository) GetListOfPersons() ([]*Person, error) {

	request := `SELECT p.ID, p.FIRSTNAME, p.MIDNAME, p.LASTNAME, p.DATE, 
	g.Gender FROM persons as p 	join Gender as g ON p.GENDER_ID = g.Id`

	rows, err := r.DB.Query(request)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := []*Person{}

	for rows.Next() {
		p := &Person{}

		err = rows.Scan(&p.ID, &p.FIRSTNAME, &p.MIDNAME, &p.LASTNAME, &p.DATE, &p.GENDER)
		if err != nil {
			return nil, err
		}

		people = append(people, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}

func (r *Repository) DeletePersonBy(personId int) (int, int, error) {
	stmt := fmt.Sprintf("DELETE FROM persons WHERE ID = %d", personId)

	result, err := r.DB.Exec(stmt)

	if err != nil {
		return 0, 0, err
	}

	id, err := result.RowsAffected()

	if err != nil {
		return 0, 0, err
	}

	return int(id), int(personId), nil
}
