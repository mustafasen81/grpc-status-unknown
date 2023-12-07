# gRPC Status of Unknown
This repository consists of a test case that reveals the gRPC status of `Unknown` in the server stream.

If the underlying connection to the client is closed before sending the first message in a server stream, it gets a grpc error having the status of `Unknown`, as `rpc error:  code = Unknown desc = connection error: desc = "transport is closing"`. After an investigation, I found that if the connection is closed at the very beginning of the stream before the server writes the headers. When it writes the headers, it gets the `ConnectionError` and tries to convert it to a grpc status error by `status.Convert(err).Err()`. Since `ConnectionError` is not a grpc status error it converts the error to grpc status of `Unknown`.

Code snippet that runs this as follows:
```
// WriteHeader sends the header metadata md back to the client.
func (t *http2Server) WriteHeader(s *Stream, md metadata.MD) error {
	s.hdrMu.Lock()
	defer s.hdrMu.Unlock()
	if s.getState() == streamDone {
		return t.streamContextErr(s)
	}

	if s.updateHeaderSent() {
		return ErrIllegalHeaderWrite
	}

	if md.Len() > 0 {
		if s.header.Len() > 0 {
			s.header = metadata.Join(s.header, md)
		} else {
			s.header = md
		}
	}
	if err := t.writeHeaderLocked(s); err != nil {
		return status.Convert(err).Err()
	}
	return nil
}
```
