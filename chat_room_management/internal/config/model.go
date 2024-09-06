package config

type Config struct {
	Server    Server
	Databases Database
	JWT       JWT
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type JWT struct {
	Key     string
	Expired int
}
