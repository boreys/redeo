package info

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Registry", func() {
	var subject *Registry

	BeforeEach(func() {
		subject = New()
		subject.Section("Server").Register("version", StringValue("1.0.1"))
		subject.Section("Server").Register("date", StringValue("2014-11-11"))
		subject.Section("Clients").Register("count", StringValue("17"))
		subject.Section("Clients").Register("total", StringValue("123456"))
		subject.Section("Empty")
	})

	It("should generate info strings", func() {
		Expect(New().String()).To(Equal(""))
		Expect(subject.String()).To(Equal("# Server\nversion:1.0.1\ndate:2014-11-11\n\n# Clients\ncount:17\ntotal:123456\n"))
	})

	It("should clear", func() {
		subject.Section("Clients").Clear()
		Expect(subject.sections[1].kvs).To(BeEmpty())
		subject.Clear()
		Expect(subject.sections).To(BeEmpty())
	})

})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redeo/info")
}
