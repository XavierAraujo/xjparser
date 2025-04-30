# 🧩 xjparser — A Simple JSON Parser in Go

This is a lightweight JSON parser implemented in Go, inspired by [John Cricket's coding challenges](https://codingchallenges.fyi/challenges/challenge-json-parser). It parses JSON strings into a structured interface with support for all major JSON types including objects, arrays, strings, numbers, booleans, and null.

---

## ✨ Features

- Parse JSON objects and arrays using a hand-crafted interface
- Represent JSON values as strongly typed Go structs
- Efficient `null` handling using a singleton `JsonNull` instance
- Easy value inspection with helper functions like `IsObject`, `IsInt`, `IsNull`, etc.
- ✅ **ANTLR4 grammar** included to tokenize and parse raw JSON input into the internal structure

---

## 📦 JSON Types

| Type         | Enum Value     | Go Type         |
|--------------|----------------|-----------------|
| Object       | `JSON_OBJECT`  | `JsonObject`    |
| Array        | `JSON_ARRAY`   | `JsonArray`     |
| Integer      | `JSON_INT`     | `JsonInt`       |
| Float        | `JSON_FLOAT`   | `JsonFloat`     |
| String       | `JSON_STR`     | `JsonStr`       |
| Boolean      | `JSON_BOOL`    | `JsonBool`      |
| Null         | `JSON_NULL`    | `JsonNull`      |

---

## 🔤 ANTLR4 Grammar

This project includes a custom ANTLR4 grammar for JSON parsing, making it easy to convert raw JSON strings into Go-native structures using your `JsonValue` interface.
To generate the Go parser from the grammar just run the following script (you'll need to have antlr4 installed in your machine):

```bash
generate-grammar.sh
```

## 🧪 How to test

Just run 

```bash
go test ./...
```

It recursively searches all project sub-folders and runs the correspondent tests

## 📝 License
MIT — feel free to use, modify, and share!
