package jwt

func (j jwt) ExtractKeys(keys []string) (map[string]string, error) {
	claims := make(map[string]string)
	token, err := j.Parser()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		claimed, err := j.Claim(token, key)
		if err != nil {
			return nil, err
		}
		claims[key] = claimed
	}
	return claims, nil
}
