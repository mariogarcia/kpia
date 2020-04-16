package api

import (
	"database/sql"
	"log"
)

// Entries are fixtures
var Entries []string = []string{
	`CREATE TABLE piweek (id UUID, year integer, month integer, from datetime, to datetime); `,
	`CREATE TABLE projects (id UUID, name text, description text, piweek_id integer);`,
	`CREATE TABLE members (id UUID, name text, profile text);`,
	`CREATE TABLE projects_members (project_id UUID, member_id UUID);`,
	`INSERT INTO piweek (id, year, month) VALUES ('45bba81a-802f-11ea-bc55-0242ac130003', 2019, 12);`,
	`INSERT INTO projects (id, name, piweek_id) VALUES ('45bbab76-802f-11ea-bc55-0242ac130003', 'Masiaventura', '45bba81a-802f-11ea-bc55-0242ac130003');`,
	`INSERT INTO members (id, name, profile) VALUES ('45bbac84-802f-11ea-bc55-0242ac130003', 'Mario', 'BACK');`,
	`INSERT INTO projects_members (project_id, member_id) VALUES ('45bbab76-802f-11ea-bc55-0242ac130003', '45bbac84-802f-11ea-bc55-0242ac130003');`,
}

// LoadFixtures loads fixtures entries
func LoadFixtures(db *sql.DB) {
	log.Println("creating database")
	for _, stmt := range Entries {
		db.Exec(stmt)
	}
}
