package yiigo

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"strings"
)

// MD5 calculate the md5 hash of a string.
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 calculate the sha1 hash of a string.
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// Hash Generate a hash value, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func Hash(t string, s string) string {
	var h hash.Hash

	switch strings.ToUpper(t) {
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	case "SHA224":
		h = sha256.New224()
	case "SHA256":
		h = sha256.New()
	case "SHA384":
		h = sha512.New384()
	case "SHA512":
		h = sha512.New()
	default:
		return s
	}

	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// AddSlashes returns a string with backslashes added before characters that need to be escaped.
func AddSlashes(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		if ch == '\'' || ch == '"' || ch == '\\' {
			buf.WriteRune('\\')
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}

// StripSlashes returns a string with backslashes stripped off. (\' becomes ' and so on.) Double backslashes (\\) are made into a single backslash (\).
func StripSlashes(s string) string {
	var buf bytes.Buffer

	l, skip := len(s), false

	for i, ch := range s {
		if skip {
			buf.WriteRune(ch)
			skip = false

			continue
		}

		if ch == '\\' {
			if i+1 < l && s[i+1] == '\\' {
				skip = true
			}

			continue
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}

// QuoteMeta returns a version of str with a backslash character (\) before every character that is among these: . \ + * ? [ ^ ] ( $ )
func QuoteMeta(s string) string {
	var buf bytes.Buffer

	for _, ch := range s {
		switch ch {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}

		buf.WriteRune(ch)
	}

	return buf.String()
}
