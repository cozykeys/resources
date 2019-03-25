#!/usr/bin/env python3

# This file is super duper hacky and was a quick and dirty way to get the
# actual coordinates of switch cutouts to use to create a PCB. To use it, first
# generate a switch layer svg using KbUtil, then strip everything out of it
# except for the switches and the groups they are in. Load the svg in inkscape,
# ungroup everything, change the document settings so that it saves to MM
# instead of px, and save a copy. Then, pull all of the transform information
# into a rotations.json file that looks like this:
# [
#     { "rotation": { "degrees": 10, "axis": { "x": -127.23654, "y": 233.15604 } } },
#     { "rotation": { "degrees": 10, "axis": { "x": -236.10779, "y": 223.63104 } } },
#     { "translation": { "x": 155.95, "y": 67.675} },
# ]
#
# Then run this script on that file

from math import sin, cos, radians
import json

svg_header = '<?xml version="1.0" encoding="UTF-8"?>\n<svg xmlns="http://www.w3.org/2000/svg" width="500mm" height="500mm" viewBox="0 0 500 500">\n'
switch_path = '<path style="fill:none;stroke:#0000ff;stroke-width:0.5" d="M -7 -7 L 7 -7 L 7 -6 L 7.8 -6 L 7.8 6 L 7 6 L 7 7 L -7 7 L -7 6 L -7.8 6 L -7.8 -6 L -7 -6 L -7 -7" />\n'
svg_footer = '</svg>'

class Transform:
    def __init__(self, translation, rotation):
        self.translation = translation
        self.rotation = rotation


class Rotation:
    def __init__(self, theta, axis):
        self.theta = theta
        self.axis = axis


class Vec2:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def to_string(self):
        return '({0},{1})'.format(self.x, self.y)

def rotate(initial_pos, rotation):
    # Shift the axis to the origin and move the item accordingly
    x = initial_pos.x - rotation.axis.x
    y = initial_pos.y - rotation.axis.y

    dx = x * cos(rotation.theta) - y * sin(rotation.theta)
    dy = y * cos(rotation.theta) + x * sin(rotation.theta)

    # Shift back after rotating
    return Vec2(dx + rotation.axis.x, dy + rotation.axis.y)


def test(point, rotation, expected):
    print('[TEST] Rotate {0} around {1} by {2} radians'.format(point.to_string(), rotation.axis.to_string(), rotation.theta))
    print('    Expected: {0}'.format(expected.to_string()))
    print('    Actual:   {0}'.format(rotate(point, rotation).to_string()))


def main():
    #test(Vec2(0.0, -2.0), Rotation(radians(-90), Vec2(0.0, 0.0)), Vec2(-2.0, 0.0))
    #test(Vec2(0.0, 0.0), Rotation(radians(-90), Vec2(2.0, 0.0)), Vec2(2.0, 2.0))

    bearings = []
    with open('/home/pewing/src/cozykeys/rotations2.json') as f:
        bearings = json.loads(f.read())

    svg = svg_header

    transforms = []
    for bearing in bearings:
        if 'rotation' in bearing:
            r = bearing['rotation']
            position = rotate(Vec2(0.0, 0.0), Rotation(radians(r['degrees']), Vec2(r['axis']['x'], r['axis']['y'])))
            transforms.append(Transform(Vec2(position.x, position.y), r['degrees']))
        else:
            t = bearing['translation']
            transforms.append(Transform(Vec2(t['x'], t['y']), None))

    for transform in transforms:
        if transform.rotation is None:
            transform_string = 'transform="translate({0},{1})"'.format(
                    transform.translation.x, transform.translation.y)
            svg = svg + '<g {0}>{1}</g>'.format(transform_string, switch_path)
        else:
            transform_string = 'transform="translate({0},{1}) rotate({2})"'.format(
                    transform.translation.x, transform.translation.y, transform.rotation)
            svg = svg + '<g {0}>{1}</g>'.format(transform_string, switch_path)

    svg = svg + svg_footer

    with open('test.svg', 'w') as f:
        f.write(svg)

    switches = []
    for transform in transforms:
        switches.append({
            "row": 0,
            "column": 0,
            "x": transform.translation.x,
            "y": transform.translation.y,
            "rotation": 0 if transform.rotation is None else transform.rotation,
            "diode_position": "right"
        })

    with open('/home/pewing/src/cozykeys/bloomer_rev4_switches2.json', 'w') as f:
        f.write(json.dumps(switches, indent=4))


if __name__ == "__main__":
    main()

