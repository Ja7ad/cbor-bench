package bench

import (
	"bytes"
	cborFx "github.com/fxamacker/cbor/v2"
	"github.com/ugorji/go/codec"
	cbor "github.com/whyrusleeping/cbor/go"
	"testing"
)

type Sample struct {
	Name  string
	Age   int
	Email string
}

func BenchmarkFxamacker_CBOREncode(b *testing.B) {
	b.ReportAllocs()
	sampleData := Sample{Name: "Alice", Age: 30, Email: "alice@example.com"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)

		err := cborFx.NewEncoder(buf).Encode(sampleData)
		if err != nil {
			b.Fatalf("Error encoding: %v", err)
		}
	}
}

func BenchmarkFxamacker_CBORDecode(b *testing.B) {
	b.ReportAllocs()
	sampleData := Sample{Name: "Alice", Age: 30, Email: "alice@example.com"}

	buf := new(bytes.Buffer)
	err := cborFx.NewEncoder(buf).Encode(sampleData)
	if err != nil {
		b.Fatalf("Error encoding: %v", err)
	}

	var decodedData Sample

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := cborFx.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(&decodedData)
		if err != nil {
			b.Fatalf("Error decoding: %v", err)
		}
	}
}

func BenchmarkUgorji_CBOREncode(b *testing.B) {
	b.ReportAllocs()
	h := new(codec.CborHandle)
	sampleData := Sample{Name: "Alice", Age: 30, Email: "alice@example.com"}
	var encodedData []byte

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc := codec.NewEncoderBytes(&encodedData, h)
		err := enc.Encode(sampleData)
		if err != nil {
			b.Fatalf("Error encoding: %v", err)
		}
	}
}

func BenchmarkUgorji_CBORDecode(b *testing.B) {
	b.ReportAllocs()
	h := new(codec.CborHandle)
	sampleData := Sample{Name: "Alice", Age: 30, Email: "alice@example.com"}
	var encodedData []byte

	enc := codec.NewEncoderBytes(&encodedData, h)
	err := enc.Encode(sampleData)
	if err != nil {
		b.Fatalf("Error encoding: %v", err)
	}

	var decodedData Sample

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dec := codec.NewDecoderBytes(encodedData, h)
		err := dec.Decode(&decodedData)
		if err != nil {
			b.Fatalf("Error decoding: %v", err)
		}
	}
}

func BenchmarkWhyrusleeping_CBOREncode(b *testing.B) {
	b.ReportAllocs()
	sampleData := Sample{Name: "Alice", Age: 30, Email: "alice@example.com"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		encoder := cbor.NewEncoder(buf)

		err := encoder.Encode(sampleData)
		if err != nil {
			b.Fatalf("Error encoding: %v", err)
		}
	}
}

func BenchmarkWhyrusleeping_CBORDecode(b *testing.B) {
	b.ReportAllocs()
	sampleData := Sample{Name: "Alice", Age: 30, Email: "alice@example.com"}

	buf := new(bytes.Buffer)
	encoder := cbor.NewEncoder(buf)
	err := encoder.Encode(sampleData)
	if err != nil {
		b.Fatalf("Error encoding: %v", err)
	}

	var decodedData Sample

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decoder := cbor.NewDecoder(bytes.NewReader(buf.Bytes()))
		err := decoder.Decode(&decodedData)
		if err != nil {
			b.Fatalf("Error decoding: %v", err)
		}
	}
}
