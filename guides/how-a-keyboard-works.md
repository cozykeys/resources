# How a Keyboard Works

The purpose of this is to walk through how a keyboard works at every level.
This includes:

- What happens when the keyboard is connected over USB?
- What does the keyboard do when a key is pressed ?
- How does the keyboard communicate with operating system?
- How does the operating system communicate with user space applications?

This document is very much a work in progress and will likely take a lot of
time to finish. Also, it only aims to answer the above questions for Linux.

# HID Devices

- Device Class Definition for Human Interface Devices (HID)
    - https://www.usb.org/sites/default/files/hid1_11.pdf

# Driver

- Linux HID driver
    - https://github.com/torvalds/linux/tree/master/drivers/hid

# User Space Applications

- libinput
    - https://gitlab.freedesktop.org/libinput/libinput

# Firmware

- Keyboard Firmware
    - QMK 
        - https://github.com/qmk/qmk_firmware
        - More featureful fork of TMK
        - Useful links:
            - https://github.com/qmk/qmk_firmware/blob/master/tmk_core/common/keyboard.h
            - https://github.com/qmk/qmk_firmware/blob/master/tmk_core/common/keyboard.c
        - TODO: What is "quantum"?
    - TMK
        - https://github.com/tmk/tmk_keyboard
