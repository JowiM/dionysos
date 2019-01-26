// go grpc docs for further reference: 
//		- https://grpc.io/docs/tutorials/basic/go.html
//		- https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
//
// Interfaces: https://gobyexample.com/interfaces
//
// protoc generate service: protoc  -I . iquiz.proto --go_out=plugins=grpc:.


package main 

import(
	"fmt"
	"log"
	"net"
	context "context"

	"google.golang.org/grpc"

	"iquiz"
	//"https://github.com/JowiM/dionysus/tree/master/src/iquiz"
)

type quizInterface struct{}

func main() {
	fmt.Println( "--- Server Starting ---")

	// Setup Server Port
	lis, err := net.Listen("tcp", ":10101")
	if err != nil {
        log.Fatalf("Could not open acceptor on port 10101: %v", err)
	}

	// Variable for interface
	var quiz quizInterface

	// Generate actual server and enable reactor
	qz_server := grpc.NewServer()
	iquiz.RegisterQuizServer(qz_server,quiz)
	log.Fatal( qz_server.Serve(lis) )
}

// Implementation of Interfaces

func (quizInterface) CheatSheet( ctx context.Context, void *iquiz.Void) (*iquiz.QuestionList, error) {

	question := &iquiz.Question{
		Id: 1,
		Question:  "Who is the only member of ZZ Top who doesn’t have a beard?",
	 	CorrectAnswer: "Frank Beard",
	 	AnswerOptions: []string{ "Jimmy Page", "Frank Zampa"},
	}

	question_list := &iquiz.QuestionList{};
	question_list.Questions = append( question_list.Questions, question )

	return question_list, nil
}