package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

const (
	Local      = "local"
	Sandbox    = "sandbox"
	Production = "production"
)

// =================================================

type SystemConfig struct {
	GoogleClientId      string
	GoogleClientIdIos   string
	FacebookAppId       string
	FirebaseAuthId      string
	SQLiteTestDB        bool
	EnableWorkerHandler bool

	SwaggerEnabled bool
	FirstTimeRun   int

	FirebaseConfigPath string
	MinioBucketName    string
	PortalDomain       string
	MinioDomain        string

	DomainApi string

	// Special configs
	CreateNecessaryConfigsForTables string

	// Postgres config
	PostgresConnectionString string

	// Redis config
	RedisConnectionAddress  string
	RedisConnectionPassword string
	RedisURL                string
	RedisPage               int

	// SMTP config
	SMTPHost      string
	SMTPPort      int
	SMTPUsername  string
	SMTPPassword  string
	SMTPFromEmail string
	SMTPFromName  string

	// Worker quantity
	WorkerQuantity int

	// Log/debug/swagger config
	PostgresLogMode bool
	DebugApi        bool
	QORSecretKey    string

	// App & Company config
	AppName          string
	OrganizationName string

	// Portal & api port & domain config
	CdnUrl                     string
	Domain                     string
	MediaDir                   string
	PortalLoggingRetentionDays int
	PortalPort                 string
	JobsPort                   string
	ApiPort                    string

	// Firebase config
	FirebaseFileName string
	FirebaseName     string

	// Portal feature config
	PortalFeatures_SystemEnabled bool
	PortalFeatures_OthersEnabled bool

	// reCAPTCHA
	ReCAPTCHA_SiteKey   string
	ReCAPTCHA_SecretKey string

	// GoAdmin
	IndexUrl string

	CameraAI_FeatureEnabled bool
	CameraAI_ClientID       string
	CameraAI_ClientSecret   string
	CameraAI_RefreshToken   string
	CameraAI_PlaceID        string

	// Phenikaa MaaS
	PhenikaaMaaS_FeatureEnabled bool

	// PHX
	PHXReceiveHanetCameraData bool

	// Scan QR code
	QrCodePrefixKey string
}

var Config *SystemConfig

func FetchEnvironmentVariables() {
	EnvType := os.Getenv("BUSMAP_ENV")
	Config = NewSystemConfig(EnvType)
}

func NewSystemConfig(env string) *SystemConfig {
	cf := SystemConfig{}
	fileConfig := cf.GetConfigFile(env)
	fmt.Printf("Load Config File: %s \n", fileConfig)
	viper.SetConfigFile(fileConfig)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Special configs
	cf.CreateNecessaryConfigsForTables = os.Getenv("CREATE_NECESSARY_CONFIGS_FOR_TABLES")

	cf.PostgresConnectionString = viper.GetString("postgres_connection_string")
	cf.SQLiteTestDB = viper.GetBool("sqlite_test_database")

	cf.CdnUrl = viper.GetString("cdn_url")

	// Redis config
	cf.RedisConnectionAddress = fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port"))
	cf.RedisConnectionPassword = viper.GetString("redis.password")
	cf.RedisURL = fmt.Sprintf("redis://:%s@%s:%s", url.QueryEscape(viper.GetString("redis.password")), viper.GetString("redis.host"), viper.GetString("redis.port"))
	cf.RedisPage = viper.GetInt("redis.page")

	// Worker quantity
	cf.WorkerQuantity = viper.GetInt("worker_quantity")

	cf.GoogleClientId = viper.GetString("auth.google_client_id")
	cf.GoogleClientIdIos = viper.GetString("auth.google_client_id_ios")
	cf.FacebookAppId = viper.GetString("auth.facebook_app_id")
	cf.FirebaseAuthId = viper.GetString("auth.firebase_auth_id")

	cf.SMTPUsername = viper.GetString("smtp.username")
	cf.SMTPPassword = viper.GetString("smtp.password")

	cf.PostgresLogMode = viper.GetBool("postgres_log_mode")
	cf.DebugApi = viper.GetBool("debug_api")
	cf.EnableWorkerHandler = viper.GetBool("enable_worker_handler")
	cf.QORSecretKey = viper.GetString("qor_secret_key")

	cf.Domain = viper.GetString("domain")
	//set default
	if cf.Domain == "" {
		cf.Domain = fmt.Sprintf("http://127.0.0.1:%v", cf.ApiPort)
	}

	cf.SwaggerEnabled = viper.GetBool("swagger_enabled")

	firtTimeRun, _ := strconv.Atoi(os.Getenv("FIRST_TIME"))
	cf.FirstTimeRun = firtTimeRun

	appName := viper.GetString("app_name")
	if appName == "" {
		appName = "bsmart-checkin"
	}
	cf.AppName = appName

	cf.FirebaseFileName = viper.GetString("firebase_file_name")
	cf.FirebaseName = viper.GetString("firebase_name")

	minioBucketName := viper.GetString("minio_bucket_name")
	if minioBucketName == "" {
		minioBucketName = "bsmartcheckin"
	}
	cf.MinioBucketName = minioBucketName

	portalDomain := viper.GetString("portal_domain")
	if portalDomain == "" {
		portalDomain = "https://portal-checkin.busmap.vn"
	}
	cf.PortalDomain = portalDomain

	minioDomain := viper.GetString("minio_domain")
	if minioDomain == "" {
		minioDomain = "https://storage.busmap.vn"
	}
	cf.MinioDomain = minioDomain

	cf.MediaDir = viper.GetString("media_dir")
	cf.DomainApi = viper.GetString("domain_api")
	cf.PortalPort = viper.GetString("portal_port")
	cf.JobsPort = viper.GetString("jobs_port")
	cf.ApiPort = viper.GetString("api_port")

	// Portal feature config
	cf.PortalFeatures_SystemEnabled = viper.GetBool("portal_features.system_enabled")
	cf.PortalFeatures_OthersEnabled = viper.GetBool("portal_features.others_enabled")

	cf.CameraAI_FeatureEnabled = viper.GetBool("camera_ai.feature_enabled")
	cf.CameraAI_ClientID = viper.GetString("camera_ai.client_id")
	cf.CameraAI_ClientSecret = viper.GetString("camera_ai.client_secret")
	cf.CameraAI_RefreshToken = viper.GetString("camera_ai.refresh_token")
	cf.CameraAI_PlaceID = viper.GetString("camera_ai.place_id")

	// App & Company config
	cf.AppName = viper.GetString("app_name")
	cf.OrganizationName = viper.GetString("organization_name")
	if cf.AppName == "" || cf.OrganizationName == "" {
		panic("Don't let app-name and organization-name be empty!")
	}

	// GoAdmin
	cf.IndexUrl = viper.GetString("index")

	// Phenikaa MaaS
	cf.PhenikaaMaaS_FeatureEnabled = viper.GetBool("phenikaa_maas.feature_enabled")

	// PHX
	cf.PHXReceiveHanetCameraData = viper.GetBool("phx.receive_hanet_camera_data_enabled")

	// Scan QR code
	cf.QrCodePrefixKey = viper.GetString("qr_code_prefix_key")

	return &cf
}

func (config *SystemConfig) GetConfigFile(env string) string {
	fileF := "zzz/config/goadmin_%s_config.json"
	switch env {
	case Local, Sandbox, Production:
		return fmt.Sprintf(fileF, env)
	default:
		return fmt.Sprintf(fileF, Local)
	}
}
