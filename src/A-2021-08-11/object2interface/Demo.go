package main

type Item struct {
	ItemName   string      `1.json:"itemName"`
	SourceHash string      `1.json:"sourceHash"`
	ItemValue  interface{} `1.json:"itemValue"`
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
