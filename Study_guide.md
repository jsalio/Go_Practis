# Guía de Estudio de Go: De Cero a Junior

Esta guía está diseñada para ayudarte a aprender Go de manera estructurada, desde los conceptos básicos hasta temas más avanzados.

## 1. Fundamentos de Go

### 1.1 Conceptos Básicos

#### Variables y Tipos de Datos
- **Declaración de variables**
  ```go
  var nombre string = "Juan"  // Declaración explícita
  edad := 25                 // Declaración corta
  ```
  - Go es un lenguaje tipado estáticamente
  - Las variables deben ser declaradas antes de usarse
  - El operador `:=` infiere el tipo automáticamente

- **Tipos básicos**
  - `int`: Números enteros (int8, int16, int32, int64)
  - `float`: Números decimales (float32, float64)
  - `string`: Cadenas de texto
  - `bool`: Valores booleanos (true/false)
  - `byte`: Alias para uint8
  - `rune`: Alias para int32 (representa un carácter Unicode)

- **Constantes**
  ```go
  const PI = 3.14159
  const (
      A = 1
      B = 2
  )
  ```
  - Valores que no pueden ser modificados
  - Deben ser inicializadas al declararse
  - Pueden ser agrupadas

- **Conversión de tipos**
  ```go
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
  ```
  - Go requiere conversión explícita entre tipos
  - No hay conversión automática que pueda causar pérdida de datos

#### Operadores
- **Operadores aritméticos**
  - `+`: Suma
  - `-`: Resta
  - `*`: Multiplicación
  - `/`: División
  - `%`: Módulo (resto)
  - `++`: Incremento
  - `--`: Decremento

- **Operadores de comparación**
  - `==`: Igual a
  - `!=`: Diferente de
  - `<`: Menor que
  - `>`: Mayor que
  - `<=`: Menor o igual que
  - `>=`: Mayor o igual que

- **Operadores lógicos**
  - `&&`: AND lógico
  - `||`: OR lógico
  - `!`: NOT lógico

- **Operadores de asignación**
  - `=`: Asignación simple
  - `+=`: Suma y asignación
  - `-=`: Resta y asignación
  - `*=`: Multiplicación y asignación
  - `/=`: División y asignación
  - `%=`: Módulo y asignación

#### Entrada/Salida Básica
- **Uso de fmt.Println y fmt.Printf**
  ```go
  fmt.Println("Hola Mundo")           // Imprime y añade salto de línea
  fmt.Printf("Hola %s\n", nombre)    // Formateo con especificadores
  ```
  - `Println`: Imprime y añade salto de línea
  - `Printf`: Formateo con especificadores (%s, %d, %f, etc.)
  - `Sprintf`: Retorna string formateado

- **Lectura de entrada del usuario**
  ```go
  var input string
  fmt.Scanln(&input)  // Lee hasta encontrar salto de línea
  ```
  - `Scanln`: Lee una línea completa
  - `Scanf`: Lee con formato específico
  - `Scan`: Lee palabras separadas por espacios

- **Formateo de strings**
  ```go
  nombre := "Juan"
  edad := 25
  mensaje := fmt.Sprintf("%s tiene %d años", nombre, edad)
  ```
  - Especificadores comunes: %s, %d, %f, %v, %T
  - Formateo de números con precisión
  - Alineación y padding

### 1.2 Estructuras de Control

#### Condicionales
- **Sentencia if/else**
  ```go
  if edad >= 18 {
      fmt.Println("Mayor de edad")
  } else {
      fmt.Println("Menor de edad")
  }
  ```
  - No requiere paréntesis
  - Las llaves son obligatorias
  - Puede incluir una declaración inicial

- **Switch statements**
  ```go
  switch dia {
  case "lunes":
      fmt.Println("Inicio de semana")
  case "viernes":
      fmt.Println("Fin de semana")
  default:
      fmt.Println("Día normal")
  }
  ```
  - No requiere `break` (implícito)
  - Puede usar `fallthrough` para continuar
  - Soporta múltiples casos

