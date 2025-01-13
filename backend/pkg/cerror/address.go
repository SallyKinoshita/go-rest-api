package cerror

var (
	DetailUnprocessableEntityPostalCode = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した郵便番号は無効な形式です。",
	}
	DetailUnprocessableEntityPrefecture = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した都道府県は無効な形式です。",
	}
	DetailUnprocessableEntityCity = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した市区町村は無効な形式です。",
	}
	DetailUnprocessableEntityStreetAndNumber = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した丁目・番地は無効な形式です。",
	}
	DetailUnprocessableEntityBuildingAndRoom = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "入力した建物名・部屋番号は無効な形式です。",
	}
)
