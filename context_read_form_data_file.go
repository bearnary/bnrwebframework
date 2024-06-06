package bnrwebframework

import (
	"bytes"
	"io"
)

func (c *Context) ReadFormDataFileBytes(key string) (string, []byte, error) {
	file, header, err := c.Context.Request.FormFile(key)
	if err != nil {
		return "", nil, err
	}
	filename := header.Filename
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", nil, err
	}
	return filename, buf.Bytes(), nil
}