- **Switch con expresión**
  ```go
  switch {
  case edad < 18:
      fmt.Println("Menor de edad")
  case edad >= 18 && edad < 65:
      fmt.Println("Adulto")
  default:
      fmt.Println("Adulto mayor")
  }
  ```
  - Evalúa expresiones booleanas
  - Útil para rangos de valores

- **Switch sin expresión (type switch)**
  ```go
  switch v := i.(type) {
  case int:
      fmt.Println("Es un entero")
  case string:
      fmt.Println("Es un string")
  default:
      fmt.Println("Tipo desconocido")
  }
  ```
  - Útil para type assertions
  - Manejo de interfaces

#### Bucles
- **For loop tradicional**
  ```go
  for i := 0; i < 10; i++ {
      fmt.Println(i)
  }
  ```
  - Sintaxis: inicialización; condición; incremento
  - Similar a C/Java

- **For como while**
  ```go
  for i < 10 {
      i++
  }
  ```
  - Solo condición
  - Equivalente a while en otros lenguajes

- **For range**
  ```go
  for index, value := range slice {
      fmt.Println(index, value)
  }
  ```
  - Itera sobre slices, maps, strings
  - Proporciona índice y valor
  - Puede ignorar valores con _

- **Break y continue**
  ```go
  for {
      if condition {
          break
      }
      if skip {
          continue
      }
  }
  ```
  - `break`: Sale del bucle
  - `continue`: Salta a la siguiente iteración

- **Labels para break/continue**
  ```go
  outer:
  for i := 0; i < 10; i++ {
      for j := 0; j < 10; j++ {
          if i == 5 && j == 5 {
              break outer
          }
      }
  }
  ```
  - Permite salir de bucles anidados
  - Útil para control de flujo complejo

#### Control de Flujo
- **Defer**
  ```go
  func ejemplo() {
      defer fmt.Println("Final")
      fmt.Println("Inicio")
  }
  ```
  - Ejecuta la función al final del scope
  - Útil para limpieza de recursos
  - Se ejecuta en orden LIFO

- **Panic y Recover**
  ```go
  func ejemplo() {
      defer func() {
          if r := recover(); r != nil {
              fmt.Println("Recuperado:", r)
          }
      }()
      panic("Error crítico")
  }
  ```
  - `panic`: Detiene la ejecución
  - `recover`: Recupera el control después de un panic
  - Útil para manejo de errores críticos

- **Goto**
  ```go
  if error {
      goto cleanup
  }
  cleanup:
      // código de limpieza
  ```
  - Raramente usado
  - Útil para manejo de errores en casos específicos
  - Considerado mala práctica en general

## 2. Funciones y Métodos

### 2.1 Funciones Básicas

#### Declaración de Funciones
- **Sintaxis básica**
  ```go
  func nombreFuncion(param1 tipo1, param2 tipo2) tipoRetorno {
      // código
      return valor
  }
  ```
  - Palabra clave `func`
  - Parámetros con tipos
  - Tipo de retorno opcional

- **Parámetros y tipos de retorno**
  ```go
  func suma(a, b int) int {
      return a + b
  }
  ```
  - Múltiples parámetros del mismo tipo
  - Retorno único o múltiple
  - Nombres de retorno opcionales

- **Múltiples valores de retorno**
  ```go
  func divide(a, b float64) (float64, error) {
      if b == 0 {
          return 0, errors.New("división por cero")
      }
      return a / b, nil
  }
  ```
  - Útil para manejo de errores
  - Pueden ser nombrados
  - Común en la biblioteca estándar

- **Funciones anónimas**
  ```go
  func() {
      fmt.Println("Función anónima")
  }()
  ```
  - Sin nombre
  - Pueden ser asignadas a variables
  - Útiles para callbacks

#### Tipos de Funciones
- **Funciones como valores**
  ```go
  var operacion func(int, int) int
  operacion = suma
  resultado := operacion(5, 3)
  ```
  - Pueden ser asignadas a variables
  - Tipos de función como tipos de datos
  - Útiles para callbacks

