package gorest

import (
	"encoding/json"
	"encoding/xml"
)

//A Marshaller represents the two functions used to marshal/unmarshal interfaces back and forth.
type Marshaller struct {
	Marshal   func(v interface{}) ([]byte, error)
	Unmarshal func(data []byte, v interface{}) error
}

var marshallers map[string]*Marshaller

//Register a Marshaller. These registered Marshallers are shared by the client or servers side usage of gorest.
func RegisterMarshaller(mime string, m *Marshaller) {
	if marshallers == nil {
		marshallers = make(map[string]*Marshaller, 0)
	}
	if _, found := marshallers[mime]; !found {
		marshallers[mime] = m
	}
}

//Get an already registered Marshaller
func GetMarshallerByMime(mime string) (m *Marshaller) {
	if marshallers == nil {
		marshallers = make(map[string]*Marshaller, 0)
	}
	m, _ = marshallers[mime]
	return
}

//Predefined Marshallers

//JSON: This makes the JSON Marshaller. The Marshaller uses pkg: json
func NewJSONMarshaller() *Marshaller {
	m := Marshaller{jsonMarshal, jsonUnMarshal}
	return &m
}
func jsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func jsonUnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

//XML
func NewXMLMarshaller() *Marshaller {
	m := Marshaller{xmlMarshal, xmlUnMarshal}
	return &m
}
func xmlMarshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}
func xmlUnMarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
