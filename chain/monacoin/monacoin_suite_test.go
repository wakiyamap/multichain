package monacoin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMonacoin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monacoin Suite")
}
