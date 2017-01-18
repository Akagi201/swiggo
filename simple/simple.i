/* File : simple.i */
%module simple

%inline %{
extern int    gcd(int x, int y);
extern double Foo;
%}
