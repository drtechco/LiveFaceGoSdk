# cmake -S . -B build

# CC=x86_64-linux-musl-gcc \
# CXX=x86_64-linux-musl-g++ \
# GOARCH=amd64 \
# GOOS=linux \
# CGO_ENABLED=1 \
# LD_LIBRARY_PATH=/Users/ttttt/GolandProjects/FaceSdk/lib \


# GOARCH=amd64 GOOS=linux CGO_ENABLED=1\
# go build  -a -gccgoflags="-v -pthread -lpthread" \
# -ldflags='-linkmode=external -extldflags "-v -static -lm -lpthread -lstdc++ -lsupc++"' \
# -v -x ./main.go

#LDFLAGS="-L/usr/local/lib"
#CPPFLAGS="-I/usr/local/include/"

#sudo cp ./libttvfaceengine1.dylib  /usr/local/lib/

LDFLAGS="-L/usr/local/lib" \
CPPFLAGS="-I/usr/local/include/" \
CGO_ENABLED=1 \
go build  -a \
--ldflags '-extldflags "-lm -lstdc++ "'\
-v -x ./main.go