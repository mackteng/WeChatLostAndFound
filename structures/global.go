package structures


type GlobalConfiguration struct {
        WeChatConfig   *Config
        DatabaseConfig *DatabaseAccessInfo
        RedisAccessInfo *RedisAccessInfo
}


func InitGlobalConfig() *GlobalConfiguration {

	configuration := NewConfig()
        configuration.RefreshAccessToken()
        database := NewDatabase()
        redis := NewRedis()	

	return &GlobalConfiguration{

		WeChatConfig : configuration,
		DatabaseConfig: database,
		RedisAccessInfo : redis,

	}


}
