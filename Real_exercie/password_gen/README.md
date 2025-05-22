# Generador de Contraseñas en Go

Este proyecto implementa un generador de contraseñas seguras en Go que permite crear contraseñas personalizadas con diferentes características.

## Características Principales

- Generación de contraseñas con longitud personalizable (8-128 caracteres)
- Opciones para incluir:
  - Letras minúsculas (siempre incluidas)
  - Letras mayúsculas
  - Números
  - Símbolos especiales
- Generación de múltiples contraseñas en una sola ejecución
- Uso de criptografía segura para la generación de números aleatorios

## Estructura del Código

### Tipos de Caracteres (CharType)

El código utiliza un sistema de bits para representar los diferentes tipos de caracteres:

```go
type CharType uint

const (
    Lowercase CharType = 1 << iota // 0001
    Uppercase                      // 0010
    Number                         // 0100
    Symbol                         // 1000
)
```

#### Explicación del Tipo Personalizado y Sistema de Bits

1. **Definición del Tipo Personalizado**:
   - `type CharType uint`: Se define un nuevo tipo basado en `uint` (entero sin signo)
   - Este tipo personalizado permite:
     - Mayor claridad en el código al usar un nombre específico para el dominio
     - Prevenir errores al mezclar tipos de datos
     - Implementar métodos específicos para este tipo si fuera necesario
     - Mejor documentación y mantenibilidad del código

2. **Sistema de Bits y Combinaciones**:
   - El uso de bits permite representar múltiples estados en un solo número
   - Cada tipo de carácter se representa con un bit específico:
     ```
     Lowercase = 0001 (1 en decimal)
     Uppercase = 0010 (2 en decimal)
     Number    = 0100 (4 en decimal)
     Symbol    = 1000 (8 en decimal)
     ```
   - Ventajas de este sistema:
     - **Eficiencia de Memoria**: Un solo número puede representar múltiples estados
     - **Operaciones Rápidas**: Las operaciones bit a bit son muy eficientes
     - **Combinaciones Fáciles**: Se pueden combinar tipos usando operadores bit a bit
     - **Verificación Simple**: Se puede verificar la presencia de un tipo usando AND (&)

3. **Uso de `iota`**:
   - `iota` es un contador especial en Go que se incrementa automáticamente
   - `1 << iota` desplaza el bit 1 a la izquierda en cada iteración
   - Esto crea automáticamente potencias de 2 (1, 2, 4, 8, ...)
   - Cada constante obtiene un valor único y no superpuesto

4. **Ejemplo de Combinaciones**:
   ```go
   // Combinación de Lowercase y Uppercase
   Lowercase | Uppercase = 0001 | 0010 = 0011 (3 en decimal)
   
   // Combinación de todos los tipos
   Lowercase | Uppercase | Number | Symbol = 0001 | 0010 | 0100 | 1000 = 1111 (15 en decimal)
   ```

5. **Verificación de Tipos**:
   ```go
   // Verificar si un tipo está presente
   if (combination & Lowercase) != 0 {
       // Lowercase está presente
   }
   ```

### Funciones Principales

#### generateAlphabet

Esta función genera un alfabeto de letras (mayúsculas o minúsculas) basado en un parámetro booleano:

```go
func generateAlphabet(uppercase bool) []string {
    var alphabet []string
    start := 'a'
    if uppercase {
        start = 'A'
    }
    for i := 0; i < 26; i++ {
        alphabet = append(alphabet, string(rune(int(start)+i)))
    }
    return alphabet
}
```

Explicación detallada:
1. `start := 'a'` - Define el punto de inicio como la letra 'a' en minúscula
2. Si `uppercase` es true, cambia el punto de inicio a 'A'
3. El bucle itera 26 veces (una por cada letra del alfabeto)
4. `string(rune(int(start)+i))` - Esta expresión:
   - Convierte `start` a entero (valor ASCII)
   - Suma el índice `i` para avanzar en el alfabeto
   - Convierte el resultado a `rune` (carácter Unicode)
   - Convierte el `rune` a string
   - Agrega el resultado al slice `alphabet`

#### RandomCombo

Esta función selecciona aleatoriamente un tipo de carácter de una lista de combinaciones:

```go
func RandomCombo(combinations []CharType) (CharType, error) {
    n, err := rand.Int(rand.Reader, big.NewInt(int64(len(combinations))))
    if err != nil {
        return 0, fmt.Errorf("error generating random number: %w", err)
    }
    return combinations[n.Int64()], nil
}
```

