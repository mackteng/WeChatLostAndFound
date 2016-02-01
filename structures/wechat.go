package structures

type WeChatInteractor interface {
	GetAccessToken() string
	GetJSApiTicket() string

	SendSystemMessage(string, string, *GlobalConfiguration) error
	SendTemplateMessage(string, string, *GlobalConfiguration) error
	SendForwardMessage(string, string, string,*GlobalConfiguration) error
	SendBulkForwardMessages([]string, string, *GlobalConfiguration) error
}
