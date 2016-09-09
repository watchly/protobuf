# protoc-gen-govalidate

This fork of `github.com/golang/protobuf` adds support for `govalidate` based validation to protobuf output files. It does this by using the `extend` functionality of protobuf to let you define a new field that is then plucked out and a `Validate()` function written. For example:



**This is ALPHA software - feedback & fixes much appreciated**



**user.proto**

```protobuf
syntax = "proto3";

package user;

// This block is *required* - it tells protoc that we want to extend it. 
// The descriptor.proto should be available as part of your protobuf3 
// installation. In the worst case, you can copy it from the github.com/google/protobuf
// project and import it locally - but I recommend having a proper 
// installation of protobuf-dev (see Installation instructions below)
import "google/protobuf/descriptor.proto";
extend google.protobuf.FieldOptions {
  string valid = 71111;
}

message User {
  string name    = 1 [(valid)="trim,truncate(256),length(3|256),alpha"];
  int32  age     = 2 [(valid)="min(0),max(120)"];
  string email   = 3 [(valid)="trim,tolower,email"];
  string website = 4 [(valid)="trim,url"];

  double rating  = 5 [(valid)="min(0.0),max(1.0)"];
  float  average = 6 [(valid)="min(0.0),max(1.0)"];
  float  stars   = 7 [(valid)="floor,min(0.0),max(5.0)"];

  // Will remove empty strings from the array, and then will
  // run what's left through trim,truncate,... and replace
  // those that change in the array
  repeated string tags = 8 [(valid)="omitempty,trim,truncate(50),alpha"];

  repeated Address addresses = 10;  // Will call .Validate() on each Address in addresses
  map<string, Phone> numbers = 11;  // Will call .Validate() on each Phone value automatically

  // types in maps are also supported, santization functions like trim & truncate
  // will update the map's key or value if there were changes
  // omitempty will rid the map of any keys who's string values are empty
  // the tags are declared like a normal go map:
  // [(valid)="[KEY_TAGS]VALUE_TAGS"]
  map<string, string> properties = 12 [(valid)="[trim,truncate(50)]omitempty,trim,truncate(255)"];
}

message Phone {
  string number = 1 [(valid)="trim,truncate(50),length(5|50)"];
}

message Address {
  string address = 1 [(valid)="trim,truncate(50),length(0|512)"];
}
```



Will be parsed and the generated code will allow you to do:



