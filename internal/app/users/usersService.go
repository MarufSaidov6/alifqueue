package users

import "github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"

type UsersService struct {
	repo types.IUsersRepository
}

func InitService(uRep types.IUsersRepository) *UsersService {
	return &UsersService{
		repo: uRep,
	}
}

func (srv *UsersService) GetStoryCash(id, count int) ([]types.StoryCashback, error) {
	return srv.repo.GetStoryCash(id, count)
}

func (srv *UsersService) GetInfoService_Provider(count int, city string) ([]types.InfoServiceProvider, error) {
	return srv.repo.GetInfoService_Provider(count, city)
}

func (srv *UsersService) GetAllServices(id int) ([]types.InfoServiceProvider, error) {
	return srv.repo.GetAllServices(id)
}

func (srv *UsersService) GetListInTouchProduct(id int, city string) ([]types.ListInTouchProduct, error) {
	return srv.repo.GetListInTouchProduct(id, city)
}

func (srv *UsersService) GetListActiveDiscounts(city string, id int) ([]types.ActiveDiscounts, error) {
	return srv.repo.GetListActiveDiscounts(city, id)
}

func (srv *UsersService) GetNewDiscounts(id int, city string) ([]types.ActiveDiscounts, error) {
	return srv.repo.GetNewDiscounts(id, city)
}
