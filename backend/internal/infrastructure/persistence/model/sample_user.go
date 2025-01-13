package persistencemodel

import (
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/uptrace/bun"
)

type SampleUser struct {
	bun.BaseModel `bun:"table:sample_user,alias:su"`

	ID              string    `bun:"id,type:char(36),pk,notnull"`                 // サンプルユーザID
	FirstName       string    `bun:"first_name,type:varchar(255),notnull"`        // 名
	LastName        string    `bun:"last_name,type:varchar(255),notnull"`         // 姓
	FirstNameKana   string    `bun:"first_name_kana,type:varchar(255),notnull"`   // 名（カナ）
	LastNameKana    string    `bun:"last_name_kana,type:varchar(255),notnull"`    // 姓（カナ）
	EmailAddress    string    `bun:"email_address,type:varchar(255),notnull"`     // メールアドレス
	Tel             string    `bun:"tel,type:varchar(255),notnull"`               // 電話番号
	BirthDate       time.Time `bun:"birth_date,type:datetime,notnull"`            // 生年月日
	PostalCode      string    `bun:"postal_code,type:varchar(255),notnull"`       // 郵便番号
	Prefecture      string    `bun:"prefecture,type:varchar(255),notnull"`        // 都道府県
	City            string    `bun:"city,type:varchar(255),notnull"`              // 市区町村
	StreetAndNumber string    `bun:"street_and_number,type:varchar(255),notnull"` // 番地
	BuildingAndRoom string    `bun:"building_and_room,type:varchar(255),notnull"` // 建物名・部屋番号
	CreatedAt       time.Time `bun:"created_at,type:datetime,notnull"`            // 作成日時
	UpdatedAt       time.Time `bun:"updated_at,type:datetime,notnull"`            // 更新日時
}

func NewSampleUser(now time.Time, sampleUser model.SampleUser) *SampleUser {
	return &SampleUser{
		ID:              sampleUser.ID().String(),
		FirstName:       sampleUser.Name().First(),
		LastName:        sampleUser.Name().Last(),
		FirstNameKana:   sampleUser.Name().FirstKana(),
		LastNameKana:    sampleUser.Name().LastKana(),
		EmailAddress:    sampleUser.EmailAddress().String(),
		Tel:             sampleUser.Tel().String(),
		BirthDate:       sampleUser.BirthDate().Time,
		PostalCode:      sampleUser.Address().PostalCode().String(),
		Prefecture:      sampleUser.Address().Prefecture(),
		City:            sampleUser.Address().City(),
		StreetAndNumber: sampleUser.Address().StreetAndNumber(),
		BuildingAndRoom: sampleUser.Address().BuildingAndRoom(),
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
