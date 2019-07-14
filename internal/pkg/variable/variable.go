package variable

const (
	Port      = "port"
	DbDialect = "db.dialect"
	DbUrl     = "db.url"
	DbMode    = "db.mode"
	JWTSecret = "token.secret"
	GinMode   = "gin.mode"

	Mode      = "mode"
	ViperMode = "viper"
	TestMode  = "test"

	UserRole  = "user"
	AdminRole = "admin"

	LimitRoom = 2

	WebPrefix      = "/chat_stranger/web"
	APIPrefix      = "/chat_stranger/api"
	HTMLGlob       = "./web/*.html"
	StaticRelative = "/chat_stranger/web/script"
	StaticRoot     = "./web/script"
)
