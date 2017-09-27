source ./port.sh
cd ..
# appname='run.go'
mode='SCRIPT'
build='go run build.go'
server='go build main.go'
nowpath=`pwd`
export GOPATH=$nowpath
#go run $appname -r $build -r $server &
$build
$server
./main -host=$port -mode=$mode &
