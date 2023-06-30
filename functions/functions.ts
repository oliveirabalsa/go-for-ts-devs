function main() {
  console.log("Hello World");
  withParams("Leonardo");
  console.log(oneReturnedValue("Leonardo"));
  console.log(twoReturnedValues("Leonardo"));
  const [name, length] = twoReturnedValues("Leonardo");
  console.log(name, length);
  console.log(withTypeReturn("Leonardo"));
}

type User = {
  name: string;
  age: number;
  email: string;
};

function withParams(name: string) {
  console.log(`Hello ${name}`);
}

function oneReturnedValue(name: string): string {
  return `Hello ${name}`;
}

function twoReturnedValues(name: string): [string, number] {
  return [name, name.length];
}

function withTypeReturn(name: string): User {
  return {
    name: name,
    age: 21,
    email: "email@email.com",
  };
}

main();
