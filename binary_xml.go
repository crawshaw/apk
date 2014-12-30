package apk

import (
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
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
//
// These are defined in:
//
//	https://android.googlesource.com/platform/frameworks/base/+/master/include/androidfw/ResourceTypes.h
//
// The rough format of the file is a resource chunk containing a sequence of
// chunks:
//
//	File Chunk Header (ResChunk_header, type XML)
//	String Pool Header (ResStringPool_header, type STRING_POOL)
//	Sequence of strings, each with the format:
//		uint16 length
//		uint16 extended_length -- only if top bit set on length
//		UTF-16LE string
//		two zero bytes
//	Resource Map
//		TODO: maybe optional? try not generating it and see what happens.
//		The [i]th 4-byte entry in the resource map corresponds with
//		the [i]th string from the string pool. The 4-bytes are a
//		Resource ID constant defined:
//			http://developer.android.com/reference/android/R.attr.html
//		This appears to be a way to map strings onto enum values.
//	Chunk: Namespace Start (ResXMLTree_node; ResXMLTree_namespaceExt)
//	Chunk: Element Start
//		ResXMLTree_node
//		ResXMLTree_attrExt
//		ResXMLTree_attribute (repeated attributeCount times)
//	Chunk: Element End
//		(ResXMLTree_node; ResXMLTree_endElementExt)
//	...
//	Chunk: Namespace End
//
// Values are encoded as little-endian.
func binaryXML(r io.Reader) ([]byte, error) {
	lr := &lineReader{r: r}
	d := xml.NewDecoder(lr)

	pool := new(binStringPool)
	depth := 0
	elements := []interface{}{}
	namespaceEnds := make(map[int]binEndNamspace)

	for {
		line := lr.line(d.InputOffset())
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			// Intercept namespace definitions.
			var attr []binAttr
			for _, a := range tok.Attr {
				if a.Name.Space == "xmlns" {
					elements = append(elements, binStartNamspace{
						line:   line,
						prefix: pool.get(a.Name.Local),
						url:    pool.get(a.Value),
					})
					namespaceEnds[depth] = binEndNamspace{
						line:   line,
						prefix: pool.get(a.Name.Local),
						url:    pool.get(a.Value),
					}
					continue
				}
				ba, err := pool.getAttr(a)
				if err != nil {
					return nil, fmt.Errorf("%d: %s: %v", line, a.Name.Local, err)
				}
				attr = append(attr, ba)
			}

			depth++
			elements = append(elements, binStartElement{
				line: line,
				ns:   pool.get(tok.Name.Space),
				name: pool.get(tok.Name.Local),
				attr: attr,
			})
		case xml.EndElement:
			elements = append(elements, binEndElement{
				line: line,
				ns:   pool.get(tok.Name.Space),
				name: pool.get(tok.Name.Local),
			})
			depth--
			if nsEnd, ok := namespaceEnds[depth]; ok {
				delete(namespaceEnds, depth)
				elements = append(elements, nsEnd)
			}
		case xml.CharData:
			s := strings.TrimSpace(string(tok))
			if s == "" {
				continue
			}
			s = "\t" + s + "\n" // TODO just for test case
			pool.get(s)
		case xml.Comment:
			// Ignored by Anroid Binary XML format.
		case xml.ProcInst:
			// Ignored by Anroid Binary XML format?
		case xml.Directive:
			// Ignored by Anroid Binary XML format.
		default:
			return nil, fmt.Errorf("apk: unexpected token: %v (%T)", tok, tok)
		}
	}

	sortPool(pool)
	size := 8 + pool.size()
	//for _, e := range elements {
	//}

	b := []byte{}
	b = appendHeader(b, headerXML, size)
	b = pool.append(b)

	return b, nil
}

type headerType uint16

const (
	headerXML            headerType = 0x0003
	headerStringPool                = 0x0001
	headerResourceMap               = 0x0180
	headerStartNamespace            = 0x0100
	headerEndNamespace              = 0x0101
	headerStartElement              = 0x0102
	headerEndElement                = 0x0103
)

func appendU16(b []byte, v uint16) []byte {
	return append(b, byte(v), byte(v>>8))
}

