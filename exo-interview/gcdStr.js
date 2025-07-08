/**
 * @param {string} str1
 * @param {string} str2
 * @return {string}
 */
var gcdOfStrings = function (str1, str2) {
    // If concatenating in different orders doesn't give the same result,
    // there's no common divisor string
    if (str1 + str2 !== str2 + str1) {
        return "";
    }

    // Helper function to find GCD of two numbers
    function gcd(a, b) {
        return b === 0 ? a : gcd(b, a % b);
    }

    // The length of the GCD string is the GCD of the lengths
    const gcdLength = gcd(str1.length, str2.length);

    // Return the substring from 0 to gcdLength
    return str1.substring(0, gcdLength);
};

console.log(gcdOfStrings("ABABAB", "ABAB"));
