package types

type Auth int

const (
	AuthLogin Auth = iota
	AuthToken
	AuthCertificate
)
