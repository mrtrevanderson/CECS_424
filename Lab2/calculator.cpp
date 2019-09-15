//Basic user input calculator for C++ with switch statement
#include <iostream>
using namespace std;

int main(int argc, char *argv[]) 
{ 
    char *str[10];
    str[0] = strtok(argv[1], " "); // Splits spaces between words in str
    str[1] = strtok (NULL, " ");
    str[2] = strtok (NULL, " ");
    //convert string to int
    int in1 = atoi(str[0]);
	int in2 = atoi(str[2]);  
	const char *op = str[1];


	switch (*op) {   
		case '+': 
			std::cout << in1 << " + " << in2 << " = " << in1 + in2 << "\n"; 
			break;  
		case '-': 
            std::cout << in1 << " - " << in2 << " = " << in1 - in2 << "\n"; 
            break; 
       	case '*': 
            std::cout << in1 << " * " << in2 << " = " << in1 * in2 << "\n"; 
            break; 
        case '/': 
            std::cout << in1 << " / " << in2 << " = " << in1 / in2 << "\n"; 
            break; 
    	}	
    return 0; 
} 







