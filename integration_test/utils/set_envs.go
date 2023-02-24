package utils

import "os"

func SetEnvs() {
	os.Setenv("PORT", "3000")
	os.Setenv("DB_HOST", "postgres_db")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "admin")
	os.Setenv("DB_NAME", "develop")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("REDIS_HOST", "redis")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("STREAM_REDIS_NAME", "myRedisStream")
	os.Setenv("CONSUMER_GROUP_REDIS_NAME", "myConsumerGroup")
	os.Setenv("JWT_KEY", "my_ultra_secret_p@ssword-for-jwt")
	os.Setenv("ADMIN_USERNAME", "admin")
	os.Setenv("ADMIN_PASSWORD", "admin")
}
