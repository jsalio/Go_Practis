// Executor function that takes a function as a parameter
function executor(a: number, b: number, local: (a: number, b: number) => number): number {
    return local(a, b);
}

// Addition function
function sumar(a: number, b: number): number {
    return a + b;
}

// Subtraction function
function restar(a: number, b: number): number {
    return a - b;
}

// Main execution
console.log("Función como parámetro");
console.log(`La suma es ${executor(8, 6, sumar)} y la resta es ${executor(8, 6, restar)}`);