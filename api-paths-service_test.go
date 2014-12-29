package gorest

type PathsService struct {
	RestService `root:"/paths-service/" consumes:"application/json" produces:"application/json" realm:"testing"`

	//Test Mixed paths with same length
	deleteMixed1 EndPoint `method:"DELETE" path:"/bool/{Bool:bool}/mix1/{Int:int}"`
	deleteMixed2 EndPoint `method:"DELETE" path:"/bool/{Bool:bool}/mix2/{Int:int}"`
	//Now check same path for different methods
	optionsMixed EndPoint `method:"OPTIONS" path:"/bool/{Bool:bool}/mix1/{Int:int}"`
	getMixed     EndPoint `method:"GET" path:"/bool/{Bool:bool}/mix1/{Int:int}" output:"string"`
	//getMixed2    EndPoint `method:"GET" path:"/bool/{Bool:bool}/mix1/{Int:int}" output:"string"`
}

func (serv PathsService) DeleteMixed1(Bool bool, Int int) {
	//Will return default response code of 200
}
func (serv PathsService) DeleteMixed2(Bool bool, Int int) {
	//Will return default response code of 200
}
func (serv PathsService) OptionsMixed(Bool bool, Int int) {
	rb := serv.ResponseBuilder()
	rb.Allow("GET")
	rb.Allow("HEAD").Allow("POST")
}
func (serv PathsService) GetMixed(Bool bool, Int int) string {
	return "Hello"
}
func (serv PathsService) GetMixed2(Bool bool, Int int) string {
	return "Hello"
}
