package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tristanoneil/badger/models"
)

var _ = Describe("Models/User", func() {
	var validUser models.User
	var invalidUser models.User

	BeforeEach(func() {
		validUser = models.User{
			Email:                "user@example.com",
			Username:             "user",
			Password:             "password",
			PasswordConfirmation: "password",
		}

		invalidUser = models.User{
			Email:                "user@example.com",
			Username:             "user",
			Password:             "password",
			PasswordConfirmation: "notpassword",
		}
	})

	Describe("Validate", func() {
		It("valid gists should be valid", func() {
			Expect(validUser.Validate()).To(Equal(true))
		})

		It("invalid gists should be invalid", func() {
			Expect(invalidUser.Validate()).To(Equal(false))
		})

		It("invalid gists should have errors", func() {
			invalidUser.Validate()
			Expect(invalidUser.Errors["Password"]).ToNot(Equal(""))
		})
	})

	Describe("GravatarURL", func() {
		It("should return a gravatar URL for a given size", func() {
			Expect(validUser.GravatarURL(72)).To(Equal("https://secure.gravatar.com/avatar/b58996c504c5638798eb6b511e6f49af?s=72"))
		})
	})

	Describe("Create", func() {
		It("creates a new user in the database", func() {
			var count int
			models.Db.Get(&count, "SELECT count(*) FROM users")
			Expect(count).To(Equal(0))

			validUser.Create()

			models.Db.Get(&count, "SELECT count(*) FROM users")
			Expect(count).To(Equal(1))
		})
	})

	Describe("UniqueEmail", func() {
		It("returns false if a user isn't unique", func() {
			Expect(invalidUser.UniqueEmail()).To(Equal(false))
		})
	})

	Describe("UniqueUsername", func() {
		It("returns false if a user isn't unique", func() {
			Expect(invalidUser.UniqueUsername()).To(Equal(false))
		})
	})

	Describe("IsValidUser", func() {
		It("returns true for valid credentials", func() {
			Expect(models.IsValidUser("user@example.com", "password")).To(Equal(true))
		})

		It("returns false for invalid credentials", func() {
			Expect(models.IsValidUser("user@example.com", "badpassword")).To(Equal(false))
		})
	})

	Describe("FindUser", func() {
		It("returns a user for a given email", func() {
			user, _ := models.FindUser("user@example.com")
			Expect(user.Email).To(Equal(validUser.Email))
		})
	})
})
