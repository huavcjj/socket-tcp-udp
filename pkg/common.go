package pkg

type Request struct {
	RequestId int    `json:"request_id"`
	Name      string `json:"name"`
}

type Response struct {
	RequestId int    `json:"request_id"`
	SayHello  string `json:"say_hello"`
}
