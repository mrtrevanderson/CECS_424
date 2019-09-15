//Basic user input calculator for Java 
public class Calculator {
	public static void main(String[] args) {
		String[] input = args[0].split(" ");
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