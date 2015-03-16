package mtproto

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Param struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"-"`
}

type Constructor struct {
	Id        int32   `json:"id,string"`
	Type      string  `json:"type"`
	Predicate string  `json:"predicate"`
	Params    []Param `json:"params"`
}

type Method struct {
	Id     int32   `json:"id,string"`
	Type   string  `json:"type"`
	Method string  `json:"method"`
	Params []Param `json:"params"`
}

type Schema struct {
	Methods          map[string]Method
	ConstructorsId   map[int32]*Constructor
	ConstructorsType map[string]*Constructor
}

type TempSchema struct {
	Methods      []Method      `json:"methods"`
	Constructors []Constructor `json:"constructors"`
}

var TLSchema *Schema

func LoadSchema(filename string) error {
	temp := TempSchema{}
	schema := Schema{
		Methods:          map[string]Method{},
		ConstructorsId:   map[int32]*Constructor{},
		ConstructorsType: map[string]*Constructor{},
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := json.NewDecoder(f).Decode(&temp); err != nil {
		return err
	}

	for _, c := range temp.Constructors {
		schema.Constructors[c.Id] = &c
		schema.Constructors[c.Type] = &c
	}

	for _, m := range temp.Methods {
		schema.Methods[m.Method] = m
	}

	TLSchema = &schema
	return nil
}
