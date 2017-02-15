
// Converts Celsius to Fahrenheit
function celsiusToFahrenheit(tempC) {
    return  tempC * 1.8 + 32;
}

// Converts Fahrenheit to Celsius
function fahrenheitToCelsius(f) {
    //f = c * 1.8 + 32
    //c = (f - 32) / 1.8
    return (f - 32.0) / 1.8;
}

// Rounds n up to the nearesrt multiple of x
function roundUpTo(n, x) {
    var amount = x * Math.ceil(n / x);
    if (amount == 0) {
        amount = x;
    }
    return amount;
}

// Rounds up to the nearest 10
function roundUpTo100(n) {
    return roundUpTo(n, 100);
}

