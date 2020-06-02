package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	var err error
	err = initBashrc("bashrc")
	if err != nil {
		log.Println(err)
	}
	err = createDatabase("pg.sql")
	if err != nil {
		log.Println(err)
	}
	err = createProjectsDir("dirs")
	if err != nil {
		log.Println(err)
	}
	err = createCRTandKey()
	if err != nil {
		log.Println(err)
	}
}

func initBashrc(configPath string) error {
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	bashrcFile, err := os.OpenFile("~/.bashrc", os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer bashrcFile.Close()
	_, err = bashrcFile.WriteString("#export for project Hamster")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		_, err := bashrcFile.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	_, err = exec.Command("source", "~/.bashrc").Output()
	if err != nil {
		return err
	}
	return nil
}

func createDatabase(configPath string) error {
	db, err := sql.Open("postgres", "user="+os.Getenv("POSTGRESUSER")+" password="+os.Getenv("POSTGRESPASSWORD")+" sslmode=disable")
	if err != nil {
		return err
	}
	_, err = db.Exec("create database " + os.Getenv("POSTGRESDB"))
	if err != nil {
		_, err = db.Exec("drop database " + os.Getenv("POSTGRESDB"))
		if err != nil {
			return err
		}
		_, err := db.Exec("create database " + os.Getenv("POSTGRESDB"))
		if err != nil {
			return err
		}
	}
	db.Close()
	db, err = sql.Open("postgres", "user="+os.Getenv("POSTGRESUSER")+" password="+os.Getenv("POSTGRESPASSWORD")+" dbname="+os.Getenv("POSTGRESDB")+" sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		_, err = db.Exec(scanner.Text())
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func createProjectsDir(configPath string) error {
	var err error
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		err = os.Mkdir(scanner.Text(), os.FileMode(0755))
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func createCRTandKey() error {
	var err error
	cmd := exec.Command("ssh-keygen", "-f", "~/hamster/ssl/rsa_id", "-N", "bus")
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("./generate.sh")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func BuildAndRunServices() (err error) {
	build := []string{
		"go build -o $GOPATH/src/github.com/matscus/Hamster/MicroServices/auth/auth $GOPATH/src/github.com/matscus/Hamster/MicroServices/auth/main.go",
		"go build -o $GOPATH/src/github.com/matscus/Hamster/MicroServices/service/service $GOPATH/src/github.com/matscus/Hamster/MicroServices/service/main.go",
		"go build -o $GOPATH/src/github.com/matscus/Hamster/MicroServices/scenario/scenario $GOPATH/src/github.com/matscus/Hamster/MicroServices/scenario/main.go",
		"go build -o $GOPATH/src/github.com/matscus/Hamster/MicroServices/check/check $GOPATH/src/github.com/matscus/Hamster/MicroServices/check/main.go",
		"go build -o $GOPATH/src/github.com/matscus/Hamster/MicroServices/admins/admins $GOPATH/src/github.com/matscus/Hamster/MicroServices/check/admins.go",
	}
	run := []string{
		"hohup $GOPATH/src/github.com/matscus/Hamster/MicroServices/auth/./auth &",
		"hohup $GOPATH/src/github.com/matscus/Hamster/MicroServices/service/./service &",
		"hohup $GOPATH/src/github.com/matscus/Hamster/MicroServices/scenario/./scenario &",
		"hohup $GOPATH/src/github.com/matscus/Hamster/MicroServices/check/./check &",
		"hohup $GOPATH/src/github.com/matscus/Hamster/MicroServices/admins/./admins &",
	}
	for i := 0; i < len(build); i++ {
		_, err = exec.Command("bash", "-c", build[i]).Output()
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(run); i++ {
		_, err = exec.Command("bash", "-c", run[i]).Output()
		if err != nil {
			return err
		}
	}
	return nil
}
