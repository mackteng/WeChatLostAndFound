package structures

type GlobalConfiguration struct {
	WeChatInteractor   WeChatInteractor
	DatabaseInteractor DatabaseInteractor
	RedisInteractor    RedisInteractor
}

/*
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
*/
