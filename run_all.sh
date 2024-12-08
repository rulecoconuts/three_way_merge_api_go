cd backend
go mod tidy

cd frontend
npm i

go run ../backend & npm run dev && kill $!