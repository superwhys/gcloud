package define

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// JwtKey
type UserClaim struct {
	Id       int
	Identity string
	Name     string
	Email    string
	jwt.StandardClaims
}

func GetenvWithDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

var JwtKey = "gcloud-key"
var MailAccount = GetenvWithDefault("MAIL_ACCOUNT", "")
var MailPassword = GetenvWithDefault("MAIL_PASSWORD", "")
var RedisPassword = GetenvWithDefault("REDIS_PASSWORD", "")
var RedisAddr = GetenvWithDefault("REDIS_ADDR", "localhost:6379")
var MySQLAddr = GetenvWithDefault("MYSQL_ADDR", "localhost:3306")
var MySQLUser = GetenvWithDefault("MYSQL_USER", "gcloud")
var MySQLPassword = GetenvWithDefault("MYSQL_PASSWORD", "toD9GEdVpDRq")

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = GetenvWithDefault("TENCENT_SECRETKEY", "RgzYHyaHJLYlAjP1ZwDQyTtiheXOVhow")
var TencentSecretID = GetenvWithDefault("TENCENT_SECRETID", "AKIDUG9EKg1ICEMUK2dbepEhvGi7EB4um2Ja")
var CosBucket = "https://gcloud-1323673364.cos.ap-chongqing.myqcloud.com"
var CosFolderName = "gcloud"
var AvatarBaseUrl = CosBucket + "/" + CosFolderName + "/avatars/"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2000-01-01 00:00:01"

var TokenExpire = 60 * 60 * 24 * 3        // 3 days
var RefreshTokenExpire = 60 * 60 * 24 * 7 // 7 days

var UserRepositoryMaxSize = 1000 * 1024 * 1024  // 1GB
var PublicRepositoryMaxSize = 500 * 1024 * 1024 // 0.5GB
var UserRepositoryMinSize = 200 * 1024 * 1024   // 200MB
