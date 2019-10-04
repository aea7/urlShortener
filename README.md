# URL shortener HTTP service

Required libraries :

go get github.com/lib/pq
go get github.com/gorilla/mux

or

go get github.com/gorilla/mux github.com/lib/pq 

Then, just run:

go run server.go

http://localhost:8080?url=www.google.com - Shorten url endpoint
-

http://localhost:8080/HcOdpM - Redirect url endpoint
-

To simulate concurrent connections with ApacheBench I ran:

ab -n 20000 -c 200 "127.0.0.1:8080/?url=www.google.com" and it showed no problems.

Also included 3 unit tests inside test/api_test.go
To run these test under test folder run:

go test
-