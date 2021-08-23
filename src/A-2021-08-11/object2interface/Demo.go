package main

type Item struct {
	ItemName   string      `json:"itemName"`
	SourceHash string      `json:"sourceHash"`
	ItemValue  interface{} `json:"itemValue"`
}

type Boy struct {
	Name string
	Age  int
}

func main() {
	// var bs = []Boy{{"XM", 12}, {"XH", 18}}
	//
	// var item = Item{"item","sdfdfds",bs}

}
