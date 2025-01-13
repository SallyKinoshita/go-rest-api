package middleware

import (
	"golang.org/x/exp/slices"
)

// クライアントユーザーの認証をスキップするパス
var clientUserSkipAuthPath = []string{
	healthCheckPath,
	// 認証未実装のため仮にこちらに追加する
	sampleUserCreatePath,
	sampleUserDeletePath,
	sampleUserGetPath,
	sampleUserListPath,
	sampleUserUpdatePath,
}

func checkPath(path string, paths []string) bool {
	return slices.Contains(paths, path)
}
