package constants

const (
	BookTableName   = "book"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Books           = "books"
	ApiServiceName  = "bookapi"
	BookServiceName = "book"
	UserServiceName = "user"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	BookServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)