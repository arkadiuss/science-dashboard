package repository


type ICoronavirusRepository interface {
	GetGlobalStats() (int, int, int, error)
}
