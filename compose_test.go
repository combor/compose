package compose

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compose", func() {
	const (
		expectedToken = "expected"
	)
	BeforeEach(func() {
		os.Setenv("COMPOSE_TOKEN", expectedToken)
	})
	Context("when the token is not empty", func() {
		It("should return a token from env variable", func() {
			token, err := GetApiToken()
			Expect(err).ToNot(HaveOccurred())
			Expect(token).To(Equal(expectedToken))

		})
	})
	Context("when the token is empty", func() {
		BeforeEach(func() {
			os.Unsetenv("COMPOSE_TOKEN")
		})
		It("should return an error ", func() {
			token, err := GetApiToken()
			Expect(err).To(HaveOccurred())
			Expect(token).To(Equal(""))
			Expect(err.Error()).To(Equal("Empty token. Please export $COMPOSE_TOKEN"))
		})
	})
})
