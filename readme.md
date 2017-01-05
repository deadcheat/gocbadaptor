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

	couch = gocbadaptor.NewDefaultCouchAdaptor()
	if err := couch.Open(connection, bucketname, password, expiry); err != nil {
		log.Println("err")
		return
	}

	var val []byte

	key := "testkey"

	val = []byte("{\"name\":\"test1\"}")

	// execute N1qlQuery
	r, _ := couch.N1qlQuery("delete from "+bucketname, nil)

	// insert
	cas, ok := couch.Insert(key, val)

	fmt.Println("insert", cas, ok)

	// get
	cas, data, ok := couch.Get(key)

	fmt.Println("get", cas, string(data), ok)

	// key2 := "testkey2"
	val = []byte("{\"name\":\"test2\"}")

	// upsert
	cas, ok = couch.Upsert(key, val)

	fmt.Println("upsert", cas, ok)

	cas, data, ok = couch.Get(key)

	fmt.Println("get", cas, string(data), ok)

	couch.Bucket().Close()

	if err := couch.OpenWithConfig(couch.Env()); err != nil {
		log.Println("err")
		return
	}

	// execute N1qlQuery
	r, err := couch.N1qlQuery("select META("+couch.Env().BucketName+").id, name from `"+couch.Env().BucketName+"` WHERE name like 'test%'", nil)
	if err != nil {
		log.Println("err")
		return
	}

	var row interface{}

	for r.Next(&row) {
		fmt.Println(row)
	}
	r.Close()

}

```