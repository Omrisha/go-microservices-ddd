package value

import (
	"errors"
	"go-ddd-services/service/shared"
)

type PersonName struct {
	givenName  string
	familyName string
}

func BuildPersonName(givenName string, familyName string) (PersonName, error) {
	wrapWithMsg := "BuildPersonName"

	if givenName == "" {
		err := errors.New("input has invalid email given name")
		err = shared.MarkAndWrapError(err, shared.ErrInputIsInvalid, wrapWithMsg)

		return PersonName{}, nil
	}

	if familyName == "" {
		err := errors.New("input has invalid email family name")
		err = shared.MarkAndWrapError(err, shared.ErrInputIsInvalid, wrapWithMsg)

		return PersonName{}, nil
	}

	PersonName := PersonName{
		givenName:  givenName,
		familyName: familyName,
	}

	return PersonName, nil
}

func RebuildPersonName(givenName string, familyName string) PersonName {
	return PersonName{
		givenName:  givenName,
		familyName: familyName,
	}
}

func (personName PersonName) GivenName() string {
	return personName.givenName
}

func (personName PersonName) FamilyName() string {
	return personName.familyName
}

func (personName PersonName) Equals(other PersonName) bool {
	if personName.GivenName() != other.GivenName() {
		return false
	}

	if personName.FamilyName() != other.FamilyName() {
		return false
	}

	return true
}
