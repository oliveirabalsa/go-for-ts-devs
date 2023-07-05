import fs from "fs";

function main() {
  // Writing to a file
  fs.writeFile("output-ts.txt", "Hello, TypeScript!", "utf8", (err) => {
    if (err) {
      console.log("Error:", err);
    } else {
      console.log("File written successfully");
    }
  });

  // Reading from a file
  fs.readFile("output-ts.txt", "utf8", (err, data) => {
    if (err) {
      console.log("Error:", err);
    } else {
      console.log("File content:", data);
    }
  });

  // Deleting a file
  // fs.unlink("output-ts.txt", (err) => {
  //   if (err) {
  //     console.log("Error:", err);
  //   } else {
  //     console.log("File deleted successfully");
  //   }
  // });
}

main();
