package bnrwebframework

// RegisterAPIVersionGroup create default api version group api/xx
func (c *defaultGin) RegisterAPIVersionGroup() RouterGroup {
	g := c.Group(defaultAPIPath)
	c.apiVersionGroup = g
	return g
}

// APIVersionGroup return api version group's RouterGroup
func (c *defaultGin) APIVersionGroup() RouterGroup {
	return c.apiVersionGroup
}
