package variable

const (
	ServiceName = "chat_stranger"
	ConfigFile  = "config"
	ConfigPath  = "configs"

	Port       = "port"
	DbDialect  = "db.dialect"
	DbUrl      = "db.url"
	DbMode     = "db.mode"
	JWTSecret  = "token.secret"
	GinMode    = "gin.mode"
	ConfigMode = "config.mode"

	ViperMode   = "viper"
	TestMode    = "test"
	DebugMode   = "debug"
	ReleaseMode = "release"

	UserRole  = "user"
	AdminRole = "admin"

	LimitRoom = 2

	WebPrefix            = "/chat_stranger/web"
	APIPrefix            = "/chat_stranger/api"
	HTMLGlob             = "./web/*.html"
	StaticRelativeScript = "/chat_stranger/web/script"
	StaticScript         = "./web/script"
	StaticRelativeStyle  = "/chat_stranger/web/style"
	StaticStyle          = "./web/style"
	StaticRelativeImg    = "/chat_stranger/web/img"
	StaticImg            = "./web/img"
)
