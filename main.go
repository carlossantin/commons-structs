package main

import authz "github.com/carlossantin/resource-policy-framework/auth"

func main() {
	r := authz.Role{}
	println(r.ID)
}
