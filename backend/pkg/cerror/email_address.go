package cerror

var (
	DetailUnprocessableEntityEmailAddress = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力したメールアドレスは無効な形式です。",
	}
)
