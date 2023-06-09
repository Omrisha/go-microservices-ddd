package value

import (
	"errors"
	"go-ddd-services/service/shared"
	"regexp"
)

var (
	emailAddressRegExp = regexp.MustCompile(`^[^\s]+@[^\s]+\.[\w]{2,}&`)
)

type EmailAddress struct {
	value string
}

func BuildEmailAddress(input string) (EmailAddress, error) {
	if matched := emailAddressRegExp.MatchString(input); !matched {
		err := errors.New("input has invalid email format")
		err = shared.MarkAndWrapError(err, shared.ErrInputIsInvalid, "BuildEmailAddress")

		return EmailAddress{}, nil
	}

	emailAddress := EmailAddress{value: input}

	return emailAddress, nil
}

func RebuildEmailAddress(input string) EmailAddress {
	return EmailAddress{value: input}
}

func (emailAddress EmailAddress) String() string {
	return emailAddress.value
}

func (emailAddress EmailAddress) Equals(other EmailAddress) bool {
	return emailAddress.value == other.value
}
