package cerror

var (
	DetailNotFoundClientUser = Detail{
		status:        StatusNotFound,
		clientMessage: "クライアントユーザーが存在しません。",
	}
	DetailConflictClientUserEmail = Detail{
		status:        StatusConflict,
		clientMessage: "既に登録済みのメールアドレスです。",
	}
	DetailUnprocessableEntityClientUserName = Detail{
		status:        StatusUnprocessableEntity,
		clientMessage: "名前は3文字以上30文字以下で入力してください。",
	}
)
