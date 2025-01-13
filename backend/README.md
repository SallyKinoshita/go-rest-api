# backend

## 命名規則 (API・メソッド・関数)

- API名は`CRUD + ドメイン名`で命名

| 命名 | 内容 |
| ------ | ----- |
| Create | 作成 |
| Update | 更新 |
| Get    | 単体取得 |
| List   | 複数取得 |
| Delete | 削除 |
| Exists | 存在チェック |
| Set    | 登録時は空だが、後ほど設定する項目 |

## 開発環境構築

```shell
cd backend
cp .env.example .env
make init
```

## DB

### 再設定 (ドロップ・マイグレーション・テストデータ挿入)

```shell
make db-recreate
```


### マイグレーションファイル自動生成 (差分反映)

```shell
make db-migrate-diff
```

### マイグレーション (反映)
※一部seederを含む。

```shell
make db-migrate-apply
```

### ドロップ

```shell
make db-drop
```

## Go

### チェック (フォーマット・静的解析)

```shell
make go-check
```

### フォーマット

```shell
make go-fmt
```

### 静的解析

```shell
make go-lint
```

### モジュール更新

```shell
make go-mod-tidy
```

### テスト

```shell
make go-test
```

## Swagger

### 準備

- `make up`を実行し、Server起動

### Swagger 更新

1. `doc/openapi.yaml`を更新
2. `make go-update-schema`実行

### Swagger Editor

- ブラウザで`http://localhost:8001/`を開く

### Swagger UI

- ブラウザで`http://localhost:8002/`を開く