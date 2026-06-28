package storage

import (
	"github.com/hashicorp/go-memdb"
)

// CreateSchema creates the database schema for go-memdb
func CreateSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"users": {
				Name: "users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"username": {
						Name:    "username",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Username"},
					},
					"email": {
						Name:    "email",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
			"progress": {
				Name: "progress",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "UserID"},
					},
				},
			},
			"challenge_progress": {
				Name: "challenge_progress",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:   "id",
						Unique: true,
						Indexer: &memdb.CompoundIndex{
							Indexes: []memdb.Indexer{
								&memdb.StringFieldIndex{Field: "UserID"},
								&memdb.StringFieldIndex{Field: "ChallengeID"},
							},
						},
					},
					"user_id": {
						Name:    "user_id",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "UserID"},
					},
					"challenge_id": {
						Name:    "challenge_id",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ChallengeID"},
					},
				},
			},
		},
	}
}
