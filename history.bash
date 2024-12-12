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
