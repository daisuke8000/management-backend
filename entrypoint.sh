# !/bin/sh
echo "waiting for mysql server"
while ! nc -z db 3306; do
  sleep 1
done
echo "mysql Connection Successfully"

exec "$@"
exec main