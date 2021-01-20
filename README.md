# CRUD

<b>This is golang API implementation for CRUD operations</b>
Stable Version v1.0.1

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