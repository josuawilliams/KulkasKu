package model

type Config struct {
	Port string
	DBUrlMigrate string
	SecretJwt string

	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}