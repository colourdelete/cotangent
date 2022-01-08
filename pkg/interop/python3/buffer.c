#include "../cotangent.h"
#include <string.h>

char* render_and_sanitize(char* src) {
	struct RenderAndSanitize_return r = RenderAndSanitize(src, strlen(src));
	return r.r0;
}
