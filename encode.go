package fakegaekey

import (
	"encoding/base64"
	"strings"

	"github.com/golang/protobuf/proto"
	pb "google.golang.org/appengine/internal/datastore"
)

func keyToProto(defaultAppID string, k *Key) *pb.Reference {
	appID := k.appID
	if appID == "" {
		appID = defaultAppID
	}
	n := 0
	for i := k; i != nil; i = i.parent {
		n++
	}
	e := make([]*pb.Path_Element, n)
	for i := k; i != nil; i = i.parent {
		n--
		e[n] = &pb.Path_Element{
			Type: &i.kind,
		}
		// At most one of {Name,Id} should be set.
		// Neither will be set for incomplete keys.
		if i.stringID != "" {
			e[n].Name = &i.stringID
		} else if i.intID != 0 {
			e[n].Id = &i.intID
		}
	}
	var namespace *string
	if k.namespace != "" {
		namespace = proto.String(k.namespace)
	}
	return &pb.Reference{
		App:       proto.String(appID),
		NameSpace: namespace,
		Path: &pb.Path{
			Element: e,
		},
	}
}

func (k *Key) Encode() string {
	ref := keyToProto("", k)

	b, err := proto.Marshal(ref)
	if err != nil {
		panic(err)
	}

	// Trailing padding is stripped.
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}
