package variable

const (
	ServiceName = "chat_stranger"
	ConfigFile  = "config"
	ConfigPath  = "configs"

	Port      = "port"
	DbDialect = "db.dialect"
	DbUrl     = "db.url"
	DbMode    = "db.mode"
	JWTSecret = "jwt.secret"
	GinMode   = "gin.mode"

	Config = "config"
	Viper  = "viper"
	Test   = "test"

	Debug = "debug"

	UserRole = "user"

	LimitRoom         = 2
	StatusRoom        = "status"
	AnyRoom           = "any"
	NextRoom          = "next"
	SameGenderRoom    = "gender"
	SameBirthYearRoom = "birth"

	FromTime = "from"

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
