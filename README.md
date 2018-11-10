# Port Scanner

Example for creating a naive TCP port scanner.

## What is this 'Naive' ?

Port scanning is a way to detect if there is a server that listen to a port
at a given address.

The current examples are checking TCP connections. The reason for scanning TCP
is simple - TCP have a mechanism to know if a port is open or not.
That mechanism is implemented using SYNchronize and called 3 way handshake.

### 3 way handshake

When reading network specifications, there are at least two characters involved
and I'm not going to stop doing that, so meet Alice and Bob :-)

Alice wishes to contact with Bob, on a given channel (port), so Alice sends `SYN`.
Bob, if wishes, answer Alice with `SYN-ACK`.
Then Alice sends `ACK`, and there is a handshake between the two, and got the
name of 3 way handshake

### TCP vs UDP

3 way handshake exists for TCP but not for UDP. UDP does not sends a handshake
and if the server answers with the proper payload, then you know that the server
is alive.
If the server does not sends any answer, there is no way to know if there is an
access for the server or not.

### The Naive way

In order to know if a port is open, there is a need to open a connection for :-)

However, the current implementation that I've created is to first resolve IP from
DNS records, then using `timeout` for a connection.
When a connection that is initiated does not have the proper 3 way handshake, it
might not return for a long time (almost forever), so the `timeout` will kick in
and we "know" that the connection is `closed`.

When an open request is back, something is filtering the request (firewall?).

The "naive" way, is guessing the meaning of each type of "error" and translate it
to the status of a request.

# License

The aim of this repo is to teach, so I have choosen the "The Unlicense" license.

```
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>
```
