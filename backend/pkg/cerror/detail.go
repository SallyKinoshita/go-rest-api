package cerror

type Detail struct {
	status        status
	clientMessage message
}

func (d *Detail) validateAndSetDetailInternal() {
	if d.status == 0 || d.clientMessage == "" {
		*d = DetailInternalServerError
	}
}

var (
	DetailBadRequest = Detail{
		status:        StatusBadRequest,
		clientMessage: messageBadRequest,
	}
	DetailUnauthorized = Detail{
		status:        StatusUnauthorized,
		clientMessage: messageUnauthorized,
	}
	DetailForbidden = Detail{
		status:        StatusForbidden,
		clientMessage: messageForbidden,
	}
	DetailNotFound = Detail{
		status:        StatusNotFound,
		clientMessage: messageNotFound,
	}
	DetailRequestTimeout = Detail{
		status:        StatusRequestTimeout,
		clientMessage: messageRequestTimeout,
	}
	DetailTooManyRequests = Detail{
		status:        StatusTooManyRequests,
		clientMessage: messageTooManyRequests,
	}
	DetailInternalServerError = Detail{
		status:        StatusInternalServerError,
		clientMessage: messageInternal,
	}
)
