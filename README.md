# NormsBackend

This is the backend support for the application. To work in the most advisable environment, download
visual studio code it is not installed already along with basic go extensions like the Go Documentation,
gofmt, and dep.

If you are new to go, run dep init and then dep ensure in order to download all necessary packages. 
Make sure you $GOPATH starts at ~/go/src/github.com/[user]/NormsBackend, else you must change the locations
of the dependencies in your imports header. The main.go is currently located in the project file rather than the src file due to issues with Heroku's root-packaging pathing. If you know how to fix this, please let me know. 

**ALSO** please create your own branch when working on your part of the code and make a pull request for a code review before pushing to the master branch. It will keep the code functioning. Of course, not everything is
currently optimize so it will take some work. 

When using to this API, make a post request to code https://codenorm.herokuapp.com + /[request]. The request types are in the main.go file and explanations can be found in their respective packages in the pkg folder. When working
on your own environment use localhost:[port] where port will appear in your terminal after you run the program. It is not necessary to build the program. Just run it with 'go run main.go'

Well that's all that I have on the top of my head. If there is anything you think I should add, feel free to make changes. :)