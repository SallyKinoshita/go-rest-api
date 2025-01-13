schema "default" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}

table "sample_user" {
  schema = schema.default

  column "id" {
    type = char(36)
    null = false
  }
  column "first_name" {
    type = varchar(255)
    null = false
  }
  column "last_name" {
    type = varchar(255)
    null = false
  }
  column "first_name_kana" {
    type = varchar(255)
    null = false
  }
  column "last_name_kana" {
    type = varchar(255)
    null = false
  }
  column "email_address" {
    type = varchar(255)
    null = false
  }
  column "tel" {
    type = varchar(255)
    null = false
  }
  column "birth_date" {
    type = datetime
    null = false
  }
  column "postal_code" {
    type = varchar(255)
    null = false
  }
  column "prefecture" {
    type = varchar(255)
    null = false
  }
  column "city" {
    type = varchar(255)
    null = false
  }
  column "street_and_number" {
    type = varchar(255)
    null = false
  }
  column "building_and_room" {
    type = varchar(255)
    null = false
  }
  column "created_at" {
    type = datetime
    null = false
  }
  column "updated_at" {
    type = datetime
    null = false
  }
  primary_key {
    columns = [column.id]
  }
}