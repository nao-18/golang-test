package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGolangTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GolangTest Suite")
}
