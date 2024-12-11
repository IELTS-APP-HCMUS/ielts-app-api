package services

import (
	"context"
)

func (s Service) GetMasterData(ctx context.Context) (err error) {
	// filter := []repositories.Clause{
	// 	func(tx *gorm.DB) {
	// 		tx.Where("status = 1")
	// 	},
	// }
	// data, err := s.masterConfigPgRepo.List(ctx, models.QueryParams{}, filter...)
	// if err != nil {
	// 	return
	// }
	// reps := models.MasterConfigMapping{}
	// for _, d := range data {
	// 	reps[d.PublicId] = d.Value
	// }
	return nil
}
