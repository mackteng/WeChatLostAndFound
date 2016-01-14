package structures

type RedisInteractor interface {
	AddMessageToQueue(string, int, string) error
	GetMessagesFromQueue(string, int) ([]string, error)

	IsDuplicateMsgID(string) (bool, error)
}
