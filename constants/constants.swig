/* File : constants.swig */
%module constants

%{
#include "constants.h"
%}

%include "constants.h"

/* The following directives also produce constants */

%constant int iconst = 37;
%constant double fconst = 3.14;