Explicación detallada:
1. `rand.Int(rand.Reader, big.NewInt(int64(len(combinations))))`:
   - `rand.Reader` es una fuente de números aleatorios criptográficamente segura
   - `big.NewInt(int64(len(combinations)))` crea un número grande que representa el límite superior
   - `rand.Int` genera un número aleatorio entre 0 y el límite superior (exclusivo)
2. El número generado se usa como índice para seleccionar un tipo de carácter de la lista de combinaciones

### Funciones de Generación de Contraseñas

#### buildMapCombination

Esta función crea una lista de tipos de caracteres basada en las preferencias del usuario:

```go
func buildMapCombination(includeUpCase, includeNumber, includeSymbols bool) []CharType {
    var comb []CharType
    comb = append(comb, Lowercase)  // Siempre incluye minúsculas
    if includeUpCase {
        comb = append(comb, Uppercase)
    }
    if includeNumber {
        comb = append(comb, Number)
    }
    if includeSymbols {
        comb = append(comb, Symbol)
    }
    return comb
}
```

Explicación detallada:
1. La función toma tres parámetros booleanos que representan las preferencias del usuario
2. Crea un slice vacío de tipo `CharType`
3. Siempre incluye `Lowercase` como tipo base
4. Agrega condicionalmente los otros tipos según las preferencias del usuario
5. Retorna el slice con los tipos seleccionados

#### BuildMap

Esta función crea un mapa que relaciona cada tipo de carácter con su función generadora correspondiente:

```go
func BuildMap() map[CharType]func() (string, error) {
    maps := make(map[CharType]func() (string, error))
    maps[Lowercase] = GetLowerCaseLetter
    maps[Uppercase] = GetUpperCaseLetter
    maps[Number] = GetNumberLetter
    maps[Symbol] = GetSymbol
    return maps
}
```

Explicación detallada:
1. Crea un mapa donde:
   - Las claves son de tipo `CharType`
   - Los valores son funciones que retornan un string y un error
2. Asocia cada tipo de carácter con su función generadora específica:
   - `Lowercase` → `GetLowerCaseLetter`
   - `Uppercase` → `GetUpperCaseLetter`
   - `Number` → `GetNumberLetter`
   - `Symbol` → `GetSymbol`
3. Este mapa permite una fácil selección de la función generadora correcta

#### ExecuteCombination

Esta función es el núcleo del generador de contraseñas, que crea una contraseña de la longitud especificada:

```go
func ExecuteCombination(limite int, operations []CharType) (string, error) {
    functionMap := BuildMap()
    pdw := ""
    for x := 1; x <= limite; x++ {
        ctype, err := RandomCombo(operations)
        if err != nil {
            return "", fmt.Errorf("error selecting character type: %w", err)
        }
        char, err := functionMap[ctype]()
        if err != nil {
            return "", fmt.Errorf("error generating character: %w", err)
        }
        pdw += char
    }
    return pdw, nil
}
```

Explicación detallada del proceso:
1. **Inicialización**:
   - Obtiene el mapa de funciones usando `BuildMap()`
   - Inicializa una cadena vacía para la contraseña

2. **Bucle de Generación**:
   - Itera `limite` veces (una por cada carácter de la contraseña)
   - En cada iteración:
     1. Selecciona aleatoriamente un tipo de carácter usando `RandomCombo`
     2. Obtiene la función generadora correspondiente del mapa
     3. Genera un carácter aleatorio usando la función seleccionada
     4. Agrega el carácter a la contraseña

3. **Manejo de Errores**:
   - Verifica errores en la selección del tipo de carácter
   - Verifica errores en la generación del carácter
   - Retorna errores descriptivos si algo falla

4. **Resultado**:
   - Retorna la contraseña generada y nil si todo fue exitoso
   - Retorna una cadena vacía y un error si algo falló

### Flujo de Generación de Contraseñas

1. El usuario especifica sus preferencias
2. `buildMapCombination` crea la lista de tipos de caracteres permitidos
3. `BuildMap` crea el mapa de funciones generadoras
4. `ExecuteCombination` genera la contraseña:
   - Selecciona aleatoriamente tipos de caracteres
   - Genera caracteres aleatorios de cada tipo
   - Combina los caracteres en una contraseña final

## Uso del Programa

1. Ejecutar el programa
2. Ingresar la longitud deseada de la contraseña
3. Responder "yes" o "no" a las preguntas sobre tipos de caracteres
4. Especificar cuántas contraseñas generar
5. El programa mostrará las contraseñas generadas

## Consideraciones de Seguridad

- El programa utiliza `crypto/rand` para generar números aleatorios criptográficamente seguros
- Las contraseñas generadas tienen una longitud mínima de 8 caracteres
- Se garantiza que cada contraseña contenga al menos un tipo de carácter seleccionado 