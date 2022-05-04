#!/bin/sh

provisionCluster() {
while true; do
    sql --insecure --host=$DATABASE_HOST:$DATABASE_PORT \
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

nodeStatus() {
    node status --stats --insecure --host=$DATABASE_HOST:$DATABASE_PORT &>/dev/null;
    local exitCode="$?";
    if [[ "$exitCode" == "0" ]]
        then provisionCluster;
    fi
    sleep 5;
    done

    echo "Provisioning completed successfully";

}

nodeStatus;
