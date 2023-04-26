# push-to-sekoia

Push to Sekoia is a program write in golang that send event logs to a specified intake in Sekoia IO. 

The code contain two functions.

The first function will read all the stdin entry "in our case the logs" and stock it in a variable.

Th second fuction will : 
post a web request directly to the Sekoia IO endpoint with the first argument of our program as a header (which is the key of the Sekoia IO intake).
create a client for printing the responses of the request.
