package bnrwebframework

func (c *Context) ReadFormDataFile(formField string) (bool, string, []byte, error) {
	fileName, fileBytes, err := c.ReadFormDataFileBytes(formField)
	foundFile := true
	if err != nil {
		if err.Error() == "http: no such file" {
			foundFile = false
		} else {
			return false, "", nil, err
		}
	}
	if foundFile && fileBytes != nil {
		return true, fileName, fileBytes, nil
	}
	return false, "", nil, nil
}
