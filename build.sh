cd website
echo BUILDING FRONTEND
npm run build 
echo MOVING BUILD FOLDER TO BACKEND FOLDER
rm -rf ../backend/build
mv ./build/ ../backend/
cd ../
cd backend
go build -o backend .
./backend
