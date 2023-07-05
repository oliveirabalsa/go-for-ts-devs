function main() {
  // Map creation and initialization
  let ages: Map<string, number> = new Map<string, number>();
  ages.set("Alice", 25);
  ages.set("Bob", 30);
  ages.set("Charlie", 35);
  ages.set("Leonardo", 28);

  // Accessing values in a map
  let aliceAge: number | undefined = ages.get("Alice");
  console.log("Alice's age:", aliceAge);

  // Checking if a key exists in a map
  if (ages.has("Bob")) {
    let bobAge: number | undefined = ages.get("Bob");
    console.log("Bob's age:", bobAge);
  } else {
    console.log("Bob's age not found");
  }

  // Deleting a key-value pair from a map
  ages.delete("Charlie");
  console.log("Charlie's age deleted");

  // Iterating over a map
  for (let [name, age] of ages) {
    console.log(`${name}'s age: ${age}`);
  }
}

main();
