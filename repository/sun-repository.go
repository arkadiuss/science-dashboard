package repository

type ISunRepository interface {
	GetSunriseSunset() (int, int)
}

type SunRepository struct {

}

func (sr SunRepository) GetSunriseSunset() (int,int) {
	return 3,4
}
