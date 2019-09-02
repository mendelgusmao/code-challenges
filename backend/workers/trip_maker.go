package workers

import (
	"log"
	"sort"
	"time"

	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/services"
	"go.etcd.io/bbolt"
)

type tripMaker struct {
	carsService     *services.CarsService
	journeysService *services.JourneysService
	tripsService    *services.TripsService
	interval        time.Duration
	start           chan bool
	stop            chan bool
}

func newTripMaker(db *bbolt.DB, interval time.Duration) *tripMaker {
	return &tripMaker{
		carsService:     services.NewCarsService(db),
		journeysService: services.NewJourneysService(db),
		tripsService:    services.NewTripsService(db),
		interval:        interval,
		start:           make(chan bool),
		stop:            make(chan bool),
	}
}

func (tm *tripMaker) Work() {
	log.Println("tripworker is waiting for command")

	for {
		select {
		case <-tm.start:
			go tm.tick()
		}
	}
}

func (tm *tripMaker) Start() {
	log.Println("starting tripmaker")
	tm.start <- true
}

func (tm *tripMaker) Stop() {
	log.Println("stopping tripmaker")
	tm.stop <- true
}

func (tm *tripMaker) tick() {
	ticker := time.NewTicker(tm.interval)

	for {
		select {
		case <-ticker.C:
			tm.makeTrips()
		case <-tm.stop:
			log.Println("tripmaker stopped")
			break
		}
	}
}

func (tm *tripMaker) makeTrips() error {
	idleJourneys, err := tm.loadIdleJourneys()

	if err != nil {
		return err
	}

	if len(idleJourneys) == 0 {
		log.Println("tripmaker: no idle journeys found. skipping")
		return nil
	}

	availableCars, err := tm.loadAvailableCars()

	if err != nil {
		return err
	}

	availableSeats := make([]int, len(availableCars))
	index := 0

	for seats := range availableCars {
		availableSeats[index] = seats
	}

	sort.Sort(sort.IntSlice(availableSeats))
	allocatedCars := make(map[int]bool)

	for _, journey := range sortableJourneys(idleJourneys) {
	AvailableSeats:
		for seats := range availableSeats {
			cars := availableCars[seats]

			for _, car := range cars {
				_, unavailable := allocatedCars[car.ID]

				if !unavailable && journey.People <= seats {
					log.Printf("tripmaker is making a trip for car#%d and journey#%d\n", car.ID, journey.ID)

					tm.tripsService.Insert(services.Trip{
						CarID:     car.ID,
						JourneyID: journey.ID,
					})

					allocatedCars[car.ID] = true

					break AvailableSeats
				}
			}
		}
	}

	return nil
}

func (tm *tripMaker) loadIdleJourneys() ([]services.Journey, error) {
	idleJourneys := make([]services.Journey, 0)
	journeys, err := tm.journeysService.All()

	if err != nil {
		return nil, err
	}

	for _, journey := range journeys {
		if _, err := tm.tripsService.FindByJourneyID(journey.ID); err != nil {
			idleJourneys = append(idleJourneys, journey)
		}
	}

	return idleJourneys, nil
}

func (tm *tripMaker) loadAvailableCars() (map[int][]services.Car, error) {
	cars, err := tm.carsService.All()

	if err != nil {
		return nil, err
	}

	partitions := make(map[int][]services.Car)

	for _, car := range cars {
		availableSeats := car.Seats

		if trips, err := tm.tripsService.FindByCarID(car.ID); err != nil {
			for _, trip := range trips {
				journey, err := tm.journeysService.Find(trip.JourneyID)

				if err != nil {

				}

				availableSeats -= journey.People
			}
		}

		if availableSeats > 0 {
			if _, ok := partitions[availableSeats]; !ok {
				partitions[availableSeats] = make([]services.Car, 0)
			}

			partitions[availableSeats] = append(partitions[availableSeats], car)
		}
	}

	return partitions, nil
}
