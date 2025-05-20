# From Zero to Junior: Go Programming Journey

Este repositorio contiene una colección de ejercicios y ejemplos prácticos para aprender Go desde cero hasta un nivel junior. El contenido está organizado en módulos progresivos que cubren los conceptos fundamentales del lenguaje.

## Estructura del Repositorio

El repositorio está organizado en los siguientes módulos:

1. **01-Conceptos_Básicos**: Fundamentos del lenguaje Go
   - Sintaxis básica
   - Variables y tipos de datos
   - Operadores
   - Entrada/Salida básica

2. **02-Estructura de control**: Control de flujo en Go
   - Condicionales (if, else, switch)
   - Bucles (for)
   - Control de flujo avanzado

3. **03-Funciones**: Funciones y métodos en Go
   - Declaración y uso de funciones
   - Parámetros y retornos
   - Funciones anónimas
   - Métodos

4. **04-structs-slices**: Estructuras de datos
   - Structs
   - Slices
   - Arrays
   - Maps

5. **05-punteros-y-concurrencia**: Conceptos avanzados
   - Punteros
   - Concurrencia
   - Goroutines
   - Channels

6. **Real_exercie**: Ejercicios prácticos reales
   - Proyectos y ejercicios integradores
   - Casos de uso reales

## 50 Ejercicios para Aprender Go desde lo Básico

Esta lista de 50 ejercicios está diseñada para un desarrollador experto en C# que desea aprender Go desde cero, cubriendo conceptos básicos e intermedios. Los ejercicios están organizados por temas para facilitar el aprendizaje progresivo.

### 1. Conceptos Básicos (1-10)

1. Hola Mundo: Escribe un programa que imprima "Hola, Mundo!" en la consola.
2. Variables: Declara variables de tipo int, string y bool. Imprime sus valores.
3. Constantes: Define una constante para el valor de pi y úsala en un cálculo simple.
4. Entrada de usuario: Lee un nombre desde la consola y salúdalo con un mensaje personalizado.
5. Tipos implícitos: Usa := para declarar una variable sin especificar el tipo e imprime su valor.
6. Operaciones aritméticas: Calcula la suma, resta, multiplicación y división de dos números ingresados por el usuario.
7. Conversión de tipos: Convierte un string ingresado a int para realizar una operación matemática.
8. Condicionales: Escribe un programa que determine si un número es par o impar.
9. Switch simple: Usa switch para imprimir el nombre de un día de la semana según un número (1-7).
10. Bucle for: Imprime los números del 1 al 10 usando un bucle for.

### 2. Estructuras de Control (11-20)

11. Bucle infinito: Crea un bucle infinito que imprima un contador y se detenga al llegar a 5 usando break.
12. Continue: Imprime solo números impares del 1 al 20 usando continue.
13. If anidado: Determina si un número es positivo, negativo o cero, y si es par o impar.
14. Switch sin expresión: Usa un switch sin expresión para verificar rangos de edad (niño, joven, adulto).
15. Bucle anidado: Imprime una tabla de multiplicar (1 al 10) usando bucles anidados.
16. Múltiplos de 3: Imprime los múltiplos de 3 entre 1 y 50.
17. Validación de entrada: Pide un número entre 1 y 100; si no es válido, sigue pidiéndolo.
18. Factorial: Calcula el factorial de un número ingresado usando un bucle.
19. Números primos: Determina si un número ingresado es primo.
20. Serie Fibonacci: Imprime los primeros 10 números de la serie Fibonacci.

### 3. Funciones (21-30)

