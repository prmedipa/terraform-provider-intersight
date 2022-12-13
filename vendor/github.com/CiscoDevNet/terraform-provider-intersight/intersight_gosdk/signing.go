/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9661
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"os"
	"strings"
	"time"
)

const (
	// Constants for HTTP signature parameters.
	// The '(request-target)' parameter concatenates the lowercased :method, an
	// ASCII space, and the :path pseudo-headers.
	HttpSignatureParameterRequestTarget string = "(request-target)"
	// The '(created)' parameter expresses when the signature was
	// created.  The value MUST be a Unix timestamp integer value.
	HttpSignatureParameterCreated string = "(created)"
	// The '(expires)' parameter expresses when the signature ceases to
	// be valid.  The value MUST be a Unix timestamp integer value.
	HttpSignatureParameterExpires string = "(expires)"
)

const (
	// Constants for HTTP headers.
	// The 'Host' header, as defined in RFC 2616, section 14.23.
	HttpHeaderHost string = "Host"
	// The 'Date' header.
	HttpHeaderDate string = "Date"
	// The digest header, as defined in RFC 3230, section 4.3.2.
	HttpHeaderDigest string = "Digest"
	// The HTTP Authorization header, as defined in RFC 7235, section 4.2.
	HttpHeaderAuthorization string = "Authorization"
)

const (
	// Specifies the Digital Signature Algorithm is derived from metadata
	// associated with 'keyId'. Supported DSA algorithms are RSASSA-PKCS1-v1_5,
	// RSASSA-PSS, and ECDSA.
	// The hash is SHA-512.
	// This is the default value.
	HttpSigningSchemeHs2019 string = "hs2019"
	// Use RSASSA-PKCS1-v1_5 with SHA-512 hash. Deprecated.
	HttpSigningSchemeRsaSha512 string = "rsa-sha512"
	// Use RSASSA-PKCS1-v1_5 with SHA-256 hash. Deprecated.
	HttpSigningSchemeRsaSha256 string = "rsa-sha256"

	// RFC 8017 section 7.2
	// Calculate the message signature using RSASSA-PKCS1-V1_5-SIGN from RSA PKCS#1 v1.5.
	// PKCSV1_5 is deterministic. The same message and key will produce an identical
	// signature value each time.
	HttpSigningAlgorithmRsaPKCS1v15 string = "RSASSA-PKCS1-v1_5"
	// Calculate the message signature using probabilistic signature scheme RSASSA-PSS.
	// PSS is randomized and will produce a different signature value each time.
	HttpSigningAlgorithmRsaPSS string = "RSASSA-PSS"
)

var supportedSigningSchemes = map[string]bool{
	HttpSigningSchemeHs2019:    true,
	HttpSigningSchemeRsaSha512: true,
	HttpSigningSchemeRsaSha256: true,
}

// HttpSignatureAuth provides HTTP signature authentication to a request passed
// via context using ContextHttpSignatureAuth.
// An 'Authorization' header is calculated by creating a hash of select headers,
// and optionally the body of the HTTP request, then signing the hash value using
// a private key which is available to the client.
//
// SignedHeaders specifies the list of HTTP headers that are included when generating
// the message signature.
// The two special signature headers '(request-target)' and '(created)' SHOULD be
// included in SignedHeaders.
// The '(created)' header expresses when the signature was created.
// The '(request-target)' header is a concatenation of the lowercased :method, an
// ASCII space, and the :path pseudo-headers.
//
// For example, SignedHeaders can be set to:
//
//	(request-target) (created) date host digest
//
// When SignedHeaders is not specified, the client defaults to a single value, '(created)',
// in the list of HTTP headers.
// When SignedHeaders contains the 'Digest' value, the client performs the following operations:
// 1. Calculate a digest of request body, as specified in RFC3230, section 4.3.2.
// 2. Set the 'Digest' header in the request body.
// 3. Include the 'Digest' header and value in the HTTP signature.
type HttpSignatureAuth struct {
	KeyId          string // A key identifier.
	PrivateKeyPath string // The path to the private key.
	Passphrase     string // The passphrase to decrypt the private key, if the key is encrypted.
	SigningScheme  string // The signature scheme, when signing HTTP requests. Supported value is 'hs2019'.
	// The signature algorithm, when signing HTTP requests.
	// Supported values are RSASSA-PKCS1-v1_5, RSASSA-PSS.
	SigningAlgorithm string
	SignedHeaders    []string // A list of HTTP headers included when generating the signature for the message.
	// SignatureMaxValidity specifies the maximum duration of the signature validity.
	// The value is used to set the '(expires)' signature parameter in the HTTP request.
	// '(expires)' is set to '(created)' plus the value of the SignatureMaxValidity field.
	// To specify the '(expires)' signature parameter, set 'SignatureMaxValidity' and add '(expires)' to 'SignedHeaders'.
	SignatureMaxValidity time.Duration
	privateKey           crypto.PrivateKey // The private key used to sign HTTP requests.
}

