package main

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
)

var availableEncodersDecoders = []availableEncoder{
	{
		Names:       []string{"utf-8", "utf8"},
		Description: "UTF-8",
		Generator:   func() EncoderDecoder { return NewUnicodeEncoderDecoder(unicode.UTF8) },
	},
	{
		Names:       []string{"utf-8-bom", "utf8bom"},
		Description: "UTF-8 with BOM",
		Generator:   func() EncoderDecoder { return NewUnicodeEncoderDecoder(unicode.UTF8BOM) },
	},
	{
		Names:       []string{"utf-16-le", "utf16le"},
		Description: "UTF-16 Little Endian",
		Generator: func() EncoderDecoder {
			return NewUnicodeEncoderDecoder(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM))
		},
	},
	{
		Names:       []string{"utf-16-le-bom", "utf16lebom"},
		Description: "UTF-16 Little Endian with BOM",
		Generator: func() EncoderDecoder {
			return NewUnicodeEncoderDecoder(unicode.UTF16(unicode.LittleEndian, unicode.UseBOM))
		},
	},
	{
		Names:       []string{"utf-16-be", "utf16be"},
		Description: "UTF-16 Big Endian",
		Generator: func() EncoderDecoder {
			return NewUnicodeEncoderDecoder(unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM))
		},
	},
	{
		Names:       []string{"utf-16-be-bom", "utf16bebom"},
		Description: "UTF-16 Big Endian with BOM",
		Generator: func() EncoderDecoder {
			return NewUnicodeEncoderDecoder(unicode.UTF16(unicode.BigEndian, unicode.UseBOM))
		},
	},
	{
		Names:       []string{"windows-874", "windows874", "win874"},
		Description: "Windows 874",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows874) },
	},
	{
		Names:       []string{"windows-1250", "windows1250", "win1250"},
		Description: "Windows 1250",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1250) },
	},
	{
		Names:       []string{"windows-1251", "windows1251", "win1251"},
		Description: "Windows 1251",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1251) },
	},
	{
		Names:       []string{"windows-1252", "windows1252", "win1252"},
		Description: "Windows 1252",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1252) },
	},
	{
		Names:       []string{"windows-1253", "windows1253", "win1253"},
		Description: "Windows 1253",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1253) },
	},
	{
		Names:       []string{"windows-1254", "windows1254", "win1254"},
		Description: "Windows 1254",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1254) },
	},
	{
		Names:       []string{"windows-1255", "windows1255", "win1255"},
		Description: "Windows 1255",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1255) },
	},
	{
		Names:       []string{"windows-1256", "windows1256", "win1256"},
		Description: "Windows 1256",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1256) },
	},
	{
		Names:       []string{"windows-1257", "windows1257", "win1257"},
		Description: "Windows 1257",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1257) },
	},
	{
		Names:       []string{"windows-1258", "windows1258", "win1258"},
		Description: "Windows 1258",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Windows1258) },
	},
	{
		Names:       []string{"codepage037", "cp037"},
		Description: "IBM Code Page 037",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage037) },
	},
	{
		Names:       []string{"codepage437", "cp437"},
		Description: "IBM Code Page 437",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage437) },
	},
	{
		Names:       []string{"codepage850", "cp850"},
		Description: "IBM Code Page 850",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage850) },
	},
	{
		Names:       []string{"codepage852", "cp852"},
		Description: "IBM Code Page 852",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage852) },
	},
	{
		Names:       []string{"codepage855", "cp855"},
		Description: "IBM Code Page 855",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage855) },
	},
	{
		Names:       []string{"codepage858", "cp858"},
		Description: "IBM Code Page 858",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage858) },
	},
	{
		Names:       []string{"codepage860", "cp860"},
		Description: "IBM Code Page 860",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage860) },
	},
	{
		Names:       []string{"codepage862", "cp862"},
		Description: "IBM Code Page 862",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage862) },
	},
	{
		Names:       []string{"codepage863", "cp863"},
		Description: "IBM Code Page 863",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage863) },
	},
	{
		Names:       []string{"codepage865", "cp865"},
		Description: "IBM Code Page 865",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage865) },
	},
	{
		Names:       []string{"codepage866", "cp866"},
		Description: "IBM Code Page 866",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage866) },
	},
	{
		Names:       []string{"codepage1047", "cp1047"},
		Description: "IBM Code Page 1047",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage1047) },
	},
	{
		Names:       []string{"codepage1140", "cp1047"},
		Description: "IBM Code Page 1140",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.CodePage1140) },
	},
	{
		Names:       []string{"iso-8859-1", "iso-latin-1", "latin1"},
		Description: "ISO 8859-1",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_1) },
	},
	{
		Names:       []string{"iso-8859-2", "iso-latin-2", "latin2"},
		Description: "ISO 8859-2",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_2) },
	},
	{
		Names:       []string{"iso-8859-3", "iso-latin-3", "latin3"},
		Description: "ISO 8859-3",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_3) },
	},
	{
		Names:       []string{"iso-8859-4", "iso-latin-4", "latin4"},
		Description: "ISO 8859-4",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_4) },
	},
	{
		Names:       []string{"iso-8859-5", "iso-latin-cyrillic", "cyrillic"},
		Description: "ISO 8859-5",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_5) },
	},
	{
		Names:       []string{"iso-8859-6", "iso-latin-arabic", "arabic"},
		Description: "ISO 8859-6",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_6) },
	},
	{
		Names:       []string{"iso-8859-7", "iso-latin-greek", "greek"},
		Description: "ISO 8859-7",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_7) },
	},
	{
		Names:       []string{"iso-8859-8", "iso-latin-hebrew", "hebrew"},
		Description: "ISO 8859-8",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_8) },
	},
	{
		Names:       []string{"iso-8859-9", "iso-latin-5", "latin5"},
		Description: "ISO 8859-9",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_9) },
	},
	{
		Names:       []string{"iso-8859-10", "iso-latin-6", "latin6"},
		Description: "ISO 8859-10",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_10) },
	},
	{
		Names:       []string{"iso-8859-13"},
		Description: "ISO 8859-13",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_13) },
	},
	{
		Names:       []string{"iso-8859-14"},
		Description: "ISO 8859-14",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_14) },
	},
	{
		Names:       []string{"iso-8859-15"},
		Description: "ISO 8859-15",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_15) },
	},
	{
		Names:       []string{"iso-8859-16"},
		Description: "ISO 8859-16",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.ISO8859_16) },
	},
	{
		Names:       []string{"koi8-r", "koi8r"},
		Description: "KOI8-R",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.KOI8R) },
	},
	{
		Names:       []string{"koi8-u", "koi8u"},
		Description: "KOI8-U",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.KOI8U) },
	},
	{
		Names:       []string{"macintosh", "mac"},
		Description: "Macintosh",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.Macintosh) },
	},
	{
		Names:       []string{"macintosh-cyrillic", "mac-cyrillic"},
		Description: "Macintosh Cyrillic",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.MacintoshCyrillic) },
	},
	{
		Names:       []string{"x-user-defined", "xud"},
		Description: "X-User-Defined",
		Generator:   func() EncoderDecoder { return NewCharmapEncoderDecoder(charmap.XUserDefined) },
	},
	{
		Names:       []string{"rds-g0", "rds"},
		Description: "FM radio RDS, G0 table",
		Generator:   func() EncoderDecoder { return NewRDSEncoderDecoder() },
	},
}
