package main_test

import (
	"fmt"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
	"github.com/tristanoneil/badger/routes"
)

var _ = Describe("UserSignup", func() {
	var (
		page   *agouti.Page
		server *httptest.Server
	)

	BeforeEach(func() {
		server = httptest.NewServer(routes.Router())

		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		defer server.Close()
		page.Destroy()
	})

	It("creating gists", func() {
		By("user signs up", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/signup", server.URL))).To(Succeed())
			Expect(page.Find("input[name=email]").Fill("jack@example.com")).To(Succeed())
			Expect(page.Find("input[name=username]").Fill("jack")).To(Succeed())
			Expect(page.Find("input[name=password]").Fill("password")).To(Succeed())
			Expect(page.Find("input[name=password_confirmation]").Fill("password")).To(Succeed())
			Expect(page.Find("input[type=submit]").Submit()).To(Succeed())
		})

		By("user visits new gist page", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/gists/new", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/gists/new", server.URL)))
		})

		By("user fills new gist form", func() {
			Expect(page.Find("input[name=title]").Fill("Sample Gist")).To(Succeed())

			Expect(page.RunScript(
				"$(elementID).val('Some content')",
				map[string]interface{}{"elementID": "textarea[name=content]"},
				nil,
			)).To(Succeed())

			Expect(page.Find("input[type=submit]").Submit()).To(Succeed())
		})

		By("user is redirected to their gists", func() {
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})
	})

	It("updating gists", func() {
		By("user vists gists page", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})

		By("user clicks on a gist", func() {
			Expect(page.FindByLink("Sample Gist").Click()).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/gists/1", server.URL)))
		})

		By("user clicks edit", func() {
			Expect(page.FindByLink("Edit").Click()).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/gists/1/edit", server.URL)))
		})

		By("user updates gist", func() {
			Expect(page.Find("input[name=title]").Fill("Sample Gist Edited")).To(Succeed())
			Expect(page.Find("input[type=submit]").Submit()).To(Succeed())
		})

		By("the gist is saved", func() {
			Expect(page.HTML()).To(ContainSubstring("Sample Gist Edited"))
		})
	})
})
