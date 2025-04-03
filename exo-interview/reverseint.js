function reverseInt(int) {
    reversed = int.toString().split("").reverse().join("");

    if (int < 0) {
        return parseInt(reversed) * -1;
    }
    return parseInt(reversed);
}

console.log(reverseInt(-123400));
