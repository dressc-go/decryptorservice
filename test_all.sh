find . -iname "*_test.go" | \
  while read tf ; do
    echo `dirname ${tf}`
  done |
  uniq |
  while read td ; do
    pushd ${td}
    go test -v
    if [ $? -ne 0 ] ; then
    	exit -1
    fi
    popd
  done
