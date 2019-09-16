var myargs = process.argv[2].split(" ");
const regex = "\\d*(\\.\\d|\\d\\.)\\d*";
if(myargs[0].match(regex) || myargs[2].match(regex)){
	var in1 = parseFloat(myargs[0]);
	var in2 = parseFloat(myargs[2]);
	var out1;
	switch(myargs[1]){
		case "+":
			out1 = in1 + in2; 
			console.log(out1);
			break;
		case "-":
			out1 = in1 - in2;
			console.log(out1);
			break;
		case "*":
			out1 = in1 * in2;
			console.log(out1);
			break;
		case "/":
			out1 = in1 / in2;
			console.log(out1); 
			break;
	}

}
else{
	var in1 = parseInt(myargs[0]);
	var in2 = parseInt(myargs[2]);
	var out;
	switch(myargs[1]){
		case "+":
			out1 = in1 + in2; 
			console.log(out1);
			break;
		case "-":
			out1 = in1 - in2;
			console.log(out1);
			break;
		case "*":
			out1 = in1 * in2;
			console.log(out1);
			break;
		case "/":
			out1 = in1 / in2;
			console.log(out1); 
			break;
	}
	
}