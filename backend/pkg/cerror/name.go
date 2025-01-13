package cerror

var (
	DetailUnprocessableEntityFirstName = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した名前は無効な形式です。",
	}
	DetailUnprocessableEntityLastName = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した名字は無効な形式です。",
	}
	DetailUnprocessableEntityFirstNameKana = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した名前（カナ）は無効な形式です。",
	}
	DetailUnprocessableEntityLastNameKana = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した名字（カナ）は無効な形式です。",
	}
)
