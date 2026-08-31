package viper

import (
	"sync"
)

type Encoder interface {
	Encode(v map[string]any) ([]byte, error)
}

type Decoder interface {
	Decode(b []byte, v map[string]any) error
}

type Codec interface {
	Encoder
	Decoder
}

type EncoderRegistry interface {
	Encoder(format string) (Encoder, error)
}

type DecoderRegistry interface {
	Decoder(format string) (Decoder, error)
}

type CodecRegistry interface {
	EncoderRegistry
	DecoderRegistry
}

func WithEncoderRegistry(r EncoderRegistry) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDecoderRegistry(r DecoderRegistry) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCodecRegistry(r CodecRegistry) Option { _ = "STUB: not implemented"; return *new(Option) }

type DefaultCodecRegistry struct {
	codecs map[string]Codec

	mu   sync.RWMutex
	once sync.Once
}

func NewCodecRegistry() *DefaultCodecRegistry { _ = "STUB: not implemented"; return nil }

func (r *DefaultCodecRegistry) init() {
	r.once.Do(func() {
		r.codecs = map[string]Codec{}
	})
}

func (r *DefaultCodecRegistry) RegisterCodec(format string, codec Codec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *DefaultCodecRegistry) Encoder(format string) (Encoder, error) {
	_ = "STUB: not implemented"
	return *new(Encoder), nil
}

func (r *DefaultCodecRegistry) Decoder(format string) (Decoder, error) {
	_ = "STUB: not implemented"
	return *new(Decoder), nil
}

func (r *DefaultCodecRegistry) codec(format string) (Codec, bool) {
	_ = "STUB: not implemented"
	return *new(Codec), false
}
