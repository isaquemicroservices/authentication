# Microservice to authentication users using golang and grpc

Command to generate protobuf
```go
$ protoc -I . protos/auth/auth.proto --go_out=plugins=grpc:./application
```

### Create folder for config.json file
```bat
$ sudo mkdir /etc/ms-auth
$ sudo touch /etc/ms-auth/config.json
$ sudo cp ./config.json /etc/ms-auth/config.json
$ sudo chmod 777 /etc/ms-auth/config.json
```
if you changed the config.json file, use the command at the bottom to update the config.json file on your computer
```bat
$ sudo cp ./config.json /etc/ms-auth/config.json
```

### command to run the test
```go
$ go test ./... --cover
```

### Command to generate test files
```go
$ go test -coverprofile cover.out 
$ go tool cover -html=cover.out -o cover.html
```

### Create user table in PostgreSQL 
```sql
CREATE TABLE public.t_users (
  id serial4 PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  email varchar(100) NOT NULL,
  passw varchar(100) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NULL
);
```
