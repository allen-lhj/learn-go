this course is to show how to call the module from the local file
and how to replce the file 

use `go mod -replace example.com/greetings=../greetings`
an then use `go mod tidy` commond to synchronize the dependencies