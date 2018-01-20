# Go utilities

Examples:

```go
dump.Byte(binary.BigEndian, uint8(50)) // data bytes && error (if bad type)
dump.ByteTo(writer, binary.BigEndian, uint8(50)) // error (if bad type) and writes data to writer (io.Writer)
```