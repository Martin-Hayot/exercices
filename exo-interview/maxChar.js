function maxChar(str) {
    const chars = {};
    let max = 0;
    let maxChar = 0;
    for (let v of str) {
        if (!chars[v]) {
            chars[v] = 1;
        } else {
            chars[v]++;
        }
    }
    for (let v in chars) {
        if (chars[v] > max) {
            max = chars[v];
            maxChar = v;
        }
    }

    console.log(chars);
    return maxChar;
}

console.log(maxChar("12220aaaa"));
