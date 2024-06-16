# İlk aşamada, Go uygulamasını derlemek için golang resmi imajını kullanın
FROM golang:1.21 as builder

# Uygulama kodlarını çalışma dizinine kopyalayın
WORKDIR /app
COPY . .

# Bağımlılıkları yükleyin ve uygulamayı derleyin
RUN go mod tidy
RUN go build -o /app/main ./cmd/main.go

# İkinci aşamada, Alpine Linux imajını kullanın
FROM alpine:latest

# Gerekli sertifikaları yükleyin
RUN apk add --no-cache ca-certificates

# Çalışma dizinini belirleyin
WORKDIR /app

# Derlenmiş uygulamayı önceki aşamadan kopyalayın
COPY --from=builder /app/main .

# Uygulamayı çalıştırmak için gerekli izinleri verin
RUN chmod +x ./main

# Port 8080'i dış dünyaya açın
EXPOSE 8080

# Uygulamayı başlatmak için CMD komutunu belirleyin
CMD ["./main"]
