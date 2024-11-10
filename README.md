[НА РУССКОМ](https://github.com/dojyaaa-n/HTTP-Golang-Audio-Server/blob/main/README_ru.md)
# HTTP Golang Audio Server

## Setup
```
go run main.go
```
## Connection
You can simply use any audio player capable of playing realtime audio streams over http connections. For example:

### 1. Using `mpv` (Terminal):
```
sudo apt install mpv
mpv http://localhost:8080/stream
```
### 2. Using `VLC Media Player` (GUI):
1. Open VLC.
2. Go to **Media** -> **Open Network Stream** (or press `Ctrl+N`).
3. In the "Network URL" field, enter your localhost URL and click play.
