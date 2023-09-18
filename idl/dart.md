# Dart generation

```
protoc --dart_out=grpc:gen/proto/dart \
    -Iapi api/app/v1/app.proto
```