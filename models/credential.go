package models

import "fmt"

type Credential struct {
	Id       int    `json:"id"`
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Credential) GetCredData() {
	fmt.Printf("Id: %v; Domain: %v; %v:%v\n", c.Id, c.Domain, c.Username, c.Password)
}
