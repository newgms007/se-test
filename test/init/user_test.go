package unit

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/newgms007/se-test/entity"
	. "github.com/onsi/gomega"
)

func TestUsername(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`username is required`, func(t *testing.T) {
		user := entity.User{
			Username: "", // X
			Email:    "testing@email.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Username is required"))
	})
}
