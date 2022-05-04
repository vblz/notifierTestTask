# Notifier test task
## The task
### The Library
Write a library that implements an HTTP notification client. A client is configured with a URL to which notifications are sent. It implements a function that takes messages and notifies about them by sending HTTP POST requests to the configured URL with the message content in the request body. This operation should be non-blocking for the caller.

A great number of messages might arrive at once, so make sure to handle spikes in notification activity and don’t overload the event-handling service or exhaust your file descriptors. But be efficient and don’t just send requests serially.

Allow the caller to handle notification failures in case any requests should fail.

### The Executable
Write a small program that uses the library above. It should read stdin and send new messages every interval (should be configurable). Each line should be interpreted as a new message that needs to be notified about.

The program should implement graceful shutdown on SIGINT.

Example usage information for clarification purposes (the solution doesn’t have to reproduce this output):
```
usage: notify --url=URL [<flags>]

Flags:
    --help              Show context-sensitive help (also try --help-long and --help-man).
    -i, --interval=5s   Notification interval
```

Example call:
```
$ notify --url http://localhost:8080/notify < messages.txt
```

## Implementation
### Assumptions
* The task doesn't say anything about server API. Assumed that any successful (2xx) response status code is good, anything else means error
* Consider an empty message correct
* Don't trim lines from STDIN, just take them as is line by line in the executable
* The executable works with string messages, but the library is made for more generic io.Reader
* Nothing is said on cost of loosing messages, in some cases probably it would be better to use some kind of persistent store, but here in-memory store is used and when app is closed all not sent messages are lost
* The order of messages is not guaranteed in the library
* As the task is expected as a ZIP archive, no CI is made, only Makefile

### Notes on the implementation
* The requirement to make a library not a package heads to two go modules. Single module would be easier in some ways
* The library implements a retry policy, the number of retries can be set
* The rate limit in RPS can be set as well
* Notify takes io.Reader, that lets any reader be passed through without reading in memory array or slice. For example, you can pass a HTTP request's body to the library directly
* In my opinion usually it's better to make APIs blocking and let caller decide how to call it 
