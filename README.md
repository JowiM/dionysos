#  Experimenting with GO

## Server

Run Server first to start playing. The server listens on port 10101.

	1. go build
	2. ./server

## Client

To compile the application simply execute 'go build'. Once server is turned off all rankings are lost

Client is a cli app with 3 commands:

	1. './client getQA' Retrieves all questions with corresponding answer
	2. './client play' Allows the user to enter name and play game. When all questions answered will give a small report.
	3. '.client ranking' Lists all the rankings  