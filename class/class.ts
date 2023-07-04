class Person {
  private name: string;
  private age: number;
  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }

  public getName(): string {
    return this.name;
  }

  public getAge(): number {
    return this.age;
  }

  public setName(name: string) {
    this.name = name;
  }

  public setAge(age: number) {
    this.age = age;
  }

  public greet(): void {
    console.log(
      `Hello, my name is ${this.name} and I am ${this.age} years old.`
    );
  }
}

function classes() {
  // Class instance
  const person = new Person("Bianca", 25);
  person.greet();

  // Accessing class methods
  const name = person.getName();
  const age = person.getAge();
  console.log(`Name: ${name}, Age: ${age}`);

  person.setName("Leo");
  person.setAge(30);
  person.greet();
}

classes();
