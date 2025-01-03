package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Filename missing.")
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	score := Score{}
	xml.Unmarshal(raw, &score)

	pibuzzer_template := `pub const TEMPO: u16 = %s;
pub const MELODY: &str = r#"
%s
"#;`
	tempo, string_score := score.to_str()
	formatted := fmt.Sprintf(pibuzzer_template, tempo, string_score)
	fmt.Println(formatted)
}

type Score struct {
	Part struct {
		Measure struct {
			Direction struct {
				Sound struct {
					Tempo string `xml:"tempo,attr"`
				} `xml:"sound"`
			} `xml:"direction"`
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

func (s *Score) to_str() (string, string) {
	var str = ""
	var tempo = s.Part.Measure.Direction.Sound.Tempo
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
					panic(fmt.Sprintf("Note duration not implemented %s", duration))
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
	return tempo, str[:len(str)-1]
}
