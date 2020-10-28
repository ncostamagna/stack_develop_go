# Grpc Golang Course

This is a companion repository for my [GRPC Golang course](http://bit.ly/grpc-golang-github)

[![course logo](https://udemy-images.udemy.com/course/480x270/1685664_10e0_4.jpg)](http://bit.ly/grpc-golang-github)

# Content

- Greeting Service
- Calculator Service
- Unary, Server Streaming, Client Streaming, BiDi Streaming
- Error Handling, Deadlines, SSL Encryption
- Blog API CRUD w/ MongoDB
<br />


# Introduccion

gRPC es free y Open-Source desarrollado por Google<br />
Definiremos los Request y Response como RPC (Remote Procedure Calls)<br />
Algunas de las ventajas son que es **moderno, rapido y eficiente**, utiliza **HTTP/2**, tiene **latencia baja, soporta streaming y es independiente del lenguaje**, es facil de utilizar **autentificacion, load balancing, logging y monitoreo**<br /><br/>
El **LIENTE** llamara directamente la funcion del **SEVIDOR**

### Server
```go
func CreateUser(User user){
    // code
}
```

### Client
```go
// code
server.CreateUser(user)
// code
```
![Events](../images/61.png)

### Extensiones
- proto lint
- vscode-proto3
<br /><br />



# Proto
Necesitamos definir el mensaje y el servicio utilizando **Protocol Buffers**<br />
```proto
syntax = "proto3";

message Greeting {
    string first_name = 1;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse {
    string result = 1;
}

service GreetService{
    rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

