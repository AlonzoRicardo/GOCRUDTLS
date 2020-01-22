#!/bin/bash

function all_users() {
    curl -X GET \
    https://localhost:8080/users
}

function delete_user() {
    curl -X DELETE \
    https://localhost:8080/user/userid
}

function update_user() {
    curl -X PUT \
    https://localhost:8080/user/userid
}

function create_user() {
    curl -X POST \
    https://localhost:8080/user/userid
}
