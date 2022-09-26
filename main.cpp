#include <iostream>
#include <stdio.h>
#include <opencv2/opencv.hpp>
#include <vector>
#include <string>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <dirent.h>
#include <unistd.h>

#include "ttvfaceengine.h"

int main(int argc, char* argv[])
{
    int ret = ttv_init("../dict");    
    if(ret != 0) {
        printf("init failed: %d\n", ret);
        return 0;
    }
    
    cv::Mat image1 = cv::imread("../test/1.jpg");

    int faceBox1[4] = { 0 };
    double landmarks1[136] = { 0 };
    ret = ttv_detect_face(image1.data, image1.cols, image1.rows, faceBox1, landmarks1);
    if(ret == 1) {
        double feature1[1024] = { 0 };
        int featureSize1 = 0;
        ttv_extract_feature(image1.data, image1.cols, image1.rows, faceBox1, landmarks1, feature1, &featureSize1);

        int faceBox2[4] = { 0 };
        double landmarks2[136] = { 0 };
        cv::Mat image2 = cv::imread("../test/2.jpg");
        ret = ttv_detect_face(image2.data, image2.cols, image2.rows, faceBox2, landmarks2);
        if(ret == 1) {
            double feature2[1024] = { 0 };
            int featureSize2 = 0;
            ttv_extract_feature(image2.data, image2.cols, image2.rows, faceBox2, landmarks2, feature2, &featureSize2);

            double score = ttv_compare_feature(feature1, feature2);
            printf("score: %f\n", score);
        }
    }
}
