import sys
#takes the equation from the command line argument and splits it up by space
arg = sys.argv[1] 
string = arg.split(' ')
#converts string to integers
in1 = int(string[0])
in2 = int(string[2])
if string[1] == '+':
	print in1 , '+' , in2, '=', in1 + in2
elif string[1] == '-':
	print in1 , '-' , in2, '=', in1 - in2
elif string[1] == '*':
	print in1 , '*' , in2, '=', in1 * in2
elif string[1] == '/':
	print in1 , '/' , in2, '=', in1 / in2	






