/* test02_pkg/src/test02_c.c */

#include <dlfcn.h>
#include <stdio.h>

void dlsym_test(const char *name) {
    dlerror(); /* clear error */
    void *retp= dlsym(NULL, name);
    const char *errp= dlerror();
    fprintf(stderr, "dlsym(\"%s\") returned %p", name, retp);
    if (errp) {
        fprintf(stderr, " dlerror=\"%s\"", errp);
    }
    fputc('\n', stderr);
}
