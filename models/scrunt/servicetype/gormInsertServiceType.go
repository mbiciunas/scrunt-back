package servicetype

import (
	"scrunt-back/models/scrunt"
)

func GormInsertServiceType(uuid string, name string, iconCode string, descShort string, descLong string) (uint, error) {
	serviceType := scrunt.ServiceType{
		Uuid:      uuid,
		Name:      name,
		IconCode:  iconCode,
		DescShort: descShort,
		DescLong:  descLong,
	}

	if err := scrunt.GormDB.Create(&serviceType).Error; err != nil {
		return 0, err
	}

	return serviceType.Id, nil
}
