# Fix makefile thing

# Installed gorilla/mux for server 
go get github.com/gorilla/mux 

# Run postgress in docker
docker run --name some-postgres -e POSTGRES_PASSWORD=goapi -p 5432:5432 -d postgres

# Connect to that postgress docker machine
telnet localhost 5432