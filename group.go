package bnrwebframework

// Group create new RouterGroup from relativePath
func (c *defaultGin) Group(relativePath string) RouterGroup {
	g := c.server.Group(relativePath)
	return &defaultRouterGroup{
		g: *g,
	}
}
