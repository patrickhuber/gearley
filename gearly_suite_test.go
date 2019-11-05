package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGearly(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gearly Suite")
}
