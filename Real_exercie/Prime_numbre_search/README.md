# Buscador de Números Primos

Este programa implementa un buscador de números primos utilizando concurrencia en Go para mejorar el rendimiento.

## Descripción

El programa busca números primos hasta un límite especificado por el usuario, utilizando múltiples goroutines para procesar diferentes rangos de números en paralelo.

## Características Principales

- Búsqueda concurrente de números primos
- Procesamiento en paralelo de diferentes rangos de números
- Medición del tiempo de ejecución
- Ordenamiento automático de resultados
- Interfaz de usuario simple por consola

## Funcionamiento

1. El programa solicita al usuario ingresar un límite máximo para la búsqueda (entre 0 y 1000)
2. Divide el rango de búsqueda en segmentos más pequeños
3. Procesa cada segmento en una goroutine separada
4. Recolecta y ordena los resultados
5. Muestra los números primos encontrados y el tiempo de ejecución

## Componentes Principales

### Funciones Principales

- `isPrime(x int)`: Verifica si un número es primo
- `calculateRange(min, max int, channel chan int, wg *sync.WaitGroup)`: Procesa un rango de números en busca de primos
- `ViewResults(channel chan int, wg *sync.WaitGroup)`: Recolecta y muestra los resultados
- `ProcessUserSearch(limit int)`: Coordina el proceso de búsqueda

### Detalle de ProcessUserSearch

La función `ProcessUserSearch` es el núcleo del programa y coordina todo el proceso de búsqueda de números primos. Su funcionamiento se puede desglosar en los siguientes pasos:

1. **Inicialización**:
   - Registra el tiempo de inicio para medir el rendimiento
   - Crea un canal para la comunicación entre goroutines
   - Define el tamaño de cada lote de procesamiento (100 números)
   - Calcula cuántos segmentos se necesitarán para procesar todo el rango

2. **Configuración de Goroutines**:
   - Inicia una goroutine para `ViewResults` que recolectará y mostrará los resultados
   - Para cada segmento del rango:
     - Calcula el rango mínimo y máximo para ese segmento
     - Inicia una goroutine con `calculateRange` para procesar ese segmento
     - Ajusta el rango máximo si excede el límite total

3. **Sincronización**:
   - Utiliza dos `WaitGroup`:
     - `calcWG`: Para esperar que todas las goroutines de cálculo terminen
     - `wg`: Para esperar que la goroutine de visualización termine
   - Cierra el canal cuando todos los cálculos han terminado

4. **Finalización**:
   - Espera a que se complete la visualización de resultados
   - Calcula y muestra el tiempo total de ejecución

Esta implementación permite un procesamiento eficiente al:
- Dividir el trabajo en segmentos manejables
- Procesar múltiples segmentos en paralelo
- Coordinar la recolección de resultados de manera ordenada
- Mantener un control preciso sobre la sincronización de goroutines

### Optimizaciones

- Uso de goroutines para procesamiento paralelo
- Optimización de la verificación de números primos
- Procesamiento por lotes de 100 números
- Uso de canales para comunicación entre goroutines

## Requisitos

- Go 1.x o superior

## Ejecución

Para ejecutar el programa:

```bash
go run main.go
```

## Ejemplo de Uso

1. Ejecuta el programa
2. Ingresa un número límite (por ejemplo, 100)
3. El programa mostrará:
   - Lista de números primos encontrados
   - Cantidad total de primos
   - Tiempo de ejecución

## Notas Técnicas

- El programa utiliza `sync.WaitGroup` para sincronización
- Implementa canales para comunicación entre goroutines
- Utiliza el paquete `math` para optimizar la verificación de números primos
- Incluye ordenamiento de resultados usando `sort.Ints` 