- **Funciones como parámetros**
  ```go
  func aplicar(f func(int) int, x int) int {
      return f(x)
  }
  ```
  - Higher-order functions
  - Callbacks
  - Patrones de diseño funcional

- **Closures**
  ```go
  func contador() func() int {
      count := 0
      return func() int {
          count++
          return count
      }
  }
  ```
  - Acceso a variables del scope exterior
  - Mantienen estado entre llamadas
  - Útiles para generadores

- **Funciones recursivas**
  ```go
  func factorial(n int) int {
      if n <= 1 {
          return 1
      }
      return n * factorial(n-1)
  }
  ```
  - Llamadas a sí mismas
  - Caso base necesario
  - Considerar límites de stack

### 2.2 Métodos y Interfaces

#### Métodos
- **Definición de métodos**
  ```go
  type Persona struct {
      Nombre string
      Edad   int
  }

  func (p Persona) Saludar() {
      fmt.Printf("Hola, soy %s\n", p.Nombre)
  }
  ```
  - Asociados a tipos
  - Receiver entre `func` y nombre
  - Acceso a campos del tipo

- **Receivers (value y pointer)**
  ```go
  func (p *Persona) CumplirAños() {
      p.Edad++
  }
  ```
  - Value receiver: copia del valor
  - Pointer receiver: referencia al valor
  - Consideraciones de performance

- **Métodos para tipos básicos**
  ```go
  type MiInt int

  func (m MiInt) Doble() MiInt {
      return m * 2
  }
  ```
  - Extender tipos básicos
  - Crear tipos personalizados
  - Añadir funcionalidad

- **Interfaces**
  ```go
  type Animal interface {
      HacerSonido() string
      Mover() string
  }
  ```
  - Conjunto de métodos
  - Implementación implícita
  - Polimorfismo

#### Características Avanzadas
- **Defer en funciones**
  ```go
  func procesarArchivo() {
      file := abrirArchivo()
      defer file.Close()
      // procesar archivo
  }
  ```
  - Limpieza de recursos
  - Manejo de errores
  - Orden de ejecución

- **Panic y recover**
  ```go
  func recuperar() {
      defer func() {
          if r := recover(); r != nil {
              fmt.Println("Recuperado:", r)
          }
      }()
      // código que puede panic
  }
  ```
  - Manejo de errores críticos
  - Recuperación de pánico
  - Logging y debugging

- **Funciones variádicas**
  ```go
  func suma(nums ...int) int {
      total := 0
      for _, n := range nums {
          total += n
      }
      return total
  }
  ```
  - Número variable de argumentos
  - Tratadas como slice
  - Útiles para utilidades

- **Funciones init()**
  ```go
  func init() {
      // inicialización
  }
  ```
  - Ejecutada automáticamente
  - Una por archivo
  - Orden de ejecución

## 3. Estructuras de Datos

### 3.1 Structs y Slices

#### Structs (Estructuras)
- **Definición de structs**
  ```go
  type Persona struct {
      Nombre string
      Edad   int
      Email  string
  }
  ```
  - Colección de campos
  - Tipos personalizados
  - Similar a clases

- **Campos y etiquetas**
  ```go
  type Usuario struct {
      Nombre string `json:"nombre"`
      Edad   int    `json:"edad"`
  }
  ```
  - Metadatos para campos
  - Serialización
  - Validación

- **Métodos para structs**
  ```go
  func (p *Persona) CumplirAños() {
      p.Edad++
  }
  ```
  - Comportamiento asociado
  - Encapsulamiento
  - OOP en Go

- **Composición de structs**
  ```go
  type Empleado struct {
      Persona
      Puesto string
  }
  ```
  - Herencia por composición
  - Reutilización de código
  - Delegación

- **Structs anónimos**
  ```go
  persona := struct {
      Nombre string
      Edad   int
  }{
      Nombre: "Juan",
      Edad:   25,
  }
  ```
  - Sin nombre
  - Útiles para casos específicos
  - JSON unmarshaling

