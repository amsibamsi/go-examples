#define MAX_PLATFORMS 16
#define MAX_DEVICES 32

#include <CL/opencl.h>

char* get_device(int platform_flag, int device_flag, cl_device_id *device);
char* create_context(cl_device_id *device, cl_context *context);
char* create_program(cl_context context, cl_device_id *device, cl_program *program);
