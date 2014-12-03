package fakegaekey_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/SeanDolphin/fakegaekey"
)

var _ = ginkgo.Describe("Key", func() {
	ginkgo.It("should properly encode keys", func() {

		for _, test := range tests {
			context := fakegaekey.Context{AppID: test.AppID}
			if test.Pass && test.NameSpace == "" {
				var parent *fakegaekey.Key
				if test.Parent != "" {
					var err error
					parent, err = fakegaekey.DecodeKey(test.Parent)
					gomega.Expect(err).To(gomega.BeNil())
				}
				encodedKey := fakegaekey.NewKey(context, "doc", test.ID, 0, parent).Encode()
				gomega.Expect(encodedKey).To(gomega.Equal(test.Key))
			}
		}
	})

	ginkgo.It("should properly encode namespaces", func() {
		for _, test := range tests {
			if test.Pass && test.NameSpace != "" {
				context := fakegaekey.Context{AppID: "kpcompassapi", Namespace: test.NameSpace}
				var parent *fakegaekey.Key
				if test.Parent != "" {
					var err error
					parent, err = fakegaekey.DecodeKey(test.Parent)
					gomega.Expect(err).To(gomega.BeNil())
				}
				key := fakegaekey.NewKey(context, "doc", test.ID, 0, parent)
				gomega.Expect(key.Encode()).To(gomega.Equal(test.Key))
				gomega.Expect(key.Namespace()).To(gomega.Equal(test.NameSpace))
			}
		}
	})

	ginkgo.It("should properly encode keys without a appID", func() {
		for _, test := range tests {
			if test.Pass && test.AppID == "" {
				context := fakegaekey.Context{Namespace: test.NameSpace}
				var parent *fakegaekey.Key
				if test.Parent != "" {
					var err error
					parent, err = fakegaekey.DecodeKey(test.Parent)
					gomega.Expect(err).To(gomega.BeNil())
				}
				key := fakegaekey.NewKey(context, "doc", test.ID, 0, parent)
				gomega.Expect(key.Encode()).To(gomega.Equal(test.Key))
				gomega.Expect(key.AppID()).To(gomega.Equal("s~"))
			}
		}
	})
})
