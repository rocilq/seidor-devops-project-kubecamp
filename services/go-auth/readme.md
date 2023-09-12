- To run local `docker-compose up` and use the postman collection
  After finish recommended `docker-compose down -v --rmi all`

- For local testing add this env value `export MONGO_URI=<YOUR TEST MONGO INSTANCE CONNECTION STRING>` and run `go test ./...` 

- To run tests in docker run `docker-compose -f docker-compose.t.yaml up`
  After finish recommended `docker-compose -f docker-compose.t.yaml down -v --rmi all`