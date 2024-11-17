package main

import "github.com/go-pdf/fpdf"

func main() {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SetAttachments(
		[]fpdf.Attachment{
			fpdf.Attachment{
				Filename: "test.xml",
				Description: "Test file",
				Content: []byte(`<?xml version="1.0" encoding="UTF-8"?>,
							<note>
								<to>Tove</to>	
								<from>Jani</from>
								<heading>Reminder</heading>
								<body>Don't forget me this weekend!</body>
							</note>`),
			},

		},

	)
	pdf.Cell(40, 10, "Hello, world")
	pdf.OutputFileAndClose("hello.pdf")
}