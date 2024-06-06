package bnrwebframework

type JWT interface {
	GenerateAccessToken(model interface{}) (*Token, error)
	GenerateRefreshToken(model interface{}) (*Token, error)
	ParseAccessToken(ctx *Context) (map[string]interface{}, error)
	ParseRefreshToken(token string) (map[string]interface{}, error)
	ExtractToken(ctx *Context) string
	ValidateTokenMiddleware(ctx *Context)
}
