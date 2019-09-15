//Basic user input calculator for Java 
public class Calculator {
	public static void main(String[] args) {
		String[] input = args[0].split(" ");
		String decimal = "\\d*(\\.\\d|\\d\\.)\\d*";
		if (input[0].matches(decimal) || input[2].matches(decimal)){
			double in1 = double.parseDouble(input[0]);
			double in2 = double.parseDouble(input[2]);
			double out;
			if (input[1].equals("+")) out = in1 + in2;
			else if(input[1].equals("-")) out = in1 - in2;
			else if(input[1].equals("*")) out = in1 * in2;
			else if(input[1].equals("/")) out = in1 / in2;
			else out = in1 / in2;
			System.out.println(args[0] + " = " + out);
		}
		else {
			int in1 = Integer.parseInt(input[0]);
			int in2 = Integer.parseInt(input[2]);
			int out;
			if (input[1].equals("+")) out = in1 + in2;
			else if(input[1].equals("-")) out = in1 - in2;
			else if(input[1].equals("*")) out = in1 * in2;
			else if(input[1].equals("/")) out = in1 / in2;
			else out = in1 / in2;
			System.out.println(args[0] + " = " + out);
	}
}

	}


