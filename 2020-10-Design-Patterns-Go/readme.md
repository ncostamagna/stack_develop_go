

# SOLID

### Single Responsibility Principle
Cada uno tiene su responsabilidad, cada metodo hace algo y tiene una responsabilidad unica

### Open Closed Principle
Se abre por extension y se cierra por modifidacion, se genera una funcion para implementar los metodos que modifican dicho objeto, sin agregarlo como metodo del mismo. Hacemos una generica para todos los tipos de modificaciones que requiere el objeto

### Liskov Substitution Principle
No se le entiende nada al ruso del orto,
Todos los metodos de la super clase seran reemplazadas por la subclase, <br />
Cada clase que hereda de otra puede usarse como su padre sin necesidad de conocer las diferencias entre ellas.

### Interface Segregation Principle
Manejamos un objeto con una interfaz, si tenemos varios aracnidos como arañas, tarantulas y escorpiones tendriamos una interfaz aracnido que tiene las clases araña, tarantula y escorpiones. Tendriamos un problema utilizando el metodo telaaraña, lo implementariamos en todos y le podriamos un panic a escorpion el principio de ISP dice que sólo deberían conocer de éste aquellos métodos que realmente usan, y no aquellos que no necesitan usar, en este caso deberiamos generar una interfaz telaaraña e implementarsela al que lo necesita. Se generan mas interfaces asegurandonos que cada objeto tenga los metodos que realmente utilice
![Events](../images/1.jpg)

### Dependency Inversion Principle
Los módulos de alto nivel no deberían depender de los módulos de bajo nivel. Ambos deberían depender de abstracciones (p.ej., interfaces).
Las abstracciones no deberían depender de los detalles. Los detalles (implementaciones concretas) deben depender de abstracciones.
![Events](../images/2.jpg)


# Builder
Extrae el contructor de una clase y le pasa un Builder directamente, sin necesidar de tener varios parametros en la instancia de la clase, se genera un Builder, se van aregando los parametros ahi y luego se le pasa el Builder a la instancia de la clase

# Factory
Clase factory que se encarga de crear diferentes tipos de clase en base a algun/os parametro, debemos devolver siempre el puntero (En caso de un struct)
```go
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}
```

# Prototype
Replicar un diseño de objeto en varios, por ejemplo si tenemos Cars, Iphone, generamos un prototipo para cada uno. Nos permite **clonar** objetos para ser utilizado en otro lado