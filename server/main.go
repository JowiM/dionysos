// go grpc docs for further reference:
//		- https://grpc.io/docs/tutorials/basic/go.html
//		- https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
//
// Interfaces: https://gobyexample.com/interfaces
//
// protoc generate service: protoc  -I . iquiz.proto --go_out=plugins=grpc:.

package main

import (
	context "context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/JowiM/dionysus/iquiz"
)

type quizInterface struct{}

var ques Questions
var A_rankings Ranking

func main() {
	log.Println("--- Server Starting on port :10101 ---")

	// Setup Server Port
	lis, err := net.Listen("tcp", ":10101")
	if err != nil {
		log.Fatalf("Could not open acceptor on port 10101: %v", err)
	}

	// Variable for interface
	var quiz quizInterface

	// Generate actual server and enable reactor
	qz_server := grpc.NewServer()
	iquiz.RegisterQuizServer(qz_server, quiz)
	log.Fatal(qz_server.Serve(lis))
}

// Implementation of Interfaces

func (quizInterface) List(ctx context.Context, void *iquiz.Void) (*iquiz.QuestionList, error) {
	log.Println("Called: [ List ]")
	return ques.GetQuestions(), nil
}

func (quizInterface) Response(ctx context.Context, req *iquiz.QuizResponse) (*iquiz.Rank, error) {
	log.Println("Called: [ Response ]")
	all_questions := ques.GetQuestions()

	correct := 0
	for _, quest := range all_questions.Questions {
		for _, a := range req.Answers {
			if quest.Id != a.Id {
				continue
			}

			if quest.CorrectAnswer == a.Answer {
				correct++
			}
		}
	}

	result := int((float64(correct) / float64(len(all_questions.Questions))) * 100)

	log.Printf("- Name: %s Points: %d \n", req.Name, result)

	ranking, totalParticipants := A_rankings.AddResult(req.Name, result)
	resp := &iquiz.Rank{
		Name:              req.Name,
		Points:            int32(result),
		Ranking:           int32(ranking),
		TotalParticipants: int32(totalParticipants),
	}

	return resp, nil
}

func (quizInterface) Rankings(ctx context.Context, void *iquiz.Void) (*iquiz.RankingList, error) {
	log.Println("Called: [ Rankings ]")
	return A_rankings.All(), nil
}
