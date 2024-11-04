In this project, i create a mini micro-services system, which has 2 service : **Identity** and **User**  
Identity Svc has 2 endpoints :**Insert User into database 1** (db1) & **Get user's information from database 1** (db1)  
Users Svc has only 1 endpoint : **Sync data**, this endpoint call to Identity Svc to get user's information, if user's infor is exist, User svc will **write to database 2** (db2)  
All the endpoints are called via Grpc  

How to run this project:  
Step 1: Run the Docker container by "docker-compose up -d"  
Step 2: Migrate database using **Goose** by command "make g-up" (Move to folder directory first)  
Step 3: Run main.go file in folder cmd  
Step 4: Testing endpoint in Postman : open Postman, create a new gRPC request, import file .proto and choose method, you can "Use Example Message" to get a right format of the request's body  
(default records in db1 : admin-admin, user-user;  
in db2 : admin2-admin2, user2-user2)  

- Docker exec postgres  
docker exec -it db1 psql -U postgres -d db1  