#### Slices
- **Creación y manipulación**
  ```go
  slice := []int{1, 2, 3}
  slice = append(slice, 4)
  ```
  - Arrays dinámicos
  - Tamaño flexible
  - Operaciones comunes

- **Capacidad y longitud**
  ```go
  slice := make([]int, 5, 10)
  ```
  - `len`: número de elementos
  - `cap`: capacidad total
  - Reasignación automática

- **Append y copy**
  ```go
  slice1 := []int{1, 2, 3}
  slice2 := make([]int, len(slice1))
  copy(slice2, slice1)
  ```
  - Añadir elementos
  - Copiar slices
  - Manejo de memoria

- **Slicing de slices**
  ```go
  slice := []int{1, 2, 3, 4, 5}
  subslice := slice[1:4]
  ```
  - Subconjuntos
  - Referencias
  - Modificaciones

- **Slices de structs**
  ```go
  type Persona struct {
      Nombre string
      Edad   int
  }
  personas := []Persona{
      {Nombre: "Juan", Edad: 25},
      {Nombre: "María", Edad: 30},
  }
  ```
  - Colecciones de objetos
  - Iteración
  - Manipulación

### 3.2 Arrays y Maps

#### Arrays
- **Arrays fijos**
  ```go
  var arr [5]int
  arr := [5]int{1, 2, 3, 4, 5}
  ```
  - Tamaño fijo
  - Tipado por tamaño
  - Valores por defecto

- **Arrays vs Slices**
  ```go
  arr := [5]int{1, 2, 3, 4, 5}
  slice := arr[:]
  ```
  - Arrays: tamaño fijo
  - Slices: tamaño dinámico
  - Conversión entre ambos

- **Arrays multidimensionales**
  ```go
  matrix := [3][3]int{
      {1, 2, 3},
      {4, 5, 6},
      {7, 8, 9},
  }
  ```
  - Matrices
  - Arrays de arrays
  - Acceso por índices

- **Conversión entre arrays y slices**
  ```go
  arr := [5]int{1, 2, 3, 4, 5}
  slice := arr[:]
  ```
  - Slicing de arrays
  - Copia de arrays
  - Manipulación

#### Maps
- **Creación y manipulación**
  ```go
  m := make(map[string]int)
  m["uno"] = 1
  ```
  - Diccionarios
  - Clave-valor
  - Tipos personalizados

- **Operaciones comunes**
  ```go
  valor, existe := m["clave"]
  delete(m, "clave")
  ```
  - Inserción
  - Eliminación
  - Búsqueda

- **Maps de structs**
  ```go
  type Persona struct {
      Nombre string
      Edad   int
  }
  personas := map[string]Persona{
      "juan": {Nombre: "Juan", Edad: 25},
  }
  ```
  - Estructuras como valores
  - Referencias
  - Modificaciones

- **Maps anidados**
  ```go
  m := map[string]map[string]int{
      "a": {"x": 1, "y": 2},
      "b": {"x": 3, "y": 4},
  }
  ```
  - Estructuras complejas
  - Acceso anidado
  - Inicialización

## 4. Punteros y Concurrencia

### 4.1 Punteros

#### Conceptos Básicos
- **Declaración y uso**
  ```go
  var p *int
  x := 42
  p = &x
  ```
  - Referencias a memoria
  - Operador de dirección
  - Desreferencia

- **Operador de dirección (&)**
  ```go
  x := 42
  p := &x
  ```
  - Obtiene dirección
  - Referencia a variable
  - Tipado

- **Operador de desreferencia (*)**
  ```go
  x := 42
  p := &x
  *p = 43
  ```
  - Accede al valor
  - Modifica valor
  - Null safety

- **Punteros a structs**
  ```go
  type Persona struct {
      Nombre string
      Edad   int
  }
  p := &Persona{Nombre: "Juan", Edad: 25}
  ```
  - Acceso a campos
  - Modificación
  - Métodos

- **Punteros a funciones**
  ```go
  var f func(int) int
  f = func(x int) int { return x * 2 }
  ```
  - Callbacks
  - Funciones como valores
  - Patrones de diseño

### 4.2 Concurrencia

