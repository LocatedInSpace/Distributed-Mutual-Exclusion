# Distributed-Mutual-Exclusion
BSDISYS1KU-20222, mandatory handin 4

##  How to run

 1. Run `main.go` in separate terminals - the amount of terminals you need to open depends on the value of `const CLIENTS` (default is 3). `main.go` takes one parameter, which is the port to use - however, by default `const OFFSET` is set to 7000, meaning that the parameter has this added. 

 We therefore recommend simply using the following commands (in three separate terminals)

    ```console
    $ go run main.go 0
    $ go run main.go 1
    $ go run main.go 2
    ```
    
Which would use port 7000, 7001, and 7002.

2. If you wish to only see stuff related to whether or not a peer is in a critical section, then simply set `const VERBOSE = false`. 

3. Chat back and forth (if multiple clients have been opened) - if you are an active chatter thats forgotten what an original question was, you can use `Up/Down` arrowkeys, or scrollwheel, to scroll in the chat.
>![Scrolling](imgs/scrolling.gif)
A new message appearing will automatically put you back down to the bottom of the chat - this setting cannot be disabled (unless you change the code)

4. Exiting/leaving the chat can be done by pressing `Control+C`


##  Stuff that might go wrong
>![Initial UI](imgs/initial_ui.png)
There is a chance that upon running the client, the UI will look all scrambled. This is solved by quickly resizing the UI, after which it will stay "normal".

Colors are not uniform between every terminal - in the ones tested, PowerShell and Windows Terminal (old & new), yellow was the main culprit of weirdness. This is, like the previous issue, a problem caused by termui, rather than any codebase issue.

The program may also crash arbitrarily upon resizing.

Long messages of text (multiline) become a little weird when scrolling, as they are treated as a single row still by termui - we've rewritten the rendering logic for list to attempt to remedy this, but it might fail for cases we haven't thought of. The chat will also feel slightly jittery with multiline messages, we decided to accept the "solution" we came up with, since the goal of the assignment wasn't UI :)