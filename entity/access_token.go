package entity

type AccessToken struct {
	Token       string `json:"accessToken"`
	ExpiredTime int64  `json:"expiredTime"`
}
