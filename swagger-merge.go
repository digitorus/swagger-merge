package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var base = flag.String("base", "", "Base swagger file name")
	var path = flag.String("path", ".", "Path to recursivly scan for swagger files")
	var match = flag.String("match", "*.swagger.json", "Format of swagger filenames")

	var version = flag.String("version", "", "Version of this API specification")
	var title = flag.String("title", "", "Title for this API")

	flag.Parse()

	var err error
	var so *SwaggerObject

	// Base properties are taken from this file
	if *base != "" {
		so, err = Read(*base)
		if err != nil {
			log.Fatal(err)
		}
	}

	if so == nil {
		so = &SwaggerObject{}
	}
	if so.Info == nil {
		so.Info = &InfoObject{}
	}

	if *version != "" {
		so.Info.Version = *version
	}

	if *title != "" {
		so.Info.Title = *title
	}

	err = filepath.Walk(*path, func(filename string, info os.FileInfo, err error) error {
		ok, err := filepath.Match(*match, info.Name())
		if err != nil {
			log.Println(err)
			return err
		}
		if ok {
			if err = so.Merge(filename); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if so.Info.Version == "" {
		so.Info.Version = time.Now().Format("20060102150405")
	}

	err = so.Output()
	if err != nil {
		log.Fatal(err)
	}
}

// Output swagger file to screen
func (so *SwaggerObject) Output() error {
	data, err := json.MarshalIndent(so, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

// Merge swagger file into current file
func (so *SwaggerObject) Merge(filename string) error {
	swf, err := Read(filename)
	if err != nil {
		return err
	}

	if so.Swagger == "" {
		so.Swagger = swf.Swagger
	}
	so.Schemes = mergeSlice(so.Schemes, swf.Schemes)
	so.Consumes = mergeSlice(so.Consumes, swf.Consumes)
	so.Produces = mergeSlice(so.Consumes, swf.Produces)

	if so.Paths == nil {
		so.Paths = swf.Paths
	} else {
		for p, d := range swf.Paths {
			if _, ok := so.Paths[p]; !ok {
				so.Paths[p] = d
			} else {
				log.Printf("Path %s already exists", p)
			}
		}
	}

	if so.Definitions == nil {
		so.Definitions = swf.Paths
	} else {
		for p, d := range swf.Definitions {
			if _, ok := so.Definitions[p]; !ok {
				so.Definitions[p] = d
			} else {
				log.Printf("Definition %s already exists", p)
			}

		}
	}

	if so.StreamDefinitions == nil {
		so.StreamDefinitions = swf.Paths
	} else {
		for p, d := range swf.StreamDefinitions {
			if _, ok := so.StreamDefinitions[p]; !ok {
				so.StreamDefinitions[p] = d
			} else {
				log.Printf("StreamDefinition %s already exists", p)
			}

		}
	}

	if so.SecurityDefinitions == nil {
		so.SecurityDefinitions = swf.SecurityDefinitions
	}

	if so.Security == nil {
		so.Security = swf.Security
	}

	if so.ExternalDocs == nil {
		so.ExternalDocs = swf.ExternalDocs
	}

	return nil
}

// Read swagger file into struct
func Read(filename string) (*SwaggerObject, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var swf *SwaggerObject
	err = json.Unmarshal(b, &swf)
	if err != nil {
		return nil, err
	}
	return swf, nil
}

func mergeSlice(target []string, source []string) []string {
    for _, s := range source {
		var present bool
		for _, t := range target {
        	if s == t {
				present = true
			}
		}
		if !present {
			target = append(target, s)
		}
    }
	
	return target
}