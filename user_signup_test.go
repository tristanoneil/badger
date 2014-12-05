package main_test

import (
	"fmt"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti/core"
	. "github.com/sclevine/agouti/dsl"
	. "github.com/sclevine/agouti/matchers"
	"github.com/tristanoneil/badger/routes"
)

var _ = Describe("UserSignup", func() {
	var (
		page   Page
		server *httptest.Server
	)

	BeforeEach(func() {
		server = httptest.NewServer(routes.Handlers())
		var err error
		page, err = agoutiDriver.Page()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		defer server.Close()
		page.Destroy()
	})

	Scenario("signing up", func() {
		Step("user visits signup route", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/signup", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/signup", server.URL)))
		})

		Step("user fills out signout form", func() {
			Fill(page.Find("input[name=email]"), "john1234567@example.com")
			Fill(page.Find("input[name=password]"), "password")
			Fill(page.Find("input[name=password_confirmation]"), "password")
			Submit(page.Find("input[type=submit]"))
		})

		Step("user is redirected to their gists", func() {
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})
	})
})
