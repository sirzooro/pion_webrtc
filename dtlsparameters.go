// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

// DTLSParameters holds information relating to DTLS configuration.
type DTLSParameters struct {
	Role         DTLSRole          `json:"role"`
	Fingerprints []DTLSFingerprint `json:"fingerprints"`

	// srtpCryptexMode is a local-only knob used by the ORTC-style API to inform
	// the DTLS transport how SRTP should be configured once the DTLS handshake
	// completes. It is not part of the DTLS signaling parameters.
	srtpCryptexMode srtpCryptexMode `json:"-"`
}

// srtpCryptexMode controls RFC 9335 RTP Header Extension Encryption ("Cryptex") usage.
type srtpCryptexMode uint8

const (
	srtpCryptexModeDisabled srtpCryptexMode = iota
	srtpCryptexModeEnabled
	srtpCryptexModeRequired
)
