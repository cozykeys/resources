# KiCad Setup

Clone useful git repositories with components/footprints:

```bash
mkdir -p ~/kicad && cd ~/kicad

# Component library by Hasu:
git clone https://github.com/tmk/kicad_lib_tmk

# Footprint library by Hasu:
git clone https://github.com/tmk/keyboard_parts.pretty

# Footprint library by /u/techieee
git clone https://github.com/egladman/keebs.pretty
```

Symlink component/footprint libraries from this repo:

```bash
ln -s $HOME/src/github/cozykeys/resources/cozy-components ~/kicad/cozy-components
ln -s $HOME/src/github/cozykeys/resources/cozy-parts.pretty ~/kicad/cozy-parts.pretty
```

Open KiCad and select the `Preferences -> Manage Symbol Libraries...` from the
menu. From the UI, add the component libraries from above:
- `cozy-components/cozy.lib`
- `kicad_lib_tmk/keyboard_parts.lib`

Press Ok to save and close the symbol libraries dialog.

Select `Preferences -> Manage Footprint Libraries...` from the menu. From the
UI, add the fooprint libraries:
- `cozy-parts.pretty`
- `keebs.pretty`
- `keyboard_parts.pretty`

Press Ok to save and close the footprint libraries dialog.
