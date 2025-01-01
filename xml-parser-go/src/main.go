package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	 raw, err := os.ReadFile("sheet.xml")
	 if err != nil {
		 panic(err)
	 }

	score := Score{}
	xml.Unmarshal(raw, &score)
	score.to_str()
	fmt.Println(score.to_str())
}

type Score struct {
	Part struct {
		Measure struct {
			Note []struct {
				Pitch struct {
					Step   string `xml:"step"`
					Octave string `xml:"octave"`
					Alter  int    `xml:"alter"`
				} `xml:"pitch"`
				Duration int    `xml:"duration"`
				Type     string `xml:"type"`
				Voice    int    `xml:"voice"`
			} `xml:"note"`
		} `xml:"measure"`
	} `xml:"part"`
}

func (s *Score) to_str() string {
	var str = ""
	for _, note := range s.Part.Measure.Note {
		if note.Voice == 1 {
			if note.Pitch.Step != "" {
				var duration = note.Type
				switch duration {
				case "half":
					duration = "2"
				case "quarter":
					duration = "4"
				case "eighth":
					duration = "8"
				case "16th":
					duration = "16"
				default:
					duration = "4"
				}

				var alter = ""
				switch note.Pitch.Alter {
				case 1:
					alter = "d"
				case -1:
					alter = "b"
				}

				var dotted = ""
				switch note.Duration {
				case 2:
					dotted = "-"
					// todo multi dotted
				}

				str = str + note.Pitch.Step + alter + note.Pitch.Octave + "," + dotted + duration + ","
			}
		}
	}

	// remove the last comma
	return str[:len(str)-1]
}
