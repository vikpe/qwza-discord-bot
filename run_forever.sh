RETRY_TIMEOUT=10

while true; do
  ./qwza
  echo "stopped, restarting in ${RETRY_TIMEOUT} seconds.."
  sleep ${RETRY_TIMEOUT}
done
