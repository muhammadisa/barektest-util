package jwt

func (j jwt) ExtractKey(key string) (string, error) {
	token, err := j.Parser()
	if err != nil {
		return blankString, err
	}
	claimed, err := j.Claim(token, key)
	if err != nil {
		return blankString, err
	}
	return claimed, nil
}
