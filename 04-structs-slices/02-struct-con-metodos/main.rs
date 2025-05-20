// Definición de la estructura Person
struct Person {
    name: String,
    edad: i32, // Usamos i32 para representar un entero similar a int en Go
}

// Implementación de métodos para Person
impl Person {
    // Método que devuelve una descripción
    fn description(&self) -> String {
        format!("El nombre de la persona es {}", self.name)
    }
}

fn main() {
    println!("Structure con métodos");

    // Crear una nueva instancia de Person
    let new_person = Person {
        edad: 16,
        name: String::from("Luis"),
    };

    // Llamar al método description
    println!("Description: {}", new_person.description());
}