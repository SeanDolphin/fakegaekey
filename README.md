fakegaekey [![GoDoc](https://godoc.org/github.com/SeanDolphin/fakegaekey?status.svg)](https://godoc.org/github.com/SeanDolphin/fakegaekey)
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