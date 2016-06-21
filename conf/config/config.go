package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

type Config struct {
	Root map[string]interface{}
}

func (cfg *Config) LoadJSON(filename string) {
	content, err := ioutil.ReadFile(filename)
	var out map[string]interface{}
	err = json.Unmarshal([]byte(content), &out)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Root = out
}

func (cfg *Config) Get(keys string) interface{} {
	var v interface{}
	for _, k := range strings.Split(keys, ".") {
		if v != nil {
			a := reflect.ValueOf(v)
			i := a.Interface()
			b := i.(map[string]interface{})
			v = b[k]
		} else {
			v = cfg.Root[k]
		}
	}
	return v
}
