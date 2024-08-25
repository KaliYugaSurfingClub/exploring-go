package second

import "fmt"

type IRequester interface {
	Request() string
}

type RealRequester struct{}
type MockRequester struct{}

func (r RealRequester) Request() string {
	return "REAL RESPONSE"
}

func (m MockRequester) Request() string {
	return "MOCK RESPONSE"
}

type Requester struct {
	IRequester
	response string
}

func (r Requester) DoSomeComplex() {
	response := r.Request()
	response += "123"

	r.response = response

	fmt.Println(response)
}

func UseStrategy() {
	rM := Requester{MockRequester{}, ""}
	rM.DoSomeComplex()

	rR := Requester{RealRequester{}, ""}
	rR.DoSomeComplex()
}
