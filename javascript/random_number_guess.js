const prompt = require('prompt-sync')({sigint: true});

const number = Math.floor(Math.random() * 101);

while (true) {
    const user_input = prompt("Enter number: ");

    if(+user_input > number) {
        console.log("Enter lower number");
    } else if(+user_input < number) {
        console.log("Enter bigger number");
    } else {
        console.log("Congratulation!...");
        break;
    }
}