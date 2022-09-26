package main

// #include <stdio.h>
// #include <errno.h>
// #include <stdlib.h>
// #include <iostream>
// #include <opencv2/opencv.hpp>
// #include <vector>
// #include <string>
// #include <stdlib.h>
// #include <time.h>
// #include <sys/time.h>
// #include <dirent.h>
// #include <unistd.h>
// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L./lib -lttvfaceengine1
// #include "./include/ttvfaceengine.h"
import "C"
import (
	"errors"
	"fmt"
	"os"
	"unsafe"

	"github.com/go-opencv/go-opencv/opencv"
)

/*
int ttv_init(char* dictPath);
int ttv_detect_face(unsigned char* bgrData, int width, int height, int* faceBox, double* landmark);
int ttv_extract_feature(unsigned char* bgrData, int width, int height, int* faceBox, double* landmark, double* pbFeature, int* featureSize);
double ttv_compare_feature(double* feat1, double* feat2);
*/

func DetectFace(image byte[]) error { 
	imageInfo:=	opencv.DecodeImageM(unsafe.Pointer(&image[0]),1)
	imageData:= imageInfo.GetData()
	int faceBox1[4] = { 0 };
    double landmarks1[136] = { 0 };
	return nil
}

func Init(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	r := C.ttv_init(cPath)
	if r != 1 {
		return errors.New("ttv init err")
	}
	return nil
}

func main() {
	dictPath := "/opt/ttv_dict"
	if len(os.Args) > 1 {
		dictPath = os.Args[1]
	}
	fmt.Printf("dict path %s\n", dictPath)
	Init(dictPath)

	//fmt.Println("Hello c, welcome to go!")
}
