package fakegaekey

type Key struct {
	kind      string
	stringID  string
	intID     int64
	parent    *Key
	appID     string
	namespace string
}

// Kind returns the key's kind (also known as entity type).
func (k *Key) Kind() string {
	return k.kind
}

// StringID returns the key's string ID (also known as an entity name or key
// name), which may be "".
func (k *Key) StringID() string {
	return k.stringID
}

// IntID returns the key's integer ID, which may be 0.
func (k *Key) IntID() int64 {
	return k.intID
}

// Parent returns the key's parent key, which may be nil.
func (k *Key) Parent() *Key {
	return k.parent
}

// AppID returns the key's application ID.
func (k *Key) AppID() string {
	return k.appID
}

// Namespace returns the key's namespace.
func (k *Key) Namespace() string {
	return k.namespace
}

// Incomplete returns whether the key does not refer to a stored entity.
// In particular, whether the key has a zero StringID and a zero IntID.
func (k *Key) Incomplete() bool {
	return k.stringID == "" && k.intID == 0
}

// valid returns whether the key is valid.
func (k *Key) valid() bool {
	if k == nil {
		return false
	}
	for ; k != nil; k = k.parent {
		if k.kind == "" || k.appID == "" {
			return false
		}
		if k.stringID != "" && k.intID != 0 {
			return false
		}
		if k.parent != nil {
			if k.parent.Incomplete() {
				return false
			}
			if k.parent.appID != k.appID || k.parent.namespace != k.namespace {
				return false
			}
		}
	}
	return true
}

type Context struct {
	Namespace string
	AppID     string
}

func NewKey(context Context, kind, stringID string, intID int64, parent *Key) *Key {
	// If there's a parent key, use its namespace.
	// Otherwise, do a fake RPC to try to get a namespace if c is a namespacedContext (or wraps one).
	var namespace string
	if parent != nil {
		namespace = parent.namespace
	} else {
		namespace = context.Namespace
	}

	return &Key{
		kind:      kind,
		stringID:  stringID,
		intID:     intID,
		parent:    parent,
		appID:     "s~" + context.AppID,
		namespace: namespace,
	}
}
