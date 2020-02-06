#!/bin/bash
DB_DIR=`pwd`/db/data
DB_IMAGE=user-sqlite-db
DB_NAME=users.db

function build_docker_db() {
    docker build -t $DB_IMAGE ./db/.
    docker images | grep $DB_IMAGE
}

function start_docker_db() {
    docker run --rm -it -v $DB_DIR/:/db $DB_IMAGE $DB_NAME
}

function dump_docker_db() {
    # Command to backup $DB_NAME to a dump.sql file on host
    docker run --rm -it -v $DB_DIR/:/db $DB_IMAGE $DB_NAME .dump >> $DB_DIR/dump.sql
}

function restore_docker_db() {
    # Before restoring the database make sure that the destination database is empty (moving current database to .old)
    mv $DB_DIR/$DB_NAME $DB_DIR/$DB_NAME.old
    cat $DB_DIR/dump.sql | docker run --rm -i -v $DB_DIR/:/db $DB_IMAGE $DB_NAME
}