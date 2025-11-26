/**
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-26
 * @fileoverview Validate end > start and display range.
 */

let startVal: number = parseInt(prompt("Enter Start Value: ") || "0");
let endVal: number = parseInt(prompt("Enter End Value: ") || "0");

while (endVal <= startVal) {
  endVal = parseInt(
    prompt(
      "Sorry, ending value must be larger. Enter another ending value: ",
    ) || "0",
  );
}

for (let counter = startVal; counter <= endVal; counter++) {
  console.log(counter);
}
console.log("\nDone.");
