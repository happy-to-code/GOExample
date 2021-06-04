package main

import "fmt"

func test(authors []Author, authorId string) (serviceIdIsNull bool, author Author) {
	for _, a := range authors {
		if a.AuthId == authorId {
			if a.ServiceId == "" {
				serviceIdIsNull = true
				author = a
			}
		}
	}

	if !serviceIdIsNull {
		for _, a := range authors {
			if a.AuthId == authorId {
				author = a
			}
		}
	}
	return
}
func main() {
	var authorList = []Author{
		{Account: "", AuthId: "", ServiceId: ""},
		{Account: "3PLyMde9qkEeAvFfxCgW7DGq3L9Cgc39A82CjhY3xQiLUpU1Tm", AuthId: "202106031007227351", ServiceId: ""},
	}

	var newAuthorList []Author
	var authorId = "202106031007227351"
	var serviceId = "sss"

	serviceIdIsNull, a := test(authorList, authorId)
	fmt.Println("aa==>", serviceIdIsNull)
	fmt.Printf("bbb==>%+v\n", a)

	var hasAdd = false
	for _, author := range authorList {
		if author.AuthId == authorId {
			if author.ServiceId == "" {
				author.ServiceId = serviceId
				newAuthorList = append(newAuthorList, author)
				hasAdd = true
			} else {
				newAuthorList = append(newAuthorList, author)
				if !hasAdd {
					author.ServiceId = serviceId
					newAuthorList = append(newAuthorList, author)
					hasAdd = true
				}
			}
		} else {
			newAuthorList = append(newAuthorList, author)
		}
	}
	// if !serviceIdIsNull {
	// 	newAuthorList = append(newAuthorList, a)
	// }

	fmt.Printf("%+v\n", newAuthorList)

}

type Author struct {
	Account   string
	AuthId    string
	ServiceId string
	// ExpireEnd   uint64
	// ExpireStart uint64
}
