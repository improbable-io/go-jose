package jose

// GetUnverifiedPayload exposes the payload without verifying it. This should be used with caution as it is possible
// the contents have been tampered with. Only use this when the payload does not need to be trusted.
func (obj JsonWebSignature) GetUnverifiedPayload() []byte {
	return obj.payload
}
