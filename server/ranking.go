package main

import (
	"log"
	"sort"

	"github.com/satori/go.uuid"

	"github.com/JowiM/dionysus/iquiz"
)

type Rank struct {
	uuid   string
	name   string
	points int
}

type Ranking struct {
	Rankings []Rank
}

func (r *Ranking) AddResult(name string, result int) (ranking int, all int) {
	log.Println(" *ADDING TO RANK* ")

	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Could not generate UUID: %v", err)
	}

	ra := Rank{
		uuid:   u2.String(),
		name:   name,
		points: result,
	}

	// Order Ranking according to points
	r.Rankings = append(r.Rankings, ra)
	sort.Slice(r.Rankings, func(i, j int) bool {
		return r.Rankings[i].points > r.Rankings[j].points
	})

	// Find Ranking in array
	rank_list := 0
	for i, val := range r.Rankings {
		if val.uuid == ra.uuid {
			rank_list = i + 1
		}
	}

	return rank_list, len(r.Rankings)
}

func (r *Ranking) All() *iquiz.RankingList {
	resp := &iquiz.RankingList{}

	total := len(r.Rankings)
	for i, val := range r.Rankings {

		rank := &iquiz.Rank{
			Name:              val.name,
			Points:            int32(val.points),
			Ranking:           int32(i),
			TotalParticipants: int32(total),
		}

		resp.Rankings = append(resp.Rankings, rank)
	}

	return resp
}
