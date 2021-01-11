package main

import (
	"flag"
	"fmt"

	"github.com/StringNote/strnote/firebase"
	"github.com/xorvercom/util/pkg/json"
)

func main() {
	var (
		cred = flag.String("cred", "admin.json", "API creditional")
		col  = flag.String("col", "", "Collection")
		data = flag.String("dat", "", "Data JSON")
		imp  = flag.Bool("imp", false, "is Export")
	)
	flag.Parse()
	if *cred == "" || *col == "" {
		flag.PrintDefaults()
		return
	}
	if *data == "" {
		data = col
	}
	fmt.Printf("cred:\"%s\", col:\"%s\", imp:%v, CommandLine:%v\n", *cred, *col, *imp, flag.CommandLine)
	firebase.SetOption(*cred)
	cl := firebase.NewCollection(*col)
	if *imp == false {
		elem := json.NewElemObject()
		for _, key := range cl.Keys() {
			if val, err := cl.Get(key); err == nil {
				elem.Put(key, json.NewElemString(val))
			}
		}
		json.SaveToJSONFile(*data+".json", elem, true)
	} else {
		if elem, err := json.LoadFromJSONFile(*data + ".json"); err == nil {
			if obj, ok := elem.AsObject(); ok {
				for _, key := range obj.Keys() {
					if str, ok := obj.Child(key).AsString(); ok {
						val := str.Text()
						fmt.Printf("    %s <- %s\n", key, val)
						cl.Set(key, val)
					}
				}
			}
		}
	}
}