// SetPrivateKey accepts a private key string and sets it.
func (h *HttpSignatureAuth) SetPrivateKey(privateKey string) error {
	return h.parsePrivateKey([]byte(privateKey))
}

// ContextWithValue validates the HttpSignatureAuth configuration parameters and returns a context
// suitable for HTTP signature. An error is returned if the HttpSignatureAuth configuration parameters
// are invalid.
func (h *HttpSignatureAuth) ContextWithValue(ctx context.Context) (context.Context, error) {
	if h.KeyId == "" {
		return nil, fmt.Errorf("Key ID must be specified")
	}
	if h.PrivateKeyPath == "" && h.privateKey == nil {
		return nil, fmt.Errorf("Private key path must be specified")
	}
	if _, ok := supportedSigningSchemes[h.SigningScheme]; !ok {
		return nil, fmt.Errorf("Invalid signing scheme: '%v'", h.SigningScheme)
	}
	m := make(map[string]bool)
	for _, h := range h.SignedHeaders {
		if strings.EqualFold(h, HttpHeaderAuthorization) {
			return nil, fmt.Errorf("Signed headers cannot include the 'Authorization' header")
		}
		m[h] = true
	}
	if len(m) != len(h.SignedHeaders) {
		return nil, fmt.Errorf("List of signed headers cannot have duplicate names")
	}
	if h.SignatureMaxValidity < 0 {
		return nil, fmt.Errorf("Signature max validity must be a positive value")
	}
	if err := h.loadPrivateKey(); err != nil {
		return nil, err
	}
	return context.WithValue(ctx, ContextHttpSignatureAuth, *h), nil
}

// GetPublicKey returns the public key associated with this HTTP signature configuration.
func (h *HttpSignatureAuth) GetPublicKey() (crypto.PublicKey, error) {
	if h.privateKey == nil {
		if err := h.loadPrivateKey(); err != nil {
			return nil, err
		}
	}
	switch key := h.privateKey.(type) {
	case *rsa.PrivateKey:
		return key.Public(), nil
	case *ecdsa.PrivateKey:
		return key.Public(), nil
	default:
		// Do not change '%T' to anything else such as '%v'!
		// The value of the private key must not be returned.
		return nil, fmt.Errorf("Unsupported key: %T", h.privateKey)
	}
}

// loadPrivateKey reads the private key from the file specified in the HttpSignatureAuth.
// The key is loaded only when privateKey is not already set.
func (h *HttpSignatureAuth) loadPrivateKey() (err error) {
	if h.privateKey != nil {
		return nil
	}
	var file *os.File
	file, err = os.Open(h.PrivateKeyPath)
	if err != nil {
		return fmt.Errorf("Cannot load private key '%s'. Error: %v", h.PrivateKeyPath, err)
	}
	defer func() {
		err = file.Close()
	}()
	var priv []byte
	priv, err = ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return h.parsePrivateKey(priv)
}

