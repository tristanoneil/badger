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
			Content: "# Some content.",
		}

		invalidGist = models.Gist{
			Title:   "",
			Content: "# Some content.",
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

	Describe("Create", func() {
		It("creates a new gist in the database", func() {
			var count int
			models.Db.Get(&count, "SELECT count(*) FROM gists;")
			Expect(count).To(Equal(0))

			validGist.Create()

			models.Db.Get(&count, "SELECT count(*) FROM gists;")
			Expect(count).To(Equal(1))
		})
	})

	Describe("Save", func() {
		It("saves a gist in the database", func() {
			gist := models.FindGistForID(1)
			gist.Title = "Hello world"
			gist.Save()
			gist = models.FindGistForID(1)
			Expect(gist.Title).To(Equal("Hello world"))
		})
	})

	Describe("Markdown", func() {
		It("returns content as markdown", func() {
			Expect(string(validGist.Markdown())).To(ContainSubstring("<h1>Some content.</h1>"))
		})
	})
})
