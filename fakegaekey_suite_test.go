package fakegaekey_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"testing"
)

type test struct {
	ID        string
	Key       string
	NameSpace string
	AppID     string
	Parent    string
	Pass      bool
}

var tests []test

func TestFakeGAEKey(t *testing.T) {
	tests = []test{
		test{
			ID:   "1",
			Key:  "agJzfnIKCxIDZG9jIgExDA",
			Pass: true,
		},
		test{
			ID:    "1",
			Key:   "ag5zfmtwY29tcGFzc2FwaXIKCxIDZG9jIgExDA",
			AppID: "kpcompassapi",
			Pass:  true,
		},
		test{
			ID:        "1",
			Key:       "ag5zfmtwY29tcGFzc2FwaXIKCxIDZG9jIgExDKIBBHRlc3Q",
			AppID:     "kpcompassapi",
			NameSpace: "test",
			Pass:      true,
		},
		test{
			ID:     "items-1",
			Key:    "agJzfnIaCxIDZG9jIgExDAsSA2RvYyIHaXRlbXMtMQw",
			Parent: "agJzfnIKCxIDZG9jIgExDA",
			Pass:   true,
		},
		test{
			ID:     "items-1",
			Key:    "ag5zfmtwY29tcGFzc2FwaXIaCxIDZG9jIgExDAsSA2RvYyIHaXRlbXMtMQw",
			Parent: "ag5zfmtwY29tcGFzc2FwaXIKCxIDZG9jIgExDA",
			AppID:  "kpcompassapi",
			Pass:   true,
		},
		test{
			ID:        "items-1",
			Key:       "ag5zfmtwY29tcGFzc2FwaXIaCxIDZG9jIgExDAsSA2RvYyIHaXRlbXMtMQyiAQR0ZXN0",
			Parent:    "ag5zfmtwY29tcGFzc2FwaXIKCxIDZG9jIgExDKIBBHRlc3Q",
			AppID:     "kpcompassapi",
			NameSpace: "test",
			Pass:      true,
		},
		test{
			ID:   "blah",
			Key:  "blah",
			Pass: false,
		},
		test{
			ID:   "blah",
			Key:  "*",
			Pass: false,
		},
	}

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Fake GAE Key")

}
