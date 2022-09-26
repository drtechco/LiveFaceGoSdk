#pragma once

#ifdef __cplusplus
extern "C" {
#endif

int ttv_init(char* dictPath);
int ttv_detect_face(unsigned char* bgrData, int width, int height, int* faceBox, double* landmark);
int ttv_extract_feature(unsigned char* bgrData, int width, int height, int* faceBox, double* landmark, double* pbFeature, int* featureSize);
double ttv_compare_feature(double* feat1, double* feat2);

#ifdef __cplusplus
}
#endif