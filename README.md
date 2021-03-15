## Description

The assessment consists of an API to be used for opening a new “current account” of already existing
customers.

Requirements
- The API will expose an endpoint which accepts the user information (customerID,
initialCredit).
- Once the endpoint is called, a new account will be opened connected to the user whose ID is
customerID.
- Also, if initialCredit is not 0, a transaction will be sent to the new account.
- Another Endpoint will output the user information showing Name, Surname, balance, and transactions of the accounts.

Bonuses
- Accounts and Transactions are different services.
- Frontend (simple one is OK).
- Attention to CI/CD

## Deploy instractions

### Dependency
You need docker and docker compose be available on your machine

Versions on my machine are
```bash
docker -v
Docker version 20.10.5, build 55c4c88
docker-compose -v
docker-compose version 1.26.2, build eefe0d31
```
I'm sure it is not mandatory to have exactly this version, but check that
- your docker compose supports version: '3.6'
- both docker and docker compose are available for current user, otherwise you should run deploy with respect to sudo user

### Deploy
First run you need to run ```make``` it will run 4 steps:
- check that docker and docker-compose are available
- copy .env.example to .env file (if .env is not already exist in dir)
- build app and front docker file
- run both app and front apps
Other run you only up command as everything is already build: ```make up```

To avoid situation when port is already allocated on your machine you should create (edit) .env file and set available port values

After successful deploy frontend application will be available by http://localhost:{FRONT_HOST_MACHINE_PORT}