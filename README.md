# session-management

I didn't know it was possible to reduce a prod ready Golang image from 45 MB to 4 MB 🤯

1. use -ldflags="-s -w" (45 -> 38)
2. use scratch instead of distroless (38 -> 17)
3. use upx to compress image (17 -> 4)

