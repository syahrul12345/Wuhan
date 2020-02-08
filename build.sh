cd website
echo BUILDING FRONTEND
npm run build 
echo MOVING BUILD FOLDER TO BACKEND FOLDER
mv ./build/ ../backend/build/
cd ../
cd backend
go build -o backend .
./backend
