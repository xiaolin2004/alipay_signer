package tools

func BuildBaseHeader(reqTime string, clientId string, keyVersion string, signatureValue string) map[string]string {
	if keyVersion == "" {
		keyVersion = "1"
	}
	signatureValue = "algorithm=RSA256,keyVersion=" + keyVersion + ",signature=" + signatureValue
	return map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
		"Request-Time": reqTime,
		"Client-Id":    clientId,
		"Key-Version":  keyVersion,
		"Signature":    signatureValue,
	}
}
