# Study Microservices Architecture
A microservice architecture – a variant of the service-oriented architecture
structural style – is an architectural pattern that arranges an application as a collection
of loosely-coupled, fine-grained services, communicating through lightweight protocols

## Architecture Diagram

![Architecture Diagram](./docs/alivers-services.png)

## Prequisite
  ### Auth Service
  - NodeJs v18.13
  - Yarn
  - Mongodb

  ### Users Service
  ### Gateway
  - Go
  - Go swagger

## How to run
  ### Auth Service
  - Install all dependencies `Yarn install`
  - Run auth service `yarn start`
  - try to test grpc with postman grpc with url `localhost:50051`

  ### Users Service
  ### Gateway
  - `make all`
  - `make run-local`
  - try to test with url `localhost:8080/health`
