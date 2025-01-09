# init module
go mod init github.com/adamtang5/youtube_go_react_fullstack_app_20241210

# get gofiber
go get github.com/gofiber/fiber/v2

# install air
go install github.com/air-verse/air@latest
# self updates go version -> 1.23.4

# run air server with hot updates
air -c .air.toml
air

# install godotenv
go get github.com/joho/godotenv

# install mongodb driver
go get go.mongodb.org/mongo-driver/mongo

# fix missing package
go get github.com/mattn/go-isatty@v0.0.20

# in ./frontend folder
npm create vite@latest .
npm install

# install chakra ui
npm i @chakra-ui/react @emotion/react @emotion/styled framer-motion

# install react icons
npm i react-icons