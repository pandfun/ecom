package config

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

// Implement the init config function like this
// func initConfig() Config {
// 	return Config{
// 		PublicHost:             "http://localhost",
// 		Port:                   "8080",
// 		DBUser:                 "user",
// 		DBPassword:             "paswd",
// 		DBAddress:              "127.0.0.1:3306",
// 		DBName:                 "db-name",
// 		JWTSecret:              "your-secret-here",
// 		JWTExpirationInSeconds: 3600 * 24 * 7,
// 	}
// }
