# Protocol
A Minecraft protocol packet encoder and decoder written for Anchor.

## Installation
```
go get github.com/AnchorMC/Protocol
```

## Documentation
https://pkg.go.dev/github.com/AnchorMC/Protocol

## Examples

### Marshal

```go
import (
    "fmt"
    
    "github.com/anchormc/protocol"
)

func main() {
    // Handshake Packet (0x00)
    // https://wiki.vg/Protocol#Handshake
    packetData, err := protocol.Marshal(
        protocol.VarInt(0x00),         // Packet ID - varint
        protocol.VarInt(757),          // Protocol Version - varint
        protocol.String("localhost"),  // Server Host - varint
        protocol.UnsignedShort(25565), // Server Port - unsigned short
        protocol.VarInt(1),            // Next State - varint
    )

    if err != nil {
        panic(err)
    }

    fmt.Println(packetData)
}
```

### Unmarshal

```go
import (
    "fmt"
    
    "github.com/anchormc/protocol"
)

func main() {
    var packetType protocol.VarInt
    var protocolVersion protocol.VarInt
    var serverHost protocol.String
    var serverPort protocol.UnsignedShort
    var nextState protocol.VarInt
    
    err := protocol.Unmarshal(
        r,           // generic io.Reader
        &packetType, // must use pointers to DataTypeReader
        &protocolVersion,
        &serverHost,
        &serverPort,
        &nextState,
    )

    if err != nil {
        panic(err)
    }

    fmt.Println(packetType, ...)
}
```

## License
[MIT License](https://github.com/AnchorMC/Protocol/blob/main/LICENSE)