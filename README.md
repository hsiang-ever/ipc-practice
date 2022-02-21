# IPC for socket, pipe, and shared memory

## Preparations

1. Install docker/docker-compsoe
2. Run go container using docker-compose  
`docker-compose up -d`
3. Open four terminals for one server and three clients
4. All terminals get into go container (`server-client-app`) 
`docker exec -it server-client-app sh`

## Execute server and client
*Note: Please get [Preparations](#preparations) ready first.*

1. All terminals go to `bin` folder
`cd /go/bin`
2. Run Server and Client
    - One terminal run **Server** first: `./integrated`
    - One terminal run **Socket Client**: `./socket`
    - One terminal run **Pipe Client**: `./pipe`
    - One terminal run **Shared Memory Client**: `./shared_memory`
3. Console will output following lines
```
Start to write data to Client using `socket`
Start to write data to Client using `shared memory`
Start to write data to Client using `pipe`
Enter number:
```
Type a series of number separated by space and press **Enter**  
4. **Socket Client**, **Pipe Client**, **Shared Memory Client** will output the result, *Mean*, *Median*, and *Mode* respectively.

## Build the source code
*Note:*  
*- Please get [Preparations](#preparations) ready first.*  
*- Executable files in `bin` folder is built using following steps. You might go following [Execute server and client](#execute-server-and-client) steps directly.*  

1. Go to respective folder in `src`
    - Build **Server**: `cd integrated`
    - Build **Socket Client**: `cd socket`
    - Build **Pipe Client**: `cd pipe`
    - Build **Shared Memory Client**: `cd shared_memory`
2. Build executable files
Run cmd `go install` in each folder.
