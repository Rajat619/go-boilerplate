package config

import "main.go/utils"

var envVarUtilities = utils.EnvVariableUtil{}

var AppConfig *ApplicationConfig

func init() {
	AppConfig = &ApplicationConfig{
		Port:        envVarUtilities.GetSanitizedEnvAsInt("APP_PORT", "3000"),
		DbConfig:    setupDb(),
		QueueConfig: setupQueueConnection(),
	}
}

func setupDb() DbConfig {
	return DbConfig{
		DbHost:     envVarUtilities.GetSanitizedEnvAsString("DB_HOST", "localhost"),
		DbPort:     envVarUtilities.GetSanitizedEnvAsInt("DB_PORT", "5432"),
		DbName:     envVarUtilities.GetSanitizedEnvAsString("DB_NAME", "rajat"),
		DbUser:     envVarUtilities.GetSanitizedEnvAsString("DB_USER", "rajat"),
		DbPassword: envVarUtilities.GetFieldFromEnv("DB_PASSWORD", ""),
		DbSslMode:  envVarUtilities.GetSanitizedEnvAsBoolean("DB_SSL_MODE", "FALSE"),
	}
}

func setupQueueConnection() QueueConfig {
	return QueueConfig{
		ConnectionName: envVarUtilities.GetSanitizedEnvAsString("QUEUE_NAME", "default"),
		Protocol:       envVarUtilities.GetSanitizedEnvAsString("REDIS_PROTOCOL", "tcp"),
		Url:            envVarUtilities.GetSanitizedEnvAsString("REDIS_URL", "localhost:6379"),
	}
}
