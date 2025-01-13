package cerror

var (
	DetailUnprocessableEntityBirthDate = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "年齢は18歳以上でなければなりません。",
	}
)
