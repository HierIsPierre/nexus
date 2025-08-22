# NEXUS CLI

## About Nexus
The purpose of this tool is to allow developers to deploy their code changes quickly so that it can be tested in a sandbox environment. There are multiple features that the nexus cli supports:
1. deploy your code to a dedicated sandbox environment
    `nexus sandbox -imagename 'unique image name' -subscription 'name of dedicated sandbox subscription' -service 'GandalfService' -service 'you can specify multiple services'`
2. deploy your code to the staging subscription
    `nexus test -imagename 'unique image name' -service 'GandalfService' -service 'can specify more than on service'`

3. deploy your code locally
    `nexus local -service 'GandalfService'`

## Before starting
1. make sure that there is a cert folder in the root folder
2. run the following command `dotnet dev-certs https --export-path ./certs/aspnetcore-dev-cert.pem --format PEM`
3. you are ready to go
