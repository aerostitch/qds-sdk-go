# Qubole Data Services SDK for Go

## Disclaimer

This is a SDK I'm currently writing for my own usage for the Qubole Data
Services API. It is not a SDK supported by Qubole and I don't work for Qubole.
I just happened to use QDS and wanted to play with their API in Go.

For offically supported SDK by Qubole, see their Java and Python SDK here:
https://github.com/qubole

The documentation about the REST API itself can be found at:
http://docs.qubole.com/en/latest/rest-api/

NOTE:
 For now this repo just contains a draft of the SDK I just begun. I use the repo
 as a way to sync my share my code between my computers.

## Installing

## Using

### Basic example

Script that will list all the schedules you currently have in the Qubole account
associated with the token you will provide.

```Go
package main

import (
  "github.com/aerostitch/qds-sdk-go"
  "flag"
  "log"
)

// Those are the arguments you can pass to the script
var (
  verbose = flag.Bool("verbose", false, "Set a verbose output")
  // Token you should use for connecting to the qubole API
  token = flag.String("auth-token", "", "Token you should use for connecting to the qubole API")
)

func main() {

  flag.Parse()

  // Checking if the tone
  if len(*token) <= 0 {
    log.Fatal("You have to provide a token for the qubole API using the --auth-token parameter!")
  }

  schedules := qds.GetAllSchedules(token)
  log.Printf("****************************************")
  log.Printf("Response: %+v", *schedules)

}
```

## Configuring credentials

## Contributing

## License

This SDK is distributed under the Apache License, Version 2.0, see LICENSE file
for more informations.
