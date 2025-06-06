#include <string.h>

#include "./entry.h"

static char modname_buf[128];
static napi_module nativeModule = {
    .nm_version = 1,
    .nm_flags = 0,
    .nm_filename = nullptr,
    .nm_register_func = InitializeModule,
    .nm_modname = modname_buf,
    .nm_priv = nullptr,
    .reserved = {0},
};

void register_module(const char *modname) {
  strcpy(nativeModule.nm_modname, modname);
  napi_module_register(&nativeModule);
}
