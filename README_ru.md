[IN ENGLISH](https://github.com/dojyaaa-n/HTTP-Golang-Audio-Server/blob/main/README.md)

# HTTP Golang Audio Server

## Запуск

```
go run main.go
```

## Подключение

Просто используйте любой аудио проигрыватель с поддержкой воспроизведения аудиостримов в реальном времени. Например:

### 1. Используя пакет `mpv` (Терминал):
```
sudo apt install mpv
mpv http://localhost:8080/stream
```
### 2. Используя `VLC Media Player` (GUI):

1. Откройте VLC.
2. Перейдите к **Медиа** -> **Открыть URL** (или нажмите `Ctrl+N`).
3. В поле "URL сети" введите URL своей локальной сети `http://localhost:8080/stream` и начните воспроизведение.
