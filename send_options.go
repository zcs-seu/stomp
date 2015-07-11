package stomp

import (
	"github.com/jjeffery/stomp/frame"
)

// SendOpt contains options for for the Conn.Send function.
var SendOpt struct {
	// Receipt specifies that the client should request acknowledgement
	// from the server before the send operation successfully completes.
	Receipt func(*Frame) error

	// NoContentType specifies that the SEND frame should not include
	// a content-length header entry. By default the content-length header
	// entry is always included, but some message brokers assign special
	// meaning to STOMP frames that do not contain a content-length
	// header entry. (In particular ActiveMQ interprets STOMP frames
	// with no content-length as being a text message)
	NoContentType func(*Frame) error

	// Header provides the opportunity to include custom header entries
	// in the SEND frame that the client sends to the server.
	Header func(header *Header) func(*Frame) error
}

func init() {
	SendOpt.Receipt = func(f *Frame) error {
		if f.Command != frame.SEND {
			return ErrInvalidCommand
		}
		id := allocateId()
		f.Set(frame.Receipt, id)
		return nil
	}

	SendOpt.Header = func(header *Header) func(*Frame) error {
		return func(f *Frame) error {
			if f.Command != frame.SEND {
				return ErrInvalidCommand
			}
			f.AddHeader(header)
			return nil
		}
	}
}
