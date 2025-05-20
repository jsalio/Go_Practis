using System;

class Program
{
    // Define a delegate type for the function parameter
    delegate int Operation(int a, int b);

    // Executor function that takes a function as a parameter
    static int Executor(int a, int b, Operation local)
    {
        return local(a, b);
    }

    static int executor2(int a, int b, Func<int,int, int> fnc)
    {
        return fnc(a,b)
    }

    // Addition function
    static int Sumar(int a, int b)
    {
        return a + b;
    }

    // Subtraction function
    static int Restar(int a, int b)
    {
        return a - b;
    }

    static void Main(string[] args)
    {
        Console.WriteLine("Función como parámetro");
        Console.WriteLine($"La suma es {Executor(8, 6, Sumar)} y la resta es {Executor(8, 6, Restar)}");
    }
}