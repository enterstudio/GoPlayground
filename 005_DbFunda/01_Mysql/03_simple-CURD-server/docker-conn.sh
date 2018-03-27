#!/bin/sh

usage() {
    echo " Usage:"
    echo "    docker-con.sh <Container Name> <user:password> <DB Name>"
    echo
    echo " Where: <Container Name> = MySQL Container"
    echo "        <user:password> = MySQL DB Username and Password default 'root:password'"
    echo "        <DB Name> = Name of the Database where we would be working"
}

if [ "$1" != "" -a "$2" != "" -a "$3" != "" ];then
container="$1"
cred="$2"
dbname="$3"
ip="$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $container)"
echo
echo "For Container '$container' the IP address is '$ip'"
echo
echo "Connection String:"
str="$cred@tcp($ip:3306)/$dbname"
echo $str
echo
export DB_CONNECTION=$str
echo "Attempting to Run the program"
echo
go run *.go
else
usage
fi