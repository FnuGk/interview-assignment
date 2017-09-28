# Solution written in GO

This is a simple solution for the problem of reading user data
from an sqlite3 database and store it in a plain JSON file written in GO.

The application can be build by running `go build` or run without building simply by calling `go run main.go`. run with the `--help` flag to see how to run, or see `run.sh`

## Architectur
The system is split into sub packages for each logic domain.
The db package abstracts access to the database. 
The model package contains the models for the working domain.
Lastly the top level main package stiches it all together and works as the entry point for the application. 

## Security Concerns
As you can pass arbitary sql queries via the command line flags, it would be possible to modify the database.

## Next Steps
*  Moving logic from the main function into a control package would give even greater seperation of concerns. 
*  If a query given via the `--query` flag does not start with `SELECT id,firstName,lastName,email` then the app will panic.
