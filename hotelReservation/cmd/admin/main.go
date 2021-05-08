package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/harlow/go-micro-services/registry"
	"github.com/harlow/go-micro-services/services/admin"
	"github.com/harlow/go-micro-services/tracing"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	// initializeDatabase()
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)

	mongo_session := initializeDatabase(result["AdminMongoAddress"])
	defer mongo_session.Close()
	serv_port, _ := strconv.Atoi(result["AdminPort"])
	serv_ip := result["AdminIP"]

	fmt.Printf("admin ip = %s, port = %d\n", serv_ip, serv_port)

	var (
		// port       = flag.Int("port", 8086, "The server port")
		jaegeraddr = flag.String("jaegeraddr", result["jaegerAddress"], "Jaeger server addr")
		consuladdr = flag.String("consuladdr", result["consulAddress"], "Consul address")
	)
	flag.Parse()

	tracer, err := tracing.Init("admin", *jaegeraddr)
	if err != nil {
		panic(err)
	}

	registry, err := registry.NewClient(*consuladdr)
	if err != nil {
		panic(err)
	}

	srv := &admin.Server{
		Tracer:       tracer,
		Registry:     registry,
		Port:         serv_port,
		IpAddr:       serv_ip,
		MongoSession: mongo_session,
	}
	log.Fatal(srv.Run())
}
