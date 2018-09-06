package databaseinit

import (
	postgres "github.com/AlifElectronicQueue/internal/pkg/databaseinit/Repository/Postgres"
	"github.com/AlifElectronicQueue/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

func CreateAuthenticationRepository(DBType string, ConVar *sqlx.DB) types.IAuthenticationRepository {
	switch DBType {
	case "postgres":
		return &postgres.AuthenticationRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}