```go
{
  user := api.User{}
  if err := proto.Unmarshal(body, &user); err != nil {
    return err
  }

  if changed, err := user.Validate(); err != nil {
    // The User struct or one of it's nested structs did not validate
    // as per the rules in the .proto file
    return err
  } else if changed {
    // one or more of the values have been sanitized as per the ruls in the
    // proto file i.e. strings have been trimmed, or ints have been min/max'd
    fmt.Println("input had to be sanitized`")
  }
}
```



For a look at the generated code for the example `proto` file above, scroll further down.



## Features

- **string validation**
  - All tag-based validators (including param validators) are supported from https://github.com/asaskevich/govalidator
- **string santization** is supported with the following in-built sanitizers:
  - `trim`: trims leading and trailing whitespace
  - `truncate(n)`: truncates the *bytes* of a string to `n` length
  - `truncaterunes(n)`: truncates the *runes* of a string to `n` length
  - `tolower` and `toupper`: exactly as you'd expect
  - Santizers can be **chained** and should be placed before the validators:
    - `trim,truncate(255),tolower,email`
  - Negated validators are supported: `trim,host,!url` would confirm the input is a host but not a url
- **int32, int64, uint32, uint64, sint32, sint64 santization**
  - `min(n)`: makes sure the number is at least `n`
  - `max(n)`; makes sure the number is at most `n`
- **float32 (proto: float), float64 (proto: double) santization**
  - `ceil`: ceils the value
  - `floor`: floors the value
  - `min(n)`: ensures the value is at least `n`
  - `max(n)`: ensures the value is at most `n`
  - **note**: float32's are casted to float64, this might not be correct or what you intend - PRs accepted :)
- **slices validation & sanitization**
  - `ignore` tag will ignore the slice entirely
  - **scalar slices** will be validated & sanitized, so integers/floats/string slices can have the same rules as a normal field, and every element will be santized & validated. The slice will be updated with sanitzed values where necessary
  - **string slices**
    - `omitempty`: empty (val == "") elements are removed
  - **struct slices**
    - Calling `.Validate` on a struct that has a field of struct-slices will call `.Validate()` on any non-nil elements of the slice
- **maps validation & sanitization**
  - `ignore` tag will ignore the map entirely
  - **scalar keys AND/OR values** will be sanitized & validated - so you can sanitize a `map[string]float` for instance so the keys are trimmed & truncated, and the values are all between 0.0 and 1.0
  - **string values**
    - `omitempty` is supported to remove any keys from the map that have an empty string value (value == "")
  - **struct values**
    - All non-nil struct values will have `.Validate()` called on them when calling on the parent struct
- **Deep validation of nested structs, maps & arrays**: calling `.Validate()` on the main type will validate all instances of a `message` nested inside of it, such as direct nesting, a map, or a slice
- **No reflection is used**, as you can from the example output below, the validation functions are called directly



## Installation

* `go get -u github.com/watchly/protobuf/{protoc,protoc-gen-govalidate}`
* And then use the following for generating the `pb.go`files:
  * `protoc --govalidate_out=plugins=grpc:. ./*.proto`
* **Notes**
  * You must have the protobuf3 development libraries installed to `import "google/protobuf/descriptor.proto"`
  * **Linux**: Grab the release from https://github.com/google/protobuf/releases and make sure you include the `include` folder in that archive
  * **Mac** simply `brew install protobuf —devel —without-python`



### Author

 [@njpatel](https://twitter.com/njpatel)



### Generated Validate Function

**user.pb.go (generated)**

```go
// generated by protoc-gen-govalidate
// ...
type User struct {
	Name    string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age     int32   `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Email   string  `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Website string  `protobuf:"bytes,4,opt,name=website" json:"website,omitempty"`
	Rating  float64 `protobuf:"fixed64,5,opt,name=rating" json:"rating,omitempty"`
	Average float32 `protobuf:"fixed32,6,opt,name=average" json:"average,omitempty"`
	Stars   float32 `protobuf:"fixed32,7,opt,name=stars" json:"stars,omitempty"`
	Tags      []string          `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
	Addresses []*Address        `protobuf:"bytes,10,rep,name=addresses" json:"addresses,omitempty"`
	Numbers   map[string]*Phone `protobuf:"bytes,11,rep,name=numbers" json:"numbers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Properties map[string]string `protobuf:"bytes,12,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *User) Validate() (bool, error) {
	changed := false
	for _, f := range m.Addresses {
		if f != nil {
			if change, err := f.Validate(); err != nil {
				return change, err
			} else if change {
				changed = change
			}
		}
	}
	age := m.Age
	if age < 0 {
		age = 0
	}
	if age > 120 {
		age = 120
	}
	changed = changed || age != m.Age
	m.Age = age

	average := m.Average
	averageAsFloat64 := float64(average)
	averageAsFloat64 = math.Min(averageAsFloat64, 0.0)
	averageAsFloat64 = math.Max(averageAsFloat64, 1.0)
	average = float32(averageAsFloat64)
	changed = changed || average != m.Average
	m.Average = average

	email := m.Email

	email = strings.TrimSpace(email)

	email = strings.ToLower(email)

	emailEmailNegated := false
	if result := govalidator.IsEmail(email); (!result && !emailEmailNegated) || (result && emailEmailNegated) {
		return email != m.Email, fmt.Errorf("%s does not validate as %s", email, "email")
	}

	changed = changed || email != m.Email
	m.Email = email

	name := m.Name

	name = strings.TrimSpace(name)

	name = func(s string, l int) string {
		runes := []rune(name)
		if len(runes) > l {
			return string(runes[:l])
		}
		return s
	}(name, 256)

	lengthNameNegated := false
	if result := govalidator.ByteLength(name, "3", "256"); (!result && !lengthNameNegated) || (result && lengthNameNegated) {
		return name != m.Name, fmt.Errorf("%s does not validate as %s", name, "length(3|256)")
	}

	alphaNameNegated := false
	if result := govalidator.IsAlpha(name); (!result && !alphaNameNegated) || (result && alphaNameNegated) {
		return name != m.Name, fmt.Errorf("%s does not validate as %s", name, "alpha")
	}

	changed = changed || name != m.Name
	m.Name = name

	// map validation k: v:
	for k, v := range m.Numbers {
		localK := k
		localV := v

		// validate message. use `ignore` to suppress this validation
		if v != nil {
			if change, err := v.Validate(); err != nil {
				return change, err
			} else if change {
				changed = change
			}
		}

		if k != localK || v != localV {
			delete(m.Numbers, k)
			m.Numbers[localK] = localV
		}
	}
	// map validation k:trim,truncate(50) v: omitempty,trim,truncate(255)
	for k, v := range m.Properties {
		localK := k
		localV := v

		localK = strings.TrimSpace(localK)

		localK = func(s string, l int) string {
			runes := []rune(localK)
			if len(runes) > l {
				return string(runes[:l])
			}
			return s
		}(localK, 50)

		// omitempty
		if strings.TrimSpace(localV) == "" {
			delete(m.Properties, k)
			continue
		}

		localV = strings.TrimSpace(localV)

		localV = func(s string, l int) string {
			runes := []rune(localV)
			if len(runes) > l {
				return string(runes[:l])
			}
			return s
		}(localV, 255)

		if k != localK || v != localV {
			delete(m.Properties, k)
			m.Properties[localK] = localV
		}
	}
	rating := m.Rating
	rating = math.Min(rating, 0.0)
	rating = math.Max(rating, 1.0)
	changed = changed || rating != m.Rating
	m.Rating = rating

	stars := m.Stars
	starsAsFloat64 := float64(stars)
	starsAsFloat64 = math.Floor(starsAsFloat64)
	starsAsFloat64 = math.Min(starsAsFloat64, 0.0)
	starsAsFloat64 = math.Max(starsAsFloat64, 5.0)
	stars = float32(starsAsFloat64)
	changed = changed || stars != m.Stars
	m.Stars = stars

	tags := m.Tags
	// newValidators trim,truncate(50),alpha

	// omitempty, filtering empty values
	tagsFiltered := tags[:0]
	for _, v := range tags {
		if strings.TrimSpace(v) != "" {
			tagsFiltered = append(tagsFiltered, v)
		}
	}
	if len(m.Tags) != len(tagsFiltered) {
		m.Tags = tagsFiltered
		tags = m.Tags
	}

	for i, v := range tags {
		local := v

		local = strings.TrimSpace(local)

		local = func(s string, l int) string {
			runes := []rune(local)
			if len(runes) > l {
				return string(runes[:l])
			}
			return s
		}(local, 50)

		alphaLocalNegated := false
		if result := govalidator.IsAlpha(local); (!result && !alphaLocalNegated) || (result && alphaLocalNegated) {
			return local != v, fmt.Errorf("%s does not validate as %s", local, "alpha")
		}

		if v != local {
			m.Tags[i] = local
			changed = true
		}
	}
	website := m.Website

	website = strings.TrimSpace(website)

	urlWebsiteNegated := false
	if result := govalidator.IsURL(website); (!result && !urlWebsiteNegated) || (result && urlWebsiteNegated) {
		return website != m.Website, fmt.Errorf("%s does not validate as %s", website, "url")
	}

	changed = changed || website != m.Website
	m.Website = website

	return changed, nil
}
```

From that point, you can call `m.Validate()` to validate the your protobuf message before you use it.



# Go support for Protocol Buffers

Google's data interchange format.
Copyright 2010 The Go Authors.
https://github.com/golang/protobuf

This package and the code it generates requires at least Go 1.4.

This software implements Go bindings for protocol buffers.  For
information about protocol buffers themselves, see

```
https://developers.google.com/protocol-buffers/
```



## Installation ##

To use this software, you must:
- Install the standard C++ implementation of protocol buffers from
  https://developers.google.com/protocol-buffers/
- Of course, install the Go compiler and tools from
  https://golang.org/
  See
  https://golang.org/doc/install
  for details or, if you are using gccgo, follow the instructions at
  https://golang.org/doc/install/gccgo
- Grab the code from the repository and install the proto package.
  The simplest way is to run `go get -u github.com/watchly/protobuf/protoc-gen-go`.
  The compiler plugin, protoc-gen-go, will be installed in $GOBIN,
  defaulting to $GOPATH/bin.  It must be in your $PATH for the protocol
  compiler, protoc, to find it.

This software has two parts: a 'protocol compiler plugin' that
generates Go source files that, once compiled, can access and manage
protocol buffers; and a library that implements run-time support for
encoding (marshaling), decoding (unmarshaling), and accessing protocol
buffers.

There is support for gRPC in Go using protocol buffers.
See the note at the bottom of this file for details.

There are no insertion points in the plugin.


## Using protocol buffers with Go ##

Once the software is installed, there are two steps to using it.
First you must compile the protocol buffer definitions and then import
them, with the support library, into your program.

To compile the protocol buffer definition, run protoc with the --go_out
parameter set to the directory you want to output the Go code to.

	protoc --go_out=. *.proto

The generated files will be suffixed .pb.go.  See the Test code below
for an example using such a file.


The package comment for the proto library contains text describing
the interface provided in Go for protocol buffers. Here is an edited
version.

==========

The proto package converts data structures to and from the
wire format of protocol buffers.  It works in concert with the
Go source code generated for .proto files by the protocol compiler.

A summary of the properties of the protocol buffer interface
for a protocol buffer variable v:

-   Names are turned from camel_case to CamelCase for export.
-   There are no methods on v to set fields; just treat
            	them as structure fields.
    - There are getters that return a field's value if set,
      and return the field's default value if unset.
      The getters work even if the receiver is a nil message.
    - The zero value for a struct is its correct initialization state.
      All desired fields must be set before marshaling.
    - A Reset() method will restore a protobuf struct to its zero state.
    - Non-repeated fields are pointers to the values; nil means unset.
      That is, optional or required field int32 f becomes F *int32.
    - Repeated fields are slices.
    - Helper functions are available to aid the setting of fields.
      Helpers for getting values are superseded by the
      GetFoo methods and their use is deprecated.
      msg.Foo = proto.String("hello") // set field
    - Constants are defined to hold the default values of all fields that
      have them.  They have the form Default_StructName_FieldName.
      Because the getter methods handle defaulted values,
      direct use of these constants should be rare.
    - Enums are given type names and maps from names to values.
      Enum values are prefixed with the enum's type name. Enum types have
      a String method, and a Enum method to assist in message construction.
    - Nested groups and enums have type names prefixed with the name of
      the surrounding message type.
    - Extensions are given descriptor names that start with E_,
      followed by an underscore-delimited list of the nested messages
      that contain it (if any) followed by the CamelCased name of the
      extension field itself.  HasExtension, ClearExtension, GetExtension
      and SetExtension are functions for manipulating extensions.
    - Oneof field sets are given a single field in their message,
      with distinguished wrapper types for each possible field value.
    - Marshal and Unmarshal are functions to encode and decode the wire format.

When the .proto file specifies `syntax="proto3"`, there are some differences:

- Non-repeated fields of non-message type are values instead of pointers.
- Getters are only generated for message and oneof fields.
    - Enum types do not get an Enum method.

Consider file test.proto, containing

```proto
	package example;

	enum FOO { X = 17; };

	message Test {
	  required string label = 1;
	  optional int32 type = 2 [default=77];
	  repeated int64 reps = 3;
	  optional group OptionalGroup = 4 {
	    required string RequiredField = 5;
	  }
	}
```

To create and play with a Test object from the example package,

```go
	package main

	import (
		"log"

		"github.com/golang/protobuf/proto"
		"path/to/example"
	)

	func main() {
		test := &example.Test {
			Label: proto.String("hello"),
			Type:  proto.Int32(17),
			Reps:  []int64{1, 2, 3},
			Optionalgroup: &example.Test_OptionalGroup {
				RequiredField: proto.String("good bye"),
			},
		}
		data, err := proto.Marshal(test)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		newTest := &example.Test{}
		err = proto.Unmarshal(data, newTest)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		// Now test and newTest contain the same data.
		if test.GetLabel() != newTest.GetLabel() {
			log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
		}
		// etc.
	}
```

## Parameters ##

To pass extra parameters to the plugin, use a comma-separated
parameter list separated from the output directory by a colon:


	protoc --go_out=plugins=grpc,import_path=mypackage:. *.proto


- `import_prefix=xxx` - a prefix that is added onto the beginning of
  all imports. Useful for things like generating protos in a
  subdirectory, or regenerating vendored protobufs in-place.
- `import_path=foo/bar` - used as the package if no input files
  declare `go_package`. If it contains slashes, everything up to the
  rightmost slash is ignored.
- `plugins=plugin1+plugin2` - specifies the list of sub-plugins to
  load. The only plugin in this repo is `grpc`.
- `Mfoo/bar.proto=quux/shme` - declares that foo/bar.proto is
  associated with Go package quux/shme.  This is subject to the
  import_prefix parameter.

## gRPC Support ##

If a proto file specifies RPC services, protoc-gen-go can be instructed to
generate code compatible with gRPC (http://www.grpc.io/). To do this, pass
the `plugins` parameter to protoc-gen-go; the usual way is to insert it into
the --go_out argument to protoc:

	protoc --go_out=plugins=grpc:. *.proto

## Compatibility ##

The library and the generated code are expected to be stable over time.
However, we reserve the right to make breaking changes without notice for the
following reasons:

- Security. A security issue in the specification or implementation may come to
  light whose resolution requires breaking compatibility. We reserve the right
  to address such security issues.
- Unspecified behavior.  There are some aspects of the Protocol Buffers
  specification that are undefined.  Programs that depend on such unspecified
  behavior may break in future releases.
- Specification errors or changes. If it becomes necessary to address an
  inconsistency, incompleteness, or change in the Protocol Buffers
  specification, resolving the issue could affect the meaning or legality of
  existing programs.  We reserve the right to address such issues, including
  updating the implementations.
- Bugs.  If the library has a bug that violates the specification, a program
  that depends on the buggy behavior may break if the bug is fixed.  We reserve
  the right to fix such bugs.
- Adding methods or fields to generated structs.  These may conflict with field
  names that already exist in a schema, causing applications to break.  When the
  code generator encounters a field in the schema that would collide with a
  generated field or method name, the code generator will append an underscore
  to the generated field or method name.
- Adding, removing, or changing methods or fields in generated structs that
  start with `XXX`.  These parts of the generated code are exported out of
  necessity, but should not be considered part of the public API.
- Adding, removing, or changing unexported symbols in generated code.

Any breaking changes outside of these will be announced 6 months in advance to
protobuf@googlegroups.com.

You should, whenever possible, use generated code created by the `protoc-gen-go`
tool built at the same commit as the `proto` package.  The `proto` package
declares package-level constants in the form `ProtoPackageIsVersionX`.
Application code and generated code may depend on one of these constants to
ensure that compilation will fail if the available version of the proto library
is too old.  Whenever we make a change to the generated code that requires newer
library support, in the same commit we will increment the version number of the
generated code and declare a new package-level constant whose name incorporates
the latest version number.  Removing a compatibility constant is considered a
breaking change and would be subject to the announcement policy stated above.

The `protoc-gen-go/generator` package exposes a plugin interface,
which is used by the gRPC code generation. This interface is not
supported and is subject to incompatible changes without notice.
