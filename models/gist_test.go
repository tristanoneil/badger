package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tristanoneil/badger/models"
)

var _ = Describe("Models/Gist", func() {
	var validGist models.Gist
	var invalidGist models.Gist

	BeforeEach(func() {
		validGist = models.Gist{
			Title:   "Some Title",
			Content: "Some content.",
		}

		invalidGist = models.Gist{
			Title:   "",
			Content: "Some content.",
		}
	})

	Describe("Validate", func() {
		It("valid gists should be valid", func() {
			Expect(validGist.Validate()).To(Equal(true))
		})

		It("invalid gists should be invalid", func() {
			Expect(invalidGist.Validate()).To(Equal(false))
		})

		It("invalid gists should have errors", func() {
			invalidGist.Validate()
			Expect(invalidGist.Errors["Title"]).ToNot(Equal(""))
		})
	})
})
