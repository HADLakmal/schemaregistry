/**
 * Copyright 2020 TryFix Engineering.
 * All rights reserved.
 * Authors:
 *    Gayan Yapa (gmbyapa@gmail.com)
 */

package main

import (
	"time"

	"github.com/HADLakmal/schemaregistry/v2"
	"github.com/tryfix/log"
)

func main() {

	// init a new schema registry instance and connect
	registry, err := schemaregistry.NewRegistry(`localhost:8089/`,
		schemaregistry.WithBackgroundSync([]string{`localhost:9092`}, `__schemas`))
	if err != nil {
		log.Fatal(err)
	}

	if err := registry.Register(`com.org.events.test.TestTwo`, 1, func(decoder schemaregistry.AvroDecoder, data []byte) (v interface{}, err error) {
		return nil, nil
	}); err != nil {
		log.Fatal(err)
	}

	if err = registry.Sync(); err != nil {
		log.Fatal(err)
	}

	log.Info(`your event is successfully registered`)

	time.Sleep(10 * time.Minute)
}
