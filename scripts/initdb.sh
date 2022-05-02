#!/usr/bin/env bash

provisionCluster() {
    while true; do
        /cockroach/cockroach sql \
        --insecure \
        --host=$DATABASE_HOST:$DATABASE_PORT
        --execute="
            CREATE USER IF NOT EXISTS '$DATABASE_USER' WITH PASSWORD '$DATABASE_PASSWORD';
            CREATE DATABASE IF NOT EXISTS '$DATABASE_NAME';
            GRANT ALL ON DATABASE '$DATABASE_NAME' TO '$DATABASE_USER';
            "
        &>/dev/null;
    local exitCode="$?";
    if [[ "$exitCode" == "0" ]]
        then break;
    fi
    sleep 5;
    done
    echo "Provisioning completed successfully";
}

provisionCluster;
