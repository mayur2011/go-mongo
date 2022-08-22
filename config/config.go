package config

type AppConfiguration struct {
	Environment string
	Token string
	Mongo MongoConfigType
}

type MongoConfigType struct {
	URI string
	User string
	Password string
}

func GetConfig() AppConfiguration{
	
	return  AppConfiguration{}
}
