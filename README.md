# CRUD

<b>This is golang module API implementation for CRUD operations</b>
<p>Stable Version v1.0.1</p>

<b>How to use this module?</b>
- go get -u github.com/amitkumardube/CRUD
OR
- Add it's packages as dependency in your program using import github.com/amitkumardube/CRUD/update_resource
- then trigger go run and go will download all the dependency and update go.mod file

The module supports below operations
- update_resource (this support string update operation)
- delete_resource (this support string delete operation)

<b> functions in this module </b>
- Delete_resource in package delete_resource - returns string , error
- Update_resource in package update_resource - returns string , error

<b>Testing steps</b>
- to test all the packages : go test -v ./...
- to check coverage : go test -v -cover ./...
- to create coverage report : go test -v -coverprofile=cover.txt ./...
- to create a html report : go tool cover -html=cover.txt -o cover.html
- to display covrage percentage on stdout for each function : go tool cover -func=cover.txt