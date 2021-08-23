package main

func main() {
	p := new(People)
	p.GetAllUser()

	var m MemberInterface = &People{}
	m.GetAllMember()
}

type UserInterface interface {
	GetAllUser() []string
	GetById(id string) string
}

type Boy struct {
}

func (b *Boy) GetAllUser() []string {

	return []string{"XM", "XH"}
}

func (b *Boy) GetById(id string) string {
	return "MM"
}

type People struct {
}
type MemberInterface interface {
	GetAllMember() []string

	GetAllUser() []string
	GetById(id string) string
}

func (p *People) GetAllMember() []string {
	return []string{"P2P", "P1P"}
}

func (p *People) GetAllUser() []string {
	return []string{"User1", "User2"}
}

func (p *People) GetById(id string) string {
	return "1p"
}
