# Microservices with Go-kit and gRPC-gateway
In this project, i create a mini micro-services system, which has 2 service : **Identity** and **User**  
Identity Svc has 2 endpoints : **Insert User into database 1** (db1) & **Get user's information from database 1** (db1)  
Users Svc has only 1 endpoint : **Sync data**, this endpoint call to Identity Svc to get user's information, if user's infor is exist, User svc will **write to database 2** (db2)  
All the endpoints are called via Grpc  
And i also create a grpc gateway, it works as a reverse proxy, which convert http request to grpc request and vice versa  

## How to run this project:     
**Step 1:** Run the Docker container by `docker-compose up -d`  
**Step 2:** Migrate database using **Goose** by command `make g-up` (Move to folder directory first)  
**Step 3:** Run main.go file in folder cmd  
**Step 4:** Run the gateway, you need to change directory to folder /gateway and run the command `go run .`  
**Step 4.5:** Run the proxy, you need to change directory to folder /proxy and run the command `go run .` (No required)  
**Step 5:** Testing endpoint in Postman :
With gRPC endpoints, open Postman, create a new gRPC request, import file .proto and choose method, you can "Use Example Message" to get a right format of the request's body  
(default records in db1 : admin-admin, user-user;  
in db2 : admin2-admin2, user2-user2)  
With HTTP endpoints, there are 3 endpoints, you can test them by Postman or curl command:  
**If you start the gateway in each service**  
1. curl --location 'http://localhost:8080/insert' \
   --header 'Content-Type: application/json' \
   --data '{
   "username" : "gateway",
   "password" : "gateway-psw"
   }'
2. curl --location 'http://localhost:8080/get/admin' (you can change 'admin' to username you want to get)
3. curl --location 'http://localhost:8088/sync/gateway' (same with 'gateway', you can change to another username)


**If you start a **common gateway in folder /gateway**, you can test the endpoint by Postman, the endpoint is :**  
    http://localhost:8089/v1/insert  
    http://localhost:8089/v1/get/admin  
    http://localhost:8089/v1/sync/gateway  


- Docker exec postgres    
docker exec -it db1 psql -U postgres -d db1  
-> select * from db1; to get data from database db1

docker exec -it db2 psql -U postgres -d db2  
-> select * from db2; to get data from database db2

Library in use :  
+ Go kit : https://github.com/go-kit/kit  
+ Goose : https://github.com/pressly/goose  
+ Grpc-gateway : https://github.com/grpc-ecosystem/grpc-gateway  
+ Postgres : https://github.com/lib/pq  
