module go-chat

go 1.22.4

replace shared => ../shared

require (
	github.com/gorilla/mux v1.8.1
	github.com/rs/zerolog v1.33.0
	shared v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
