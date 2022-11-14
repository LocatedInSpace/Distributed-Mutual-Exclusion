# Distributed-Mutual-Exclusion
BSDISYS1KU-20222, mandatory handin 4

##  How to run

 1. Run `main.go` in separate terminals - the amount of terminals you need to open depends on the value of `const CLIENTS` (default is 3). `main.go` takes one parameter, which is the port to use - however, by default `const OFFSET` is set to 7000, meaning that the parameter has this added. 

 We therefore recommend simply using the following commands (in three separate terminals)

    ```
    $ go run main.go 0
    $ go run main.go 1
    $ go run main.go 2
    ```

Which would use port 7000, 7001, and 7002. You must not use the same port for multiple peers.

2. If you wish to only see stuff related to whether or not a peer is in a critical section, then simply set `const VERBOSE = false`. 

3. A peer will create a log.txt in the root directory under the name `peer(*)-log.txt`, with `*` being the bound port.



##  Stuff that might go wrong
The program crashes with an error regarding the index used for `p.clients` not existing/valid. While this is very unlikely, due to the sleep on line 129 - if it does happen, try extending the sleep untill each peer manages to have fully dialed up a connection to each other peer. This can be verified by making sure you see the message `Connected to all clients :)` on each peer, before the first gRPC call.