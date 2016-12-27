# gocbadaptor

## install

```
go get github.com/deadcheat/gocbadaptor
```

or using glide

```
glide get github.com/deadcheat/gocbadaptor
```

## how to use

see a sample code below
```
package main

import (
	"fmt"

	"github.com/deadcheat/gocbadaptor"
)

func main() {
	connection := "couchbase://dev-couch"
	bucketname := "couch_bucket"
	password := ""
	var expiry uint32 = 100

	var couch gocbadaptor.CouchBaseAdaptor

	couch = gocbadaptor.NewDefaultCouchAdaptor().Open(&connection, &bucketname, &password, expiry)

	var val []byte

	key := "testkey"

	val = []byte("test1")

	// execute N1qlQuery
	r, err := couch.N1qlQuery("delete from "+bucketname, nil)

	// insert
	cas, ok := couch.Insert(key, val)

	fmt.Println(cas, ok)

	// get
	cas, data, ok := couch.Get(key)

	fmt.Println(cas, string(data), ok)

	val = []byte("test2")

	// upsert
	cas, ok = couch.Upsert(key, val)

	fmt.Println(cas, ok)

	cas, data, ok = couch.Get(key)

	fmt.Println(cas, string(data), ok)

	// execute N1qlQuery
	r, err = couch.N1qlQuery("select * from "+bucketname, nil)

	var res []byte
	r.Next(&res)
	fmt.Println(res, err, r.Metrics().ResultCount)

}

```