# Indice
- [Introduccion](#introduccion)
- [Instalaciones](#instalaciones)
- [Go](#go)
- [ES6](#es6)
- [Sincronia vs Asicnronia](#sincronia-vs-asicnronia)
- [Base de datos](#base-de-datos)
- [Arquitectura Apis](#arquitectura-apis)
- [Microservicios](#microservicios)


<br />

# Introduccion

### Que es Go?
se le llama Golang tambien, cuando uno pone Go en Google aparecen mil cosas, lenguaje de google. Resolver problemas internos de google, necesitaba algo con mucha velocidad parecido a c++<br />
Lenguaje fuertemente tipado, pensado para aprovechar los ultimos avances en hardearem multiprocesadores,
 enorme cantidad de memoria y paralelismo. Aprovecha mucho el paralelismo<br />
Lenguaje compilado, genera un biranio.<br />
Obliga al desarrollador a realizar buenas practicas<br />
**Lenguaje ideal para desarrollos grandes con miles y miles de usuarios**<br />
- Facil de entender y claro
- Traducido a c++
- Las funciones de Go pueden devolver mas de un valor
- Se pueden desarrollar instrucciones Sync como Async
- Programacion Async mas clara que NodeJS (Promesas)
- Solo existe **for** para interacciones (No existe while)
- NO ES ORIENTADO A OBJETOS
- Scope se resuelve con el nombre de las variables, metodos o funciones
   - Si empieza con minuscula es privada
   - Si empieza con mayuscula se exporta a otro scope


<br />

# Instalaciones

Instalamos Go de https://golang.org/ <br />

Instalamos la extension **Go, Go Outliner y Go Autotest (chequeando nuestro programa)** de visual studio code <br />

# Go

Ejecuciones para correr el programa

```sh
# Corremos el archivo main, lo compila en memoria y lo ejecuta
go run main.go 

go build main.go # Genero el ejecutable
```

### Variables
Se inicializan en cero, blanco o false<br />
Si no uso las variables el programa no corre
```go
var numero2 int

// toma el tipo de datos del valor que le asigno
// se crean nuevas variables automaticamente
numero6, numero7, numero8, numero9 := 4, 12, 5, "Texto"

numero2 = int("2") // convierto a 2 entero
texto = fmt.Sprintf("%d", numero)

```
Hay una libreria para convertir que se llama **strconv**<br />

### Funciones
Todo con funciones, no existen metodos (son funciones en si misma)<br />
Alojamos un dato en **_** cuando no vamos a usarlo<br />

### Vectores

**slice**: Vectores dinamicos, puedo ampiar la dimension en tiempo de ejecucion
```go
var vector [5]int // vector, no puedo ampliar mas
var vector_slice []int //slice, puedo apliar dinamicamente
elementos[3:] //de la posicion 3 hasta el ultimo, igual que python

elementos := make(tipo,largo,capacidad maxima)
len(elementos) //tamaño
cap(elementos) //capacidad
```
![Events](images/00001.png)