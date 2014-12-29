package gorest

import (
	"testing"
)

func testPaths(t *testing.T) {
	deleteMixed1(t)
	deleteMixed2(t)
	optionsMixed(t)
	getMixed(t)
}

func deleteMixed1(t *testing.T) {
	//deleteMixed1 EndPoint `method:"DELETE" 	path:"/bool/{Bool:bool}/mix1/{Int:int}"`
	//*******************************

	rb, _ := NewRequestBuilder(RootPath + "paths-service/bool/true/mix1/5" + xrefStr)
	rb.AddCookie(cook)
	res, _ := rb.Delete()
	AssertEqual(res.StatusCode, 200, "Delete ResponseCode", t)
}
func deleteMixed2(t *testing.T) {
	//deleteMixed2 EndPoint `method:"DELETE" 	path:"/bool/{Bool:bool}/mix2/{Int:int}"`
	//*******************************

	rb, _ := NewRequestBuilder(RootPath + "paths-service/bool/true/mix2/5" + xrefStr)
	rb.AddCookie(cook)
	res, _ := rb.Delete()
	AssertEqual(res.StatusCode, 200, "Delete ResponseCode", t)
}

func optionsMixed(t *testing.T) {
	//optionsMixed EndPoint `method:"OPTIONS" path:"/bool/{Bool:bool}/mix1/{Int:int}"`
	//*******************************

	strArr := make([]string, 0)
	rb, _ := NewRequestBuilder(RootPath + "paths-service/bool/true/mix1/5" + xrefStr)
	rb.AddCookie(cook)
	res, _ := rb.Options(&strArr)
	AssertEqual(res.StatusCode, 200, "Options ResponseCode", t)
	AssertEqual(len(strArr), 3, "Options - slice length", t)
	if len(strArr) == 3 {
		AssertEqual(strArr[0], GET, "Options", t)
		AssertEqual(strArr[1], HEAD, "Options", t)
		AssertEqual(strArr[2], POST, "Options", t)
	}
}

func getMixed(t *testing.T) {
	//getMixed     EndPoint `method:"GET" 	path:"/bool/{Bool:bool}/mix1/{Int:int}"`
	//*******************************

	rb, _ := NewRequestBuilder(RootPath + "paths-service/bool/true/mix1/5" + xrefStr)
	rb.AddCookie(cook)
	//GET string

	res, _ := rb.Get(&str, 200)
	AssertEqual(res.StatusCode, 200, "Get string ResponseCode", t)
	AssertEqual(str, "Hello", "Get getMixed", t)

}
