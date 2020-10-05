TODO List:

- [x] Fix diode positions
- [ ] Column traces
- [ ] Row traces
- [ ] RGB traces
- [ ] Drill hole to expose reset button

LED & Resistor
```
    (module MX_LED (layer Front)
      (tedit 5A362466)
      (at 73.661 38.325 -10)
      (fp_text reference D_SW9 (at 0 3.302) (layer B.SilkS) hide
        (effects (font (size 1 1) (thickness 0.2032)) (justify mirror)))
      (fp_text value LED (at 0 6.858) (layer B.SilkS) hide
        (effects (font (size 1 1) (thickness 0.2032)) (justify mirror)))
      (fp_text user - (at -1.27 3.25) (layer B.SilkS)
        (effects (font (size 1 1) (thickness 0.15)) (justify mirror)))
      (fp_text user - (at -1.27 3.25) (layer F.SilkS)
        (effects (font (size 1 1) (thickness 0.15))))
      (fp_text user + (at 1.27 3.25) (layer B.SilkS)
        (effects (font (size 1 1) (thickness 0.15)) (justify mirror)))
      (fp_text user + (at 1.27 3.25) (layer F.SilkS)
        (effects (font (size 1 1) (thickness 0.15))))
      (pad 1 thru_hole rect (at -1.27 5.08 -10) (size 1.905 1.905) (drill 0.9906) (layers *.Cu *.Mask))
      (pad 2 thru_hole circle (at 1.27 5.08 -10) (size 1.905 1.905) (drill 0.9906) (layers *.Cu *.SilkS *.Mask))
    )

    (module R_0805 (layer Front)
      (tedit 5CED8C8D)
      (at 50 25 -10)
      (descr "SMT, 2012, 0805")
      (tags "SMT, 2012, 0805")
      (attr smd)
      (fp_text reference R_0805 (at -2.2 0 80) (layer F.SilkS)
        (effects (font (size 0.7 0.7) (thickness 0.15))))
      (fp_text value R01 (at 0 1.55 -10) (layer F.SilkS)
        (effects (font (size 0.7 0.7) (thickness 0.15))))
      (fp_line (start -1.5 -0.85) (end 1.5 -0.85) (layer F.SilkS) (width 0.2))
      (fp_line (start 1.5 -0.85) (end 1.5 0.85) (layer F.SilkS) (width 0.2))
      (fp_line (start -1.5 -0.85) (end -1.5 0.85) (layer F.SilkS) (width 0.2))
      (fp_line (start -1.5 0.85) (end 1.5 0.85) (layer F.SilkS) (width 0.2))
      (pad 2 smd rect (at 0.95 0 -10) (size 0.7 1.3) (layers F.Cu F.Paste F.Mask)
        (clearance 0.1))
      (pad 1 smd rect (at -0.95 0 -10) (size 0.7 1.3) (layers F.Cu F.Paste F.Mask)
        (clearance 0.1))
      (model ${KISYS3DMOD}/Resistor_SMD.3dshapes/R_0805_2012Metric.wrl
        (at (xyz 0 0 0))
        (scale (xyz 1 1 1))
        (rotate (xyz 0 0 0)))
    )
```