func appendU32(b []byte, v uint32) []byte {
	return append(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

func appendHeader(b []byte, typ headerType, size int) []byte {
	b = appendU16(b, uint16(typ))
	b = appendU16(b, 8)
	b = appendU16(b, uint16(size))
	b = appendU16(b, 0)
	return b
}

// Attributes of the form android:key are mapped to resource IDs, which are
// embedded into the Binary XML format.
//
// http://developer.android.com/reference/android/R.attr.html
var resourceCodes = map[string]uint32{
	"versionCode":   0x0101021b,
	"versionName":   0x0101021c,
	"minSdkVersion": 0x0101020c,
	"label":         0x01010001,
	"hasCode":       0x0101000c,
	"debuggable":    0x0101000f,
	"name":          0x01010003,
	"configChanges": 0x0101001f,
	"value":         0x01010024,
}

// http://developer.android.com/reference/android/R.attr.html#configChanges
var configChanges = map[string]uint32{
	"mcc":                0x0001,
	"mnc":                0x0002,
	"locale":             0x0004,
	"touchscreen":        0x0008,
	"keyboard":           0x0010,
	"keyboardHidden":     0x0020,
	"navigation":         0x0040,
	"orientation":        0x0080,
	"screenLayout":       0x0100,
	"uiMode":             0x0200,
	"screenSize":         0x0400,
	"smallestScreenSize": 0x0800,
	"layoutDirection":    0x2000,
	"fontScale":          0x40000000,
}

type lineReader struct {
	off   int64
	lines []int64
	r     io.Reader
}

func (r *lineReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		if p[i] == '\n' {
			r.lines = append(r.lines, r.off+int64(i))
		}
	}
	r.off += int64(n)
	return n, err
}

func (r *lineReader) line(pos int64) int {
	return sort.Search(len(r.lines), func(i int) bool {
		return pos < r.lines[i]
	}) + 1
}

type bstring struct {
	ind uint32
	str string
	enc []byte // 2-byte length, utf16le, 2-byte zero
}

type binStringPool struct {
	s []*bstring
	m map[string]*bstring
}

func (p *binStringPool) get(str string) *bstring {
	if p.m == nil {
		p.m = make(map[string]*bstring)
	}
	res := p.m[str]
	if res != nil {
		return res
	}
	res = &bstring{
		ind: uint32(len(p.s)),
		str: str,
	}
	p.s = append(p.s, res)
	p.m[str] = res

	if len(str)>>16 > 0 {
		panic(fmt.Sprintf("string lengths over 1<<15 not yet supported, got len %d for string that starts %q", len(str), str[:100]))
	}
	res.enc = appendU16(nil, uint16(len(str)))
	for _, w := range utf16.Encode([]rune(str)) {
		res.enc = appendU16(res.enc, w)
	}
	res.enc = appendU16(res.enc, 0)
	return res
}

func (p *binStringPool) getAttr(attr xml.Attr) (binAttr, error) {
	a := binAttr{
		ns:   p.get(attr.Name.Space),
		name: p.get(attr.Name.Local),
	}
	if attr.Name.Space != "http://schemas.android.com/apk/res/android" {
		a.data = p.get(attr.Value)
		return a, nil
	}

	// Some android attributes have interesting values.
	switch attr.Name.Local {
	case "versionCode", "minSdkVersion":
		v, err := strconv.Atoi(attr.Value)
		if err != nil {
			return binAttr{}, err
		}
		a.data = int(v)
	case "hasCode", "debuggable":
		v, err := strconv.ParseBool(attr.Value)
		if err != nil {
			return binAttr{}, err
		}
		a.data = v
	case "configChanges":
		v := uint32(0)
		for _, c := range strings.Split(attr.Value, "|") {
			v |= configChanges[c]
		}
		a.data = v
	default:
		a.data = p.get(attr.Value)
	}
	return a, nil
}

const stringPoolPreamble = 0 +
	8 + // chunk header
	4 + // string count
	4 + // style count
	4 + // flags
	4 + // strings start
	4 + // styles start
	0

func (p *binStringPool) size() int {
	strLens := 0
	for _, s := range p.s {
		//log.Printf("len=%02d, enc=%02d value=%q", len(s.str), len(s.enc), s.str)
		strLens += len(s.enc)
	}
	return stringPoolPreamble + 4*len(p.s) + strLens + 2
}

var sortPool = func(p *binStringPool) { sort.Sort(p) }

func (b *binStringPool) Len() int           { return len(b.s) }
func (b *binStringPool) Less(i, j int) bool { return b.s[i].str < b.s[j].str }
func (b *binStringPool) Swap(i, j int) {
	b.s[i], b.s[j] = b.s[j], b.s[i]
	b.s[i].ind, b.s[j].ind = b.s[j].ind, b.s[i].ind
}

func (p *binStringPool) append(b []byte) []byte {
	stringsStart := uint32(stringPoolPreamble + 4*len(p.s))
	b = appendU16(b, uint16(headerStringPool))
	b = appendU16(b, 0x1c) // chunk header size
	b = appendU16(b, uint16(p.size()))
	b = appendU16(b, 0)
	b = appendU32(b, uint32(len(p.s)))
	b = appendU32(b, 0) // style count
	b = appendU32(b, 0) // flags
	b = appendU32(b, stringsStart)
	b = appendU32(b, 0) // styles start

	off := 0
	for _, bstr := range p.s {
		b = appendU32(b, uint32(off))
		off += len(bstr.enc)
	}
	for _, bstr := range p.s {
		b = append(b, bstr.enc...)
	}
	b = appendU16(b, 0)
	return b
}

type binStartElement struct {
	line int
	ns   *bstring
	name *bstring
	attr []binAttr
}

func (b *binStartElement) size() int {
	return 8 + // chunk header
		4 + // line number
		4 + // comment
		4 + // ns
		4 + // name
		len(b.attr)*(4+4+4+4+4)
}

type binAttr struct {
	ns   *bstring
	name *bstring
	data interface{} // either int (INT_DEC) or *bstring (STRING)
}

func (a *binAttr) append(b []byte) []byte {
	b = appendU32(b, a.ns.ind)
	b = appendU32(b, a.name.ind)
	b = appendU32(b, 0xffffffff) // raw value
	b = appendU16(b, 8)          // size
	b = appendU16(b, 0)          // unused padding
	switch v := a.data.(type) {
	case int:
		b = append(b, 0x10) // INT_DEC
		b = appendU32(b, uint32(v))
	case bool:
		b = append(b, 0x12) // INT_BOOLEAN
		if v {
			b = appendU32(b, 1)
		} else {
			b = appendU32(b, 0)
		}
	case uint32:
		b = append(b, 0x10) // TODO double check configChanges
		b = appendU32(b, uint32(v))
	default:
		panic(fmt.Sprintf("unexpected attr type: %T (%v)", v, v))
	}
	return b
}

type binEndElement struct {
	line int
	ns   *bstring
	name *bstring
	attr []binAttr
}

func (*binEndElement) size() int {
	return 8 + // chunk header
		4 + // line number
		4 + // comment
		4 + // ns
		4 // name
}

type binStartNamspace struct {
	line   int
	prefix *bstring
	url    *bstring
}

type binEndNamspace struct {
	line   int
	prefix *bstring
	url    *bstring
}
