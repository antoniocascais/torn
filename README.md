# Intro

This program will call Torn's API each 30 seconds to check if the chain timer is below 200 seconds.
If that condition is met, it will execute $PROGRAM with args $PROGRAM_ARGS.

# Usage
You can download the binary form `releases` or compile the code yourself: 
`go build -o main main.go`

To run it, execute `./main API_KEY PROGRAM PROGRAM_ARGS`

* API_KEY: your torn api key
* PROGRAM: the program that is going to be executed if the chain timer is below 200 seconds
* PROGRAM_ARGS: the args to be passed to the program

For example, if you want to ring an alarm when the chain timer is low:
`./main APYKEY paplay /usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga`
