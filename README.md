fakegaekey [![GoDoc](https://godoc.org/github.com/SeanDolphin/fakegaekey?status.svg)](https://godoc.org/github.com/SeanDolphin/fakegaekey) [![wercker status](https://app.wercker.com/status/e4b5c15da5315d343d1a2ebfa37ca175/s "wercker status")](https://app.wercker.com/project/bykey/e4b5c15da5315d343d1a2ebfa37ca175)
==========

Creates a Fake Google App Engine datastore.Key.  This can be used for outside scripts that need to be able to encode keys and store them.


## Usage

~~~ go
// main.go
package main

import (
	"github.com/SeanDolphin/faekgaekey"
)


func main() {
  	context := fakegaekey.Context{AppID: "myappid", Namespace: "test"}
  	createdKey := fakegaekey.NewKey(context, "doc", "someID", 0, nil)

  	// Do something with your key


  	decodedKey, err := fakegaekey.DecodeKey("randomencodekeystring")
  	//Do something with your.
}

~~~