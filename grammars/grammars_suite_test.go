package grammars_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGrammars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grammars Suite")
}
