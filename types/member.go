package types

type MemberId struct {
	Cname	string
	Name 	string
}

type Member struct {
	Mid		MemberId
	Host	string
	Port 	string
}
