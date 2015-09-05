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

	It("signing up", func() {
		By("user visits signup route", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/signup", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/signup", server.URL)))
		})

		By("user fills out signout form", func() {
			Expect(page.Find("input[name=email]").Fill("john@example.com")).To(Succeed())
			Expect(page.Find("input[name=username]").Fill("john")).To(Succeed())
			Expect(page.Find("input[name=password]").Fill("password")).To(Succeed())
			Expect(page.Find("input[name=password_confirmation]").Fill("password")).To(Succeed())
			Expect(page.Find("input[type=submit]").Submit()).To(Succeed())
		})

		By("user is redirected to their gists", func() {
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})
	})

	It("logging in", func() {
		By("user visits login route", func() {
			Expect(page.Navigate(fmt.Sprintf("%s/login", server.URL))).To(Succeed())
			Expect(page).To(HaveURL(fmt.Sprintf("%s/login", server.URL)))
		})

		By("user fills out login form", func() {
			Expect(page.Find("input[name=email]").Fill("john@example.com")).To(Succeed())
			Expect(page.Find("input[name=password]").Fill("password")).To(Succeed())
			Expect(page.Find("input[type=submit]").Submit()).To(Succeed())
		})

		By("user is redirected to their gists", func() {
			Expect(page).To(HaveURL(fmt.Sprintf("%s/", server.URL)))
		})
	})
})
