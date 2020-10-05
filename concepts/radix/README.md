# CozyKeys Radix

This repository contains the design files for the CozyKeys Radix keyboard. This
keyboard is still early in the design phase and is a work in progress.

![Radix Keyboard Render](./render.svg)

## TODO List
- [ ] Add in-switch LEDs
    - [ ] Reverse engineer how the Iris supports these
    - [ ] Update the PCB (And generation tools/scripts)
    - [ ] Find a suitable SMD LED part number

## Resources

### Keyboard PCB Design Tutorial

https://github.com/ruiqimao/keyboard-pcb-guide

### OLED Diplays

Adafruit has two different 128x32 OLED display modules but I believe the I2C
one is preferable. Other keyboards that have integrated SSD1306 displays have
used I2C display modules.

```
- 0.91" 128x32 Adafruit Display Modules
    - I2C
        - Schematic: https://github.com/adafruit/Adafruit-128x32-I2C-OLED-Breakout-PCB
        - Product Page: https://www.adafruit.com/product/931
        - Underlying Display Component: UG-2832HSWEG02
            - Datasheet: https://cdn-shop.adafruit.com/datasheets/UG-2832HSWEG02.pdf
    - SPI
        - Schematic: https://github.com/adafruit/Adafruit-128x32-SPI-OLED-breakout-board-PCB
        - Product Page: https://www.adafruit.com/product/661
        - Underlying Display Module: UG-2832HSWEG04
            - Datasheet: https://cdn-shop.adafruit.com/datasheets/UG-2832HSWEG04.pdf
```

