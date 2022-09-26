# wget https://github.com/opencv/opencv_contrib/archive/refs/tags/4.5.5.zip
# unzip ./4.5.5.zip
#brew install eigen
#brew install harfbuzz


export PKG_CONFIG_PATH="/opt/homebrew/Cellar/zlib/1.2.12_1/lib/pkgconfig"
 cmake -D CMAKE_BUILD_TYPE=Release -DBUILD_opencv_legacy=OFF -DOPENCV_GENERATE_PKGCONFIG=ON \
 -DAPPLE_FRAMEWORK=OFF -DOPENCV_EXTRA_MODULES_PATH=./opencv_contrib-4.5.5/modules  -S . -B build


 cmake \
 -DWITH_OPENJPEG=OFF \
 -DWITH_IPP=OFF \
 -D BUILD_ZLIB=OFF \
 -D CMAKE_BUILD_TYPE=RELEASE \
 -D CMAKE_INSTALL_PREFIX=/usr/local \
 -D OPENCV_EXTRA_MODULES_PATH=$(pwd)/../opencv_contrib-4.5.5/modules \
 -D PYTHON3_EXECUTABLE=$CONDA_PREFIX/bin/python3 \
 -D BUILD_opencv_python2=OFF \
 -D BUILD_opencv_python3=ON \
 -D INSTALL_PYTHON_EXAMPLES=ON \
 -D INSTALL_C_EXAMPLES=OFF \
 -D OPENCV_ENABLE_NONFREE=ON \
 -D BUILD_EXAMPLES=ON ..
