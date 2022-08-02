package config

type ClientConfig struct {
	OpenSecret     string `json:"OpenSecret" bson:"OpenSecret"`
	OpenPrivateKey string `json:"OpenPrivateKey" bson:"OpenPrivateKey"`
	OpenPublicKey  string `json:"OpenPublicKey" bson:"OpenPublicKey"`
}
