Our initial step in starting the project was to install MySQL on my computer, after which we created 
a new database called "goapi". We created a table called time_log in this database, with fields for 
the timestamp (timestamp) and the primary key (id). Next, we created a file, such as main.go, and 
started the project for the Go program. We included Go's time package to ensure precise time zone 
conversion for Toronto. The go-sql-driver/mysql driver was used to connect the Go program to the 
MySQL database. We added the timestamp to the time_log database after every API request to get 
the current time.
We created a Dockerfile for the Go application. The Dockerfile includes instructions to build the 
application within the container. We executed docker build to create a Docker image, streamlining 
deployment by encapsulating the Go API and its dependencies.
