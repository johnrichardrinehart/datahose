package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	schema "github.com/johnrichardrinehart/datahose/pkg/schema"
	sinksPkg "github.com/johnrichardrinehart/datahose/pkg/sinks"
)

var (
	sinksString      = flag.String("sinks", "stdout,data.out", "comma-separated list of data destinations (stdout, <filename>.out, <filename>.unix, <IP>:<PORT>")
	schemaFileString = flag.String("schema", "data.schema", "a path to a valid schema file")
)

func main() {
	flag.Parse()
	schemaFileBytes, err := ioutil.ReadFile(*schemaFileString)
	if err != nil {
		log.Fatal(err)
	}
	userSchema, err := schema.Parse(string(schemaFileBytes))
	if err != nil {
		log.Fatal(err)
	}

	var sinks []sinksPkg.Sink
	for _, sinkString := range strings.Split(*sinksString, ",") {
		sink, err := sinksPkg.StringToSink(sinkString)
		if err != nil {
			log.Fatal(err)
		}
		sinks = append(sinks, sink)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			for _, w := range sinks {
				if file, ok := w.(*os.File); ok {
					file.Close()
					os.Remove(file.Name())
				} else {
					w.Close()
				}
			}
			os.Exit(1)
		}
	}()

	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 500)
		for _, sink := range sinks {
			bytes, err := json.Marshal(userSchema.Generate())
			if err != nil {
				log.Fatal(err)
			}
			sink.Write(bytes)
		}
	}
}
