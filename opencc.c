#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "opencc/opencc.h"

const char *Convert(const char *input, const char *config) {
	if(strlen(config) > 16) {
		return 0;
	}

	char configFile[256] = "/usr/share/opencc/";
	strcat(configFile, config);
	strcat(configFile, ".json");

	opencc_t p = opencc_new(configFile);
	char *out = opencc_convert(p, input);
	out[strlen(input)] = '\0';

	opencc_delete(p);

	return out;
}

void Convert_free_string(char *p) {
	opencc_free_string(p);
}

void* Opencc_New(const char *configFile) {
	return opencc_new(configFile);
}

void Opencc_Delete(void *id) {
	opencc_delete(id);
}

const char *Opencc_Convert(void *id, const char *input) {
	char *output = opencc_convert(id, input);
	output[strlen(input)] = '\0';
	return output;
}

void Opencc_Free_String(char *p) {
	opencc_free_string(p);
}

