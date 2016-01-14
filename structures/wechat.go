package structures

type WeChatInteractor interface {
	GetAccessToken() string

	SendSystemMessage(string, string, *GlobalConfiguration) error
	SendTemplateMessage(string, string, *GlobalConfiguration) error
	SendForwardMessage(string, string, int, *GlobalConfiguration) error
	SendBulkForwardMessages([]string, string, *GlobalConfiguration) error
}
