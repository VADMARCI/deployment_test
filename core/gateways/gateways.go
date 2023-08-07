package gateways

type RemoteFileGatewayInterface interface {
	Upload(name string, body []byte, contentType string) error
	GetPublicSignedPathForObject(name string) string
	GetPublicPathForObject(name string) string
	GetPutUrl(name, mimeType string) (string, error)
	Delete(name string) error
}

type HashGatewayInterface interface {
	GenerateHash(length int) string
}
