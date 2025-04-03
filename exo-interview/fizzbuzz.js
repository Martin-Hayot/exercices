function fizzbuzz(int) {
    if (int <= 0) {
        return "Enter a number higher than 0";
    }
    let i = 1;
    while (i <= int) {
        if (i % 3 == 0 && i % 5 == 0) {
            console.log("fizzbuzz");
        } else if (i % 3 == 0) {
            console.log("fizz");
        } else if (i % 5 == 0) {
            console.log("buzz");
        } else {
            console.log(i);
        }
        i++;
    }
}

fizzbuzz(15);