// parsePrivateKey decodes privateKey byte array to crypto.PrivateKey type.
func (h *HttpSignatureAuth) parsePrivateKey(priv []byte) error {
	pemBlock, _ := pem.Decode(priv)
	if pemBlock == nil {
		// No PEM data has been found.
		return fmt.Errorf("File '%s' does not contain PEM data", h.PrivateKeyPath)
	}
	var privKey []byte
	var err error
	if x509.IsEncryptedPEMBlock(pemBlock) {
		// The PEM data is encrypted.
		privKey, err = x509.DecryptPEMBlock(pemBlock, []byte(h.Passphrase))
		if err != nil {
			// Failed to decrypt PEM block. Because of deficiencies in the encrypted-PEM format,
			// it's not always possible to detect an incorrect password.
			return err
		}
	} else {
		privKey = pemBlock.Bytes
	}
	switch pemBlock.Type {
	case "RSA PRIVATE KEY":
		if h.privateKey, err = x509.ParsePKCS1PrivateKey(privKey); err != nil {
			return err
		}
	case "EC PRIVATE KEY", "PRIVATE KEY":
		// https://tools.ietf.org/html/rfc5915 section 4.
		if h.privateKey, err = x509.ParsePKCS8PrivateKey(privKey); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Key '%s' is not supported", pemBlock.Type)
	}
	return nil
}

