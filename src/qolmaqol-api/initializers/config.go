package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Environment   string `mapstructure:"ENV"`
	InternalDBURL string `mapstructure:"INTERNAL_DB_URL"`
	ExternalDBURL string `mapstructure:"EXTERNAL_DB_URL"`

	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`

	FirebaseStorageFile string `mapstructure:"FIREBASE_STORAGE_FILE"`
	FirebaseBucketName  string `mapstructure:"FIREBASE_BUCKET_NAME"`

	IPServiceUrl string `mapstructure:"IP_SERVICE_URL"`
}

func LoadConfig(path string) (config DBConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
