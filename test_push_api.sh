#!/usr/bin/env bash
set -o x

export DB_HOST="https://push2new.databox.com"
if [ -n "$1" ]; then DB_HOST=$1; fi

export DB_TOKEN="adxg1kq5a4g04k0wk0s4wkssow8osw84"
if [ -n "$2" ]; then DB_TOKEN=$2; fi


echo "Pushing without Base Auth token"
curl -sk -XPOST \
    -H "Content-Type: application/json" \
    $DB_HOST | python -m json.tool

echo "Pushing without Content-Type application/json"
curl -sk -XPOST -u $DB_TOKEN: \
    $DB_HOST | python -m json.tool

 echo "Pushing without using POST"
curl -sk -u $DB_TOKEN: \
    $DB_HOST

echo "Last pushes"
curl -sk -XPOST -u $DB_TOKEN: \
    -H "Content-Type: application/json" \
    $DB_HOST/lastpushes | python -m json.tool

curl -sk -XPOST -u $DB_TOKEN: \
    -H "Content-Type: application/json" \
    $DB_HOST/lastpushes/1 | python -m json.tool

echo "Pushing with empty data"
curl -sk -XPOST -u $DB_TOKEN: \
    -H "Content-Type: application/json" \
    $DB_HOST | python -m json.tool
