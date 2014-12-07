package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBadger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Badger Suite")
}

var _ = BeforeSuite(func() {
	var err error
	Expect(err).NotTo(HaveOccurred())
})
