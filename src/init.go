package main

func init() {
	generators = make(map[string]func() EncoderDecoder)
	for _, enc := range availableEncodersDecoders {
		for _, name := range enc.Names {
			generators[name] = enc.Generator
		}
	}
}
