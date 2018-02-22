echo ""
export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run main.go

echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run main.go

echo ""
export GOMAXPROCS=4
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run main.go

echo ""
export GOMAXPROCS=8
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run main.go

echo ""
export GOMAXPROCS=16
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run main.go
