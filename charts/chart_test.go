package charts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chart", func() {
	Describe("Add", func() {
		Context("when set exists", func() {
			Context("when state exists", func() {
				It("returns false", func() {
					Expect(true).To(BeTrue())
				})
			})
			Context("when state doesn't exist", func() {
				It("adds state and returns true", func() {

				})
			})
		})
		Context("when set doesn't exist", func() {
			It("creates set", func() {

			})
		})
	})
})
