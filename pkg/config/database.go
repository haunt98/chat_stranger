package config

type DatabaseReader interface {
	GetDBMS() string
	GetSource() string
}
