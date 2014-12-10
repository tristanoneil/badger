package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tristanoneil/badger/models"

	"testing"
)

func TestBadger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Badger Suite")
}

var _ = BeforeSuite(func() {
	models.ResetDB()
	models.MigrateDB()

	var err error
	Expect(err).NotTo(HaveOccurred())
})
