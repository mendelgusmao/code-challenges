package workers

import "gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/services"

type sortableJourneys []services.Journey

func (j sortableJourneys) Len() int {
	return len(j)
}

func (j sortableJourneys) Swap(a, b int) {
	j[a], j[b] = j[b], j[a]
}

func (j sortableJourneys) Less(a, b int) bool {
	return j[a].People < j[b].People
}
