package apk

import (
	"bytes"
	"testing"
)

func TestBinaryXML(t *testing.T) {
	got, err := binaryXML(bytes.NewBufferString(input))
	if err != nil {
		t.Fatal(err)
	}
	_ = got
	t.Error("TODO")
}

// Hexdump of output generated by the Android SDK's ant build system.
// Annotated after studying Android source code.
var output = []byte{
	/* 0000 */ 0x03, 0x00, 0x08, 0x00, //  chunk header XML
	/* 0004 */ 0x78, 0x07, 0x00, 0x00, //  chunk size 1912

	/* 0008 */ 0x01, 0x00, 0x1c, 0x00, //  chunk header STRING_POOL
	/* 000c */ 0x00, 0x04, 0x00, 0x00, //  chunk size 1024
	/* 0010 */ 0x1f, 0x00, 0x00, 0x00, //  string count 31
	/* 0014 */ 0x00, 0x00, 0x00, 0x00, //  style count 0
	/* 0018 */ 0x00, 0x00, 0x00, 0x00, //  flags (none set means UTF-16)
	/* 001c */ 0x98, 0x00, 0x00, 0x00, //  strings_start 0x98+0x08 = 0xa0
	/* 0020 */ 0x00, 0x00, 0x00, 0x00, //  styles_start (none)
	/* 0024 */ 0x00, 0x00, 0x00, 0x00, //  string offset [0x00] (from strings_start)
	/* 0028 */ 0x1a, 0x00, 0x00, 0x00, //  string offset [0x01]
	/* 002c */ 0x34, 0x00, 0x00, 0x00, //  string offset [0x02]
	/* 0030 */ 0x52, 0x00, 0x00, 0x00, //  string offset [0x03]
	/* 0034 */ 0x60, 0x00, 0x00, 0x00, //  string offset [0x04]
	/* 0038 */ 0x72, 0x00, 0x00, 0x00, //  string offset [0x05]
	/* 003c */ 0x8a, 0x00, 0x00, 0x00, //  string offset [0x06]
	/* 0040 */ 0x96, 0x00, 0x00, 0x00, //  string offset [0x07]
	/* 0044 */ 0xb4, 0x00, 0x00, 0x00, //  string offset [0x08]
	/* 0048 */ 0xc2, 0x00, 0x00, 0x00, //  string offset [0x09]
	/* 004c */ 0xd4, 0x00, 0x00, 0x00, //  string offset [0x0a]
	/* 0050 */ 0x2c, 0x01, 0x00, 0x00, //  string offset [0x0b]
	/* 0054 */ 0x30, 0x01, 0x00, 0x00, //  string offset [0x0c]
	/* 0058 */ 0x42, 0x01, 0x00, 0x00, //  string offset [0x0d]
	/* 005c */ 0x56, 0x01, 0x00, 0x00, //  string offset [0x0e]
	/* 0060 */ 0x7e, 0x01, 0x00, 0x00, //  string offset [0x0f]
	/* 0064 */ 0x88, 0x01, 0x00, 0x00, //  string offset [0x10]
	/* 0068 */ 0x9c, 0x01, 0x00, 0x00, //  string offset [0x11]
	/* 006c */ 0xb6, 0x01, 0x00, 0x00, //  string offset [0x12]
	/* 0070 */ 0xcc, 0x01, 0x00, 0x00, //  string offset [0x13]
	/* 0074 */ 0xe0, 0x01, 0x00, 0x00, //  string offset [0x14]
	/* 0078 */ 0x18, 0x02, 0x00, 0x00, //  string offset [0x15]
	/* 007c */ 0x2a, 0x02, 0x00, 0x00, //  string offset [0x16]
	/* 0080 */ 0x40, 0x02, 0x00, 0x00, //  string offset [0x17]
	/* 0084 */ 0x6c, 0x02, 0x00, 0x00, //  string offset [0x18]
	/* 0088 */ 0x7e, 0x02, 0x00, 0x00, //  string offset [0x19]
	/* 008c */ 0x9c, 0x02, 0x00, 0x00, //  string offset [0x1a]
	/* 0090 */ 0xc6, 0x02, 0x00, 0x00, //  string offset [0x1b]
	/* 0094 */ 0xd6, 0x02, 0x00, 0x00, //  string offset [0x1c]
	/* 0098 */ 0x0e, 0x03, 0x00, 0x00, //  string offset [0x1d]
	/* 009c */ 0x22, 0x03, 0x00, 0x00, //  string offset [0x1e]
	/* 00a0 */ 0x0b, 0x00, 0x76, 0x00, //  [0x00] len=11 value="versionCode"
	/* 00a4 */ 0x65, 0x00, 0x72, 0x00,
	/* 00a8 */ 0x73, 0x00, 0x69, 0x00,
	/* 00ac */ 0x6f, 0x00, 0x6e, 0x00,
	/* 00b0 */ 0x43, 0x00, 0x6f, 0x00,
	/* 00b4 */ 0x64, 0x00, 0x65, 0x00,
	/* 00b8 */ 0x00, 0x00,
	/* 00ba */ 0x0b, 0x00, //  [0x01] len=11 value="versionName"
	/* 00bc */ 0x76, 0x00, 0x65, 0x00,
	/* 00c0 */ 0x72, 0x00, 0x73, 0x00,
	/* 00c4 */ 0x69, 0x00, 0x6f, 0x00,
	/* 00c8 */ 0x6e, 0x00, 0x4e, 0x00,
	/* 00cc */ 0x61, 0x00, 0x6d, 0x00,
	/* 00d0 */ 0x65, 0x00, 0x00, 0x00,
	/* 00d4 */ 0x0d, 0x00, 0x6d, 0x00, //  [0x02] len=13 value="minSdkVersion"
	/* 00d8 */ 0x69, 0x00, 0x6e, 0x00,
	/* 00dc */ 0x53, 0x00, 0x64, 0x00,
	/* 00e0 */ 0x6b, 0x00, 0x56, 0x00,
	/* 00e4 */ 0x65, 0x00, 0x72, 0x00,
	/* 00e8 */ 0x73, 0x00, 0x69, 0x00,
	/* 00ec */ 0x6f, 0x00, 0x6e, 0x00,
	/* 00f0 */ 0x00, 0x00,
	/* 00f2 */ 0x05, 0x00, //  [0x03] len=5 value="label"
	/* 00f4 */ 0x6c, 0x00, 0x61, 0x00,
	/* 00f8 */ 0x62, 0x00, 0x65, 0x00,
	/* 00fc */ 0x6c, 0x00, 0x00, 0x00,
	/* 0100 */ 0x07, 0x00, 0x68, 0x00, //  [0x04] len=7 value="hasCode"
	/* 0104 */ 0x61, 0x00, 0x73, 0x00,
	/* 0108 */ 0x43, 0x00, 0x6f, 0x00,
	/* 010c */ 0x64, 0x00, 0x65, 0x00,
	/* 0110 */ 0x00, 0x00,
	/* 0112 */ 0x0a, 0x00, //  [0x05] len=10 value="debuggable" (NOT IN ORIGINAL)
	/* 0114 */ 0x64, 0x00, 0x65, 0x00,
	/* 0118 */ 0x62, 0x00, 0x75, 0x00,
	/* 011c */ 0x67, 0x00, 0x67, 0x00,
	/* 0120 */ 0x61, 0x00, 0x62, 0x00,
	/* 0124 */ 0x6c, 0x00, 0x65, 0x00,
	/* 0128 */ 0x00, 0x00,
	/* 012a */ 0x04, 0x00, //  [0x06] len=4 value="name"
	/* 012c */ 0x6e, 0x00, 0x61, 0x00,
	/* 0130 */ 0x6d, 0x00, 0x65, 0x00,
	/* 0134 */ 0x00, 0x00,
	/* 0136 */ 0x0d, 0x00, //  [0x07] len=13 value="configChanges"
	/* 0138 */ 0x63, 0x00, 0x6f, 0x00,
	/* 013c */ 0x6e, 0x00, 0x66, 0x00,
	/* 0140 */ 0x69, 0x00, 0x67, 0x00,
	/* 0144 */ 0x43, 0x00, 0x68, 0x00,
	/* 0148 */ 0x61, 0x00, 0x6e, 0x00,
	/* 014c */ 0x67, 0x00, 0x65, 0x00,
	/* 0150 */ 0x73, 0x00, 0x00, 0x00,
	/* 0154 */ 0x05, 0x00, 0x76, 0x00, //  [0x08] len=5 value="value"
	/* 0158 */ 0x61, 0x00, 0x6c, 0x00,
	/* 015c */ 0x75, 0x00, 0x65, 0x00,
	/* 0160 */ 0x00, 0x00,
	/* 0162 */ 0x07, 0x00, //  [0x09] len=7 value="android"
	/* 0164 */ 0x61, 0x00, 0x6e, 0x00,
	/* 0168 */ 0x64, 0x00, 0x72, 0x00,
	/* 016c */ 0x6f, 0x00, 0x69, 0x00,
	/* 0170 */ 0x64, 0x00, 0x00, 0x00,
	/* 0174 */ 0x2a, 0x00, 0x68, 0x00, //  [0x0a] len=42 value="http://schemas.android.com/apk/res/android"
	/* 0178 */ 0x74, 0x00, 0x74, 0x00,
	/* 017c */ 0x70, 0x00, 0x3a, 0x00,
	/* 0180 */ 0x2f, 0x00, 0x2f, 0x00,
	/* 0184 */ 0x73, 0x00, 0x63, 0x00,
	/* 0188 */ 0x68, 0x00, 0x65, 0x00,
	/* 018c */ 0x6d, 0x00, 0x61, 0x00,
	/* 0190 */ 0x73, 0x00, 0x2e, 0x00,
	/* 0194 */ 0x61, 0x00, 0x6e, 0x00,
	/* 0198 */ 0x64, 0x00, 0x72, 0x00,
	/* 019c */ 0x6f, 0x00, 0x69, 0x00,
	/* 01a0 */ 0x64, 0x00, 0x2e, 0x00,
	/* 01a4 */ 0x63, 0x00, 0x6f, 0x00,
	/* 01a8 */ 0x6d, 0x00, 0x2f, 0x00,
	/* 01ac */ 0x61, 0x00, 0x70, 0x00,
	/* 01b0 */ 0x6b, 0x00, 0x2f, 0x00,
	/* 01b4 */ 0x72, 0x00, 0x65, 0x00,
	/* 01b8 */ 0x73, 0x00, 0x2f, 0x00,
	/* 01bc */ 0x61, 0x00, 0x6e, 0x00,
	/* 01c0 */ 0x64, 0x00, 0x72, 0x00,
	/* 01c4 */ 0x6f, 0x00, 0x69, 0x00,
	/* 01c8 */ 0x64, 0x00, 0x00, 0x00,
	/* 01cc */ 0x00, 0x00, 0x00, 0x00, //  [0x0b] len=0 (sigh)
	/* 01d0 */ 0x07, 0x00, 0x70, 0x00, //  [0x0c] len=7 value="package"
	/* 01d4 */ 0x61, 0x00, 0x63, 0x00,
	/* 01d8 */ 0x6b, 0x00, 0x61, 0x00,
	/* 01dc */ 0x67, 0x00, 0x65, 0x00,
	/* 01e0 */ 0x00, 0x00,
	/* 01e2 */ 0x08, 0x00, //  [0x0d] len=8 value="manifest"
	/* 01e4 */ 0x6d, 0x00, 0x61, 0x00,
	/* 01e8 */ 0x6e, 0x00, 0x69, 0x00,
	/* 01ec */ 0x66, 0x00, 0x65, 0x00,
	/* 01f0 */ 0x73, 0x00, 0x74, 0x00,
	/* 01f4 */ 0x00, 0x00,
	/* 01f6 */ 0x12, 0x00, //  [0x0e] len=12 value="com.zentus.balloon"
	/* 01f8 */ 0x63, 0x00, 0x6f, 0x00,
	/* 01fc */ 0x6d, 0x00, 0x2e, 0x00,
	/* 0200 */ 0x7a, 0x00, 0x65, 0x00,
	/* 0204 */ 0x6e, 0x00, 0x74, 0x00,
	/* 0208 */ 0x75, 0x00, 0x73, 0x00,
	/* 020c */ 0x2e, 0x00, 0x62, 0x00,
	/* 0210 */ 0x61, 0x00, 0x6c, 0x00,
	/* 0214 */ 0x6c, 0x00, 0x6f, 0x00,
	/* 0218 */ 0x6f, 0x00, 0x6e, 0x00,
	/* 021c */ 0x00, 0x00,
	/* 021e */ 0x03, 0x00, //  [0x0f] len=3 value="1.0"
	/* 0220 */ 0x31, 0x00, 0x2e, 0x00,
	/* 0224 */ 0x30, 0x00, 0x00, 0x00,
	/* 0228 */ 0x08, 0x00, 0x75, 0x00, //  [0x10] len=8 value="uses-sdk"
	/* 022c */ 0x73, 0x00, 0x65, 0x00,
	/* 0230 */ 0x73, 0x00, 0x2d, 0x00,
	/* 0234 */ 0x73, 0x00, 0x64, 0x00,
	/* 0238 */ 0x6b, 0x00, 0x00, 0x00,
	/* 023c */ 0x0b, 0x00, 0x61, 0x00, //  [0x11] len=11 value="application"
	/* 0240 */ 0x70, 0x00, 0x70, 0x00,
	/* 0244 */ 0x6c, 0x00, 0x69, 0x00,
	/* 0248 */ 0x63, 0x00, 0x61, 0x00,
	/* 024c */ 0x74, 0x00, 0x69, 0x00,
	/* 0250 */ 0x6f, 0x00, 0x6e, 0x00,
	/* 0254 */ 0x00, 0x00,
	/* 0256 */ 0x09, 0x00, //  [0x12] len=9 value="Balloon世界" (UTF16-LE, 0x4e16 is "16 4e", etc)
	/* 0258 */ 0x42, 0x00, 0x61, 0x00,
	/* 025c */ 0x6c, 0x00, 0x6c, 0x00,
	/* 0260 */ 0x6f, 0x00, 0x6f, 0x00,
	/* 0264 */ 0x6e, 0x00, 0x16, 0x4e,
	/* 0268 */ 0x4c, 0x75, 0x00, 0x00,
	/* 026c */ 0x08, 0x00, 0x61, 0x00, //  [0x13] len=8 value="activity"
	/* 0270 */ 0x63, 0x00, 0x74, 0x00,
	/* 0274 */ 0x69, 0x00, 0x76, 0x00,
	/* 0278 */ 0x69, 0x00, 0x74, 0x00,
	/* 027c */ 0x79, 0x00, 0x00, 0x00,
	/* 0280 */ 0x1a, 0x00, 0x61, 0x00, //  [0x14] len=26 value="android.app.NativeActivity"
	/* 0284 */ 0x6e, 0x00, 0x64, 0x00,
	/* 0288 */ 0x72, 0x00, 0x6f, 0x00,
	/* 028c */ 0x69, 0x00, 0x64, 0x00,
	/* 0290 */ 0x2e, 0x00, 0x61, 0x00,
	/* 0294 */ 0x70, 0x00, 0x70, 0x00,
	/* 0298 */ 0x2e, 0x00, 0x4e, 0x00,
	/* 029c */ 0x61, 0x00, 0x74, 0x00,
	/* 02a0 */ 0x69, 0x00, 0x76, 0x00,
	/* 02a4 */ 0x65, 0x00, 0x41, 0x00,
	/* 02a8 */ 0x63, 0x00, 0x74, 0x00,
	/* 02ac */ 0x69, 0x00, 0x76, 0x00,
	/* 02b0 */ 0x69, 0x00, 0x74, 0x00,
	/* 02b4 */ 0x79, 0x00, 0x00, 0x00,
	/* 02b8 */ 0x07, 0x00, 0x42, 0x00, //  [0x15] len=7 value="Balloon"
	/* 02bc */ 0x61, 0x00, 0x6c, 0x00,
	/* 02c0 */ 0x6c, 0x00, 0x6f, 0x00,
	/* 02c4 */ 0x6f, 0x00, 0x6e, 0x00,
	/* 02c8 */ 0x00, 0x00,
	/* 02ca */ 0x09, 0x00, //  [0x16] len=9 value="meta-data"
	/* 02cc */ 0x6d, 0x00, 0x65, 0x00,
	/* 02d0 */ 0x74, 0x00, 0x61, 0x00,
	/* 02d4 */ 0x2d, 0x00, 0x64, 0x00,
	/* 02d8 */ 0x61, 0x00, 0x74, 0x00,
	/* 02dc */ 0x61, 0x00, 0x00, 0x00,
	/* 02e0 */ 0x14, 0x00, 0x61, 0x00, //  [0x17] len=20
	/* 02e4 */ 0x6e, 0x00, 0x64, 0x00,
	/* 02e8 */ 0x72, 0x00, 0x6f, 0x00,
	/* 02ec */ 0x69, 0x00, 0x64, 0x00,
	/* 02f0 */ 0x2e, 0x00, 0x61, 0x00,
	/* 02f4 */ 0x70, 0x00, 0x70, 0x00,
	/* 02f8 */ 0x2e, 0x00, 0x6c, 0x00,
	/* 02fc */ 0x69, 0x00, 0x62, 0x00,
	/* 0300 */ 0x5f, 0x00, 0x6e, 0x00,
	/* 0304 */ 0x61, 0x00, 0x6d, 0x00,
	/* 0308 */ 0x65, 0x00, 0x00, 0x00,
	/* 030c */ 0x07, 0x00, 0x62, 0x00, //  [0x18] len=7 value="balloon"
	/* 0310 */ 0x61, 0x00, 0x6c, 0x00,
	/* 0314 */ 0x6c, 0x00, 0x6f, 0x00,
	/* 0318 */ 0x6f, 0x00, 0x6e, 0x00,
	/* 031c */ 0x00, 0x00,
	/* 031e */ 0x0d, 0x00, //  [0x19] len=13 value="intent-filter"
	/* 0320 */ 0x69, 0x00, 0x6e, 0x00,
	/* 0324 */ 0x74, 0x00, 0x65, 0x00,
	/* 0328 */ 0x6e, 0x00, 0x74, 0x00,
	/* 032c */ 0x2d, 0x00, 0x66, 0x00,
	/* 0330 */ 0x69, 0x00, 0x6c, 0x00,
	/* 0334 */ 0x74, 0x00, 0x65, 0x00,
	/* 0338 */ 0x72, 0x00, 0x00, 0x00,
	/* 033c */ 0x13, 0x00, 0x09, 0x00, //  [0x1a] len=19 value="here is some text"
	/* 0340 */ 0x68, 0x00, 0x65, 0x00,
	/* 0344 */ 0x72, 0x00, 0x65, 0x00,
	/* 0348 */ 0x20, 0x00, 0x69, 0x00,
	/* 034c */ 0x73, 0x00, 0x20, 0x00,
	/* 0350 */ 0x73, 0x00, 0x6f, 0x00,
	/* 0354 */ 0x6d, 0x00, 0x65, 0x00,
	/* 0358 */ 0x20, 0x00, 0x74, 0x00,
	/* 035c */ 0x65, 0x00, 0x78, 0x00,
	/* 0360 */ 0x74, 0x00, 0x0a, 0x00,
	/* 0364 */ 0x00, 0x00,
	/* 0366 */ 0x06, 0x00, //  [0x1b] len=6 value="action"
	/* 0368 */ 0x61, 0x00, 0x63, 0x00,
	/* 036c */ 0x74, 0x00, 0x69, 0x00,
	/* 0370 */ 0x6f, 0x00, 0x6e, 0x00,
	/* 0374 */ 0x00, 0x00,
	/* 0376 */ 0x1a, 0x00, //  [0x1c] len=26 value="android.in..."
	/* 0378 */ 0x61, 0x00, 0x6e, 0x00,
	/* 037c */ 0x64, 0x00, 0x72, 0x00,
	/* 0380 */ 0x6f, 0x00, 0x69, 0x00,
	/* 0384 */ 0x64, 0x00, 0x2e, 0x00,
	/* 0388 */ 0x69, 0x00, 0x6e, 0x00,
	/* 038c */ 0x74, 0x00, 0x65, 0x00,
	/* 0390 */ 0x6e, 0x00, 0x74, 0x00,
	/* 0394 */ 0x2e, 0x00, 0x61, 0x00,
	/* 0398 */ 0x63, 0x00, 0x74, 0x00,
	/* 039c */ 0x69, 0x00, 0x6f, 0x00,
	/* 03a0 */ 0x6e, 0x00, 0x2e, 0x00,
	/* 03a4 */ 0x4d, 0x00, 0x41, 0x00,
	/* 03a8 */ 0x49, 0x00, 0x4e, 0x00,
	/* 03ac */ 0x00, 0x00,
	/* 03ae */ 0x08, 0x00, //  [0x1d] len=8 value="category"
	/* 03b0 */ 0x63, 0x00, 0x61, 0x00,
	/* 03b4 */ 0x74, 0x00, 0x65, 0x00,
	/* 03b8 */ 0x67, 0x00, 0x6f, 0x00,
	/* 03bc */ 0x72, 0x00, 0x79, 0x00,
	/* 03c0 */ 0x00, 0x00,
	/* 03c2 */ 0x20, 0x00, //  [0x1e] len=32 value="android.in..."
	/* 03c4 */ 0x61, 0x00, 0x6e, 0x00,
	/* 03c8 */ 0x64, 0x00, 0x72, 0x00,
	/* 03cc */ 0x6f, 0x00, 0x69, 0x00,
	/* 03d0 */ 0x64, 0x00, 0x2e, 0x00,
	/* 03d4 */ 0x69, 0x00, 0x6e, 0x00,
	/* 03d8 */ 0x74, 0x00, 0x65, 0x00,
	/* 03dc */ 0x6e, 0x00, 0x74, 0x00,
	/* 03e0 */ 0x2e, 0x00, 0x63, 0x00,
	/* 03e4 */ 0x61, 0x00, 0x74, 0x00,
	/* 03e8 */ 0x65, 0x00, 0x67, 0x00,
	/* 03ec */ 0x6f, 0x00, 0x72, 0x00,
	/* 03f0 */ 0x79, 0x00, 0x2e, 0x00,
	/* 03f4 */ 0x4c, 0x00, 0x41, 0x00,
	/* 03f8 */ 0x55, 0x00, 0x4e, 0x00,
	/* 03fc */ 0x43, 0x00, 0x48, 0x00,
	/* 0400 */ 0x45, 0x00, 0x52, 0x00,
	/* 0404 */ 0x00, 0x00,
	/* 0406 */ 0x00, 0x00,
	// End of STRING_POOL.

	/* 0408 */ 0x80, 0x01, 0x08, 0x00, // chunk header XML_RESOURCE_MAP
	/* 040c */ 0x2c, 0x00, 0x00, 0x00, // chunk size 44
	/* 0410 */ 0x1b, 0x02, 0x01, 0x01, // 0x0101021b = versionCode
	/* 0414 */ 0x1c, 0x02, 0x01, 0x01, // 0x0101021c = versionName
	/* 0418 */ 0x0c, 0x02, 0x01, 0x01, // 0x0101020c = minSdkVersion
	/* 041c */ 0x01, 0x00, 0x01, 0x01, // 0x01010001 = label
	/* 0420 */ 0x0c, 0x00, 0x01, 0x01, // 0x0101000c = hasCode
	/* 0424 */ 0x0f, 0x00, 0x01, 0x01, // 0x0101000f = debuggable
	/* 0428 */ 0x03, 0x00, 0x01, 0x01, // 0x01010003 = name
	/* 042c */ 0x1f, 0x00, 0x01, 0x01, // 0x0101001f = configChanges
	/* 0430 */ 0x24, 0x00, 0x01, 0x01, // 0x01010024 = value

	/* 0434 */ 0x00, 0x01, 0x10, 0x00, // chunk header XML_START_NAMESPACE
	/* 0438 */ 0x18, 0x00, 0x00, 0x00, // chunk size 24
	/* 043c */ 0x07, 0x00, 0x00, 0x00, // line number
	/* 0440 */ 0xff, 0xff, 0xff, 0xff, // comment string reference
	/* 0444 */ 0x09, 0x00, 0x00, 0x00, // prefix [0x09]="android"
	/* 0448 */ 0x0a, 0x00, 0x00, 0x00, // url [0x0a]="http://schemas..."

	// Start XML_START_ELEMENT
	/* 044c */ 0x02, 0x01, 0x10, 0x00, // chunk header XML_START_ELEMENT
	/* 0450 */ 0x60, 0x00, 0x00, 0x00, // chunk size 96
	/* 0454 */ 0x07, 0x00, 0x00, 0x00, // line number
	/* 0458 */ 0xff, 0xff, 0xff, 0xff, // comment ref
	/* 045c */ 0xff, 0xff, 0xff, 0xff, // ns (start ResXMLTree_attrExt)
	/* 0460 */ 0x0d, 0x00, 0x00, 0x00, // name [0x0d]="manifest"
	/* 0464 */ 0x14, 0x00, // attribute start
	/* 0466 */ 0x14, 0x00, // attribute size
	/* 0468 */ 0x03, 0x00, // attribute count
	/* 046a */ 0x00, 0x00, // ID index    (1-based, 0 means none)
	/* 046c */ 0x00, 0x00, // class index (1-based, 0 means none)
	/* 046e */ 0x00, 0x00, // style index (1-based, 0 means none)
	// ResXMLTree_attribute[0]
	/* 0470 */ 0x0a, 0x00, 0x00, 0x00, // ns [0x0a]="http://schemas..."
	/* 0474 */ 0x00, 0x00, 0x00, 0x00, // name [0x00]=="versionCode"
	/* 0478 */ 0xff, 0xff, 0xff, 0xff, // rawValue
	/* 047c */ 0x08, 0x00, // Res_value size
	/* 047e */ 0x00, // Res_value padding
	/* 047f */ 0x10, // Res_value dataType (INT_DEC)
	/* 0480 */ 0x01, 0x00, 0x00, 0x00, // Res_value data
	// ResXMLTree_attribute[1]
	/* 0484 */ 0x0a, 0x00, 0x00, 0x00, // ns [0x0a]="http://schemas..."
	/* 0488 */ 0x01, 0x00, 0x00, 0x00, // name [0x01]="versionName"
	/* 048c */ 0x0f, 0x00, 0x00, 0x00, // rawValue
	/* 0490 */ 0x08, 0x00, // Res_value size
	/* 0492 */ 0x00, // Res_value padding
	/* 0493 */ 0x03, // Res_value dataType (STRING)
	/* 0494 */ 0x0f, 0x00, 0x00, 0x00, // Res_value data [0x0f]="1.0"
	// ResXMLTree_attribute[2]
	/* 0498 */ 0xff, 0xff, 0xff, 0xff, // ns none
	/* 049c */ 0x0c, 0x00, 0x00, 0x00, // name [0x0c]="package"
	/* 04a0 */ 0x0e, 0x00, 0x00, 0x00, // rawValue
	/* 04a4 */ 0x08, 0x00, // Res_value size
	/* 04a6 */ 0x00, // Res_value padding
	/* 04a7 */ 0x03, // Res_value dataType (STRING)
	/* 04a8 */ 0x0e, 0x00, 0x00, 0x00, // Res_value data [0x0e]="com.zentus..."
	// End XML_START_ELEMENT

	// Start XML_START_ELEMENT
	/* 04ac */ 0x02, 0x01, 0x10, 0x00, // chunk header XML_START_ELEMENT
	/* 04b0 */ 0x38, 0x00, 0x00, 0x00, // chunk size 56
	/* 04b4 */ 0x0d, 0x00, 0x00, 0x00, // line number
	/* 04b8 */ 0xff, 0xff, 0xff, 0xff, // comment
	/* 04bc */ 0xff, 0xff, 0xff, 0xff, // ns
	/* 04c0 */ 0x10, 0x00, 0x00, 0x00, // name [0x10]="uses-sdk"
	/* 04c4 */ 0x14, 0x00, 0x14, 0x00,
	/* 04c8 */ 0x01, 0x00, 0x00, 0x00,
	/* 04cc */ 0x00, 0x00, 0x00, 0x00,
	/* 04d0 */ 0x0a, 0x00, 0x00, 0x00,
	/* 04d4 */ 0x02, 0x00, 0x00, 0x00,
	/* 04d8 */ 0xff, 0xff, 0xff, 0xff,
	/* 04dc */ 0x08, 0x00, 0x00, 0x10,
	/* 04e0 */ 0x09, 0x00, 0x00, 0x00,
	// End XML_START_ELEMENT

	// Start XML_END_ELEMENT
	/* 04e4 */ 0x03, 0x01, 0x10, 0x00, // chunk header XML_END_ELEMENT
	/* 04e8 */ 0x18, 0x00, 0x00, 0x00, // chunk size 24
	/* 04ec */ 0x0d, 0x00, 0x00, 0x00, // line number
	/* 04f0 */ 0xff, 0xff, 0xff, 0xff, // comment
	/* 04f4 */ 0xff, 0xff, 0xff, 0xff, // ns
	/* 04f8 */ 0x10, 0x00, 0x00, 0x00, // name [0x10]="uses-sdk"
	// End XML_END_ELEMENT

	/* 04fc */ 0x02, 0x01, 0x10, 0x00,
	/* 0500 */ 0x60, 0x00, 0x00, 0x00,
	/* 0504 */ 0x0e, 0x00, 0x00, 0x00,
	/* 0508 */ 0xff, 0xff, 0xff, 0xff,
	/* 050c */ 0xff, 0xff, 0xff, 0xff,
	/* 0510 */ 0x11, 0x00, 0x00, 0x00,
	/* 0514 */ 0x14, 0x00, 0x14, 0x00,
	/* 0518 */ 0x03, 0x00, 0x00, 0x00,
	/* 051c */ 0x00, 0x00, 0x00, 0x00,
	/* 0520 */ 0x0a, 0x00, 0x00, 0x00,
	/* 0524 */ 0x03, 0x00, 0x00, 0x00,
	/* 0528 */ 0x12, 0x00, 0x00, 0x00,
	/* 052c */ 0x08, 0x00, 0x00, 0x03,
	/* 0530 */ 0x12, 0x00, 0x00, 0x00,
	/* 0534 */ 0x0a, 0x00, 0x00, 0x00,
	/* 0538 */ 0x04, 0x00, 0x00, 0x00,
	/* 053c */ 0xff, 0xff, 0xff, 0xff,
	/* 0540 */ 0x08, 0x00, 0x00, 0x12,
	/* 0544 */ 0x00, 0x00, 0x00, 0x00,
	/* 0548 */ 0x0a, 0x00, 0x00, 0x00,
	/* 054c */ 0x05, 0x00, 0x00, 0x00,
	/* 0550 */ 0xff, 0xff, 0xff, 0xff,
	/* 0554 */ 0x08, 0x00, 0x00, 0x12,
	/* 0558 */ 0xff, 0xff, 0xff, 0xff,
	/* 055c */ 0x02, 0x01, 0x10, 0x00,
	/* 0560 */ 0x60, 0x00, 0x00, 0x00,
	/* 0564 */ 0x0f, 0x00, 0x00, 0x00,
	/* 0568 */ 0xff, 0xff, 0xff, 0xff,
	/* 056c */ 0xff, 0xff, 0xff, 0xff,
	/* 0570 */ 0x13, 0x00, 0x00, 0x00,
	/* 0574 */ 0x14, 0x00, 0x14, 0x00,
	/* 0578 */ 0x03, 0x00, 0x00, 0x00,
	/* 057c */ 0x00, 0x00, 0x00, 0x00,
	/* 0580 */ 0x0a, 0x00, 0x00, 0x00,
	/* 0584 */ 0x03, 0x00, 0x00, 0x00,
	/* 0588 */ 0x15, 0x00, 0x00, 0x00,
	/* 058c */ 0x08, 0x00, 0x00, 0x03,
	/* 0590 */ 0x15, 0x00, 0x00, 0x00,
	/* 0594 */ 0x0a, 0x00, 0x00, 0x00,
	/* 0598 */ 0x06, 0x00, 0x00, 0x00,
	/* 059c */ 0x14, 0x00, 0x00, 0x00,
	/* 05a0 */ 0x08, 0x00, 0x00, 0x03,
	/* 05a4 */ 0x14, 0x00, 0x00, 0x00,
	/* 05a8 */ 0x0a, 0x00, 0x00, 0x00,
	/* 05ac */ 0x07, 0x00, 0x00, 0x00,
	/* 05b0 */ 0xff, 0xff, 0xff, 0xff,
	/* 05b4 */ 0x08, 0x00, 0x00, 0x11,
	/* 05b8 */ 0xa0, 0x00, 0x00, 0x00,
	/* 05bc */ 0x02, 0x01, 0x10, 0x00,
	/* 05c0 */ 0x4c, 0x00, 0x00, 0x00,
	/* 05c4 */ 0x12, 0x00, 0x00, 0x00,
	/* 05c8 */ 0xff, 0xff, 0xff, 0xff,
	/* 05cc */ 0xff, 0xff, 0xff, 0xff,
	/* 05d0 */ 0x16, 0x00, 0x00, 0x00,
	/* 05d4 */ 0x14, 0x00, 0x14, 0x00,
	/* 05d8 */ 0x02, 0x00, 0x00, 0x00,
	/* 05dc */ 0x00, 0x00, 0x00, 0x00,
	/* 05e0 */ 0x0a, 0x00, 0x00, 0x00,
	/* 05e4 */ 0x06, 0x00, 0x00, 0x00,
	/* 05e8 */ 0x17, 0x00, 0x00, 0x00,
	/* 05ec */ 0x08, 0x00, 0x00, 0x03,
	/* 05f0 */ 0x17, 0x00, 0x00, 0x00,
	/* 05f4 */ 0x0a, 0x00, 0x00, 0x00,
	/* 05f8 */ 0x08, 0x00, 0x00, 0x00,
	/* 05fc */ 0x18, 0x00, 0x00, 0x00,
	/* 0600 */ 0x08, 0x00, 0x00, 0x03,
	/* 0604 */ 0x18, 0x00, 0x00, 0x00,
	/* 0608 */ 0x03, 0x01, 0x10, 0x00,
	/* 060c */ 0x18, 0x00, 0x00, 0x00,
	/* 0610 */ 0x12, 0x00, 0x00, 0x00,
	/* 0614 */ 0xff, 0xff, 0xff, 0xff,
	/* 0618 */ 0xff, 0xff, 0xff, 0xff,
	/* 061c */ 0x16, 0x00, 0x00, 0x00,
	/* 0620 */ 0x02, 0x01, 0x10, 0x00,
	/* 0624 */ 0x24, 0x00, 0x00, 0x00,
	/* 0628 */ 0x13, 0x00, 0x00, 0x00,
	/* 062c */ 0xff, 0xff, 0xff, 0xff,
	/* 0630 */ 0xff, 0xff, 0xff, 0xff,
	/* 0634 */ 0x19, 0x00, 0x00, 0x00,
	/* 0638 */ 0x14, 0x00, 0x14, 0x00,
	/* 063c */ 0x00, 0x00, 0x00, 0x00,
	/* 0640 */ 0x00, 0x00, 0x00, 0x00,
	/* 0644 */ 0x04, 0x01, 0x10, 0x00,
	/* 0648 */ 0x1c, 0x00, 0x00, 0x00,
	/* 064c */ 0x13, 0x00, 0x00, 0x00,
	/* 0650 */ 0xff, 0xff, 0xff, 0xff,
	/* 0654 */ 0x1a, 0x00, 0x00, 0x00,
	/* 0658 */ 0x08, 0x00, 0x00, 0x00,
	/* 065c */ 0x00, 0x00, 0x00, 0x00,
	/* 0660 */ 0x02, 0x01, 0x10, 0x00,
	/* 0664 */ 0x38, 0x00, 0x00, 0x00,
	/* 0668 */ 0x15, 0x00, 0x00, 0x00,
	/* 066c */ 0xff, 0xff, 0xff, 0xff,
	/* 0670 */ 0xff, 0xff, 0xff, 0xff,
	/* 0674 */ 0x1b, 0x00, 0x00, 0x00,
	/* 0678 */ 0x14, 0x00, 0x14, 0x00,
	/* 067c */ 0x01, 0x00, 0x00, 0x00,
	/* 0680 */ 0x00, 0x00, 0x00, 0x00,
	/* 0684 */ 0x0a, 0x00, 0x00, 0x00,
	/* 0688 */ 0x06, 0x00, 0x00, 0x00,
	/* 068c */ 0x1c, 0x00, 0x00, 0x00,
	/* 0690 */ 0x08, 0x00, 0x00, 0x03,
	/* 0694 */ 0x1c, 0x00, 0x00, 0x00,
	/* 0698 */ 0x03, 0x01, 0x10, 0x00,
	/* 069c */ 0x18, 0x00, 0x00, 0x00,
	/* 06a0 */ 0x15, 0x00, 0x00, 0x00,
	/* 06a4 */ 0xff, 0xff, 0xff, 0xff,
	/* 06a8 */ 0xff, 0xff, 0xff, 0xff,
	/* 06ac */ 0x1b, 0x00, 0x00, 0x00,
	/* 06b0 */ 0x02, 0x01, 0x10, 0x00,
	/* 06b4 */ 0x38, 0x00, 0x00, 0x00,
	/* 06b8 */ 0x16, 0x00, 0x00, 0x00,
	/* 06bc */ 0xff, 0xff, 0xff, 0xff,
	/* 06c0 */ 0xff, 0xff, 0xff, 0xff,
	/* 06c4 */ 0x1d, 0x00, 0x00, 0x00,
	/* 06c8 */ 0x14, 0x00, 0x14, 0x00,
	/* 06cc */ 0x01, 0x00, 0x00, 0x00,
	/* 06d0 */ 0x00, 0x00, 0x00, 0x00,
	/* 06d4 */ 0x0a, 0x00, 0x00, 0x00,
	/* 06d8 */ 0x06, 0x00, 0x00, 0x00,
	/* 06dc */ 0x1e, 0x00, 0x00, 0x00,
	/* 06e0 */ 0x08, 0x00, 0x00, 0x03,
	/* 06e4 */ 0x1e, 0x00, 0x00, 0x00,
	/* 06e8 */ 0x03, 0x01, 0x10, 0x00,
	/* 06ec */ 0x18, 0x00, 0x00, 0x00,
	/* 06f0 */ 0x16, 0x00, 0x00, 0x00,
	/* 06f4 */ 0xff, 0xff, 0xff, 0xff,
	/* 06f8 */ 0xff, 0xff, 0xff, 0xff,
	/* 06fc */ 0x1d, 0x00, 0x00, 0x00,
	/* 0700 */ 0x03, 0x01, 0x10, 0x00,
	/* 0704 */ 0x18, 0x00, 0x00, 0x00,
	/* 0708 */ 0x17, 0x00, 0x00, 0x00,
	/* 070c */ 0xff, 0xff, 0xff, 0xff,
	/* 0710 */ 0xff, 0xff, 0xff, 0xff,
	/* 0714 */ 0x19, 0x00, 0x00, 0x00,
	/* 0718 */ 0x03, 0x01, 0x10, 0x00,
	/* 071c */ 0x18, 0x00, 0x00, 0x00,
	/* 0720 */ 0x18, 0x00, 0x00, 0x00,
	/* 0724 */ 0xff, 0xff, 0xff, 0xff,
	/* 0728 */ 0xff, 0xff, 0xff, 0xff,
	/* 072c */ 0x13, 0x00, 0x00, 0x00,
	/* 0730 */ 0x03, 0x01, 0x10, 0x00,
	/* 0734 */ 0x18, 0x00, 0x00, 0x00,
	/* 0738 */ 0x19, 0x00, 0x00, 0x00,
	/* 073c */ 0xff, 0xff, 0xff, 0xff,
	/* 0740 */ 0xff, 0xff, 0xff, 0xff,
	/* 0744 */ 0x11, 0x00, 0x00, 0x00,
	/* 0748 */ 0x03, 0x01, 0x10, 0x00,
	/* 074c */ 0x18, 0x00, 0x00, 0x00,
	/* 0750 */ 0x1a, 0x00, 0x00, 0x00,
	/* 0754 */ 0xff, 0xff, 0xff, 0xff,
	/* 0758 */ 0xff, 0xff, 0xff, 0xff,
	/* 075c */ 0x0d, 0x00, 0x00, 0x00,
	/* 0760 */ 0x01, 0x01, 0x10, 0x00,
	/* 0764 */ 0x18, 0x00, 0x00, 0x00,
	/* 0768 */ 0x1a, 0x00, 0x00, 0x00,
	/* 076c */ 0xff, 0xff, 0xff, 0xff,
	/* 0770 */ 0x09, 0x00, 0x00, 0x00,
	/* 0774 */ 0x0a, 0x00, 0x00, 0x00,
}

const input = `<?xml version="1.0" encoding="utf-8"?>
<!--
Copyright 2014 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->
<manifest
	xmlns:android="http://schemas.android.com/apk/res/android"
	package="com.zentus.balloon"
	android:versionCode="1"
	android:versionName="1.0">

	<uses-sdk android:minSdkVersion="9" />
	<application android:label="Balloon世界" android:hasCode="false">
	<activity android:name="android.app.NativeActivity"
		android:label="Balloon"
		android:configChanges="orientation|keyboardHidden">
		<meta-data android:name="android.app.lib_name" android:value="balloon" />
		<intent-filter>
			here is some text
			<action android:name="android.intent.action.MAIN" />
			<category android:name="android.intent.category.LAUNCHER" />
		</intent-filter>
	</activity>
	</application>
</manifest>`
