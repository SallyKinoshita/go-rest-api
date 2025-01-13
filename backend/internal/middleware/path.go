package middleware

const (
	// Sample User Endpoints
	sampleUserBasePath         = "/api/v1/sample-user"
	sampleUserCreatePath       = sampleUserBasePath + "/create"
	sampleUserDeletePath       = sampleUserBasePath + "/delete"
	sampleUserGetPath          = sampleUserBasePath + "/get"
	sampleUserListPath         = sampleUserBasePath + "/list"
	sampleUserUpdatePath       = sampleUserBasePath + "/update"
	sampleUserTokenRefreshPath = sampleUserBasePath + "/token-refresh" // 中身は未実装だが、middlewareで参照するため仮のパスを指定

	// Health Check Endpoint
	healthCheckPath = "/api/v1/health-check"
)
