package buckets

import (
	"os"
)

func InitBuckets() {

}

func MinioEnvEndpoint() string {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		endpoint = "s3.amazonaws.com"
	}
	return endpoint
}

func MinioEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}