// SignRequest signs the request using HTTP signature.
// See https://datatracker.ietf.org/doc/draft-cavage-http-signatures/
//
// Do not add, remove or change headers that are included in the SignedHeaders
// after SignRequest has been invoked; this is because the header values are
// included in the signature. Any subsequent alteration will cause a signature
// verification failure.
// If there are multiple instances of the same header field, all
// header field values associated with the header field MUST be
// concatenated, separated by a ASCII comma and an ASCII space
// ', ', and used in the order in which they will appear in the
// transmitted HTTP message.
func SignRequest(
	ctx context.Context,
	r *http.Request,
	auth HttpSignatureAuth) error {

	if auth.privateKey == nil {
		return fmt.Errorf("Private key is not set")
	}
	now := time.Now()
	date := now.UTC().Format(http.TimeFormat)
	// The 'created' field expresses when the signature was created.
	// The value MUST be a Unix timestamp integer value. See 'HTTP signature' section 2.1.4.
	created := now.Unix()

	var h crypto.Hash
	var err error
	var prefix string
	var expiresUnix float64

	if auth.SignatureMaxValidity < 0 {
		return fmt.Errorf("Signature validity must be a positive value")
	}
	if auth.SignatureMaxValidity > 0 {
		e := now.Add(auth.SignatureMaxValidity)
		expiresUnix = float64(e.Unix()) + float64(e.Nanosecond())/float64(time.Second)
	}
	// Determine the cryptographic hash to be used for the signature and the body digest.
	switch auth.SigningScheme {
	case HttpSigningSchemeRsaSha512, HttpSigningSchemeHs2019:
		h = crypto.SHA512
		prefix = "SHA-512="
	case HttpSigningSchemeRsaSha256:
		// This is deprecated and should no longer be used.
		h = crypto.SHA256
		prefix = "SHA-256="
	default:
		return fmt.Errorf("Unsupported signature scheme: %v", auth.SigningScheme)
	}
	if !h.Available() {
		return fmt.Errorf("Hash '%v' is not available", h)
	}

	// Build the "(request-target)" signature header.
	var sb bytes.Buffer
	fmt.Fprintf(&sb, "%s %s", strings.ToLower(r.Method), r.URL.EscapedPath())
	if r.URL.RawQuery != "" {
		// The ":path" pseudo-header field includes the path and query parts
		// of the target URI (the "path-absolute" production and optionally a
		// '?' character followed by the "query" production (see Sections 3.3
		// and 3.4 of [RFC3986]
		fmt.Fprintf(&sb, "?%s", r.URL.RawQuery)
	}
	requestTarget := sb.String()
	sb.Reset()

	// Build the string to be signed.
	signedHeaders := auth.SignedHeaders
	if len(signedHeaders) == 0 {
		signedHeaders = []string{HttpSignatureParameterCreated}
	}
	// Validate the list of signed headers has no duplicates.
	m := make(map[string]bool)
	for _, h := range signedHeaders {
		m[h] = true
	}
	if len(m) != len(signedHeaders) {
		return fmt.Errorf("List of signed headers must not have any duplicates")
	}
	hasCreatedParameter := false
	hasExpiresParameter := false
	for i, header := range signedHeaders {
		header = strings.ToLower(header)
		var value string
		switch header {
		case strings.ToLower(HttpHeaderAuthorization):
			return fmt.Errorf("Cannot include the 'Authorization' header as a signed header.")
		case HttpSignatureParameterRequestTarget:
			value = requestTarget
		case HttpSignatureParameterCreated:
			value = fmt.Sprintf("%d", created)
			hasCreatedParameter = true
		case HttpSignatureParameterExpires:
			if auth.SignatureMaxValidity.Nanoseconds() == 0 {
				return fmt.Errorf("Cannot set '(expires)' signature parameter. SignatureMaxValidity is not configured.")
			}
			value = fmt.Sprintf("%.3f", expiresUnix)
			hasExpiresParameter = true
		case "date":
			value = date
			r.Header.Set(HttpHeaderDate, date)
		case "digest":
			// Calculate the digest of the HTTP request body.
			// Calculate body digest per RFC 3230 section 4.3.2
			bodyHash := h.New()
			if r.Body != nil {
				// Make a copy of the body io.Reader so that we can read the body to calculate the hash,
				// then one more time when marshaling the request.
				var body io.Reader
				body, err = r.GetBody()
				if err != nil {
					return err
				}
				if _, err = io.Copy(bodyHash, body); err != nil {
					return err
				}
			}
			d := bodyHash.Sum(nil)
			value = prefix + base64.StdEncoding.EncodeToString(d)
			r.Header.Set(HttpHeaderDigest, value)
		case "host":
			value = r.Host
			r.Header.Set(HttpHeaderHost, r.Host)
		default:
			var ok bool
			var v []string
			canonicalHeader := textproto.CanonicalMIMEHeaderKey(header)
			if v, ok = r.Header[canonicalHeader]; !ok {
				// If a header specified in the headers parameter cannot be matched with
				// a provided header in the message, the implementation MUST produce an error.
				return fmt.Errorf("Header '%s' does not exist in the request", canonicalHeader)
			}
			// If there are multiple instances of the same header field, all
			// header field values associated with the header field MUST be
			// concatenated, separated by a ASCII comma and an ASCII space
			// `, `, and used in the order in which they will appear in the
			// transmitted HTTP message.
			value = strings.Join(v, ", ")
		}
		if i > 0 {
			fmt.Fprintf(&sb, "\n")
		}
		fmt.Fprintf(&sb, "%s: %s", header, value)
	}
	if expiresUnix != 0 && !hasExpiresParameter {
		return fmt.Errorf("SignatureMaxValidity is specified, but '(expired)' parameter is not present")
	}
	msg := []byte(sb.String())
	msgHash := h.New()
	if _, err = msgHash.Write(msg); err != nil {
		return err
	}
	d := msgHash.Sum(nil)

	var signature []byte
	switch key := auth.privateKey.(type) {
	case *rsa.PrivateKey:
		switch auth.SigningAlgorithm {
		case HttpSigningAlgorithmRsaPKCS1v15:
			signature, err = rsa.SignPKCS1v15(rand.Reader, key, h, d)
		case "", HttpSigningAlgorithmRsaPSS:
			signature, err = rsa.SignPSS(rand.Reader, key, h, d, nil)
		default:
			return fmt.Errorf("Unsupported signing algorithm: '%s'", auth.SigningAlgorithm)
		}
	case *ecdsa.PrivateKey:
		signature, err = key.Sign(rand.Reader, d, h)
	case ed25519.PrivateKey: // requires go 1.13
		signature, err = key.Sign(rand.Reader, msg, crypto.Hash(0))
	default:
		return fmt.Errorf("Unsupported private key")
	}
	if err != nil {
		return err
	}

	sb.Reset()
	for i, header := range signedHeaders {
		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(strings.ToLower(header))
	}
	headers_list := sb.String()
	sb.Reset()
	fmt.Fprintf(&sb, `Signature keyId="%s",algorithm="%s",`, auth.KeyId, auth.SigningScheme)
	if hasCreatedParameter {
		fmt.Fprintf(&sb, "created=%d,", created)
	}
	if hasExpiresParameter {
		fmt.Fprintf(&sb, "expires=%.3f,", expiresUnix)
	}
	fmt.Fprintf(&sb, `headers="%s",signature="%s"`, headers_list, base64.StdEncoding.EncodeToString(signature))
	authStr := sb.String()
	r.Header.Set(HttpHeaderAuthorization, authStr)
	return nil
}
