package discount

import (
	"errors"

	"github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"
)

type DiscountService struct {
	repo types.IDiscountRepository
}

func InitService(aRep types.IDiscountRepository) *DiscountService {
	return &DiscountService{
		repo: aRep,
	}
}

func (srv *DiscountService) CreateDiscount(dis types.Discount) error {
	if srv.repo.DiscountExist(dis.Name, dis.ServiceProviderID) {
		return errors.New("DISCOUNT ALREADY EXISTS!")
	}
	return srv.repo.CreateDiscount(dis)
}

func (srv *DiscountService) DiscountsList() ([]types.Discount, error) {
	return srv.repo.DiscountsList()
}

func (srv *DiscountService) DiscountsActiveList() ([]types.Discount, error) {
	return srv.repo.DiscountsActiveList()
}

func (srv *DiscountService) DiscountsSoonList() ([]types.Discount, error) {
	return srv.repo.DiscountsSoonList()
}

func (srv *DiscountService) DiscountsPastList(pastDis types.PastDiscount) ([]types.Discount, error) {
	return srv.repo.DiscountsPastList(pastDis)
}
