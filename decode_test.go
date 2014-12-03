package fakegaekey_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"bitbucket.org/kpcompass/fakegaekey"
)

var _ = ginkgo.Describe("Decoding", func() {
	ginkgo.It("should propery decode keys", func() {
		for _, test := range tests {
			if test.Pass {
				key, err := fakegaekey.DecodeKey(test.Key)

				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(key.Kind()).To(gomega.Equal("doc"))
				gomega.Expect(key.StringID()).To(gomega.Equal(test.ID))
				gomega.Expect(key.IntID()).To(gomega.BeEquivalentTo(0))
				if test.Parent != "" {
					gomega.Expect(key.Parent().Encode()).To(gomega.Equal(test.Parent))
				}
			}
		}
	})

	ginkgo.It("should fail if the key is invalid", func() {
		for _, test := range tests {
			if !test.Pass {
				_, err := fakegaekey.DecodeKey(test.Key)
				gomega.Expect(err).ToNot(gomega.BeNil())
			}
		}
	})
})
