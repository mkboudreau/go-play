
if [ ! -n $WERCKER_GO_BINDATA_INPUT_DIRECTORY ]; then
	error 'Please specify input directory'
	exit 1
fi

go get github.com/jteeuwen/go-bindata



go-bindata $WERCKER_GO_BINDATA_INPUT_DIRECTORY
