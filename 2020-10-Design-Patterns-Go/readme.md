

# SOLID

### Single Responsibility Principle
Cada uno tiene su responsabilidad, cada metodo hace algo y tiene una responsabilidad unica

### Open Closed Principle
Se abre por extension y se cierra por modifidacion, se genera una funcion para implementar los metodos que modifican dicho objeto, sin agregarlo como metodo del mismo. Hacemos una generica para todos los tipos de modificaciones que requiere el objeto


# Builder
Extrae el contructor de una clase y le pasa un Builder directamente, sin necesidar de tener varios parametros en la instancia de la clase, se genera un Builder, se van aregando los parametros ahi y luego se le pasa el Builder a la instancia de la clase