21. Función simple: Escribe una función que reciba dos int y devuelva su suma.
22. Valores múltiples: Crea una función que devuelva el cuadrado y el cubo de un número.
23. Parámetros nombrados: Escribe una función que use valores de retorno nombrados para calcular área y perímetro de un rectángulo.
24. Función recursiva: Implementa una función recursiva para calcular el factorial.
25. Función como parámetro: Pasa una función como parámetro para aplicar una operación (suma o resta) a dos números.
26. Cierre: Crea un cierre que genere una secuencia de números incrementales.
27. Variádica: Escribe una función variádica que calcule el promedio de una lista de números.
28. Error handling: Crea una función que divida dos números y maneje la división por cero con un error.
29. Defer: Usa defer para cerrar un recurso simulado (como un archivo) después de usarlo.
30. Panic y recover: Simula un error con panic y usa recover para manejarlo.

### 4. Estructuras y Slices (31-40)

31. Struct básico: Define un struct para representar una persona (nombre, edad) e imprime sus datos.
32. Método en struct: Agrega un método a un struct para devolver una descripción de la persona.
33. Slice básico: Crea un slice de string, agrégale elementos e imprime su contenido.
34. Slice dinámico: Escribe un programa que permita al usuario agregar nombres a un slice hasta que ingrese "salir".
35. Copia de slice: Crea una copia de un slice y modifica la copia sin afectar el original.
36. Map básico: Usa un map para almacenar nombres y edades, e imprime sus valores.
37. Eliminar de map: Escribe una función que elimine una clave de un map.
38. Slice de structs: Crea un slice de structs Producto (nombre, precio) y calcula el precio total.
39. Ordenar slice: Ordena un slice de números en orden ascendente usando el paquete sort.
40. Filtrar slice: Filtra un slice de números para conservar solo los mayores a un valor dado.

### 5. Punteros y Concurrencia Básica (41-50)

41. Puntero simple: Usa un puntero para modificar el valor de una variable dentro de una función.
42. Puntero en struct: Modifica un campo de un struct usando un puntero.
43. Nil check: Escribe una función que verifique si un puntero es nil antes de usarlo.
44. Goroutine básica: Lanza una goroutine que imprima un mensaje 10 veces.
45. WaitGroup: Usa sync.WaitGroup para esperar que varias goroutines terminen.
46. Channel simple: Crea un canal para enviar y recibir un mensaje entre dos goroutines.
47. Channel con buffer: Usa un canal con buffer para enviar varios números y recibirlos en orden.
48. Select: Usa select para leer de dos canales alternadamente.
49. Mutex: Usa un sync.Mutex para proteger una variable compartida entre goroutines.
50. Productor-consumidor: Implementa un patrón productor-consumidor con un canal para procesar una lista de tareas.

### Consejos para el Aprendizaje

- **Entorno**: Usa Go Playground (play.golang.org) o instala Go localmente con go install.
- **Documentación**: Consulta la documentación oficial en golang.org/doc.
- **Práctica diaria**: Dedica 30-60 minutos diarios a resolver 2-3 ejercicios.
- **Proyectos pequeños**: Después de los ejercicios, intenta crear una CLI simple (como un administrador de tareas) para aplicar lo aprendido.
- **Comunidad**: Únete a foros como el subreddit r/golang o el canal de Slack de Gophers para dudas.

## Requisitos Previos

- Go 1.x instalado en tu sistema
- Un editor de código (recomendado: VSCode con extensiones para Go)
- Conocimientos básicos de programación

## Cómo Usar Este Repositorio

1. Clona este repositorio:
   ```bash
   git clone https://github.com/tu-usuario/from-zero-to-junior.git
   ```

2. Navega a través de los módulos en orden secuencial
3. Cada módulo contiene ejemplos y ejercicios prácticos
4. Completa los ejercicios en cada sección antes de avanzar al siguiente módulo

## Contribuciones

Las contribuciones son bienvenidas. Si encuentras algún error o tienes sugerencias para mejorar el contenido, por favor:

1. Abre un issue describiendo el problema o sugerencia
2. Crea un pull request con tus cambios

## Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## Contacto

Si tienes preguntas o sugerencias, no dudes en abrir un issue en el repositorio. 