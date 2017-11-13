package main

import (
	"github.com/dghubble/oauth1"
	"flag"
	"log"
	"os"
	"github.com/jackmanlabs/errors"
	"encoding/json"
)

func main() {


var configPath	*string = flag.String("config","","The path of the configuration file. Specifying a non-existent file will cause an empty one to be created.")
	flag.Parse()

	if *configPath == ""{
		flag.Usage()
		log.Fatal("A config file path must be specified. If the file you specify does not exist, a new one will be created for you which you may populate.")
	}


	configFile, err := os.Open(*configPath)
	if os.IsNotExist(err){
		configFile,err = os.Create(*configPath)
		if err != nil{
			log.Print(err)
			log.Print("Unable to create the new, empty config file.")
			return
		}
		defer configFile.Close()

		enc := json.NewEncoder(configFile)
		enc.SetIndent("","\t")
		err = enc.Encode(Config{})
		if err != nil{
			log.Print(err)
			log.Print("Failure to write empty configuration to new file.")
			return
		}

		log.Print("New, empty config file successfully created: ", *configPath)
		return

	}else if err != nil{
		log.Fatal(errors.Stack(err))
	}
	defer configFile.Close()

	var config Config
	dec := json.NewDecoder(configFile)
	err = dec.Decode(&config)
	if err != nil{
		log.Print(err)
		log.Print("Failure to read existing configuration.")
		return
	}


	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	oauthToken := oauth1.NewToken(config.Token,config.TokenSecret)

	oauthClient := oauth1.NewClient(nil, oauthConfig,oauthToken)
}


type Config struct{
	ConsumerKey string
	ConsumerSecret string
	Token string
	TokenSecret string
}