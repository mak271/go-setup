package stat

import (
	"go/adv-demo/pkg/db"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Database: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	// если нет статистики за сегодня по ссылке - создаём
	// если есть - увеличиваем на 1
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Database.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Database.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Database.Save(&stat)
	}
}

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	query := repo.Database.Table("stats").
		Select(selectQuery).
		Session(&gorm.Session{})

	// Сессию можно использовать для динамического добавления параметров запроса

	// if true {
	// 	query.Where("count > 10")
	// }

	query.Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
