Passargad User is a simple base code that has 5 basic APIs:
create user
update user
delete user
login
get user info

to run this project you need to run commands as well:
make proto
make config
docker compose up

API Postman Json files are in postman folder and call APIS.

Project template:
this project supports both GRPC and rest call and there is a delivery layer that
validates Inputs and in case of business logic calls usecase layer that handles logic.
also in entities package database models and request models are defined.
all configs are set in config package that loads config.json file.
middlewares are defined in app/middleware directory.
database queries are defined in repository layer.
this project uses postgresql as database, uses jaeger for tracing and
uses jwt for authentication.
GRPC protobuf request response models are defined in proto directory.