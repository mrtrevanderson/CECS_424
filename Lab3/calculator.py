import sys
import re
#takes the equation from the command line argument and splits it up by space
arg = sys.argv[1] 
string = arg.split(' ')
decimal1 = re.findall(r"\d*\.\d*|\d*\.\d*", string[0])
decimal2 = re.findall(r"\d*\.\d*|\d*\.\d*", string[2])

# If arg has no floats then use integer math
if (len(decimal1) == 0) and (len(decimal2) == 0): 
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
# If arg has floats then use decimal math
else:
	in1 = float(string[0])
	in2 = float(string[2])
	if string[1] == '+':
		print in1 , '+' , in2, '=', in1 + in2
	elif string[1] == '-':
		print in1 , '-' , in2, '=', in1 - in2
	elif string[1] == '*':
		print in1 , '*' , in2, '=', in1 * in2
	elif string[1] == '/':
		print in1 , '/' , in2, '=', in1 / in2	





