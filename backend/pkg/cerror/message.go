package cerror

type message string

const (
	messageBadRequest      message = "リクエストの形式が不正です。"
	messageUnauthorized    message = "ログインされていません。ログインし直してください。"
	messageForbidden       message = "アクセス権限がありません。"
	messageNotFound        message = "リソースが見つかりません。"
	messageRequestTimeout  message = "リクエストがタイムアウトしました。ネットワークを確認し、再試行してください。"
	messageTooManyRequests message = "リクエストが多すぎます。しばらく時間を置いてから再試行してください。"
	messageInternal        message = "エラーが発生しました。再試行しても解消しない場合は、「お問合せ」からご連絡ください。"
)
