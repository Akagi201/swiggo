/* File : simple.i */
%module simple
%{
#include "simple.h"
%}

%inline %{
extern int gcd(int x, int y);
extern double Foo;
%}

