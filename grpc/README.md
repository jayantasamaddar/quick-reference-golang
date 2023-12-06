# Table of Contents

- [Table of Contents](#table-of-contents)
- [Setup](#setup)
- [Protocol Buffers](#protocol-buffers)
  - [Protobuf: Introduction](#protobuf-introduction)
  - [Protobuf: JSON vs Proto Definition](#protobuf-json-vs-proto-definition)
  - [Protobuf: Usage](#protobuf-usage)
    - [Usage: When to use Protocol Buffers?](#usage-when-to-use-protocol-buffers)
    - [Usage: When not to use Protocol Buffers?](#usage-when-not-to-use-protocol-buffers)
  - [Protobuf: Style Guide](#protobuf-style-guide)
  - [Protobuf: Protocol Buffers Definition](#protobuf-protocol-buffers-definition)
- [References](#references)

---

# Setup

1. [Install the Go Programming Language](../golang/README.md#installation)
2. Install VS-Code and basic VS Code Extensions
3. Install VS Code Extensions for working with Protocol Buffers
   - `vscode-proto3`
   - `Proto Lint`
4. [Install Protolint](https://github.com/yoheimuta/protolint)
5. [Install Protocol Buffers Compiler](https://grpc.io/docs/protoc-installation/)
6. [Install Go plugins for the protocol compiler](https://grpc.io/docs/languages/go/quickstart/)

---

# Protocol Buffers

## Protobuf: Introduction

Protocol Buffers or Protobuf for short are language-neutral, platform-neutral extensible mechanisms for serializing structured data.

**Elements of Protocol Buffers**:

- Definition language (`.proto` files)
- Proto Compiler (`protoc`): Reads proto files and generates code for specific programming language.
- Runtime libraries: Programming language specific runtime libraries.
- Serialized Data: Data to be sent across networks.

---

## Protobuf: JSON vs Proto Definition

Let's first take how a Proto Definition looks against the JSON Schema.

**JSON Schema**:

```json
{
  "id": 34512,
  "first_name": "Jayanta",
  "last_name": "Samaddar",
  "active": true
}
```

**Proto Definition**:

```proto
syntax = "proto3";

message Character {
    uint32 id = 1;
    string first_name = 2;
    string last_name = 3;
    boolean active = 4;
}
```

**Pros of using Protocol Buffers over JSON**

<!--prettier-ignore-->
| JSON Schema                                                     | Proto Definition                                                                        |
| --------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| JSON is serialized as text-based.                               | Data in Protobuf is serialized as binary making it smaller than JSON/compressed JSON.   |
| Data type is not strict and not enforced.                       | Data type is strict. Protobuf schema is enforced.                                       |
| Need to use JSON representation in the form of a Java class or Go struct. | Proto compiler generates efficient code in supported languages (Java, Go, Python etc. ) |
| JSON doesn't support schema evolution.                          | Protobuf supports schema evolution (both forwards and backwards compatibility).         |
| Comments are not supported.                                     | Supports comments.                                                                      |

**Pros of JSON over Protocol Buffers**

| JSON Schema                                                     | Proto Definition                               |
| --------------------------------------------------------------- | ---------------------------------------------- |
| Easy to read (human readable text serialization).               | Need to convert binary data to human readable. |
| Wide programming language support (e.g. JavaScript front-ends). | Limited programming language support.          |
| More community support.                                         | Less community support.                        |

---

## Protobuf: Usage

### Usage: When to use Protocol Buffers?

- **Service to Service communication**: Protocol buffers are a great fit for service to service communication, due to its smaller size which means efficient network usage.
  - **Same schema definition across multiple programming languages**: Same data specification for Service A written in Go, Service B written in Java etc.
- **Used with gRPC**

### Usage: When not to use Protocol Buffers?

- **Programming languages that do not support (or limited support) protobuf**: E.g. JavaScript (Might improve over time)
- **When sharing API with external party**:
  - Less adoption for protobuf / gRPC
  - Need to share protobuf schema file to be able to work with protobuf, and this means you need to manage and share the proto files with party outside your organization, which means this can have non-technical issues like legal or company confidentiality rules.
- **Data that exceed few megabytes**:
  - Although protobuf binary serialization is smaller compared to JSON, Protobuf is not designed to handle data larger than a few megabytes.
  - This means a text-based data will be a good fit for Protobuf but media files like images, or video should not use Protobuf.

---

## Protobuf: Style Guide

**Standard Formatting conventions**:

- File name should be of Lower Snake Case with `.proto` extension. E.g. `lower_snake_case.proto`
- Each line should be at most 80 characters long.
- Indentation: 2 spaces
- Double quotes for string (e.g. "This is a string")
- Tips:
  - Use Protolint extension for VS Code

**Order of the file content, starting from top**:

1. The license, if applicable. Put this license on comment (can be multi-line comment)
2. Brief description, overview of the file
3. Proto syntax
4. Package, if any. Package name in `lowercase` with no whitespace.
5. Imports, if any (Sort imports alphabetically)
6. Proto file options
7. Everything else, basically the protocol buffer definitions. Common ordering is:
   - enums: Name in `PascalCase`, Value in `UPPER_SNAKE_CASE`. Zero value enum should have `UNSPECIFIED` as suffix.
   - messages: Name in `PascalCase`, Filed Name in `lower_snake_case`
   - services: Name in `PascalCase`, RPC method name is `PascalCase`

```proto
/* Limited License, only for internal mycompany.com usage */

/* Just a simple proto file to demonstrate style guide */

syntax = proto3;

package course.styleguide;

import "proto/styleguide/alpha_file.proto";
import "proto/styleguide/beta_file.proto";
import "proto/styleguide/gamma_file.proto";

option optimize_for = SPEED;
option go_package = "/protogen/styleguide;styleguide";
option java_package = "com.mycompany.proto.generated";

enum Distance {
    DISTANCE_UNSPECIFIED = 0;
    DISTANCE_NEAR = 1;
    DISTANCE_FAR = 2;
}

message MyProtoMessage {
    uint32 my_proto_id = 1 [json_name = "MyProtoId"];
    AlphaMessage alpha = 2;
    BetaMessage beta = 3;
    repeated string tags = 4;
}

service MyProtoService {
    rpc: DoSomething(MyProtoMessage) returns (GammaMessage);
}
```

> **Note**: This above is the Google Style Guide and it's simple. A more detailed Style Guide is the [Uber Protobuf Style Guide](https://github.com/uber/prototool/blob/dev/style/README.md).

---

## Protobuf: Protocol Buffers Definition

**Example**:

```proto
/* In `hello.proto` */
syntax = "proto3"

option go_package = "my-protobuf/protogen/basic";

message Hello {
    string name = 1;
}
```

**Compile**:

```bash
# Generate to module defined by go_package
protoc --java-out=. --go-out=. .proto/basic/*.proto

# Generate to a specified go_opt module name
protoc --go_opt module=hello-protobuf --go_out=. ./proto/basic/*.proto
```

The generated source code is a data structure in a chosen language. This means a Java, Kotlin, C# or Python class or a Go struct. This represents the data to be transferred over a network. The generated code is optimized and ready to be used with Protobuf.

---

# References

- [Uber Protobuf Style Guide](https://github.com/uber/prototool/blob/dev/style/README.md)
- [Protolint](https://github.com/yoheimuta/protolint)
- [Buf](https://github.com/bufbuild/buf)
