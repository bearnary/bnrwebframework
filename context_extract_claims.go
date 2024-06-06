package bnrwebframework

import (
	"encoding/json"
)

func (c *Context) ParseJWT(jc JWT, output interface{}) error {
	// validate token
	claims, err := jc.ParseAccessToken(c)
	if err != nil {
		return err
	}

	// convert  map[string]interface{} to struct authenmodels.TokenData
	jwtJSON, err := json.Marshal(claims)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jwtJSON, output); err != nil {
		return err
	}
	return nil
}
