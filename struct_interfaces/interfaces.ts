interface PersonInterface {
  name: string;
  age: number;
}

function interfaces() {
  // Struct initialization
  const person: PersonInterface = { name: "Alice", age: 25 };
  console.log("Person:", person);

  // Struct field access
  console.log("Name:", person.name);
  console.log("Age:", person.age);
}

interfaces();
