[IN ENGLISH]([https://github.com/dojyaaa-n/HTTP-Golang-Audio-Server/blob/main/README.md](https://github.com/dojyaaa-n/HTTP-Golang-Audio-Server))

# Передача аудио через HTTP сервер на Golang

## Запуск

```
go run main.go
```

## Подключение

Введите адресс локального сервера `http://localhost:8080/stream` в поле поисковика вашего браузера, ожидая начала проигрывания или используйте любой аудио проигрыватель с поддержкой воспроизведения аудиостримов в реальном времени. Например:

### 1. Используя пакет `mpv` (Терминал):
```
sudo apt install mpv
mpv http://localhost:8080/stream
```
### 2. Используя `VLC Media Player` (GUI):

1. Откройте VLC.
2. Перейдите к **Медиа** -> **Открыть сетевой поток** (или нажмите `Ctrl+N`).
3. В поле "URL сети" введите URL своей локальной сети `http://localhost:8080/stream` и начните воспроизведение.
