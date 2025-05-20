// Definición de la clase Person
class Person {
    constructor(public name: string, public edad: number) {}

    // Método para obtener la descripción
    description(): string {
        return `El nombre de la persona es ${this.name}`;
    }
}

// Función principal
function main() {
    console.log("Structure con métodos");
    const newPerson = new Person("Luis", 16);
    console.log(`Description: ${newPerson.description()}`);
}

// Ejecutar la función principal
main();