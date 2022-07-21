package config

// Config 設定檔
type Config struct {
	Host        string
	Port        string
	Account     string
	Password    string
	DBName      string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifeTime int
}
