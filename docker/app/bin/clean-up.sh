#!/usr/bin/env bash
echo 'Runing migrations...'
/cloud-native/migrate up > /dev/null 2>&1 &
echo 'Deleting mysql-client...'
apk del mysql-client
echo 'Start application...'
/cloud-native/app