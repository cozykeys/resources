# kbutil

The `kbutil` command line tool provides a way to translate keyboard layout
information from an XML format that is easy to work with into other useful file
types.

Currently, it only supports generating SVG files; however, these can have
several purposes:
* Prototype keyboard layouts faster than raw SVG
* Generate SVG files that can be uploaded to Ponoko, Sculpteo, BigBlueSaw, or
  other online laser cutting services
    * Easily convert path styles in the generated SVG files to adhere to the
      style expectations of the various laser cutting services
* Draw legends and keycap borders for prototyping keycap sets

## Development

The tool is developed as a .NET Core command line application. To build the
tool, the .NET Core 2.0+ SDK must be installed.

To build from a shell using the `dotnet` CLI:
```bash
dotnet build -c Release src/KbUtil/KbUtil.sln
```

## Usage

## Generating SVGs

TODO

### Keycap Overlays
It can be hard to tell just from the switch cutouts whether or not keys are too
close together. By specifying the `--keycap-overlays` command line option when
running the tool, the resulting SVG will contain additional paths outlining
what the keycaps will look like on the keyboard

### Visual Styles
Ponoko expects all the lines to adhere to a specific style where the stroke
width is 0.01mm wide. This makes the lines very difficult to see, so by
specifying the `--visual-switch-cutouts` command line option, all paths will be
drawn with a more visual style.

### Keycap Legends
Keycap legends don't need to be in the SVG when it is sent to a laser cutting
service; however, they can be useful during layout development. So, the
`--keycap-legends` command line option will enable legends being written to the
resulting SVG.
gen-svg --keycap-overlays --visual-switch-cutouts --keycap-legends ../../../../../../Ergo87.xml
