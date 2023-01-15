/* mypow.c */

#include <math.h>

double mypow (double base, long p, long q)
{
    double retval;

    if (q==0) return NAN;
    if (p==0) return 1.0;
    if (p==q) return base;

    if (q<0) {
       p= -p;
       q= -q;
    }

    while (p%2==0 && q%2==0) {
       p/=2;
       q/=2;
    }

    if (base<0 && p%2==0) {
        base= -base;
    }

    if (base<0) {
        if (q%2==0) return NAN;
        retval= - pow (-base, (double)p/(double)q);

    } else if (base==0.0) {
        if (p<0) return NAN;
        retval= 0.0;

    } else {
        retval= pow (base, (double)p/(double)q);
    }

    return retval;
}
