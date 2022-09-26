package main

// #include <stdio.h>
// #include <errno.h>
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

	"gocv.io/x/gocv"
)

/*
int ttv_init(char* dictPath);
int ttv_detect_face(unsigned char* bgrData, int width, int height, int* faceBox, double* landmark);
int ttv_extract_feature(unsigned char* bgrData, int width, int height, int* faceBox, double* landmark, double* pbFeature, int* featureSize);
double ttv_compare_feature(double* feat1, double* feat2);
*/

func DetectFace(image []byte) error {
	imageInfo, err := gocv.IMDecode(image, gocv.IMReadAnyColor)
	if err != nil {
		return err
	}
	imageData, err := imageInfo.DataPtrUint8()
	if err != nil {
		return err
	}
	faceBox := make([]int, 4)
	landmarks := make([]float64, 136)
	imageDataPtr := unsafe.Pointer(&imageData[0])
	defer C.free(imageDataPtr)
	faceBoxPtr := unsafe.Pointer(&faceBox[0])
	defer C.free(faceBoxPtr)
	landmarksPtr := unsafe.Pointer(&landmarks[0])
	defer C.free(landmarksPtr)
	imageDataPtrC := (*C.uchar)(imageDataPtr)
	faceBoxPtrC := (*C.int)(imageDataPtr)
	landmarksPtrC := (*C.double)(imageDataPtr)
	C.ttv_detect_face(imageDataPtrC, C.int(imageInfo.Cols()), C.int(imageInfo.Rows()), faceBoxPtrC, landmarksPtrC)
	return nil
}

func Init(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	r := C.ttv_init(cPath)
	if r != 0 {
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
