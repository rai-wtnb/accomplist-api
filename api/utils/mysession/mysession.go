package mysession

import (
    "crypto/rand"
    "encoding/base64"
    "io"
)

func NewSessionID() string {
    b := make([]byte, 64)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return base64.URLEncoding.EncodeToString(b)
}
