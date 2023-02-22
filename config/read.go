package config

const (
	ConfigType    string = "yaml"    // 配置文件类型
	ConfigDirPath string = "configs" // 配置文件路径
)

func ReadServerConfig() (*Server, error) {
	var c Server
	err := ReadConfigs(ConfigDirPath+"/server.yaml", ConfigType, "Server", &c)
	if err != nil {
		return &c, err
	}
	return &c, nil
}

func ReadJwtConfig() (*Jwt, error) {
	var c Jwt
	err := ReadConfigs(ConfigDirPath+"/server.yaml", ConfigType, "Jwt", &c)
	if err != nil {
		return &c, err
	}
	return &c, nil
}

func ReadDbConfig() (*DataBase, error) {
	var c DataBase
	err := ReadConfigs(ConfigDirPath+"/database.yaml", ConfigType, "DataBase", &c)
	if err != nil {
		return &c, err
	}
	return &c, nil
}

func ReadGenConfig() (*Gen, error) {
	var c Gen
	err := ReadConfigs(ConfigDirPath+"/database.yaml", ConfigType, "Gen", &c)
	if err != nil {
		return &c, err
	}
	return &c, nil
}
