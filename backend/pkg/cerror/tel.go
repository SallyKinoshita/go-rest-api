package cerror

var (
	DetailUnprocessableEntityTel = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した電話番号は無効な形式です。",
	}
)