#### Concurrencia Básica
- **Goroutines**
  ```go
  go func() {
      // código concurrente
  }()
  ```
  - Threads ligeros
  - Múltiples goroutines
  - Sincronización

- **Channels**
  ```go
  ch := make(chan int)
  go func() {
      ch <- 42
  }()
  valor := <-ch
  ```
  - Comunicación entre goroutines
  - Buffered/unbuffered
  - Select

- **Select**
  ```go
  select {
  case msg1 := <-ch1:
      fmt.Println(msg1)
  case msg2 := <-ch2:
      fmt.Println(msg2)
  }
  ```
  - Múltiples channels
  - Timeout
  - Default case

- **WaitGroups**
  ```go
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
      defer wg.Done()
      // trabajo
  }()
  wg.Wait()
  ```
  - Sincronización
  - Espera de goroutines
  - Conteo

- **Mutex**
  ```go
  var mu sync.Mutex
  mu.Lock()
  // código crítico
  mu.Unlock()
  ```
  - Exclusión mutua
  - Protección de datos
  - Deadlock prevention

#### Patrones de Concurrencia
- **Worker pools**
  ```go
  jobs := make(chan int, 100)
  results := make(chan int, 100)
  for w := 1; w <= 3; w++ {
      go worker(w, jobs, results)
  }
  ```
  - Procesamiento paralelo
  - Control de recursos
  - Balanceo de carga

- **Pipeline**
  ```go
  func pipeline(input <-chan int) <-chan int {
      output := make(chan int)
      go func() {
          for v := range input {
              output <- v * 2
          }
          close(output)
      }()
      return output
  }
  ```
  - Procesamiento en etapas
  - Transformación de datos
  - Composición

- **Fan-in/Fan-out**
  ```go
  func fanOut(input <-chan int, workers int) []<-chan int {
      channels := make([]<-chan int, workers)
      for i := 0; i < workers; i++ {
          channels[i] = worker(input)
      }
      return channels
  }
  ```
  - Distribución de trabajo
  - Agregación de resultados
  - Paralelismo

