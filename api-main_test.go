package gorest

import (
	"io/ioutil"
	"log"
	"net/http"
	//"net/http/httptest"
	"runtime"
	"testing"
)

var MUX_ROOT = "/home/now/the/future/"
var RootPath = "http://localhost:8787" + MUX_ROOT
var globalTestScope *testing.T //This is used to do Asserts inside the service implementations

func TestInit(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	globalTestScope = t
	log.Println("Starting tests")
	log.SetOutput(ioutil.Discard) //Toggle comment in-out to see log output

	RegisterRealmAuthorizer("testing", TestingAuthorizer)
	RegisterServiceOnPath(MUX_ROOT, new(TypesService))
	RegisterServiceOnPath(MUX_ROOT, new(PathsService))
	RegisterServiceOnPath(MUX_ROOT, new(StressService))

	http.Handle(MUX_ROOT, Handle())

	//http.HandleFunc(MUX_ROOT, HandleFunc)
	//httptest.NewServer(Handle())
	//server.Start()

	go http.ListenAndServe(":8787", nil)
	//go ServeStandAlone(8787)

}

func TestDataTransmition(t *testing.T) {
	testTypes(t)

}
func TestPaths(t *testing.T) {
	testPaths(t)
}

func TestStress(t *testing.T) {
	testStress(t)
}

func TestServiceMeta(t *testing.T) {
	if meta, found := restManager.serviceTypes["code.google.com/p/gorest/TypesService"]; !found {
		t.Error("Service Not registered correctly")
	} else {
		AssertEqual(meta.consumesMime, "application/json", "Service consumesMime", t)
		AssertEqual(meta.producesMime, "application/json", "Service producesMime", t)
		AssertEqual(meta.root, MUX_ROOT+"types-service/", "Service root", t)

	}

}
