# Fetch Receipt Processor
Containarized Go API for handling customer purrchase reciepts.
## Installation

### Requirements
- Clone of the repository
- Docker

### Setting Up

1. Clone the repository.
2. Execute `docker build . -t go-containerized:latest` (can provide any name inplace of go-containerized ).
3. Execute `docker run -e PORT=9000 -p 9000:8080 go-containerized:latest` (can provide any port number inplace of 9000 ).

The API is up and running. 
### Testing
The program runs programmed test cases while building during setup. But you can manually test by using following:
1. Execute in a seperate terminal in same folder `curl localhost:9000/receipts/process --include --header "Content-Type: application/json" -d @test1.json --request "POST"` (Will return an ID for test1.json ).
2. Execute `curl localhost:9000/receipts/process --include --header "Content-Type: application/json" -d @test2.json --request "POST"` (Will return an ID for test2.json ).
3. Execute `curl localhost:9000/receipts/__ID__/points` ( Replace `__ID__` by ID you recieved from either of the eariler POST request  ). Using ID for test1 should return 28 and ID for test2 should return 109.

You can send any other POST and GET requests on `localhost:9000/receipts/process` and `localhost:9000/receipts/__ID__/points` respectively.

