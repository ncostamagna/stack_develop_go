# Indice
- [Introduccion](#introduccion)
- [Instalaciones](#instalaciones)
- [NodeJS](#nodejs)
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


### VSCode
- **Bracket Pair Colorizer**: Coloca las {} de colores
- **JavaScript (ES6) code snippets**: ES6
- **TypeScript Importer**: Nos ayuda con todas las importaciones de JS

Puedo modificar y agregarle comportamientos a NodeJS modificando el motor V8 en C++, logicamente C++ es mucho mas potente que JS, ya que JS solo fue pensado para web
<br /><br /><br />

# NodeJS
NodeJS se maneja con un solo hilo
<br />
**package.json** -> nodemon
```sh
npm install -D nodemon
```
```json
{
    "scripts":{
        "start":"node index.js",
        "dev":"nodemon index.js"
    }
}
```
```sh
npm run dev
```
<br />
podemos ejecutar el siguiente comando para salir de un procesamiento, en lugar de return

```javascript
process.exit(0);
```
<br />

### Blocking vs Non Blocking I/O
**Blocking**: Cuando programamos de manera lineal, para que comience a ejecutar la linea 3 tiene que esperar
a que termine la 2, va de manera secuencial<br />
**Non Blocking**: Cuando trabajamos con programacion en paralelo, con los callbacks, no hay que esperar a la 2 para que empiece la 3

### Arquitectura NodeJS

![Events](images/00001.png)

Cuando existen callbacks y funcionalidades en paralelo
![Events](images/00002.png)

Primero ejecuta el Main, luego los callbacks que estan en el node api pasan a la cola de callbacks a medida que ya pueden ser ejecutados, cuando ya se pueden ejecutar pasan de la cola de callbacks a la pila
![Events](images/00003.png)


# ES6
### Expresiones
Cualquier cosa que genere un valor primitivo, cuando coloco 'Hi', 6, 4.5, etc..<br />
Esto me devuelve valores primitivos:<br />
- **String**
- **Number**
- **Boolean**
- **Array**
- **Set**
- **Map**
- **Function**
- **Symbol**
- **Object**: Los demas tipos heredan de Object, todo esta en Object
- **undefined**
- **null**

### Let vs Var
Cambia donde vive la variable<br />
**let** no se puede volver a inicializar y viven en ambitos o scopes diferentes

### Funciones
Dentro de un objeto puedo generar una funcion de la siguiente manera desde ES6
```javascript
const person = {
  name: "Nahuel",
  getFullName(){
      // .....
  }
}
```
las funciones siempre retornan **undefined** si yo no le indico que retorne nada<

### Arrow Function
```javascript
let suma = (a, b) => a + b //Retorna suma, en una linea
```

### CallBacks
La mejor practica es utilizar los **callbacks** lo menos posible, cuando tenemos muchas y surge un error
aveces es dificil encontrarlo y requiere buscar en el codigo<br /><br />
**SIEMPRE** que genero un **callback** el primer parametro debe ser un **ERROR**
```javascript
if (!book){
    const error = new Error();
    error.message = "Book not found";
    return callback(error);
}
callback(null, book);
```
**callback hel**: cuando empiezo a generar callbacks uno adentro del otro y el codigo se hace cada vez 
menos mantenible, es la anidacion de callbacks

### Async / Await
- async: devuelve un elemento promesa
- await: ejecucion de una funcion async sea pausada hasta que una promesa sea terminada
Cuando utiliamos **async/await** tenemos que utilizar un metodo **try-catch**, tienes un codigo mas limpio y mejor control de las excepciones

```javascript
let getInformacion = async (id) => {
  let empleado = await getEmpleado(id);
  let resp = await getSalario(empleado);

  return `${resp.nombre} - ${resp.salario}`;
}

getInformacion(3)
  .then(mensaje => console.log(mensaje))
  .catch(err => console.log(err))

```
**Promesas**: Las promesas siguen corriendo al hacer un reject, luego debemos hacer un return para que no siga corriendo
<br /><br />

# Sincronia vs Asicnronia
JavaScript **NO ES ASINCRONICO**, NodeJS hace las cosas de forma asincrona pero V8 se ejecuta sincronicamente,
JavaScript **ES SINCRONO**, se ejecuta linea por linea<br />
Si embargo, NodeJS es Asincrono, mientras V8 lo covierte en lenguaje maquina, NodeJS podria estar haciendo otras cosas a la vez< br />

### Eventos
Algo que ha sucedido en nuestra aplicacion que podemos responder<br />
2 tipos de eventos: **de sistema** (Libuv / C++ Core) y **personalizados** (JavaScript Core / Event Emitter)<br />
JS es **Non-Blocking I/O**: capacidad de realizar acciones sin pausar o detener otras
![Events](images/events.png)
<br /><br />

# Base de datos

### Relacionales (Sequelize)
**Sequelize**: ORM para db relacionales
```sh
npm i sequelize sequelize-cli
node_modules/.bin/sequelize init # Inicializar proyecto sequelize

# Creamos el modelo, pluraliza las tablas, le agrega s al final
node_modules/.bin/sequelize model:generate  --name Contact --attributes firstname:string, lastname:string

# migrate para crear las tablas
node_modules/.bin/sequelize db:migrate

# genera script seed para cargar datos, podemos hacer insert en up
node_modules/.bin/sequelize seed:generate --name seed-contact

# cargamos los datos detallados en el archivo anterior
node_modules/.bin/sequelize db:seed:all
```

**Seeders**: Crear data dummy o datos de prueba
<br /><br /><br />

# Arquitectura Apis

### Patrones de Arquitectura

https://medium.com/@maniakhitoccori/los-10-patrones-comunes-de-arquitectura-de-software-d8b9047edf0b


**N Capas**: Cada capa proporciona servicio a la capa superior, por lo general se utilizan las siguientes capas: presentacion, aplicacion, negocio y datos.<br />

**Container** dentro de startup configuramos las inyecciones de dependencia
```javascript
const { createContainer, asClass, asValue, asFunction } = require("awilix");

// services
const { HomeService } = require("../services");

// controllers
const { } = require("../controllers");

// routes
const { } = require("../routes/index.routes");
const Routes = require("../routes");

// models
const { User, Comment, Idea } = require("../models");

// repositories
const { } = require("../repositories");

const container = createContainer();

container.register({
    HomeService: asClass(HomeService).singleton(),
    UserService: asClass(UserService).singleton(),
    CommentService: asClass(CommentService).singleton(),
    IdeaService: asClass(IdeaService).singleton(),
    AuthService: asClass(AuthService).singleton()
  });

module.exports = container;
```
**awilix** nos ayuda a mantener todas las inyecciones de dependencia<br />
**helmet** nos ayuda con las brechas de seguridad en nuestra aplicacion<br />
**compression** nos ayuda a comprimir nuestras peticiones http para que sea mas rapida<br />
**express-async-error** nos ayuda a capturar los errores asincronos en los middlewares<br />
**mongoose-autopopulate** me ayuda a traer y relacionar objetos del modelo<br />
**yargs** nos ayuda cuando queremos ejecutar node por linea de comando y usar parametros<br />
**request & axios**
  - request trabaja en base a callbacks
  - axios trabaja en base a promesas

**ARQUITECRURA DE NUESTRA APLICACION**
![Arquitectura](images/arquitectura.png)

### Autentificacion
Generamos el jwt helper, definiendole el **JWT Secret** y el **tiempo de expiracion**
```javascript
const { sign } = require("jsonwebtoken");
const { JWT_SECRET } = require("../config");

module.exports.generateToken = function(user) {
  return sign({ user }, JWT_SECRET, { expiresIn: "4h" });
};
```
Luego creamos nuestro servicio de autentificacion **auth.service**<br />
Puedo autentificar tambien peticiones **get**, pasando el Token como header

### Unit Test
#### Untilizaremos Jest
Implementaremos pruebas unitarias a los repositorios y servicios (unicamente en este proyecto)<br />
debemos crear un archivo que se llame **jest.config.js**, cremos un directorio de **unit** para las pruebas unitarias y de **mocks** para todos los datos falsos que vamos a usar para las pruebas, ahi crearemos objetos con data dummy<br /><br />
Debemos modificar nuestro script de test en **package.json**
```json
  "scripts": {
    "start": "node index.js",
    "dev": "nodemon index.js",
    "test": "jest"
  },
```
para correr los test, va a tomar todos los archivos que tengan la palabra test
```sh
npm run test
```

### Swagger
Proyecto open source que nos ayuda a implementar la documentacion de nuestras apis<br />
agregamos en config/swagger los 2 archivos Json<br /><br />
Links para implementar:
- https://medium.com/@diegopm2000/creando-un-api-rest-con-swagger-node-c880bdac04a5
- https://github.com/swagger-api/swagger-node

**Generar proyecto Express**
```sh
npm i -g express-generator
express --views=hbs nombre-proyecto
cd nombre-proyecto
npm i
```

# Microservicios

**Arquitectura Monolitica**: Un unico ejecutable logico<br />
Si utilizo MVC estoy implementando una arquitectura monolitica<br />


|       Monolitico              |         Microservicios         |
|-------------------------------|--------------------------------|
| **Ventajas**                  | **Ventajas**                   |
| Simple desarrollar            | Responsabilidad unica          |
| Facil para trabajar solo      | Persistente                    |
| Simple de probar y desplegar  | Resiliencia, API bien definida |
|                               | facil escalar                  |
|                               | Potencial confiable            |
|                               | Combinacion Tecnologias        |
|                               | Deploys trasnparentes          |
| **Desventajas**               | **Desventajas**                |
| No son flexibles              | Dificiles de testear           |
| Dificil de mantener           | Posibilidad de duplicidad      |
| Dificil de escalar            | Integracion de Informacion     |
| Potencial poco confiable      | Equipos especializados         |

**Netflix** para solucionar el colapso que tuvo implemento microservicios, y para probar que al romperse servicios su aplicacion seguia funcionando usaron **monkey chaos**<br /><br/>
Para correr el proyecto debemos traer la imagen de nginx
```sh
docker pull nginx
```

# CI/CD
Este paso se divide en 3 tareas:
- continuous integration
- continuous delivery: Manual antes de subir a Prod
- continuous deployment: Automaticamente se implementa en Prod

Cuando hacemos test unitarios, debemos validarlo desde la integracion continua que este todo ok automaticamente, antes de pasar a otra rama<br />
Tenemos herramientas como **travisCD**, **Azure DevOps**, **Jenkis**, etc.. Que nos dicen cuando un cambio no debera unirse por un determinado error<br /><br />
Un **pipeline** es una secuencia de pasoso a seguir en un proceso de DevOps<br /><br />
**TravisCI**
- Creamos un usuario en https://travis-ci.org/dashboard
- Creamos un archivo .travis.yml
- debemos instalar **travis CLI**

```yaml
language: node_js
node_js:
  - stable # version estable de NodeJS
deploy:
  provider: heroku # el deploy lo vamos a hacer con heroku
  app: miaplicacion_nodejs # nombre en heroku
```
```sh
# Ecryp token que tenemos en heroku y se lo paso a nuestro archivo
travis encrypt $(heroku auth:token) --add deploy.api_key
```
Esto me genera el token en mi archivo
```yaml
language: node_js
node_js:
  - stable # version estable de NodeJS
deploy:
  provider: heroku # el deploy lo vamos a hacer con heroku
  app: miaplicacion_nodejs # nombre en heroku
  api_key:
    secure: asfdskfdbsbfdfbsjffds
```
En el proveedor (en este caso heroku) configuramos todas las variables<br /><br />
Luego:
  - vamos a la pagina oficial https://travis-ci.org
  - vamos al simbolto de + para agregar algo nuevo
  - selecciono cual de todos los servicios que genere voy a usar
  