package bnrwebframework

import "fmt"

// Start begin gin server service
func (c *defaultGin) Start(port string) error {
	server := fmt.Sprintf(":%v", port)
	return c.server.Run(server)
}
