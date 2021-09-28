package main

type availableEncoder struct {
	Names       []string
	Description string
	Generator   func() EncoderDecoder
}

var generators map[string]func() EncoderDecoder

func getEncoderDecoder(enc string) EncoderDecoder {
	if gen, ok := generators[enc]; ok {
		return gen()
	}
	return nil
}
