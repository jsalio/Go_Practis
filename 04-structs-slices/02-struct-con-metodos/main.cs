using System;

namespace PersonApp
{
    // Definición de la clase Person
    public class Person
    {
        public string Name { get; set; }
        public int Edad { get; set; }

        // Constructor para inicializar los campos
        public Person(string name, int edad)
        {
            Name = name;
            Edad = edad;
        }

        // Método para obtener la descripción
        public string Description()
        {
            return $"El nombre de la persona es {Name}";
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Structure con métodos");
            var newPerson = new Person("Luis", 16);
            Console.WriteLine($"Description: {newPerson.Description()}");
        }
    }
}