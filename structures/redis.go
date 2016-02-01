package structures

type RedisInteractor interface {
	AddMessageToQueue(string, string, string) error
	GetMessagesFromQueue(string, string) ([]string, error)

	IsDuplicateMsgID(string) (bool, error)
}
