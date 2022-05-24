package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"broho-push.go/url_broker"

	"github.com/kirsle/configdir"
	"google.golang.org/grpc"
)

type Config struct {
	BindAddress string `json:"bind-address"`
	BindPort    string `json:"bind-port"`
}

func load_config() Config {
	configPath := configdir.LocalConfig("broho")
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}

	// Deal with a JSON configuration file in that folder.
	configFile := filepath.Join(configPath, "config.json")
	var config Config

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		config = Config{"127.0.0.1", "50051"}
		fh, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		encoder := json.NewEncoder(fh)
		encoder.Encode(&config)
	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		decoder := json.NewDecoder(fh)
		decoder.Decode(&config)
	}
	return config
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(filepath.Base(os.Args[0]), "requires at least one URL as argument")
		os.Exit(1)
	} else {
		var config Config = load_config()
		var endpoint string = config.BindAddress + ":" + config.BindPort

		var channel *grpc.ClientConn
		channel, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Could not connect: %s\n", err)
		}
		defer channel.Close()

		client := url_broker.NewURLBrokerServiceClient(channel)
		re := regexp.MustCompile(`https?://.*`)

		for _, arg := range os.Args[1:] {
			if re.MatchString(arg) {
				fmt.Printf("Pushing URL %s\n", arg)
				response, err := client.PushURL(context.Background(), &url_broker.URL{Url: arg})
				if err != nil {
					fmt.Printf("Error occurred when pushing URL %s: %s\n", arg, err)
				} else if response.ResponseCode != 0 {
					fmt.Printf("Error occurred when pushing URL %s (response code: %d)\n", arg, response.ResponseCode)
				}
			} else {
				fmt.Printf("Invalid URL: %s\n", arg)
			}
		}
	}
}

func URLBrokerServiceClient(channel *grpc.ClientConn) {
	panic("unimplemented")
}
