#!/bin/bash

# Example use: all_users | jq
function all_users() {
    curl -X GET \
    https://localhost:8080/users
}

function delete_user() {
    UUID=$1
    
    curl -X DELETE \
    https://localhost:8080/user/$UUID
}

function update_user() {
    TEST_UUID=$1
    TEST_UNAME="test_user_name_udpated"
    TEST_EMAIL="test_user_udpated@example.com"
    TEST_PW="123456_udpated"
    
    curl -X PUT \
    https://localhost:8080/user \
    -H 'Content-Type: application/json' \
    -d '{
    "UUID": "'$TEST_UUID'",
	"name": "'$TEST_UNAME'",
	"email": "'$TEST_EMAIL'",
	"password": "'$TEST_PW'"
    }'
}

# Example use: create_user
# Example use: create_user && all_users | jq
# Example use: create_user myusername myemail mypassword && all_users | jq
function create_user() {
    TEST_UNAME=${1:="test_user_name"}
    TEST_EMAIL=${2:="test_user@example.com"}
    TEST_PW=${3:="123456"}
    
    curl -X POST \
    https://localhost:8080/user \
    -H 'Content-Type: application/json' \
    -d '{
	"name": "'$TEST_UNAME'",
	"email": "'$TEST_EMAIL'",
	"password": "'$TEST_PW'"
    }'
}
