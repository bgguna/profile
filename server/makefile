# Project
current_dir	=	${PWD}
bin_dir		=	${current_dir}/bin
out_dir		=	${current_dir}/out
db_dir		=	${current_dir}/storage
database	=	${db_dir}/bgguna.db
binary		=   server
project		=   bogdan.guna.me	
main		=	main.go
coverage	=	coverage.out

all: clean format build test

# Go Commands
gobuild		=	go build
gorun		=	go run
goget		=	go goget
gofmt		=	go fmt
golint		=	golint
gotest		=	go test
gocover		=	go tool cover
ldflags		=	"-s -w"
modules		=	./...

deep-clean:
	@ echo "Running a deep clean operation..."
	rm -rf ${bin_dir}
	@ echo "Step 1/4: Deleted binaries."
	rm -rf ${out_dir}
	@ echo "Step 2/4: Deleted output files."
	rm -rf ${coverage}
	@ echo "Step 3/4: Deleted coverage files."
	rm -rf ${database}
	@ echo "Step 4/4: Deleted database."
	@ echo "Done deep cleaning."

clean:
	@ echo "Cleaning project repository..."
	rm -rf ${bin_dir}
	rm -rf ${out_dir}
	rm -rf ${coverage}
	@ echo "Done cleaning."

db:
	@ echo "Checking database..."
	@ if [ -f "${database}" ]; then (echo "Database already exists: ${database}") else (echo "Database not found. Creating database from file..."; sqlite3 ${database} < ${db_dir}/contact.sql; echo "Done setting up the database: ${database}") fi

format:
	@ echo "Formatting code..."
	@ $(gofmt) $(modules)
	@ echo "Done formatting."

lint:
	@ echo "Linting the code to Go standards..."
	@ $(golint) $(modules)
	@ echo "Done linting."

run-dev:
	@ echo "Running in dev mode..."
	@ make db
	$(gorun) ${main}

build:
	@ echo "Running build command..."
	#@ echo "Step 1/3: TESTS"
	#@ make test
	@ echo "Step 2/3: BUILD BINARIES"
	$(gobuild) $(modules)
	$(gobuild) -o ${bin_dir}/${binary} ./
	@ echo "Done building binaries."
	@ echo "Step 3/3: DATABASES"
	@ make db

run:
	@ echo "Running binary..."
	${bin_dir}/${binary}

test:
	@ echo "Running tests with package level coverage..."
	$(gotest) $(modules) -v -cover
	@ echo "Done running tests."

coverage:
	@ echo "Running tests and printing coverage report..."
	@ -$(gotest) $(modules) -race -v -covermode=atomic -coverprofile $(coverage)
	@ $(gocover) -func $(coverage)
	@ echo "Done writing test coverage report."

help:
	@ echo "\
	Manual for Server\n\
	\n\
	Available make commands:\n\
	------------------------------\n\
	* all:		formats, cleans, builds and tests the binaries\n\
	* build:	builds binaries\n\
	* run-dev:	runs the server in development mode\n\
	* run:		runs the server\n\
	* format:	formats the code to the Golang standards\n\
	* lint:		highlights smells in the code\n\
	* test:		runs the tests with coverage for all packages in the project\n\
	* coverage	runs the tests and writes the coverage to file\n\
	* clean:	deletes the generated binaries and output files (log and health report)\n\
	* help:		prints this menu again\n\
	\n\
	Done."
