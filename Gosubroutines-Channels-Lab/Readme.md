### Task
You are given an array of links and your task is to iterate through the array of links and execute the HTTP GET request by using each link (Don't worry, HTTP GET Request code is present in the boilerplate code).

The following conditions should satisfied.

```
1. All the requests should execute concurrently in their own go-subroutine(child subroutine).

2. Use channels for Communication b/w subroutines...

3. Upon completion of each child subroutine, relay the status of the request to the main subroutine through channels.

Send -> "Website is up and running" if the request is successfully executed.

Otherwise, Send -> "Website is down" if any error occurs.

4. At the end of the main subroutine print the message sent by child-subroutines.
```