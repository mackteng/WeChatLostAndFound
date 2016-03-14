package structures

type DatabaseInteractor interface {
	
	AddUser(OpenID string) error
	RegisterTag(OpenID string, TagID string, Info ItemInfo) error
	FindTag(FinderOpenID string, TagID string) error
	FindCorrespondingUser(OpenID string) (string, string, error)

	//GetActiveTag(OpenID string) (string, error)
	ChangeActiveTag(OpenID, TagID string) error
	GetAllOwnedItems(OpenID string) ([]ItemInfo,  error)
	DeleteTag(OpenID string, TagID string) error
	DeleteFinder(OpenID, TagID string) error
	
	//GetAllFoundItems(OpenID string) ([]string, error)
}
