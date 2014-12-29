package gorest

import (
	"testing"
)

func testStress(t *testing.T) {
	loop1(t)
	loop2(t)
}

func loop1(t *testing.T) {
	//loop1 EndPoint `method:"DELETE" 	path:"/loop1/{Bool:bool}/mix1/{Int:int}"`
	//*******************************

	ch := make(chan int)
	fn := func(ch chan int, count int) {
		rb, _ := NewRequestBuilder(RootPath + "stress-service/loop1/true/mix1/5" + xrefStr)
		rb.AddCookie(cook)
		res, _ := rb.Delete()
		AssertEqual(res.StatusCode, 200, "Delete ResponseCode", t)
		ch <- count
	}

	runs := 200
	total := 0

	for i := 0; i < runs; i++ {
		go fn(ch, i)
		total = total + i
	}

	for i := 0; i < runs; i++ {
		total = total - <-ch
	}

	AssertEqual(total, 0, "testStress --> loop1", t)

}

func loop2(t *testing.T) {
	//loop2() EndPoint `method:"OPTIONS" path:"/loop2/{Bool:bool}/mix1/{Int:int}"`
	//*******************************

	ch := make(chan int)
	fn := func(ch chan int, count int) {
		strArr := make([]string, 0)
		rb, _ := NewRequestBuilder(RootPath + "stress-service/loop2/true/mix1/5" + xrefStr)
		rb.AddCookie(cook)
		res, _ := rb.Options(&strArr)
		AssertEqual(res.StatusCode, 200, "testStress --> loop2: Options ResponseCode", t)
		AssertEqual(len(strArr), 3, "testStress --> loop2: Options - slice length", t)
		if len(strArr) == 3 {
			AssertEqual(strArr[0], GET, "testStress --> loop2: Options", t)
			AssertEqual(strArr[1], HEAD, "testStress --> loop2: Options", t)
			AssertEqual(strArr[2], POST, "testStress --> loop2: Options", t)
		}
		ch <- count
	}

	runs := 200
	total := 0

	for i := 0; i < runs; i++ {
		go fn(ch, i)
		total = total + i
	}

	for i := 0; i < runs; i++ {
		total = total - <-ch
	}

	AssertEqual(total, 0, "testStress --> loop2", t)

}
