function reverse(str) {
    res = "";
    i = str.length;
    while (i != 0) {
        res += str[i - 1];
        i -= 1;
    }
    return res;
}

console.log(reverse("Greetings!"));
