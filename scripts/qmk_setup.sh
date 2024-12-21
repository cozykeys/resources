#!/usr/bin/env bash

# (2024-12-21) TODO: I don't really remember what this was but it was a local
# file I had in my QMK repo. I vaguely recall having to run this to set things
# up before building/flashing my firmware. Next time we go through setting up
# QMK we should see if this is still necessary. If it is, better document what
# it's for. If it's not, delete it to avoid confusion.

function yell () { >&2 echo "$*";  }
function die () { yell "$*"; exit 1; }
function try () { "$@" || die "Command failed: $*"; }

SCRIPT_PATH="$( realpath "$0" )"
SCRIPT_DIR="$( dirname "$SCRIPT_PATH" )"

QMK_REPO_DIR="$HOME/src/qmk_firmware"

# Missing or outdated udev rules for 'atmel-dfu' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'kiibohd' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'stm32' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'bootloadhid' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'usbasploader' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'massdrop' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'caterina' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Missing or outdated udev rules for 'hid-bootloader' boards. Run:
sudo cp /home/pewing/src/github/pcewing/qmk_firmware/util/udev/50-qmk.rules /etc/udev/rules.d/

# Detected ModemManager without the necessary udev rules. Please either disable
# it or set the appropriate udev rules if you are using a Pro Micro.
