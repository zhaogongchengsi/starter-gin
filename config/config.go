package config

type Server struct {
	Port   int
	Host   string
	Prefix string
	Mode   string
	Static string
}

type DataBase struct {
	Username string
	Password string
	Port     int
	Url      string
	DbName   string
	Charset  string

	MaxIdleConns int
	MaxOpenConns int

	TablePrefix   string
	SingularTable bool
}

type Jwt struct {
	SigningKey string
	Issuer     string
}

func ReadConfigs(path, t, key string, value any) error {
	vp, err := InitViper(path, t)
	err = vp.UnmarshalKey(key, value)
	return err
}
