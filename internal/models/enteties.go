package models

import "fmt"

type Person struct {
	ID        int
	FIRSTNAME string
	MIDNAME   string
	LASTNAME  string
	DATE      string
	GENDER    string
}

func (p *Person) NewPerson(ID int,
	FIRSTNAME string,
	MIDNAME string,
	LASTNAME string,
	DATE string,
	GENDER string) *Person {

	pers := new(Person)
	pers.ID = ID
	pers.FIRSTNAME = FIRSTNAME
	pers.MIDNAME = MIDNAME
	pers.LASTNAME = LASTNAME
	pers.GENDER = GENDER
	pers.DATE = DATE

	return pers
}

func (p *Person) GetShortName(person Person) string {
	return fmt.Sprintf("%s %s %s", person.LASTNAME, cutFirstLetter(p.FIRSTNAME), cutFirstLetter(p.MIDNAME))
}

func cutFirstLetter(name string) string {
	return fmt.Sprintf("%c.", name[0])
}

type Sportsman struct {
	ID         int
	Person     Person
	Categories []Category
	Grade      SportGrade
	Region     Region
}

type Referee struct {
	ID     int
	Person Person
	Grade  RefereeGrade
	Region Region
	Role   RefereeRole
}

type Category struct {
	ID   int
	Name string
}

type SportGrade struct {
	ID   int
	Name string
}

type Region struct {
	ID   int
	Name string
	//LogoURL string
}

type GENDER struct {
	Id   int
	Name string
}

type RefereeGrade struct {
	ID   int
	Name string
}

type RefereeRole struct {
	ID   int
	Name string
}
