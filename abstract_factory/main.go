package main

import "fmt"

type UserIeterface interface {
	getUserRole() string
}

type Member struct {
}

func (u *Member) getUserRole() string {
	return "普通用户"
}

type VIP struct {
}

func (u *VIP) getUserRole() string {
	return "VIP用户"
}

type SVIP struct {
}

func (u *SVIP) getUserRole() string {
	return "超级VIP用户"
}

// 抽象工厂
type UserFactory interface {
	CreateUser() UserIeterface
}

type MemberUserFactory struct {
}

func (fact *MemberUserFactory) CreateUser() UserIeterface {
	return new(Member)
}

type VIPUserFactory struct {
}

func (fact *VIPUserFactory) CreateUser() UserIeterface {
	return new(VIP)
}

type SVIPUserFactory struct {
}

func (fact *SVIPUserFactory) CreateUser() UserIeterface {
	return new(SVIP)
}

func main() {

	var ufact UserFactory = new(MemberUserFactory)
	fmt.Println(ufact.CreateUser().getUserRole())

	ufact = new(SVIPUserFactory)
	fmt.Println(ufact.CreateUser().getUserRole())

	ufact = new(VIPUserFactory)
	fmt.Println(ufact.CreateUser().getUserRole())
}
