package main

import (
	"regate-core/core/modules"
	"database/sql"
	"fmt"
	"time"

	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	_ "github.com/lib/pq"


	"github.com/spf13/viper"
)

func init() {
	// viper.SetConfigFile(`config.json`)
	
	viper.SetConfigFile(`/home/regate/regate-core/app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Println(loc)
	}
	host := viper.GetString(`database.host`)
	port := viper.GetInt(`database.port`)
	user := viper.GetString(`database.user`)
	password := viper.GetString(`database.pass`)
	dbname := viper.GetString(`database.name`)
	time.Local = loc
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	creds := credentials.NewStaticCredentials(viper.GetString("AWS_ID"), viper.GetString("AWS_SECRET"), "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("sa-east-1"),
		Credentials: creds,
	})
	if err != nil {
		log.Println("FAIL TO CONNECT AWS",err)
	}
	modules.InitModules(db,sess)

	defer db.Close()
}