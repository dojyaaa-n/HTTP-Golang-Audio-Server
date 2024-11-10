[НА РУССКОМ](https://github.com/dojyaaa-n/HTTP-Golang-Audio-Server/blob/main/README_ru.md)
# Audio transfer via HTTP server in Golang

## Setup
```
go run main.go
```
## Connection
Enter the address of your local `http://localhost:8080/stream` server into your browser's search bar and wait for it to start playing, or use any audio player that supports real-time audio streaming. For example:

### 1. Using `mpv` (Terminal):
```
sudo apt install mpv
mpv http://localhost:8080/stream
```
### 2. Using `VLC Media Player` (GUI):
1. Open VLC.
2. Go to **Media** -> **Open Network Stream** (or press `Ctrl+N`).
3. In the "Network URL" field, enter `http://localhost:8080/stream` and click play.
