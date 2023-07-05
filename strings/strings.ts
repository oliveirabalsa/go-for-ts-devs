function main() {
  // String concatenation
  const str1: string = "Hello";
  const str2: string = "World";
  const result: string = str1 + " " + str2;
  console.log(result);

  // String length
  const str: string = "Hello, World!";
  const length: number = str.length;
  console.log("Length:", length);

  // String splitting
  const split: string[] = str.split(" ");
  console.log("Split result:", split);

  // String substring
  const substr: string = str.substring(0, 5);
  console.log("Substring:", substr);

  // String contains
  const contains: boolean = str.includes("World");
  console.log("Contains:", contains);
}

main();
