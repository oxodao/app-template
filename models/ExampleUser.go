package models

type ExampleUser struct {
	Uid      string `db:"USER_ID"`
	Username string `db:"USER_NAME"`
}

func (eu ExampleUser) GetTableName() string {
	return "EXAMPLE_USER"
}

func (eu ExampleUser) GetCreationScript() string {
	return `
		CREATE TABLE EXAMPLE_USER(
			USER_ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			USER_NAME VARCHAR(255) NOT NULL
		);
	`
}
