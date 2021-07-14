BUILD_PATH?=build
MAIN_OUTPUT_FILE=test
MAIN_PROGRAM=${BUILD_PATH}/${MAIN_OUTPUT_FILE}


# make run 
run:  plug_1.so
	go build  -o=${MAIN_PROGRAM} main.go
	[ -f "${MAIN_PROGRAM}" ] && cd ${BUILD_PATH} && sh -c ./${MAIN_OUTPUT_FILE}; cd -

plug_1.so: clean
	go build -buildmode=plugin -o=${BUILD_PATH}/plug_1.so plug_1.go

pre_build: 
	([ ! -d "${BUILD_PATH}" ] && mkdir ${BUILD_PATH}) || echo "ok."
	

clean: pre_build
	[ -d "${BUILD_PATH}" ] && rm -rf ${BUILD_PATH}

.Phony: run clean pre_build


