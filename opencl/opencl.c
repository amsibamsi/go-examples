#include <stdio.h>
#include <CL/cl.h>
#include "opencl.h"

char* get_error(cl_int error) {
  switch(error) {
    case CL_INVALID_VALUE: return "CL_INVALID_VALUE";
    case CL_OUT_OF_HOST_MEMORY: return "CL_OUT_OF_HOST_MEMORY";
    case CL_INVALID_PLATFORM: return "CL_INVALID_PLATFORM";
    case CL_INVALID_DEVICE_TYPE: return "CL_INVALID_DEVICE_TYPE";
    case CL_DEVICE_NOT_FOUND: return "CL_DEVICE_NOT_FOUND";
    case CL_OUT_OF_RESOURCES: return "CL_OUT_OF_RESOURCES";
    case CL_INVALID_PROPERTY: return "CL_INVALID_PROPERTY";
    case CL_DEVICE_NOT_AVAILABLE: return "CL_DEVICES_NOT_AVAILABLE";
    default: return "Unknown";
  }
}

char* get_device(int platform_flag, int device_flag, cl_device_id *device) {
  cl_uint err;
  cl_uint num_platforms;
  err = clGetPlatformIDs(MAX_PLATFORMS, NULL, &num_platforms);
  if (err != CL_SUCCESS) {
    return get_error(err);
  }
  printf("Found %d platforms of max %d\n", num_platforms, MAX_PLATFORMS);
  cl_platform_id platforms[num_platforms];
  err = clGetPlatformIDs(num_platforms, &platforms[0], NULL);
  if (err != CL_SUCCESS) {
    return get_error(err);
  }
  int i;
  for (i=0; i<num_platforms; i++) {
    size_t name_size;
    err = clGetPlatformInfo(platforms[i], CL_PLATFORM_NAME, 0, NULL, &name_size);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    char name[name_size];
    err = clGetPlatformInfo(platforms[i], CL_PLATFORM_NAME, name_size, name, NULL);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    size_t vendor_size;
    err = clGetPlatformInfo(platforms[i], CL_PLATFORM_VENDOR, 0, NULL, &vendor_size);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    char vendor[vendor_size];
    err = clGetPlatformInfo(platforms[i], CL_PLATFORM_VENDOR, vendor_size, vendor, NULL);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    size_t version_size;
    err = clGetPlatformInfo(platforms[i], CL_PLATFORM_VERSION, 0, NULL, &version_size);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    char version[version_size];
    err = clGetPlatformInfo(platforms[i], CL_PLATFORM_VERSION, version_size, version, NULL);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    printf("Platform #%d: %s, %s, %s\n", i, name, vendor, version);
  }
  if (platform_flag < 0 || platform_flag >= num_platforms) {
    return "Invalid platform number";
  }
  cl_platform_id platform_id = platforms[platform_flag];
  printf("Using platform #%d\n", platform_flag);
  cl_uint num_devices;
  err = clGetDeviceIDs(platform_id, CL_DEVICE_TYPE_ALL, MAX_DEVICES, NULL, &num_devices);
  if (err != CL_SUCCESS) {
    return get_error(err);
  }
  printf("Found %d devices of max %d\n", num_devices, MAX_DEVICES);
  cl_device_id devices[num_devices];
  err = clGetDeviceIDs(platform_id, CL_DEVICE_TYPE_ALL, num_devices, &devices[0], NULL);
  if (err != CL_SUCCESS) {
    return get_error(err);
  }
  for (i=0; i<num_devices; i++) {
    size_t name_size;
    err = clGetDeviceInfo(devices[i], CL_DEVICE_NAME, 0, NULL, &name_size);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    char name[name_size];
    err = clGetDeviceInfo(devices[i], CL_DEVICE_NAME, name_size, &name, NULL);
    if (err != CL_SUCCESS) {
      return get_error(err);
    }
    printf("Device #%d: %s\n", i, name);
  }
  if (device_flag < 0 || device_flag >= num_devices) {
    return "Invalid device number";
  }
  printf("Using device #%d\n", device_flag);
  *device = devices[device_flag];
  return NULL;
}

char* create_context(cl_device_id *device, cl_context *context) {
  cl_int err;
  *context = clCreateContext(NULL, 1, device, NULL, NULL, &err);
  if (err != CL_SUCCESS) {
    return get_error(err);
  }
  return NULL;
}