- **Context**
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  ```
  - Cancelación
  - Timeout
  - Valores

- **Timeout y cancelación**
  ```go
  select {
  case <-time.After(5 * time.Second):
      fmt.Println("timeout")
  case <-ctx.Done():
      fmt.Println("cancelado")
  }
  ```
  - Control de tiempo
  - Cancelación
  - Recursos

#### Sincronización
- **Mutex y RWMutex**
  ```go
  var mu sync.RWMutex
  mu.RLock()
  // lectura
  mu.RUnlock()
  ```
  - Exclusión mutua
  - Lectura/escritura
  - Performance

- **Atomic operations**
  ```go
  var value atomic.Value
  value.Store(42)
  v := value.Load()
  ```
  - Operaciones atómicas
  - Thread safety
  - Performance

- **Once**
  ```go
  var once sync.Once
  once.Do(func() {
      // inicialización
  })
  ```
  - Inicialización única
  - Thread safety
  - Lazy loading

- **Cond**
  ```go
  var cond sync.Cond
  cond.L.Lock()
  for !ready {
      cond.Wait()
  }
  ```
  - Espera condicional
  - Señalización
  - Broadcast

- **Semaphore**
  ```go
  sem := make(chan struct{}, 3)
  sem <- struct{}{} // acquire
  <-sem            // release
  ```
  - Control de recursos
  - Límites de concurrencia
  - Rate limiting

## 5. Proyecto Práctico: Gestor de Tareas

### 5.1 Implementación

#### Estructuras de Datos
- **Structs para tareas**
  ```go
  type Task struct {
      TaskName string
      Status   string
      Date     string
  }
  ```
  - Modelo de datos
  - Campos necesarios
  - Validación

- **Slices para almacenamiento**
  ```go
  var tasks []Task
  ```
  - Colección de tareas
  - Operaciones CRUD
  - Persistencia

- **JSON para persistencia**
  ```go
  type Task struct {
      TaskName string `json:"taskName"`
      Status   string `json:"status"`
      Date     string `json:"date"`
  }
  ```
  - Serialización
  - Deserialización
  - Almacenamiento

#### Funcionalidades
- **Crear tareas**
  ```go
  func AddTask(tasks *[]Task) {
      // implementación
  }
  ```
  - Entrada de usuario
  - Validación
  - Almacenamiento

- **Marcar tareas como completadas**
  ```go
  func markTaskAsComplete(tasks *[]Task) {
      // implementación
  }
  ```
  - Búsqueda
  - Modificación
  - Estado

- **Listar tareas**
  ```go
  func print_tasks(tasks *[]Task) {
      // implementación
  }
  ```
  - Iteración
  - Formateo
  - Presentación

- **Eliminar tareas**
  ```go
  func RemoveTask(tasks *[]Task) {
      // implementación
  }
  ```
  - Búsqueda
  - Eliminación
  - Actualización

- **Persistencia de datos**
  ```go
  func saveTasksToFile(tasks *[]Task) error {
      // implementación
  }
  ```
  - Archivos
  - JSON
  - Manejo de errores

### 5.2 Buenas Prácticas

#### Desarrollo
- **Manejo de errores**
  ```go
  if err != nil {
      return fmt.Errorf("error: %v", err)
  }
  ```
  - Propagación
  - Logging
  - Recuperación

- **Documentación de código**
  ```go
  // AddTask añade una nueva tarea a la lista
  func AddTask(tasks *[]Task) {
      // implementación
  }
  ```
  - Comentarios
  - Documentación
  - Ejemplos

- **Estructura del proyecto**
  ```
  /task_man
    /cmd
    /internal
    /pkg
    /test
  ```
  - Organización
  - Modularidad
  - Mantenibilidad

- **Testing**
  ```go
  func TestAddTask(t *testing.T) {
      // implementación
  }
  ```
  - Unit tests
  - Integration tests
  - Benchmarks

#### Concurrencia
- **Manejo de race conditions**
  ```go
  var mu sync.Mutex
  mu.Lock()
  defer mu.Unlock()
  ```
  - Sincronización
  - Protección
  - Detección

- **Deadlock prevention**
  ```go
  // Orden consistente de locks
  mu1.Lock()
  mu2.Lock()
  defer mu2.Unlock()
  defer mu1.Unlock()
  ```
  - Orden de locks
  - Timeout
  - Detección

- **Resource management**
  ```go
  defer file.Close()
  defer db.Close()
  ```
  - Limpieza
  - Cierre
  - Recursos

- **Testing concurrent code**
  ```go
  func TestConcurrent(t *testing.T) {
      // implementación
  }
  ```
  - Race detector
  - Stress testing
  - Coverage

## Recursos Adicionales

1. **Documentación Oficial**
   - [Go Documentation](https://golang.org/doc/)
   - [Go by Example](https://gobyexample.com/)
   - [Go Tour](https://tour.golang.org/)

2. **Herramientas**
   - Go Playground
   - Go Modules
   - Go Testing
   - Go Lint

3. **Comunidad**
   - Stack Overflow
   - Go Forum
   - Go Slack
   - GitHub Discussions

## Consejos para el Estudio

1. **Práctica Regular**
   - Escribe código todos los días
   - Completa los ejercicios de cada sección
   - Experimenta con el código de ejemplo

2. **Proyectos Personales**
   - Crea pequeños proyectos
   - Implementa las características aprendidas
   - Contribuye a proyectos open source

3. **Revisión**
   - Revisa el código regularmente
   - Refactoriza cuando sea necesario
   - Aprende de los errores

4. **Comunidad**
   - Participa en foros
   - Comparte tu código
   - Aprende de otros desarrolladores

## Siguientes Pasos

1. Comienza con los conceptos básicos
2. Practica con los ejercicios de cada sección
3. Implementa el proyecto de gestión de tareas
4. Explora temas más avanzados
5. Contribuye a la comunidad

Recuerda que el aprendizaje es un proceso continuo. Esta guía te proporciona una base sólida, pero siempre hay más que aprender en el mundo de Go. 