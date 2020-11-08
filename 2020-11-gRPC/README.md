# Indice
- [Content](#content)
- [Instalacion](#instalacion)
- [Introduccion](#introduccion)
- [Proto](#proto)
- [Protocol Buffers](#protocol-buffers)
- [gRPC Languages](#grpc-anguages)
- [HTTP/2](#http/2)
- [Escalabilidad](#escalabilidad)
- [Errors Code](#errors-code)
- [Deadline](#deadline)

# Content

- Greeting Service
- Calculator Service
- Unary, Server Streaming, Client Streaming, BiDi Streaming
- Error Handling, Deadlines, SSL Encryption
- Blog API CRUD w/ MongoDB
<br />

# Instalacion

```sh
curl -OL https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip
unzip protoc-3.13.0-linux-x86_64.zip -d protoc3
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/
sudo chown ncostamagna /usr/local/bin/protoc
sudo chown -R ncostamagna /usr/local/include/google
```
```sh
go get -u google.golang.org/grpc

# El instructor lo hace con este
go get -u github.com/golang/protobuf/tree/master/protoc-gen-go

# Yo encontre este
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc folder/folderpb/file.proto --go_out=plugins=grpc:.


# Como no me funcionaba instale esto
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/golang/protobuf/protoc-gen-go

protoc -I . templatespb/templates.proto --go_out=plugins=grpc:.
```
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

# message -> data, request and response
message Greeting {
    string first_name = 1;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse {
    string result = 1;
}

# en el service seria como definir el gRPC endpoint
service GreetService{
    rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

# Protocol Buffers
Diferencias entre JSON y Protocol Buffer
- JSON: CPU intensive, because the format is human readable
- Pro Buff: is less CPU intensive, mas cerca del codigo maquinal al ser binario<br />
![Events](../images/63.png)<br />
![Events](../images/69.png)
<br />
Porque Protocol Buffer
- Facil de escribir la definicion de mensajes
- La definicion de la API es independiente a la implementacion
- Todo el codigo gordo se genera automaticamente en base a un siemple .proto

# gRPC Languages
Implementaciones de gRPC en los siguientes lenguajes:
- java: puto gRPC en Java
- go: puro gRPC en go
- C: puro gRPQ en C
- C++, Python, Ruby, objective C, PHP, C# y el resto en C <br />
Por mas que tengamos microservicios en diferentes lenguajes podemos comunicarlos igual con gRPC<br />
![Events](../images/64.png)

# HTTP/2
diferencias entre http/2 y http/1.1: https://imagekit.io/demo/http2-vs-http1
<br />

**HTTP/1.1**
- Por cada request hace una nueva conexion TCP
- No soporta header compression
- text, muy facil para hacer debugging y ver la data
<br />

![Events](../images/65.png)<br />

**HTTP/2**
- Mucho mas rapido
- El client y el server pueden pushear mensajes en paralelo por la misma conexion TCP
- Grandioso para reducir latencia
- Server puede pushear streams, multiples mensajes
- Soporta header compression
- Es binario
- Es seguro (SSL no es requerido)
<br />

![Events](../images/66.png)<br />

### 4 Types if API in gRPC
- Unary -> traditional API (HTTP REST)<br />
![Events](../images/67.png)<br />
![Events](../images/68.png)<br />
![Events](../images/70.png)


# Escalabilidad
- gRPC server es asincronico por defecto
- No tenemos bloqueos de hilos en los request
- por lo tanto por cada gRPC server podemos servir millones de request en paralelo
- gRPC Client puede ser asincronico o sincronico (bloking)
- gRPC Client puede realizar load balancing

# Errors Code
gRPC maneja un standart de codigos de errores
<br />
documentacion: https://grpc.io/docs/guides/error/
<br />
demo: http://avi.im/grpc-errors/

# Deadline
Es recomendable siempre definir un tiempo de deadline en todos nuestros clients <br />
El server debera chequear si el deadline se excedio y debera cancelar el trabajo que esta haciendo<br />
https://grpc.io/blog/deadlines/
<br />

![Events](../images/71.png)