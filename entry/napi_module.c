#include <string.h>

#include "./entry.h"

static char modname_buf[128];
static napi_module nativeModule = {
    .nm_version = 1,
    .nm_flags = 0,
    .nm_filename = NULL,
    .nm_register_func = InitializeModule,
    .nm_modname = modname_buf,
    .nm_priv = NULL,
    .reserved = {0},
};

void register_module(const char *modname) {
  strcpy(modname_buf, modname);
  napi_module_register(&nativeModule);
}
