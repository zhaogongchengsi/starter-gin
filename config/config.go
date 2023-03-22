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
	UploadDir  string
}

type HttpsConfig struct {
	CertFile string
	KeyFile  string
}

type DataBase struct {
	DbType   string
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

type Captcha struct {
	Width, Height, Length, DotCount int
	MaxSkew                         float64
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

type Config struct {
	AppId    string
	Server   Server
	DataBase DataBase
	Jwt      Jwt
	Gen      Gen
	Redis    Redis
	Captcha  Captcha
	Zap      Zap
}
