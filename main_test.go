package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"github.com/tristanoneil/badger/models"

	"testing"
)

func TestBadger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Badger Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	models.ResetDB()
	models.MigrateDB()

	var err error

	agoutiDriver = agouti.PhantomJS()

	Expect(err).NotTo(HaveOccurred())
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	agoutiDriver.Stop()
})
