package value_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"go-ddd-services/service/customeraccounts/application/domain/customer/value"
	"testing"
)

func TestPersonName_Equals(t *testing.T) {
	Convey("Given a PersonName", t, func() {
		personName := value.RebuildPersonName("Omri", "Shapira")

		Convey("When it is compared with an identical PersonName", func() {
			identicalPersonName := value.RebuildPersonName(personName.GivenName(), personName.FamilyName())
			isEqual := personName.Equals(identicalPersonName)

			Convey("Then it should be equal", func() {
				So(isEqual, ShouldBeTrue)
			})
		})

		Convey("When it is compared with another PersonName with different givenName", func() {
			differentPersonName := value.RebuildPersonName("Amit", personName.FamilyName())
			isEqual := personName.Equals(differentPersonName)

			Convey("Then it should not be equal", func() {
				So(isEqual, ShouldBeFalse)
			})
		})

		Convey("When it is compared with another PersonName with different familyName", func() {
			differentPersonName := value.RebuildPersonName(personName.GivenName(), "Jackson")
			isEqual := personName.Equals(differentPersonName)

			Convey("Then it should not be equal", func() {
				So(isEqual, ShouldBeFalse)
			})
		})
	})
}
