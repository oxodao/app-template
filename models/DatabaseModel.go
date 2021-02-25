package models

type DatabaseModel interface {
	GetTableName() string
	GetCreationScript() string
}
