package config

type Server struct {
	Port       int
	Host       string
	Prefix     string
	Mode       string
	StaticName string
	Static     string
	IndexHtml  string
	Https      HttpsConfig
}

type HttpsConfig struct {
	CertFile string
	KeyFile  string
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
	ExpiresAt  int
}

type Gen struct {
	OutPath           string
	FieldNullable     bool
	FieldCoverable    bool
	FieldSignable     bool
	FieldWithIndexTag bool
	FieldWithTypeTag  bool
}

type Redis struct {
	Db       int
	Addr     string
	Password string
}

type Config struct {
	Server   Server
	DataBase DataBase
	Jwt      Jwt
	Gen      Gen
	Redis    Redis
}
