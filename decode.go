package fakegaekey

import (
	"encoding/base64"
	"errors"
	"strings"

	pb "code.google.com/p/appengine-go/appengine_internal/datastore"
	"code.google.com/p/goprotobuf/proto"
)

var ErrInvalidKey = errors.New("datastore: invalid key")

func protoToKey(r *pb.Reference) (k *Key, err error) {
	appID := r.GetApp()
	namespace := r.GetNameSpace()
	for _, e := range r.Path.Element {
		k = &Key{
			kind:      e.GetType(),
			stringID:  e.GetName(),
			intID:     e.GetId(),
			parent:    k,
			appID:     appID,
			namespace: namespace,
		}
		if !k.valid() {
			return nil, ErrInvalidKey
		}
	}
	return
}

func DecodeKey(encoded string) (*Key, error) {
	// Re-add padding.
	if m := len(encoded) % 4; m != 0 {
		encoded += strings.Repeat("=", 4-m)
	}

	b, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	ref := new(pb.Reference)
	if err := proto.Unmarshal(b, ref); err != nil {
		return nil, err
	}

	return protoToKey(ref)
}
