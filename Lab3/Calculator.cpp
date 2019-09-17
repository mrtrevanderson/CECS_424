#include <vector>
#include <iostream>
#include <string>
#include <sstream>
using namespace std;

const vector<string> explode(const string& s, const char& c)
{
	string buff{ "" };
	vector<string> v;

	for (auto n : s)
	{
		if (n != c) buff += n; else
			if (n == c && buff != "") { v.push_back(buff); buff = ""; }
	}
	if (buff != "") v.push_back(buff);

	return v;
}

int main() {
	string input;
	getline(cin, input);

	vector<string> v{ explode(input, ' ') };

	double in1 = stod(v[0].c_str());
	double in2 = stod(v[2].c_str());
	double out;


	if (v[1] == "+") out = in1 + in2;
	else if (v[1] == "-") out = in1 - in2;
	else if (v[1] == "*") out = in1 * in2;
	else out = in1 / in2;

	cout << input << " = " << out;

	getchar();
	getchar();
	return 0;
}