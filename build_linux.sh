#/bin/bash
TG=decryptorservice_x86_64
mkdir -p build/${TG}
go build -o build/${TG}/decryptorservice decryptorservice.go 
cp README.md build/${TG}
cp LICENSE build/${TG}
pushd build
tar -czf ${TG}.tar.gz ${TG}
sha256sum ${TG}.tar.gz > ${TG}.sha256
popd
