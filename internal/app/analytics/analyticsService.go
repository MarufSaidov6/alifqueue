package analytics

import "github.com/AlifAcademy/ClientLoyaltyProgram/internal/pkg/types"

type AnalyticsService struct {
	repo types.IAnalyticsRepository
}

func InitService(aRep types.IAnalyticsRepository) *AnalyticsService {
	return &AnalyticsService{
		repo: aRep,
	}
}

/*func (srv *AnalyticsService) TransactionAmount() error {
	return srv.repo.TransactionAmount()
}

func (srs *AnalyticsService) GetMostHihgCashBack(count int) ([]types.NewServiceProvider, error) {
	return srs.repo.GetMostHihgCashBack(count)
}
*/
