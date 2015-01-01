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
		server = httptest.NewServer(routes.Router())

		var err error
		page, err = agoutiDriver.Page()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		defer server.Close()
		page.Destroy()
	})

	Scenario("creating gists", func() {
		Step("user signs up", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/signup", server.URL))).To(Succeed())
			Fill(page.Find("input[name=email]"), "jack@example.com")
			Fill(page.Find("input[name=username]"), "jack")
			Fill(page.Find("input[name=password]"), "password")
			Fill(page.Find("input[name=password_confirmation]"), "password")
			Submit(page.Find("input[type=submit]"))
		})

		Step("user visits new gist page", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/gists/new", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/gists/new", server.URL)))
		})

		Step("user fills new gist form", func() {
			Fill(page.Find("input[name=title]"), "Sample Gist")

			Expect(page.RunScript(
				"$(elementID).val('Some content')",
				map[string]interface{}{"elementID": "textarea[name=content]"},
				nil,
			)).To(Succeed())

			Submit(page.Find("input[type=submit]"))
		})

		Step("user is redirected to their gists", func() {
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})
	})

	Scenario("updating gists", func() {
		Step("user vists gists page", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})

		Step("user clicks on a gist", func() {
			Click(page.FindByLink("Sample Gist"))
			Expect(page).To(HaveURL(fmt.Sprintf("%s/gists/1", server.URL)))
		})

		Step("user clicks edit", func() {
			Click(page.FindByLink("Edit"))
			Expect(page).To(HaveURL(fmt.Sprintf("%s/gists/1/edit", server.URL)))
		})

		Step("user updates gist", func() {
			Fill(page.Find("input[name=title]"), "Sample Gist Edited")
			Submit(page.Find("input[type=submit]"))
		})

		Step("the gist is saved", func() {
			Expect(page.HTML()).To(ContainSubstring("Sample Gist Edited"))
		})
	})
})
