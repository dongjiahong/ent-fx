package sealed

import "gorm.io/gorm"

type Sealed struct {
	Name       string
	DeadLine   int
	Partition  int
	SectorID   int
	MinerID    int
	Ticket     string
	Seed       string
	CachePath  string
	SealedPath string
}

func GetSealed(db *gorm.DB, name string) (*Sealed, error) {
	var s Sealed
	if err := db.Where("name = ?", name).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil

}
