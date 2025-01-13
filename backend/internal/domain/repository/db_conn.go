package repository

import "github.com/uptrace/bun"

type DBConn interface {
	NewInsert() *bun.InsertQuery
	NewUpdate() *bun.UpdateQuery
	NewSelect() *bun.SelectQuery
	NewDelete() *bun.DeleteQuery
}
