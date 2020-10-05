# How to Create a Kicad Module from a PNG File

## Creating the Full-Size Kicad Module

Create a bitmap image file using the software of your choice. For example,
create an SVG in Inkscape and export a PNG.

For reasons discussed below, it's better to export a larger image and then
scale the module down. In my experience, when exporting an image with a 4:1
aspect ratio, a size of 2048x512 worked well.

With the image file in hand, run Kicad and open the `Bitmap to Component
Converter` tool. On the right-hand side of the window, click the `Load Bitmap`
button and select the image file.

In the `Board Layer for Outline` options, select the `Front silk screen` option.

Click the `Export` button to write the result to a `kicad_mod` file. This
should be placed in a directory that is loaded in Kicad as a footprint library.

From the home Kicad window, open the `Footprint Editor` tool. Select the newly
created module and ensure it looks as expected.

As hinted at above, Kicad's bitmap converter tool struggles with smaller
bitmaps. This becomes obvious when trying to convert small images containing
text, in which case the font will look terrible.

## Scaling Down Kicad Modules

A script exists for scaling down Kicad modules. It was taken from David
Siroky's blog here:
http://www.smallbulb.net/kicad_mod-scaler

I've committed a copy of the script to this repository with Python 3 support:
[scripts/scale_kicad_mod.py](../../scripts/scale_kicad_mod.py)

The script is very easy to use:
```bash
scripts/scale_kicad_mod.py 0.25 < image.kicad_mod > image-small.kicad_mod
```

My workflow for scaling modules is:
    - 1. Open the `pcbnew` tool in Kicad
    - 2. Scale the module using the script
    - 3. Place the scaled module on the PCB
    - 4. Check if the size is adequate
        - 5. If not, delete the module from the PCB and return to step 2
    - 6. Done!
