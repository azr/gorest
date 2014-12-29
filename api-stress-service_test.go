package gorest

import (
	"time"
)

type StressService struct {
	RestService `root:"/stress-service/" consumes:"application/json" produces:"application/json" realm:"testing"`

	//Test Mixed paths with same length
	loop1 EndPoint `method:"DELETE" path:"/loop1/{Bool:bool}/mix1/{Int:int}"`

	//Now check same path for different methods
	loop2 EndPoint `method:"OPTIONS" path:"/loop2/{Bool:bool}/mix1/{Int:int}"`
}

func (serv StressService) Loop1(Bool bool, Int int) {
	<-time.After(2 * time.Second)
}

func (serv StressService) Loop2(Bool bool, Int int) {
	rb := serv.ResponseBuilder()
	rb.Allow("GET")
	rb.Allow("HEAD").Allow("POST")
	<-time.After(2 * time.Second)
}
