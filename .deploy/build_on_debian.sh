CGO_ENABLED=1 \
go build ./ \
--ldflags '-extldflags "-lm -lstdc++ "'\
-v -x ./main.go


