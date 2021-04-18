# Snarl Server and Client

## Running the server
To run the server, extract the executable from the GitHub release in the directory of your choosing.
You should then be able to run the executable with the expected flags in the command line. The usage is
as follows:
```
  -address string
        tells the server what ip address to listen on (default "127.0.0.1")
  -clients int
        tells the server how many clients to wait for. Default is 4 (default 1)
  -levels string
        tells the server which levels file to use. Default is ./snarl.levels (default "snarl.levels")
  -observe
        launches a local observer if toggled
  -port int
        tells the server what por to listen on (default 45678)
  -wait int
        used to determine the amount of time to wait for players to register from booting the server (default 60)
```

## Running the client
The client should be extracted, like the server, and then run from the command line. It can be given a port and
address as arguments:
```
  -address string
        tells the client what IP address to connect over (default "127.0.0.1")
  -port int
        tells the client what port to connect over (default 45678)
```

### Using the client
To use the client, make sure that the terminal is in focus. Use the arrow keys to adjust the move that you would
like to make. When you have selected a move (the terminal will tell you what the current selected move is), press
enter.