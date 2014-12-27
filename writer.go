// Package apk provides support for writing APK archives.
//
// TODO(crawshaw): implement
//
// APK is the archival format used for Android apps. It is a ZIP archive with
// three extra files:
//
//	META-INF/MANIFEST.MF
//	META-INF/CERT.SF
//	META-INF/CERT.RSA
//
// The MANIFEST.MF comes from the Java JAR archive format. It is a list of
// files included in the archive along with a SHA1 hash, for example:
//
//	Name: lib/armeabi/libbasic.so
//	SHA1-Digest: ntLSc1eLCS2Tq1oB4Vw6jvkranw=
//
// CERT.SF is a similar manifest. For each entry in the manifest it includes
// a base64-encoded SHA1 digest RSA signed SHA1 key, for example:
//
//	Name: lib/armeabi/libbasic.so
//	SHA1-Digest: Q7NAS6uzrJr6WjePXSGT+vvmdiw=
//
// For debugging, the equivalent SHA1-Digest can be generated with OpenSSL:
//
//	echo -en "Name: lib/armeabi/libbasic.so\r\nSHA1-Digest: ntLSc1eLCS2Tq1oB4Vw6jvkranw=\r\n\r\n" | openssl sha1 -binary | openssl base64
//
// Note the \r\n nonsense.
//
// CERT.RSA is an RSA signature block made of CERT.SF. Verify it with:
//
//	openssl smime -verify -in CERT.RSA -inform DER -content CERT.SF cert.pem
//
// The APK format imposes two extra restrictions on the ZIP format. First,
// it is uncompressed. Second, each contained file is 4-byte aligned. This
// allows the Android OS to mmap contents without unpacking the archive.
package apk

// Note: to make life a little harder, Android Studio stores the RSA key used
// for signing in an Oracle Java proprietary keystore format, JKS. For example,
// the generated debug key is in ~/.android/debug.keystore, and can be
// extracted using the JDK's keytool utility:
//
//	keytool -importkeystore -srckeystore ~/.android/debug.keystore -destkeystore ~/.android/debug.p12 -deststoretype PKCS12
//
// Once in standard PKCS12, the key can be converted to PEM for use in the
// Go crypto packages:
//
//	openssl pkcs12 -in ~/.android/debug.p12 -nocerts -nodes -out ~/.android/debug.pem
//
// Fortunately for debug builds, all that matters is that the APK is signed.
// The choice of key is unimportant, so we can generate one for normal builds.
// For production builds, we can ask users to provide a PEM file.

import (
	"archive/zip"
	"crypto/rsa"
	"hash"
	"io"
)

type manifestEntry struct {
	name string
	sha1 hash.Hash
}

type Writer struct {
	w zip.Writer

	manifest []manifestEntry

	cur struct {
		path string
		w    io.Writer
		sha1 hash.Hash
	}
}

func (w *Writer) Create(name string) (io.Writer, error) {
	return nil, nil
}

func (w *Writer) Close() error {
	return nil
}

func NewWriter(w io.Writer, key *rsa.PrivateKey) *Writer {
	return nil
}
