cd server
go build
cd ..

cd client
go build
cd ..

move client\client.exe .
move server\simpleserver.exe .
