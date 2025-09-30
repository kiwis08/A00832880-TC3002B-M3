# Implementación de Estructuras de Datos en Go

Este proyecto contiene implementaciones completas de tres estructuras de datos fundamentales en Go:

1. **Stack (Pila)** - Estructura de datos Last-In-First-Out (LIFO)
2. **Queue (Cola)** - Estructura de datos First-In-First-Out (FIFO)  
3. **Hash Table (Tabla Hash)** - Mapeo clave-valor con operaciones promedio O(1)

## Como ejecutarlo

### Ejecutar el Programa Principal

Para correr el demo de las estructuras de datos:

```bash
# Navegar al directorio del proyecto
cd Tarea1

# Ejecutar el programa principal
go run main.go hash_table.go queue.go stack.go
```

### Ejecutar Tests

Para ejecutar todos los tests:

```bash
# Ejecutar todos los tests
go test

# Ejecutar tests con salida detallada
go test -v

# Ejecutar tests de una estructura específica
go test -run TestArrayStack
go test -run TestArrayQueue  
go test -run TestArrayHashTable
```

## Casos de Prueba (Test Cases)

### 1. Stack (Pila) - `stack_test.go`

#### **Estado Inicial**
- Verifica que una pila nueva esté vacía (`IsEmpty()` retorna `true`)
- Verifica que el tamaño inicial sea 0 (`Size()` retorna `0`)
- Verifica que `Peek()` en pila vacía retorne `nil`
- Verifica que `Pop()` en pila vacía retorne `nil`

#### **Operaciones Push y Peek**
- Prueba `Push()` con elementos individuales
- Verifica que `IsEmpty()` retorne `false` después de push
- Verifica que `Size()` se incremente correctamente
- Verifica que `Peek()` retorne el último elemento agregado (LIFO)

#### **Operación Pop**
- Verifica que `Pop()` retorne el elemento más recientemente agregado
- Verifica que `Size()` se decremente correctamente
- Verifica que `Peek()` retorne el nuevo elemento superior después de pop
- Verifica que la pila quede vacía después de popear todos los elementos

#### **Múltiples Operaciones**
- Agrega múltiples elementos: `["a", "b", "c", "d", "e"]`
- Verifica el orden LIFO al hacer pop: `["e", "d", "c", "b", "a"]`
- Verifica que la pila quede vacía al final

#### **Operación Clear**
- Agrega elementos a la pila
- Ejecuta `Clear()`
- Verifica que la pila quede vacía y con tamaño 0
- Verifica que `Peek()` retorne `nil` después de clear

#### **Diferentes Tipos de Datos**
- Prueba con `string`, `int`, `float64`, y `bool`
- Verifica que se mantenga el orden LIFO con diferentes tipos

### 2. Queue (Cola) - `queue_test.go`

#### **Estado Inicial**
- Verifica que una cola nueva esté vacía (`IsEmpty()` retorna `true`)
- Verifica que el tamaño inicial sea 0 (`Size()` retorna `0`)
- Verifica que `Front()` en cola vacía retorne `nil`
- Verifica que `Dequeue()` en cola vacía retorne `nil`

#### **Operaciones Enqueue y Front**
- Prueba `Enqueue()` con elementos individuales
- Verifica que `IsEmpty()` retorne `false` después de enqueue
- Verifica que `Size()` se incremente correctamente
- Verifica que `Front()` retorne el primer elemento agregado (FIFO)

#### **Operación Dequeue**
- Verifica que `Dequeue()` retorne el primer elemento agregado
- Verifica que `Size()` se decremente correctamente
- Verifica que `Front()` retorne el nuevo elemento frontal después de dequeue
- Verifica que la cola quede vacía después de desencolar todos los elementos

#### **Múltiples Operaciones**
- Agrega múltiples elementos: `["a", "b", "c", "d", "e"]`
- Verifica el orden FIFO al hacer dequeue: `["a", "b", "c", "d", "e"]`
- Verifica que la cola quede vacía al final

#### **Operación Clear**
- Agrega elementos a la cola
- Ejecuta `Clear()`
- Verifica que la cola quede vacía y con tamaño 0
- Verifica que `Front()` retorne `nil` después de clear

#### **Diferentes Tipos de Datos**
- Prueba con `string`, `int`, `float64`, y `bool`
- Verifica que se mantenga el orden FIFO con diferentes tipos

#### **Operaciones Alternadas**
- Prueba secuencias de enqueue/dequeue intercalados
- Verifica que se mantenga el orden FIFO en operaciones complejas

### 3. Hash Table (Tabla Hash) - `hash_table_test.go`

#### **Estado Inicial**
- Verifica que una tabla hash nueva tenga tamaño 0 (`Len()` retorna `0`)
- Verifica que `Get()` en clave inexistente retorne `false`
- Verifica que `Contains()` en clave inexistente retorne `false`
- Verifica que `Delete()` en clave inexistente retorne `false`
- Verifica que `Keys()` y `Values()` retornen slices vacíos

#### **Operaciones Básicas Put y Get**
- Prueba `Put()` con pares clave-valor
- Verifica que `Len()` se incremente correctamente
- Verifica que `Contains()` retorne `true` para claves existentes
- Verifica que `Get()` retorne el valor correcto

#### **Actualización de Clave Existente**
- Actualiza el valor de una clave existente
- Verifica que `Len()` no cambie
- Verifica que `Get()` retorne el nuevo valor
- Verifica que otras claves no se vean afectadas

#### **Operación Delete**
- Verifica que `Delete()` retorne `true` para claves existentes
- Verifica que `Len()` se decremente correctamente
- Verifica que `Contains()` retorne `false` para claves eliminadas
- Verifica que `Get()` retorne `false` para claves eliminadas
- Verifica que otras claves no se vean afectadas

#### **Diferentes Tipos de Claves String**
- Prueba con claves que representan diferentes tipos: `"string_key"`, `"42"`, `"3.14"`, `"true"`
- Verifica que todas las claves se manejen correctamente

#### **Operaciones Keys y Values**
- Agrega múltiples pares clave-valor
- Verifica que `Keys()` retorne todas las claves agregadas
- Verifica que `Values()` retorne todos los valores agregados
- Verifica que el orden no importe (usando mapas para verificación)

#### **Operación Clear**
- Ejecuta `Clear()`
- Verifica que `Len()` sea 0
- Verifica que `Keys()` y `Values()` retornen slices vacíos

#### **Colisiones de Hash**
- Agrega múltiples claves que pueden producir colisiones
- Verifica que todas las claves se manejen correctamente a pesar de las colisiones

#### **Tabla Hash con Capacidad Personalizada**
- Crea una tabla hash con capacidad específica
- Verifica que funcione correctamente con la capacidad dada

## Comportamiento Esperado

- **Stack**: Comportamiento LIFO - el último elemento agregado es el primero en ser removido
- **Queue**: Comportamiento FIFO - el primer elemento agregado es el primero en ser removido  
- **Hash Table**: Mapeo clave-valor con operaciones promedio O(1), maneja colisiones correctamente

## Características de las Implementaciones

### Stack
- Implementado usando slice de Go
- El "top" de la pila es el último elemento del slice
- Operaciones `Push()` y `Pop()` son O(1)

### Queue
- Implementado usando slice de Go con índices front y rear
- Operaciones `Enqueue()` y `Dequeue()` son O(1) amortizado
- Maneja eficientemente las operaciones de cola

### Hash Table
- Implementado usando separate chaining (slice de slices)
- Función hash personalizada para diferentes tipos de datos
- Manejo automático de colisiones
- Operaciones promedio O(1) para put, get y delete
