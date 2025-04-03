function palindrome(str) {
    reversed = str.split("").reverse().join("");

    return str === reversed;
}

// function palindrome(str) {
//     if (str.length == 0) {
//         return false;
//     }
//     let middle = str.length / 2;
//     let first = "";
//     let second = "";
//     let i = 0;
//     for (let v in str) {
//         if (i < middle) {
//             first += str[v];
//             i++;
//         } else if (i >= middle) {
//             second += str[v];
//             i++;
//         }
//     }

//     second = second.split("").reverse().join("");

//     console.log(first, second);
//     if (first == second) {
//         return true;
//     }

//     return false;
// }

console.log(palindrome("HAllAH"));
