var array = process.argv[2].split(" ");

var in1 = parseInt(array[0], 10);
var in2 = parseInt(array[2], 10);

if (array[1] == '+')
	console.log(in1 + in2);
else if (array[1] == '-')
	console.log(in1 - in2);
else if (array[1] == '*')
	console.log(in1 * in2);
else if (array[1] == '/')
	console.log(in1 / in2);