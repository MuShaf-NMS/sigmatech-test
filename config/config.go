package config

var Conf = &Config{
	Allowed_Ext: []string{"jpg", "jpeg", "png"},
}

type Config struct {
	App_Mode    string
	App_Port    string
	DB_User     string
	DB_Pass     string
	DB_Name     string
	DB_Host     string
	DB_Port     string
	Secret_Key  string
	Allowed_Ext []string
	Assets_Link string
	Assets_Dir  string
}

func LoadConfig() *Config {
	Conf.App_Mode = getVariable("APP_MODE")
	Conf.App_Port = getVariable("APP_PORT")
	Conf.DB_User = getVariable("DB_USER")
	Conf.DB_Pass = getVariable("DB_PASS")
	Conf.DB_Name = getVariable("DB_NAME")
	Conf.DB_Host = getVariable("DB_HOST")
	Conf.DB_Port = getVariable("DB_PORT")
	Conf.Secret_Key = getVariable("SECRET_KEY")
	Conf.Assets_Link = getVariable("ASSETS_LINK")
	Conf.Assets_Dir = getVariable("ASSETS_DIR")
	return Conf
}
