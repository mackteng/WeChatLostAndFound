package sysmsg

const (
	WELCOME_MESSAGE = "歡迎來到失物招領平台!"

	REGISTER_SUCCESS        = "您的物品已成功被註冊"
	REGISTER_FAIL           = "您的物品未成功被註冊"
	ITEM_ALREADY_REGISTERED = "該TAG已被註冊 - 您是否找到此物品?"

	FIND_SUCCESS = "您與失主已聯繫上，請立即傳送訊息連絡失主！"
	FIND_FAIL    = "未能找到施主，您是否要註冊此物品？"

	OWNER_LIMIT_REACHED  = "註冊物品數量已達上限"
	FINDER_LIMIT_REACHED = "註冊物品數量已達上限"

	CHANNEL_CHANGE      = "已成功切換至頻道　"
	CHANNEL_CHANGE_FAIL = "未能切換至頻道　"

	SAME_CHANNEL = "已在該頻道上！"
)
