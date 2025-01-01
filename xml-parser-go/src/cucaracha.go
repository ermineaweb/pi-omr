package main

var Cucaracha = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE score-partwise PUBLIC "-//Recordare//DTD MusicXML 4.0.1 Partwise//EN" "http://www.musicxml.org/dtds/partwise.dtd">
<score-partwise version="4.0.1">
  <work>
    <work-number>29 LA CUCARACHA</work-number>
  </work>
  <movement-number>1</movement-number>
  <movement-title>[Audiveris detected movement]</movement-title>
  <identification>
    <creator type="composer">Amérique du Sud</creator>
    <encoding>
      <software>Audiveris 5.3-alpha</software>
      <supports type="yes" element="print" attribute="new-system" value="yes"></supports>
      <supports type="yes" element="print" attribute="new-page" value="yes"></supports>
      <software>ProxyMusic 4.0.1</software>
      <encoding-date>2024-12-30</encoding-date>
    </encoding>
    <source>/input/cucaracha.png</source>
    <miscellaneous>
      <miscellaneous-field name="source-file">/input/cucaracha.png</miscellaneous-field>
      <miscellaneous-field name="source-sheet-1">1 2 3</miscellaneous-field>
    </miscellaneous>
  </identification>
  <defaults>
    <scaling>
      <millimeters>7.1120</millimeters>
      <tenths>40</tenths>
    </scaling>
    <page-layout>
      <page-height>945</page-height>
      <page-width>1200</page-width>
      <page-margins type="both">
        <left-margin>80</left-margin>
        <right-margin>80</right-margin>
        <top-margin>80</top-margin>
        <bottom-margin>80</bottom-margin>
      </page-margins>
    </page-layout>
    <lyric-font font-family="Default" font-size="10"></lyric-font>
  </defaults>
  <credit page="1">
    <credit-words default-x="56" default-y="889" font-family="serif" font-size="8">26</credit-words>
  </credit>
  <credit page="1">
    <credit-words default-x="441" default-y="861" font-family="serif" font-size="14">29 LA CUCARACHA</credit-words>
  </credit>
  <credit page="1">
    <credit-type>composer</credit-type>
    <credit-words default-x="963" default-y="813" font-family="serif" font-size="9">Amérique du Sud</credit-words>
  </credit>
  <part-list>
    <score-part id="P1">
      <part-name>Piano</part-name>
      <score-instrument id="P1-I1">
        <instrument-name>Acoustic Grand Piano</instrument-name>
      </score-instrument>
      <midi-instrument id="P1-I1">
        <midi-channel>1</midi-channel>
        <midi-program>1</midi-program>
        <volume>78</volume>
      </midi-instrument>
    </score-part>
  </part-list>
  <!--= = = = = = = = = = = = = = = = = = = = = = = = = = = = =-->
  <part id="P1">
    <!--=======================================================-->
    <measure number="1" width="263">
      <print>
        <system-layout>
          <system-margins>
            <left-margin>50</left-margin>
            <right-margin>-21</right-margin>
          </system-margins>
          <top-system-distance>95</top-system-distance>
        </system-layout>
        <staff-layout number="2">
          <staff-distance>76</staff-distance>
        </staff-layout>
        <measure-numbering>system</measure-numbering>
      </print>
      <attributes>
        <divisions>2</divisions>
        <staves>2</staves>
        <clef number="1">
          <sign>G</sign>
          <line>2</line>
        </clef>
        <clef number="2">
          <sign>G</sign>
          <line>2</line>
        </clef>
        <staff-details print-object="yes"></staff-details>
      </attributes>
      <direction placement="above">
        <direction-type>
          <symbol xsi:type="formatted-text"></symbol>
        </direction-type>
        <sound tempo="120"></sound>
      </direction>
      <direction placement="above">
        <direction-type>
          <symbol xsi:type="formatted-text" default-y="33" relative-x="-2" font-family="serif" font-size="9">Rapide et joyeux (J;112 env.)</symbol>
        </direction-type>
        <staff>1</staff>
      </direction>
      <note default-x="42">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-36">down</stem>
        <staff>1</staff>
        <notations>
          <slur type="continue" number="1" bezier-x2="3" bezier-y2="-5" placement="below" default-x="0" default-y="-35"></slur>
          <slur type="stop" number="1" bezier-x="2" bezier-y="-6" default-x="9" default-y="-30"></slur>
        </notations>
      </note>
      <direction placement="below">
        <direction-type>
          <dynamics default-y="-67" relative-x="-6">
            <mf/>
          </dynamics>
        </direction-type>
        <staff>1</staff>
        <sound dynamics="89"></sound>
      </direction>
      <note default-x="73">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="11">up</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="104">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="11">up</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="139">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-56">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="166">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-55">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="196">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-45">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="227">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-44">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>8</duration>
      </backup>
      <forward>
        <duration>2</duration>
        <voice>5</voice>
        <staff>2</staff>
      </forward>
      <direction placement="below">
        <direction-type>
          <dynamics default-y="-78" relative-x="-1">
            <p/>
          </dynamics>
        </direction-type>
        <staff>2</staff>
        <sound dynamics="56"></sound>
      </direction>
      <note default-x="73">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-9">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="139">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="139">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="139">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="196">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="196">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="196">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="2" width="170">
      <note default-x="16">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-33">down</stem>
        <staff>1</staff>
      </note>
      <note default-x="70">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>half</type>
        <stem default-y="-43">down</stem>
        <staff>1</staff>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="17">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-18">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="70">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="70">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="70">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="117">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="117">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="117">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="3" width="206">
      <note default-x="17">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-36">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="49">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-35">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="80">
        <pitch>
          <step>A</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-34">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="108">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-35">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="138">
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-45">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="169">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-45">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="17">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-14">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="81">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="81">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="81">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="138">
        <pitch>
          <step>C</step>
          <alter>1</alter>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <accidental>sharp</accidental>
        <stem default-y="-15">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="4" width="173">
      <note default-x="14">
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-38">down</stem>
        <staff>1</staff>
      </note>
      <note default-x="66">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>half</type>
        <stem default-y="-48">down</stem>
        <staff>1</staff>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="14">
        <pitch>
          <step>D</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-7">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="66">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="66">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="66">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="114">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-63">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="114">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-63">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="114">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-63">down</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="5" width="200">
      <note default-x="14">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="11">up</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="44">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="11">up</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="76">
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-62">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="105">
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-62">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="134">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-52">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="164">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-52">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="15">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-19">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="76">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="76">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="76">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="135">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="135">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="135">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-65">down</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="6" width="193">
      <print new-system="yes">
        <system-layout>
          <system-margins>
            <left-margin>-7</left-margin>
            <right-margin>-22</right-margin>
          </system-margins>
          <system-distance>122</system-distance>
        </system-layout>
        <staff-layout number="2">
          <staff-distance>76</staff-distance>
        </staff-layout>
      </print>
      <attributes>
        <staff-details print-object="yes"></staff-details>
      </attributes>
      <note default-x="54">
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-41">down</stem>
        <staff>1</staff>
      </note>
      <note default-x="101">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>half</type>
        <stem default-y="-51">down</stem>
        <staff>1</staff>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="54">
        <pitch>
          <step>D</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-6">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="101">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="101">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="101">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="146">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="146">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="146">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-64">down</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="7" width="183">
      <note default-x="16">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-37">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="46">
        <pitch>
          <step>A</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-37">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="72">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-42">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="99">
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-42">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="128">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-48">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="154">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-49">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="17">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-16">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="73">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="73">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="73">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="128">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-19">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="8" width="179">
      <note default-x="16">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-44">down</stem>
        <staff>1</staff>
      </note>
      <note default-x="64">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-56">down</stem>
        <staff>1</staff>
      </note>
      <note default-x="90">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="14">up</stem>
        <staff>1</staff>
        <notations>
          <slur type="start" number="2" bezier-x="29" bezier-y="20" placement="above" default-x="16" default-y="21"></slur>
        </notations>
      </note>
      <note default-x="118">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="10">up</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="143">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="10">up</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="17">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-12">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="65">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-69">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="65">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-69">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="65">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-69">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="120">
        <rest>
          <display-step>B</display-step>
          <display-octave>4</display-octave>
        </rest>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="9" width="171">
      <note default-x="12">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-55">down</stem>
        <staff>1</staff>
        <notations>
          <slur type="stop" number="2" bezier-x="-21" bezier-y="31" default-x="9" default-y="6"></slur>
          <slur type="start" number="2" bezier-x="10" bezier-y="1" placement="below" default-x="10" default-y="-52"></slur>
        </notations>
      </note>
      <note default-x="59">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-48">down</stem>
        <staff>1</staff>
        <notations>
          <slur type="stop" number="2" bezier-x="-9" bezier-y="-3" default-x="-9" default-y="-47"></slur>
        </notations>
      </note>
      <note default-x="85">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-56">down</stem>
        <staff>1</staff>
        <notations>
          <slur type="start" number="2" bezier-x="23" bezier-y="28" placement="above" default-x="8" default-y="-2"></slur>
        </notations>
      </note>
      <note default-x="112">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-54">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="140">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-54">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="12">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-13">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="60">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="60">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="60">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="112">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-18">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="10" width="159">
      <note default-x="17">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-45">down</stem>
        <staff>1</staff>
        <notations>
          <articulations>
            <staccato default-y="8" placement="above"></staccato>
          </articulations>
          <slur type="stop" number="2" bezier-x="-30" bezier-y="17" default-x="0" default-y="17"></slur>
        </notations>
      </note>
      <note default-x="64">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>half</type>
        <stem default-y="-34">down</stem>
        <staff>1</staff>
        <notations>
          <articulations>
            <tenuto default-y="21" placement="above"></tenuto>
          </articulations>
        </notations>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="17">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-14">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="64">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="64">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="64">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-66">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="107">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-18">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="11" width="185">
      <note default-x="13">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-58">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="44">
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-59">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="71">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-57">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="98">
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-58">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="124">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-55">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="153">
        <pitch>
          <step>C</step>
          <alter>1</alter>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <accidental>sharp</accidental>
        <stem default-y="-55">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="14">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-14">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="72">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="72">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="72">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="125">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-20">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="12" width="247">
      <print new-system="yes">
        <system-layout>
          <system-margins>
            <left-margin>-6</left-margin>
            <right-margin>-22</right-margin>
          </system-margins>
          <system-distance>122</system-distance>
        </system-layout>
        <staff-layout number="2">
          <staff-distance>76</staff-distance>
        </staff-layout>
      </print>
      <attributes>
        <staff-details print-object="yes"></staff-details>
      </attributes>
      <note default-x="53">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <tie type="start"></tie>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-50">down</stem>
        <staff>1</staff>
        <notations>
          <tied type="start" orientation="over"></tied>
        </notations>
      </note>
      <note default-x="108">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <tie type="stop"></tie>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="26">up</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
        <notations>
          <tied type="stop"></tied>
        </notations>
      </note>
      <note default-x="143">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="24">up</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
        <notations>
          <slur type="start" number="2" bezier-x="33" bezier-y="-19" placement="below" default-x="8" default-y="-44"></slur>
        </notations>
      </note>
      <note default-x="176">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="11">up</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="209">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="11">up</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="53">
        <pitch>
          <step>B</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-15">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="109">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="109">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="109">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="176">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-18">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="13" width="213">
      <direction placement="below">
        <direction-type>
          <wedge type="diminuendo" spread="11" default-x="14" default-y="-52"></wedge>
        </direction-type>
        <staff>1</staff>
      </direction>
      <note default-x="18">
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="20">up</stem>
        <staff>1</staff>
        <notations>
          <slur type="stop" number="2" bezier-x="-31" bezier-y="-26" default-x="0" default-y="-32"></slur>
        </notations>
      </note>
      <direction>
        <direction-type>
          <wedge type="stop" spread="0" default-x="-11"></wedge>
        </direction-type>
        <staff>1</staff>
      </direction>
      <note default-x="77">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-52">down</stem>
        <staff>1</staff>
      </note>
      <note default-x="109">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="16">up</stem>
        <staff>1</staff>
      </note>
      <note default-x="142">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="13">up</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="175">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="12">up</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="19">
        <pitch>
          <step>D</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-7">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="77">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="77">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="77">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="143">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-17">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="14" width="190">
      <note default-x="17">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-50">down</stem>
        <staff>1</staff>
        <notations>
          <articulations>
            <staccato default-y="7" placement="above"></staccato>
          </articulations>
        </notations>
      </note>
      <note default-x="78">
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>4</duration>
        <voice>1</voice>
        <type>half</type>
        <stem default-y="-43">down</stem>
        <staff>1</staff>
        <notations>
          <articulations>
            <tenuto default-y="15" placement="above"></tenuto>
          </articulations>
        </notations>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="17">
        <pitch>
          <step>D</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-6">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="77">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="77">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="77">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="130">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-16">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="15" width="224">
      <note default-x="18">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-36">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="50">
        <pitch>
          <step>A</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-35">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="82">
        <pitch>
          <step>G</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-41">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="118">
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-42">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <note default-x="150">
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-49">down</stem>
        <staff>1</staff>
        <beam number="1">begin</beam>
      </note>
      <note default-x="184">
        <pitch>
          <step>D</step>
          <octave>5</octave>
        </pitch>
        <duration>1</duration>
        <voice>1</voice>
        <type>eighth</type>
        <stem default-y="-50">down</stem>
        <staff>1</staff>
        <beam number="1">end</beam>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="19">
        <pitch>
          <step>B</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-16">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="83">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="83">
        <chord/>
        <pitch>
          <step>B</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="83">
        <chord/>
        <pitch>
          <step>F</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-67">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="150">
        <pitch>
          <step>G</step>
          <octave>3</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="-19">up</stem>
        <staff>2</staff>
      </note>
    </measure>
    <!--=======================================================-->
    <measure number="16" width="194">
      <note default-x="18">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-52">down</stem>
        <staff>1</staff>
        <notations>
          <articulations>
            <staccato default-y="7" placement="above"></staccato>
          </articulations>
        </notations>
      </note>
      <note default-x="71">
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <stem default-y="-55">down</stem>
        <staff>1</staff>
        <notations>
          <articulations>
            <tenuto default-y="9" placement="above"></tenuto>
          </articulations>
        </notations>
      </note>
      <note default-x="128">
        <rest>
          <display-step>B</display-step>
          <display-octave>4</display-octave>
        </rest>
        <duration>2</duration>
        <voice>1</voice>
        <type>quarter</type>
        <staff>1</staff>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="19">
        <rest>
          <display-step>B</display-step>
          <display-octave>4</display-octave>
        </rest>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <staff>2</staff>
      </note>
      <note default-x="72">
        <pitch>
          <step>G</step>
          <octave>4</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="25">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="72">
        <chord/>
        <pitch>
          <step>C</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="25">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="72">
        <chord/>
        <pitch>
          <step>E</step>
          <octave>5</octave>
        </pitch>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <stem default-y="25">up</stem>
        <staff>2</staff>
      </note>
      <note default-x="128">
        <rest>
          <display-step>B</display-step>
          <display-octave>4</display-octave>
        </rest>
        <duration>2</duration>
        <voice>5</voice>
        <type>quarter</type>
        <staff>2</staff>
      </note>
      <backup>
        <duration>6</duration>
      </backup>
      <note default-x="20">
        <pitch>
          <step>C</step>
          <octave>4</octave>
        </pitch>
        <duration>4</duration>
        <voice>6</voice>
        <type>half</type>
        <stem default-y="-86">down</stem>
        <staff>2</staff>
      </note>
      <note default-x="128">
        <rest>
          <display-step>B</display-step>
          <display-octave>3</display-octave>
        </rest>
        <duration>2</duration>
        <voice>6</voice>
        <type>quarter</type>
        <staff>2</staff>
      </note>
      <barline location="right">
        <bar-style>light-light</bar-style>
      </barline>
    </measure>
  </part>
</score-partwise>`)
