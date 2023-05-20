#!/bin/bash
# build the binary and apply rudimentary testing to make sure the binary works
#
go build decryptorservice.go
pushd testdata
../decryptorservice &
echo $!
jobs
bash curl_tls_base640oeap.sh | \
  grep "Txt\":\"Contrary to popular belief, Lorem Ipsum is not simply"
RES=$?
if [ ${RES} -ne 0 ]; then
  echo "ERROR: Server did not response correcly"
  kill $!
  exit ${RES}
fi
bash curl_tls_base64garbage.sh | \
  grep "Errmsg\":\"decryption failed"
RES=$?
if [ ${RES} -ne 0 ]; then
  echo "ERROR: Server did not response correcly"
  kill $!
  exit ${RES}
fi

kill $!
exit ${RES}