go build decryptorservice.go
pushd testdata
../decryptorservice &
echo $!
jobs
bash curl_tls_base640oeap.sh | \
  grep "Lorem Ipsum is not simply"
RES=$?
if [ ${RES} -ne 0 ]; then
  echo "ERROR: Server did not response correcly"
fi
bash curl_tls_base640garbage.sh | \
  grep "decryption error"
RES=$?
if [ ${RES} -ne 0 ]; then
  echo "ERROR: Server did not response correcly"
fi

kill $!
exit ${RES}