package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/opctl/util/containerprovider"
	"github.com/opspec-io/opctl/util/pubsub"
)

var _ = Context("core", func() {
	Context("New", func() {
		It("should return Core", func() {
			/* arrange/act/assert */
			Expect(
				New(
					new(pubsub.Fake),
					new(containerprovider.Fake),
				),
			).Should(Not(BeNil()))
		})
	})
})
