# Hamster
## Backend from system automation perfomance testing


used as a  frontend  https://github.com/kolesnikovm/automation-front

used as a database postgres

## Install 

### install postgres

`docker run --name postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres`

then you can create a user and a database of your choice, or use the default values.

### Customise project config
By default, the creation of project directories is done in the user directory, if you need to change the location,
then you need to edit the bashrc and dirs project files, specifying the paths you need.

### Rolling
auto create project dirs, generate cert and rsa key, create DataBase structs, build and run all services/
and create default user:
username: god
password: Ab123456

`cd cd $GOPATH/src/github.com/matscus/Hamster`
`go run Rolling.go`

### Install apache jmeter

`cd ~/hamster/guns`
`wget https://apache-mirror.rbc.ru/pub/apache//jmeter/binaries/apache-jmeter-5.3.tgz`
`tar  -xvzf  apache-jmeter-5.3.tgz && rm apache-jmeter-5.3.tgz`

### Check install
sent users data from auth service. password in base64
`curl --insecure -d '{"user": "god","password": "QWIxMjM0NTY="}' -X POST https://172.18.0.1:10000/api/v1/auth/new `
