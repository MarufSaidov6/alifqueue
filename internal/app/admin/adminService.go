package admin

import "github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"

type AdminService struct {
	repo types.IAdminRepository
}

func InitService(adRep types.IAdminRepository) *AdminService {
	return &AdminService{
		repo: adRep,
	}
}
