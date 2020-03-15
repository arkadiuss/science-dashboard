package service

import (
	"github.com/arkadiuss/science-dashboard/repository"
	"github.com/arkadiuss/science-dashboard/models"
	"time"
	"fmt"
)

type IDashboardService interface {
	GetDashboard(string) (models.Dashboard, error)
}

type DashboardService struct {
	sunRepository repository.ISunRepository
	locationRepository repository.ILocationRepository 
	iissRepository repository.IISSRepository
	coronavirusRepository repository.ICoronavirusRepository
}

func GetDashboardService(sr repository.ISunRepository, lr repository.ILocationRepository, ir repository.IISSRepository, cr repository.ICoronavirusRepository) *DashboardService {
	return &DashboardService { sr, lr, ir, cr }
}

func (ds *DashboardService) getCoordinates( location string, locCh chan models.Location, errCh chan error) {
	loc, err := ds.locationRepository.GetCoordinates(location)
	if err != nil {
		errCh <- err
	} else {
		locCh <- loc
	}
}


func (ds *DashboardService) getSunriseSunset(location models.Location, ssCh chan struct {sunrise string; sunset string}, errCh chan error) {
	sunrise, sunset, err := ds.sunRepository.GetSunriseSunset(location)
	if err != nil {
		errCh <- err
	} else {
		ssCh <- struct { sunrise string; sunset string }{ sunrise, sunset }
	}
}

func (ds *DashboardService) getISSLocation(locCh chan models.Location, errCh chan error) {
	loc, err := ds.iissRepository.GetCurrentLocation()
	if err != nil {
		errCh <- err
	} else {
		locCh <- loc
	}
}

func (ds *DashboardService) getISSPasses(location models.Location, passCh chan []time.Time, errCh chan error) {
	ps, err := ds.iissRepository.GetNextPasses(location)
	if err != nil {
		errCh <- err
	} else {
		passCh <- ps
	}
}

func (ds *DashboardService) getCoronavirusStats(stCh chan struct { cases int; deaths int; recovered int}, errCh chan error) {
	a, d, r, err := ds.coronavirusRepository.GetGlobalStats()
	if err != nil {
		errCh <- err
	} else {
		stCh <- struct {cases int; deaths int; recovered int} { a, d, r }
	}
}

func (ds *DashboardService) GetDashboard(location string) (models.Dashboard, error) {
	d := models.Dashboard{}
	errCh := make(chan error, 5)
	locCh := make(chan models.Location)
	passCh := make(chan []time.Time)
	ssCh := make(chan struct {sunrise string; sunset string})
	stCh := make(chan struct {cases int; deaths int; recovered int})
	var err error
	go ds.getCoordinates(location, locCh, errCh)
	select {
		case d.Location = <-locCh: {}
		case err = <-errCh: {
			fmt.Println(err)
			return models.Dashboard {}, err
		}
	}
	go ds.getSunriseSunset(d.Location, ssCh, errCh)
	go ds.getISSLocation(locCh, errCh)
	go ds.getISSPasses(d.Location, passCh, errCh)
	go ds.getCoronavirusStats(stCh, errCh)
	for i:=0;i<4;i++ {
		select {
			case ss := <-ssCh: {
				d.Sunrise = ss.sunrise
				d.Sunset = ss.sunset
			}
			case d.ISSLocation = <-locCh: {}
			case passes := <-passCh: {
				d.ISSNextPass = int(passes[0].Sub(time.Now())/1e9/60)
			}
			case st := <-stCh: {
				d.CoronavirusActiveCases = st.cases - (st.deaths + st.recovered)
				d.CoronavirusDeathRecoveredRatio = float64(st.deaths)/float64(st.recovered)
			}
			case err = <-errCh: {
				fmt.Println(err)
			}
		}
	}
	return d, err
}

