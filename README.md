# Automatic Timing System


## Analysing Logic of the Business
According to the [requirement-1](./docs/requirements/Backend_take_home_test.pdf) and [requirements-2](./docs/requirements/Requirements_of_the_task.png), summarized the following tasks:

#### Tasks
- [ ] Protocol between Timing Service and Test-Client

    - [ ] Types of the timing points which are finish corridor and finish line
    - [ ] The IDs of the athletes
    - [ ] Precision of the time records is to second
    - [ ] Documents
- [ ] Timing Service

    - [ ] Receives IDs of the athletes from Test-Client in Real-time
    - [ ] Receoves Types of the timing points
    - [ ] Get athletes' full names by their IDs
    - [ ] Get athletes' start numbers by their IDs
    - [ ] Sends start numbers when the athlete enters the finish corridor
    - [ ] Sends full names when the athlete crosses the finish line
    - [ ] Sends the finish time when the athlete crosses the finish line
    - [ ] Database
        - [ ] IDs of the athletes
        - [ ] Start numbers of the athletes
        - [ ] Full name of the athletes
    - [ ] Documents

- [ ] Test-Client
    - [ ] Suppose that 10 athletes are there
    - [ ] Suppose that there are 10 priorities from 1 to 10 and the priorities are as factors for the speeds of the athletes
    - [ ] Documents

- [ ] Web-Client

    - [ ] Displays SmartNumber and Name in the table form when the athlete enters the finish corridor
    - [ ] Displays the finish time in the athlete's row of the table when the athlete crosses the finish line
    - [ ] Only displays the latest record for each athlete? Need to confirm!
    - [ ] Displays data without refreshing browser
    - [ ] Documents

#### Extra Tasks
- [ ] The order of the displaying of the data need to change, when athlete B go over athelete A in the finishi corridor.

- [ ] The Web-Client should stop to communicate with server, when the window of the browser is in the backgorund.

- [ ] The Web-Client should resume connection with the server and communite with it, when the window of the browser is in the foreground.

- [x] Proving a loading animation to solve the problem that the data will display incorrect for a while when the browser window become actived from deactivated. 

#### Bonus Tasks

- [ ] Docker Compose File


## Choice of the libraries


## Parts of the project


## References

#### Timing Mats

- [Race Timing RFID Mats - How to Setup & Use](https://www.youtube.com/watch?v=MnkCDdUjP5w)



#### Golang

- [The Go Programming Language Specification](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [How to write Go code](https://golang.org/doc/code.html)
- [Wiki of Go in the Github](https://github.com/golang/go/wiki)
- [Go database/sql tutorial](http://go-database-sql.org/index.html)
- [SQLDrivers - Wiki of Go in the Github](https://github.com/golang/go/wiki/SQLDrivers)
- [SQL Interface - Wiki of Go in the Github](https://github.com/golang/go/wiki/SQLInterface)
- [Golang, ORMs, and why I am still not using one](http://www.hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html)

