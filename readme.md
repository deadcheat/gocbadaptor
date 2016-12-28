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

	couch = gocbadaptor.NewDefaultCouchAdaptor().Open(connection, bucketname, password, expiry)

	var val []byte

	key := "testkey"

	val = []byte("{\"name\":\"test1\"}")

	// execute N1qlQuery
	r, _ := couch.N1qlQuery("delete from "+bucketname, nil)

	// insert
	cas, ok := couch.Insert(key, val)

	fmt.Println(cas, ok)

	// get
	cas, data, ok := couch.Get(key)

	fmt.Println(cas, string(data), ok)

	key2 := "testkey2"
	val = []byte("{\"name\":\"test2\"}")

	// upsert
	cas, ok = couch.Upsert(key2, val)

	pp.Println(cas, ok)

	cas, data, ok = couch.Get(key)

	fmt.Println(cas, string(data), ok)

	couch.Bucket().Close()

	couch = gocbadaptor.NewDefaultCouchAdaptor().OpenWithConfig(couch.Env())

	// execute N1qlQuery
	r, _ = couch.N1qlQuery("select META("+bucketname+").id, name from `"+bucketname+"` WHERE name like 'test%'", nil)

	var row interface{}

	for r.Next(&row) {
		pp.Println(row)
	}
	r.Close()
	pp.Println(r.Metrics().ResultSize)

}

```