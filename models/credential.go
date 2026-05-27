package models

import "fmt"

type Credential struct {
	Id       int
	Domain   string
	Username string
	Password string
}

func (c *Credential) GetCredData() {
	fmt.Printf("Domain: %v; %v:%v\n", c.Domain, c.Username, c.Password)
}
