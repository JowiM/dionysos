package main

import (
	"fmt"
	"sort"
	"log"

	"github.com/satori/go.uuid"

	"iquiz"
)

type Rank struct {
	uuid string
	name string
	points int
}

type Ranking struct {
	Rankings []Rank
}

func (r *Ranking) AddResult( name string, result int ) ( ranking int, all int ) {
	fmt.Printf( "---- ADD RANKING")

	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Could not generate UUID: %v", err)
	}

	ra := Rank{
		uuid: u2.String(),
		name: name,
		points: result,
	}

	r.Rankings = append( r.Rankings, ra )
	sort.Slice(r.Rankings, func(i, j int) bool {
		fmt.Printf( " Points I: %d - Points J: %d \n", r.Rankings[i].points, r.Rankings[j].points )
  		return r.Rankings[i].points < r.Rankings[j].points 
	})


	rank_list := 0
	for i, val := range r.Rankings {
		if val.uuid == ra.uuid {
			fmt.Println( "--- FOUND RANKING %d \n", i)
			rank_list = i + 1
		}
	}

	return rank_list, len(r.Rankings)
}

func (r *Ranking) All() (*iquiz.RankingList) {
	resp := &iquiz.RankingList {}

	total := len(r.Rankings)
	fmt.Printf( "--- TOTAL: %d \n", total)
	for i, val := range r.Rankings {

		rank := &iquiz.Rank{
			Name: val.name,
			Points: int32(val.points),
			Ranking: int32(i),
			TotalParticipants: int32(total),
		}
		
		resp.Rankings = append( resp.Rankings, rank )
	}

	return resp
}