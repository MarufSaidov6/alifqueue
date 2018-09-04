package databaseinit

import (
	postgres "github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/databaseinit/Repository/Postgres"
	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
	"github.com/jmoiron/sqlx"
)

func CreateAnalyticsRepository(DBType string, ConVar *sqlx.DB) types.IAnalyticsRepository {
	switch DBType {
	case "postgres":
		return &postgres.AnalyticsRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}

func CreateAdminRepository(DBType string, ConVar *sqlx.DB) types.IAdminRepository {
	switch DBType {
	case "postgres":
		return &postgres.AdminRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}

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

func CreateCashbackRepository(DBType string, ConVar *sqlx.DB) types.ICashbackRepository {
	switch DBType {
	case "postgres":
		return &postgres.CashbackRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}
func CreateCatalogueRepository(DBType string, ConVar *sqlx.DB) types.ICatalogueRepository {
	switch DBType {
	case "postgres":
		return &postgres.CatalogueRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}

func CreateDiscountRepository(DBType string, ConVar *sqlx.DB) types.IDiscountRepository {
	switch DBType {
	case "postgres":
		return &postgres.DiscountRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}

func CreateUsersRepository(DBType string, ConVar *sqlx.DB) types.IUsersRepository {
	switch DBType {
	case "postgres":
		return &postgres.UsersRepository{
			DB: ConVar,
		}
	case "mysql":
		return nil
	default:
		return nil
	}
}
