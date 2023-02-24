package config

type Server struct {
	Port       int
	Host       string
	Prefix     string
	Mode       string
	StaticName string
	Static     string
	IndexHtml  string
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

type Config struct {
	Server
	DataBase
	Jwt
	Gen
}
