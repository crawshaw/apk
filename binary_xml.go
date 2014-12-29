package apk

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
)

// binaryXML converts XML into Android's undocumented binary XML format.
//
// The best source of information on this format seems to be the source code
// in AOSP frameworks-base. Android "resource" types seem to describe the
// encoded bytes, in particular:
//
//	ResChunk_header
//	ResStringPool_header
//	ResXMLTree_node
//	ResXMLTree_attribute TODO(crawshaw): really?
//
// These are defined in:
//
//	https://android.googlesource.com/platform/frameworks/base/+/master/include/androidfw/ResourceTypes.h
//
// The rough format of the file is a resource chunk containing a sequence of
// chunks:
//
//	ResChunk_header (type XML)
//	ResStringPool_header
//	Sequence of strings, each with the format:
//		uint16 length
//		uint16 extended_length -- only if top bit set on length
//		UTF-16LE string
//		two zero bytes
//	Resource Map
//		The [i]th 4-byte entry in the resource map corresponds with
//		the [i]th string from the string pool. The 4-bytes are a
//		Resource ID constant defined:
//			http://developer.android.com/reference/android/R.attr.html
//		This appears to be a way to map strings onto enum values.
//	Chunk: Namespace Start
//	Chunk: Element Start
//	Chunk: Element End
//	...
//	Chunk: Namespace End
//
// Values are encoded as little-endian.
func binaryXML(r io.Reader) ([]byte, error) {
	d := xml.NewDecoder(r)

	for {
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch tok.(type) {
		case xml.StartElement:
			log.Printf("StartElement: %v", tok)
		case xml.EndElement:
			log.Printf("EndElement: %v", tok)
		case xml.CharData:
			log.Printf("CharData: %v", tok)
		case xml.Comment:
			log.Printf("Comment: %v", tok)
		case xml.ProcInst:
			log.Printf("ProcInst: %v", tok)
		case xml.Directive:
			log.Printf("Directive: %v", tok)
		default:
			return nil, fmt.Errorf("apk: unexpected token: %v (%T)", tok, tok)
		}
	}

	buf := new(bytes.Buffer)

	//buf.WriteByte([]byte{0x03, 0x00}) // magic
	//buf.WriteByte([]byte{0x08, 0x00}) //

	return buf.Bytes(), nil
}
