#!/usr/bin/env python3

# A lot of this is ridiculously hacky and makes major assumptions but it's
# still better than drawing traces by hand. :)

import json
import math

# Lookup table for net ids
net_ids = {
    '""': 0, 'N-5V-0': 1, 'N-GND-0': 2, 'N-MOSFET-0': 3,

    'N-row-0': 4, 'N-row-1': 5, 'N-row-2': 6, 'N-row-3': 7, 'N-row-4': 8,

    'N-col-0':   9, 'N-col-1':  10, 'N-col-2':  11, 'N-col-3': 12, 'N-col-4': 13,
    'N-col-5':  14, 'N-col-6':  15, 'N-col-7':  16, 'N-col-8': 17, 'N-col-9': 18,
    'N-col-10': 19, 'N-col-11': 20, 'N-col-12': 21,

    'N-diode-0-0':  22, 'N-diode-0-1':  23, 'N-diode-0-2':  24, 'N-diode-0-3':  25,
    'N-diode-0-4':  26, 'N-diode-0-5':  27, 'N-diode-0-6':  28, 'N-diode-0-7':  29,
    'N-diode-0-8':  30, 'N-diode-0-9':  31, 'N-diode-0-10': 32, 'N-diode-0-11': 33,
    'N-diode-0-12': 34, 'N-diode-1-0':  35, 'N-diode-1-1':  36, 'N-diode-1-2':  37,
    'N-diode-1-3':  38, 'N-diode-1-4':  39, 'N-diode-1-5':  40, 'N-diode-1-6':  41,
    'N-diode-1-7':  42, 'N-diode-1-8':  43, 'N-diode-1-9':  44, 'N-diode-1-10': 45,
    'N-diode-1-11': 46, 'N-diode-1-12': 47, 'N-diode-2-0':  48, 'N-diode-2-1':  49,
    'N-diode-2-2':  50, 'N-diode-2-3':  51, 'N-diode-2-4':  52, 'N-diode-2-5':  53,
    'N-diode-2-6':  54, 'N-diode-2-7':  55, 'N-diode-2-8':  56, 'N-diode-2-9':  57,
    'N-diode-2-10': 58, 'N-diode-2-11': 59, 'N-diode-2-12': 60, 'N-diode-3-0':  61,
    'N-diode-3-1':  62, 'N-diode-3-2':  63, 'N-diode-3-3':  64, 'N-diode-3-4':  65,
    'N-diode-3-5':  66, 'N-diode-3-7':  67, 'N-diode-3-8':  68, 'N-diode-3-9':  69,
    'N-diode-3-10': 70, 'N-diode-3-11': 71, 'N-diode-3-12': 72, 'N-diode-4-0':  73,
    'N-diode-4-1':  74, 'N-diode-4-2':  75, 'N-diode-4-3':  76, 'N-diode-4-4':  77,
    'N-diode-4-5':  78, 'N-diode-4-7':  79, 'N-diode-4-8':  80, 'N-diode-4-9':  81,
    'N-diode-4-10': 82, 'N-diode-4-11': 83, 'N-diode-4-12': 84,

    'N-LED-0-0':   85, 'N-LED-0-1':   86, 'N-LED-0-2':  87,  'N-LED-0-3':   88,
    'N-LED-0-4':   89, 'N-LED-0-5':   90, 'N-LED-0-6':  91,  'N-LED-0-7':   92,
    'N-LED-0-8':   93, 'N-LED-0-9':   94, 'N-LED-0-10': 95,  'N-LED-0-11':  96,
    'N-LED-0-12':  97, 'N-LED-1-0':   98, 'N-LED-1-1':  99,  'N-LED-1-2':  100,
    'N-LED-1-3':  101, 'N-LED-1-4':  102, 'N-LED-1-5':  103, 'N-LED-1-6':  104,
    'N-LED-1-7':  105, 'N-LED-1-8':  106, 'N-LED-1-9':  107, 'N-LED-1-10': 108,
    'N-LED-1-11': 109, 'N-LED-1-12': 110, 'N-LED-2-0':  111, 'N-LED-2-1':  112,
    'N-LED-2-2':  113, 'N-LED-2-3':  114, 'N-LED-2-4':  115, 'N-LED-2-5':  116,
    'N-LED-2-6':  117, 'N-LED-2-7':  118, 'N-LED-2-8':  119, 'N-LED-2-9':  120,
    'N-LED-2-10': 121, 'N-LED-2-11': 122, 'N-LED-2-12': 123, 'N-LED-3-0':  124,
    'N-LED-3-1':  125, 'N-LED-3-2':  126, 'N-LED-3-3':  127, 'N-LED-3-4':  128,
    'N-LED-3-5':  129, 'N-LED-3-7':  130, 'N-LED-3-8':  131, 'N-LED-3-9':  132,
    'N-LED-3-10': 133, 'N-LED-3-11': 134, 'N-LED-3-12': 135, 'N-LED-4-0':  136,
    'N-LED-4-1':  137, 'N-LED-4-2':  138, 'N-LED-4-3':  139, 'N-LED-4-4':  140,
    'N-LED-4-5':  141, 'N-LED-4-7':  142, 'N-LED-4-8':  143, 'N-LED-4-9':  144,
    'N-LED-4-10': 145, 'N-LED-4-11': 146, 'N-LED-4-12': 147,
    
    'N-RGB-D0': 148, 'N-RGB-D1': 149, 'N-RGB-D2':  150, 'N-RGB-D3':  151,
    'N-RGB-D4': 152, 'N-RGB-D5': 153, 'N-RGB-D6':  154, 'N-RGB-D7':  155,
    'N-RGB-D8': 156, 'N-RGB-D9': 157, 'N-RGB-D10': 158, 'N-RGB-D11': 159,

    'N-LED-PWM': 160,
}

def print_segment(start_x, start_y, end_x, end_y, net_id, layer="Front"):
    start_str   = '(start {x} {y})'.format(x=start_x, y=start_y)
    end_str     = '(end {x} {y})'.format(x=end_x, y=end_y)
    width_str   = '(width 0.2032)'
    layer_str   = '(layer {})'.format(layer)
    net_str     = '(net {net_id})'.format(net_id=net_id)
    if net_id is not None:
        print('(segment {0} {1} {2} {3} {4})'.format(
            start_str, end_str, width_str, layer_str, net_str))
    else:
        print('(segment {0} {1} {2} {3})'.format(
            start_str, end_str, width_str, layer_str))

def print_via(x, y, net_id):
  at_str     = '(at {} {})'.format(x, y)
  size_str   = '(size 0.889)'
  drill_str  = '(drill 0.635)'
  layers_str = '(layers Front Back)'
  net_str    = '(net {})'.format(net_id)
  print('(via {} {} {} {} {})'.format(at_str, size_str, drill_str, layers_str, net_str))

class Vec2:
    def __init__(self, x = 0, y = 0):
        self.x = x
        self.y = y
    
    def __str__(self):
        return "({0}, {1})".format(self.x, self.y)


class Segment:
    def __init__(self, start = Vec2(), end = Vec2()):
        self.start = start
        self.end = end

    def theta(self):
        dx = self.end.x - self.start.x
        dy = self.end.y - self.start.y

        # Early out if the segment is parallel to an axis.
        if abs(dy) < 0.001 and dx > 0:
            return (0.0 * math.pi)
        if abs(dx) < 0.001 and dy > 0:
            return (0.5 * math.pi)
        if abs(dy) < 0.001 and dx < 0:
            return (1.0 * math.pi)
        if abs(dx) < 0.001 and dy < 0:
            return (1.5 * math.pi)

        theta = math.atan(dy / dx)

        # Because slope doesn't take direction into account, we have to manually adjust theta for segments whose
        # direction points into quadrants 2 and 3.
        if ((dx < 0 and dy > 0) or (dx < 0 and dy < 0)):
            theta += math.pi

        return theta - (math.pi * 2)  * math.floor(theta / (math.pi * 2))
    
    def __str__(self):
        return "[ {0}, {1} ]".format(self.start, self.end)

    def print(self, net_id=None, layer="Front"):
        print_segment(self.start.x, self.start.y, self.end.x, self.end.y, net_id, layer)


def bar(x, y, dx, dy, rotation):
    m = math.sqrt(math.pow(dx, 2) + math.pow(dy, 2))
    s = Segment(Vec2(0, 0), Vec2(dx, dy))
    theta = s.theta() + math.radians(rotation)

    x = x + m * math.cos(theta)
    y = y + m * math.sin(theta)
    return Vec2(x, y)


def foo(switch, dx, dy):
    return bar(switch['x'], switch['y'], dx, dy, switch['rotation'])


def calculate_column_segments(switches):
    dx = 7.065
    dy = 2.54

    segments = {}
    for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
        s0 = switches[0][c]
        s1 = switches[4][c]
        segments[c] = Segment(foo(s0, dx, dy), foo(s1, dx, dy))
    return segments


def print_column_segments(column_segments):
    for c in column_segments:
        column_segments[c].print(net_ids['N-col-{}'.format(c)], "Front")


def print_switch_to_column_segments(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
            s = switches[r][c]
            Segment(foo(s, 3.81, 2.54), foo(s, 7.065, 2.54)).print(net_ids['N-col-{}'.format(c)], "Front")

    s = switches[0][6]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 7.065, 2.54)
    v3 = foo(s, 7.065, -7.7)
    Segment(v1, v2).print(net_ids['N-col-6'], "Front")
    Segment(v2, v3).print(net_ids['N-col-6'], "Front")

    s = switches[1][6]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 7.065, 2.54)
    v3 = Vec2(162.255, 125.095)
    Segment(v1, v2).print(net_ids['N-col-6'], "Front")
    Segment(v2, v3).print(net_ids['N-col-6'], "Front")
    print_via(v3.x, v3.y, net_ids['N-col-6'])

    s = switches[2][6]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 7.065, 2.54)
    v3 = Vec2(202.093, 125.095)
    Segment(v1, v2).print(net_ids['N-col-6'], "Front")
    Segment(v2, v3).print(net_ids['N-col-6'], "Front")
    print_via(v3.x, v3.y, net_ids['N-col-6'])

    v1 = Vec2(162.255, 125.095)
    v2 = Vec2(202.093, 125.095)
    Segment(v1, v2).print(net_ids['N-col-6'], "Back")

    v1 = Vec2(182.065, 108.54)
    v2 = Vec2(182.065, 125.095)
    Segment(v1, v2).print(net_ids['N-col-6'], "Front")
    print_via(v2.x, v2.y, net_ids['N-col-6'])

    v1 = Vec2(182.065, 98.3)
    v2 = Vec2(167.64, 98.3)
    v3 = Vec2(165.735, 96.52)
    v4 = Vec2(165.735, 76.38)
    v5 = Vec2(167.38, 76.38)
    Segment(v1, v2).print(net_ids['N-col-6'], "Front")
    print_via(v2.x, v2.y, net_ids['N-col-6'])
    Segment(v2, v3).print(net_ids['N-col-6'], "Back")
    Segment(v3, v4).print(net_ids['N-col-6'], "Back")
    Segment(v4, v5).print(net_ids['N-col-6'], "Back")


def print_led_to_resistor_segments(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,6,7,8,9,10,11,12]:
            if r in switches and c in switches[r]:
                s = switches[r][c]
                Segment(foo(s, 1.27, -5.08), foo(s, 4.445, -5.08)).print(net_ids['N-LED-{}-{}'.format(r, c)], "Back")


def print_switch_to_diode_segments(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,6,7,8,9,10,11,12]:
            if r in switches and c in switches[r]:
                s = switches[r][c]
                Segment(foo(s, -2.54, 5.08), foo(s, -2.54, 1.27)).print(net_ids['N-diode-{}-{}'.format(r, c)], "Back")
                Segment(foo(s, -2.54, 1.27), foo(s, -4.445, -2.146)).print(net_ids['N-diode-{}-{}'.format(r, c)], "Back")


def calculate_mosfet_column_segments(switches):
    dx = 8.335
    dy = -3.18

    segments = {}
    for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
        s0 = switches[0][c]
        s1 = switches[4][c]
        segments[c] = Segment(foo(s0, dx, dy), foo(s1, dx, dy))
    return segments


def print_mosfet_column_segments(mosfet_column_segments):
    for c in mosfet_column_segments:
        mosfet_column_segments[c].print(net_ids['N-MOSFET-0'], "Front")


def print_switch_to_mosfet_column_segments(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
            s = switches[r][c]
            Segment(foo(s, 4.445, -3.18), foo(s, 8.335, -3.18)).print(net_ids['N-MOSFET-0'], "Back")
    
    s = switches[0][6]
    Segment(foo(s, 4.445, -3.18), foo(s, 8.335, -3.18)).print(net_ids['N-MOSFET-0'], "Back")

    s = switches[1][6]
    Segment(foo(s, 4.445, -3.18), foo(s, 8.335, -3.18)).print(net_ids['N-MOSFET-0'], "Back")

    s = switches[2][6]
    Segment(foo(s, 4.445, -3.18), foo(s, 8.335, -3.18)).print(net_ids['N-MOSFET-0'], "Back")


def print_switch_to_mosfet_column_vias(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
            s = switches[r][c]
            v = foo(s, 8.335, -3.18)
            print_via(v.x, v.y, net_ids['N-MOSFET-0'])

    s = switches[0][6]
    v = foo(s, 8.335, -3.18)
    print_via(v.x, v.y, net_ids['N-MOSFET-0'])

    s = switches[1][6]
    v = foo(s, 8.335, -3.18)
    print_via(v.x, v.y, net_ids['N-MOSFET-0'])

    s = switches[2][6]
    v = foo(s, 8.335, -3.18)
    print_via(v.x, v.y, net_ids['N-MOSFET-0'])


def calculate_vcc_column_segments(switches):
    dx = -7.065
    dy = -3.81

    segments = {}
    for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
        s0 = switches[0][c]
        s1 = switches[4][c]
        segments[c] = Segment(foo(s0, dx, dy), foo(s1, dx, dy))
    return segments


def print_vcc_column_segments(vcc_column_segments):
    for c in vcc_column_segments:
        vcc_column_segments[c].print(net_ids['N-5V-0'], "Front")


def print_switch_to_vcc_column_segments(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,6,7,8,9,10,11,12]:
            if r in switches and c in switches[r]:
                s = switches[r][c]
                Segment(foo(s, -1.27, -5.08), foo(s, -4.445, -3.81)).print(net_ids['N-5V-0'], "Front")
                Segment(foo(s, -4.445, -3.81), foo(s, -7.065, -3.81)).print(net_ids['N-5V-0'], "Front")


def print_diode_to_row_components(switches):
    for r in [0,1,2,3,4]:
        for c in [0,1,2,3,4,5,6,7,8,9,10,11,12]:
            if r in switches and c in switches[r]:
                s = switches[r][c]
                net_id = net_ids['N-row-{}'.format(r)]

                v1 = foo(s, -4.445, -5.474)
                v2 = foo(s, -7.065, -5.474)
                v3 = foo(s, -7.065, 2.54)
                v4 = foo(s, -5.08, 2.54)
                v5 = foo(s, -1.905, 2.54)
                v6 = foo(s, 0.635, 5.08)
                v7 = foo(s, 5.08, 5.08)

                Segment(v1, v2).print(net_id, "Back")
                Segment(v2, v3).print(net_id, "Back")
                Segment(v3, v4).print(net_id, "Back")
                Segment(v4, v5).print(net_id, "Front")
                Segment(v5, v6).print(net_id, "Front")
                Segment(v6, v7).print(net_id, "Front")

                print_via(v4.x, v4.y, net_id)
                print_via(v7.x, v7.y, net_id)

    v1 = Vec2(170.555, 100.526)
    v2 = Vec2(169.92, 99.06)
    v3 = Vec2(169.92, 86.36)
    Segment(v1, v2).print(net_ids['N-row-0'], "Back")
    Segment(v2, v3).print(net_ids['N-row-0'], "Back")

    v1 = Vec2(150.28408, 126.699296)
    v2 = Vec2(150.679, 124.46)
    v3 = Vec2(150.679, 121.285)
    print_via(v2.x, v2.y, net_ids['N-row-1'])
    Segment(v1, v2).print(net_ids['N-row-1'], "Back")
    Segment(v2, v3).print(net_ids['N-row-1'], "Front")
    v4 = Vec2(155.575, 121.285)
    v5 = Vec2(155.575, 92.71)
    v6 = Vec2(155.575, 87.63)
    Segment(v3, v4).print(net_ids['N-row-1'], "Front")
    Segment(v4, v5).print(net_ids['N-row-1'], "Front")
    print_via(v5.x, v5.y, net_ids['N-row-1'])
    Segment(v5, v6).print(net_ids['N-row-2'], "Back")
    print_via(v6.x, v6.y, net_ids['N-row-2'])

    v1 = Vec2(190.960979, 128.243029)
    v2 = Vec2(190.63, 126.365)
    v3 = Vec2(190.63, 121.285)
    print_via(v2.x, v2.y, net_ids['N-row-2'])
    Segment(v1, v2).print(net_ids['N-row-2'], "Back")
    Segment(v2, v3).print(net_ids['N-row-2'], "Front")
    v4 = Vec2(194.31, 121.285)
    v5 = Vec2(194.31, 92.71)
    v6 = Vec2(194.31, 88.9)
    Segment(v3, v4).print(net_ids['N-row-2'], "Front")
    Segment(v4, v5).print(net_ids['N-row-2'], "Front")
    print_via(v5.x, v5.y, net_ids['N-row-2'])
    Segment(v5, v6).print(net_ids['N-row-2'], "Back")
    print_via(v6.x, v6.y, net_ids['N-row-2'])


def print_row_attachment_segments(switches):
    for r in [0,1,2,3,4]:
        net_id = net_ids['N-row-{}'.format(r)]

        for c in [0,1,2,3,4]:
            s1 = switches[r][c]
            s2 = switches[r][c+1]

            v1 = foo(s1, 5.08, 5.08)
            v2 = foo(s2, -7.065, 2.54)

            Segment(v1, v2).print(net_id, "Back")

        for c in [7,8,9,10,11]:
            s1 = switches[r][c]
            s2 = switches[r][c+1]

            v1 = foo(s1, 5.08, 5.08)
            v2 = foo(s2, -7.065, 2.54)

            Segment(v1, v2).print(net_id, "Back")

    for i in range(0,5):
        net_id = net_ids['N-row-{}'.format(i)]
        x = 169.92 + (i * 2.54)
        v1 = Vec2(x, 78.92)
        v2 = Vec2(x, 86.36 + (i * 1.27))
        Segment(v1, v2).print(net_id, "Back")
        print_via(v2.x, v2.y, net_id)

    row = 0
    net_id = net_ids['N-row-0']
    v1 = Vec2(152.303, 60.396)
    v2 = bar(v1.x, v1.y, 5, 0, 10)
    v3 = bar(v2.x, v2.y, 0, 17.78, 10)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = bar(v3.x, v3.y, 1.27, 0, 10)
    v5 = Vec2(154.092, 86.36)
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Front")
    v6 = Vec2(169.92, 86.36)
    Segment(v5, v6).print(net_id, "Front")

    row = 1
    net_id = net_ids['N-row-1']
    v1 = Vec2(148.995, 79.157)
    v2 = bar(v1.x, v1.y, 5, 0, 10)
    v3 = Vec2(152.578, 87.63)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = Vec2(172.46, 87.63)
    Segment(v3, v4).print(net_id, "Front")

    row = 2
    net_id = net_ids['N-row-2']
    v1 = Vec2(145.687, 97.917)
    v2 = bar(v1.x, v1.y, 5, 0, 10)
    v3 = Vec2(152.355, 88.9)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = Vec2(175, 88.9)
    Segment(v3, v4).print(net_id, "Front")

    row = 3
    net_id = net_ids['N-row-3']
    v1 = Vec2(142.379, 116.678)
    v2 = bar(v1.x, v1.y, 5, 0, 10)
    v3 = bar(v2.x, v2.y, 0, -17.78, 10)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = bar(v3.x, v3.y, 1.27, 0, 10)
    v5 = Vec2(153.42, 90.17)
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Front")
    v6 = Vec2(177.54, 90.17)
    Segment(v5, v6).print(net_id, "Front")

    row = 4
    net_id = net_ids['N-row-4']
    v1 = Vec2(139.071, 135.439)
    v2 = bar(v1.x, v1.y, 5, 0, 10)
    v3 = bar(v2.x, v2.y, 0, -17.78, 10)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = bar(v3.x, v3.y, 1.27, 0, 10)
    v5 = bar(v4.x, v4.y, 0, -17.78, 10)
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Front")
    v6 = bar(v5.x, v5.y, 1.27, 0, 10)
    v7 = Vec2(154.486, 91.44)
    Segment(v5, v6).print(net_id, "Front")
    Segment(v6, v7).print(net_id, "Front")
    v8 = Vec2(180.08, 91.44)
    Segment(v7, v8).print(net_id, "Front")

    row = 0
    net_id = net_ids['N-row-0']
    va = Vec2(192.772961, 61.264241)
    vb = Vec2(195.3014, 58.239236)
    Segment(va, vb).print(net_id, "Back")
    v1 = Vec2(175 + (175 - 152.303), 60.396)
    v2 = bar(v1.x, v1.y, -5, 0, -10)
    v3 = bar(v2.x, v2.y, 0, 17.78, -10)
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = bar(v3.x, v3.y, -1.27, 0, -10)
    v5 = Vec2(175 + (175 - 154.092), 86.36)
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Front")
    v6 = Vec2(169.92, 86.36)
    Segment(v5, v6).print(net_id, "Front")

    row = 1
    net_id = net_ids['N-row-1']
    va = Vec2(196.080961, 80.025241)
    vb = Vec2(198.6094, 77.000236)
    Segment(va, vb).print(net_id, "Back")
    v1 = Vec2(175 + (175 - 148.995), 79.157)
    v2 = bar(v1.x, v1.y, -5, 0, -10)
    v3 = Vec2(175 + (175 - 152.578), 87.63)
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = Vec2(172.46, 87.63)
    Segment(v3, v4).print(net_id, "Front")

    row = 2
    net_id = net_ids['N-row-2']
    va = Vec2(199.388961, 98.785241)
    vb = Vec2(201.9174, 95.760236)
    Segment(va, vb).print(net_id, "Back")
    v1 = Vec2(175 + (175 - 145.687), 97.917)
    v2 = bar(v1.x, v1.y, -5, 0, -10)
    v3 = Vec2(175 + (175 - 152.355), 88.9)
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = Vec2(175, 88.9)
    Segment(v3, v4).print(net_id, "Front")

    row = 3
    net_id = net_ids['N-row-3']
    va = Vec2(202.696961, 117.546241)
    vb = Vec2(205.2254, 114.521236)
    Segment(va, vb).print(net_id, "Back")
    v1 = Vec2(175 + (175 - 142.379), 116.678)
    v2 = bar(v1.x, v1.y, -5, 0, -10)
    v3 = bar(v2.x, v2.y, 0, -17.78, -10)
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = bar(v3.x, v3.y, -1.27, 0, -10)
    v5 = Vec2(175 + (175 - 153.42), 90.17)
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Front")
    v6 = Vec2(177.54, 90.17)
    Segment(v5, v6).print(net_id, "Front")

    row = 4
    net_id = net_ids['N-row-4']
    #va = Vec2(206.004961, 136.307241)
    va = Vec2(207.255667, 136.086708)
    vb = Vec2(208.5334, 133.282236)
    Segment(va, vb).print(net_id, "Back")
    v2 = Vec2(207.255667, 136.086708)
    v3 = bar(v2.x, v2.y, 0, -17.78, -10)
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    v4 = bar(v3.x, v3.y, -2.54, 0, -10)
    v5 = bar(v4.x, v4.y, 0, -17.78, -10)
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Front")
    v6 = bar(v5.x, v5.y, -1.27, 0, -10)
    v7 = Vec2(175 + (175 - 154.486), 91.44)
    Segment(v5, v6).print(net_id, "Front")
    Segment(v6, v7).print(net_id, "Front")
    v8 = Vec2(180.08, 91.44)
    Segment(v7, v8).print(net_id, "Front")



def print_column_attachment_segments(switches):
    x = 7.7
    y1 = 9.525 - 1.27
    y2 = 9.525
    y3 = 9.525 + 1.27

    r = 1
    for c in [0,1,2,3,4,5]:
        s = switches[r][c]

        v1 = foo(s, -x, y1)
        v2 = foo(s, x, y1)
        v3 = foo(s, -x, y2)
        v4 = foo(s, x, y2)
        v5 = foo(s, -x, y3)
        v6 = foo(s, x, y3)

        if c > 5:
            Segment(v1, v2).print(net_ids['N-col-0'], "Back")
        if c > 4:
            Segment(v3, v4).print(net_ids['N-col-1'], "Back")
        if c > 3:
            Segment(v5, v6).print(net_ids['N-col-2'], "Back")

    for c in [7,8,9,10,11,12]:
        s = switches[r][c]

        v1 = foo(s, -x, y1)
        v2 = foo(s, x, y1)
        v3 = foo(s, -x, y2)
        v4 = foo(s, x, y2)
        v5 = foo(s, -x, y3)
        v6 = foo(s, x, y3)

        if c < 7:
            Segment(v1, v2).print(net_ids['N-col-7'], "Back")
        if c < 8:
            Segment(v3, v4).print(net_ids['N-col-8'], "Back")
        if c < 9:
            Segment(v5, v6).print(net_ids['N-col-9'], "Back")

    r = 2
    for c in [0,1,2,3,4,5]:
        s = switches[r][c]

        v1 = foo(s, -x, y1)
        v2 = foo(s, x, y1)
        v3 = foo(s, -x, y2)
        v4 = foo(s, x, y2)
        v5 = foo(s, -x, y3)
        v6 = foo(s, x, y3)

        if c > 2:
            Segment(v1, v2).print(net_ids['N-col-3'], "Back")
        if c > 1:
            Segment(v3, v4).print(net_ids['N-col-4'], "Back")
        if c > 0:
            Segment(v5, v6).print(net_ids['N-col-5'], "Back")

        
    for c in [7,8,9,10,11,12]:
        s = switches[r][c]

        v1 = foo(s, -x, y1)
        v2 = foo(s, x, y1)
        v3 = foo(s, -x, y2)
        v4 = foo(s, x, y2)
        v5 = foo(s, -x, y3)
        v6 = foo(s, x, y3)

        if c < 10:
            Segment(v1, v2).print(net_ids['N-col-10'], "Back")
        if c < 11:
            Segment(v3, v4).print(net_ids['N-col-11'], "Back")
        if c < 12:
            Segment(v5, v6).print(net_ids['N-col-12'], "Back")

    r = 1
    for c in [0,1,2,3,4]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]

        v1 = foo(s1, x, y1)
        v2 = foo(s2, -x, y1)
        v3 = foo(s1, x, y2)
        v4 = foo(s2, -x, y2)
        v5 = foo(s1, x, y3)
        v6 = foo(s2, -x, y3)

        if c >= 5:
            Segment(v1, v2).print(net_ids['N-col-0'], "Back")
        if c >= 4:
            Segment(v3, v4).print(net_ids['N-col-1'], "Back")
        if c >= 3:
            Segment(v5, v6).print(net_ids['N-col-2'], "Back")

    for c in [7,8,9,10,11]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]

        v1 = foo(s1, x, y1)
        v2 = foo(s2, -x, y1)
        v3 = foo(s1, x, y2)
        v4 = foo(s2, -x, y2)
        v5 = foo(s1, x, y3)
        v6 = foo(s2, -x, y3)

        if c < 7:
            Segment(v1, v2).print(net_ids['N-col-7'], "Back")
        if c < 8:
            Segment(v3, v4).print(net_ids['N-col-8'], "Back")
        if c < 9:
            Segment(v5, v6).print(net_ids['N-col-9'], "Back")

    r = 2
    for c in [0,1,2,3,4]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]

        v1 = foo(s1, x, y1)
        v2 = foo(s2, -x, y1)
        v3 = foo(s1, x, y2)
        v4 = foo(s2, -x, y2)
        v5 = foo(s1, x, y3)
        v6 = foo(s2, -x, y3)

        if c >= 2:
            Segment(v1, v2).print(net_ids['N-col-3'], "Back")
        if c >= 1:
            Segment(v3, v4).print(net_ids['N-col-4'], "Back")
        if c >= 0:
            Segment(v5, v6).print(net_ids['N-col-5'], "Back")

    for c in [7,8,9,10,11]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]

        v1 = foo(s1, x, y1)
        v2 = foo(s2, -x, y1)
        v3 = foo(s1, x, y2)
        v4 = foo(s2, -x, y2)
        v5 = foo(s1, x, y3)
        v6 = foo(s2, -x, y3)

        if c < 10:
            Segment(v1, v2).print(net_ids['N-col-10'], "Back")
        if c < 11:
            Segment(v3, v4).print(net_ids['N-col-11'], "Back")
        if c < 12:
            Segment(v5, v6).print(net_ids['N-col-12'], "Back")

    s = switches[1][3]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y3)
    v3 = foo(s, x, y3)
    Segment(v1, v2).print(net_ids['N-col-2'], "Back")
    Segment(v2, v3).print(net_ids['N-col-2'], "Back")

    s = switches[1][4]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y2)
    v3 = foo(s, x, y2)
    Segment(v1, v2).print(net_ids['N-col-1'], "Back")
    Segment(v2, v3).print(net_ids['N-col-1'], "Back")

    s = switches[1][5]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y1)
    v3 = foo(s, x, y1)
    Segment(v1, v2).print(net_ids['N-col-0'], "Back")
    Segment(v2, v3).print(net_ids['N-col-0'], "Back")

    s = switches[2][0]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y3)
    v3 = foo(s, x, y3)
    Segment(v1, v2).print(net_ids['N-col-5'], "Back")
    Segment(v2, v3).print(net_ids['N-col-5'], "Back")

    s = switches[2][1]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y2)
    v3 = foo(s, x, y2)
    Segment(v1, v2).print(net_ids['N-col-4'], "Back")
    Segment(v2, v3).print(net_ids['N-col-4'], "Back")

    s = switches[2][2]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y1)
    v3 = foo(s, x, y1)
    Segment(v1, v2).print(net_ids['N-col-3'], "Back")
    Segment(v2, v3).print(net_ids['N-col-3'], "Back")

    s = switches[1][7]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y1)
    v3 = foo(s, -x, y1)
    Segment(v1, v2).print(net_ids['N-col-7'], "Back")
    Segment(v2, v3).print(net_ids['N-col-7'], "Back")

    s = switches[1][8]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y2)
    v3 = foo(s, -x, y2)
    Segment(v1, v2).print(net_ids['N-col-8'], "Back")
    Segment(v2, v3).print(net_ids['N-col-8'], "Back")

    s = switches[1][9]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y3)
    v3 = foo(s, -x, y3)
    Segment(v1, v2).print(net_ids['N-col-9'], "Back")
    Segment(v2, v3).print(net_ids['N-col-9'], "Back")

    s = switches[2][10]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y1)
    v3 = foo(s, -x, y1)
    Segment(v1, v2).print(net_ids['N-col-10'], "Back")
    Segment(v2, v3).print(net_ids['N-col-10'], "Back")

    s = switches[2][11]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y2)
    v3 = foo(s, -x, y2)
    Segment(v1, v2).print(net_ids['N-col-11'], "Back")
    Segment(v2, v3).print(net_ids['N-col-11'], "Back")

    s = switches[2][12]
    v1 = foo(s, 3.81, 2.54)
    v2 = foo(s, 3.81, y3)
    v3 = foo(s, -x, y3)
    Segment(v1, v2).print(net_ids['N-col-12'], "Back")
    Segment(v2, v3).print(net_ids['N-col-12'], "Back")


    dx = 161.925
    y = 73.84 + 1.27
    v1 = Vec2(151.024, 82.739)
    v2 = Vec2(dx, 82.739)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx - 1.27, y - 1.27)
    v5 = Vec2(167.38, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-0'], "Back")
    Segment(v2, v3).print(net_ids['N-col-0'], "Back")
    Segment(v3, v4).print(net_ids['N-col-0'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-0'])
    Segment(v4, v5).print(net_ids['N-col-0'], "Front")

    dx += 0.635
    y -= 2.54
    v1 = Vec2(150.803, 83.989)
    v2 = Vec2(dx, 83.989)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx - 1.27, y - 1.27)
    v5 = Vec2(167.38, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-1'], "Back")
    Segment(v2, v3).print(net_ids['N-col-1'], "Back")
    Segment(v3, v4).print(net_ids['N-col-1'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-1'])
    Segment(v4, v5).print(net_ids['N-col-1'], "Front")

    dx += 0.635
    y -= 2.54
    v1 = Vec2(150.582, 85.24)
    v2 = Vec2(dx, 85.24)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx - 1.27, y - 1.27)
    v5 = Vec2(167.38, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-2'], "Back")
    Segment(v2, v3).print(net_ids['N-col-2'], "Back")
    Segment(v3, v4).print(net_ids['N-col-2'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-2'])
    Segment(v4, v5).print(net_ids['N-col-2'], "Front")

    dx += 0.635
    y -= 2.54
    v1 = Vec2(147.716, 101.499)
    v2 = Vec2(dx, 101.499)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx - 1.27, y - 1.27)
    v5 = Vec2(167.38, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-3'], "Back")
    Segment(v2, v3).print(net_ids['N-col-3'], "Back")
    Segment(v3, v4).print(net_ids['N-col-3'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-3'])
    Segment(v4, v5).print(net_ids['N-col-3'], "Front")

    dx += 0.635
    y -= 2.54
    v1 = Vec2(147.496, 102.749)
    v2 = Vec2(dx, 102.749)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx - 1.27, y - 1.27)
    v5 = Vec2(167.38, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-4'], "Back")
    Segment(v2, v3).print(net_ids['N-col-4'], "Back")
    Segment(v3, v4).print(net_ids['N-col-4'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-4'])
    Segment(v4, v5).print(net_ids['N-col-4'], "Front")

    dx += 0.635
    y -= 2.54
    v1 = Vec2(147.274, 104.000)
    v2 = Vec2(dx, 104.000)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx - 1.27, y - 1.27)
    v5 = Vec2(167.38, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-5'], "Back")
    Segment(v2, v3).print(net_ids['N-col-5'], "Back")
    Segment(v3, v4).print(net_ids['N-col-5'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-5'])
    Segment(v4, v5).print(net_ids['N-col-5'], "Front")


    dx = 188.075
    y = 73.84 + 1.27
    v1 = Vec2(198.976, 82.739)
    v2 = Vec2(dx, 82.739)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx + 1.27, y - 1.27)
    v5 = Vec2(182.62, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-7'], "Back")
    Segment(v2, v3).print(net_ids['N-col-7'], "Back")
    Segment(v3, v4).print(net_ids['N-col-7'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-7'])
    Segment(v4, v5).print(net_ids['N-col-7'], "Front")

    dx -= 0.635
    y -= 2.54
    v1 = Vec2(199.197, 83.989)
    v2 = Vec2(dx, 83.989)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx + 1.27, y - 1.27)
    v5 = Vec2(182.62, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-8'], "Back")
    Segment(v2, v3).print(net_ids['N-col-8'], "Back")
    Segment(v3, v4).print(net_ids['N-col-8'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-8'])
    Segment(v4, v5).print(net_ids['N-col-8'], "Front")

    dx -= 0.635
    y -= 2.54
    v1 = Vec2(199.418, 85.24)
    v2 = Vec2(dx, 85.24)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx + 1.27, y - 1.27)
    v5 = Vec2(182.62, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-9'], "Back")
    Segment(v2, v3).print(net_ids['N-col-9'], "Back")
    Segment(v3, v4).print(net_ids['N-col-9'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-9'])
    Segment(v4, v5).print(net_ids['N-col-9'], "Front")

    dx -= 0.635
    y -= 2.54
    v1 = Vec2(202.284, 101.499)
    v2 = Vec2(dx, 101.499)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx + 1.27, y - 1.27)
    v5 = Vec2(182.62, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-10'], "Back")
    Segment(v2, v3).print(net_ids['N-col-10'], "Back")
    Segment(v3, v4).print(net_ids['N-col-10'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-10'])
    Segment(v4, v5).print(net_ids['N-col-10'], "Front")

    dx -= 0.635
    y -= 2.54
    v1 = Vec2(202.505, 102.749)
    v2 = Vec2(dx, 102.749)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx + 1.27, y - 1.27)
    v5 = Vec2(182.62, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-11'], "Back")
    Segment(v2, v3).print(net_ids['N-col-11'], "Back")
    Segment(v3, v4).print(net_ids['N-col-11'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-11'])
    Segment(v4, v5).print(net_ids['N-col-11'], "Front")

    dx -= 0.635
    y -= 2.54
    v1 = Vec2(202.726, 104.000)
    v2 = Vec2(dx, 104.000)
    v3 = Vec2(dx, y)
    v4 = Vec2(dx + 1.27, y - 1.27)
    v5 = Vec2(182.62, y - 1.27)
    Segment(v1, v2).print(net_ids['N-col-12'], "Back")
    Segment(v2, v3).print(net_ids['N-col-12'], "Back")
    Segment(v3, v4).print(net_ids['N-col-12'], "Back")
    print_via(v4.x, v4.y, net_ids['N-col-12'])
    Segment(v4, v5).print(net_ids['N-col-12'], "Front")


def print_mosfet_attachment_segments(switches):
    r = 4
    for c in [1,2,3,4,5,7,8,9,10,11,12]:
        s = switches[r][c]
        v1 = foo(s, -8.335, 7.7)
        v2 = foo(s, 8.335, 7.7)
        Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
        s = switches[r][c]
        v1 = foo(s, 8.335, 7.7)
        v2 = foo(s, 8.335, -3.18)
        Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    for c in [0,1,2,3,4]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]
        v1 = foo(s1, 8.335, 7.7)
        v2 = foo(s2, -8.335, 7.7)
        Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    for c in [7,8,9,10,11]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]
        v1 = foo(s1, 8.335, 7.7)
        v2 = foo(s2, -8.335, 7.7)
        Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    for r in [1,2]:
        s = switches[r][6]
        v1 = foo(s, -8.335, 7.7)
        v2 = foo(s, 8.335, 7.7)
        Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    s1 = switches[4][5]
    s2 = switches[1][6]
    v1 = foo(s1, 8.335, 7.7)
    v2 = foo(s2, -8.335, 7.7)
    Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    s1 = switches[1][6]
    s2 = switches[2][6]
    v1 = foo(s1, 8.335, 7.7)
    v2 = foo(s2, -8.335, 7.7)
    Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    s1 = switches[2][6]
    s2 = switches[4][7]
    v1 = foo(s1, 8.335, 7.7)
    v2 = foo(s2, -8.335, 7.7)
    Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    v1 = Vec2(162.471574, 131.177669)
    v2 = Vec2(160.582282, 141.892377)
    Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    v1 = Vec2(203.945171, 128.282954)
    v2 = Vec2(205.834464, 138.997662)
    Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")

    v1 = Vec2(183.335, 102.82)
    v2 = Vec2(183.335, 141.892377)
    Segment(v1, v2).print(net_ids['N-MOSFET-0'], "Front")



def print_vcc_attachment_segments(switches):
    r = 0
    for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
        s = switches[r][c]
        if c == 0 or c == 7:
            v1 = foo(s, -7.065, -7.065)
        else:
            v1 = foo(s, -7.7, -7.065)
        if c == 12:
            v2 = foo(s, -7.065, -7.065)
        else:
            v2 = foo(s, 7.7, -7.065)
        Segment(v1, v2).print(net_ids['N-5V-0'], "Front")

    for c in [0,1,2,3,4,5,7,8,9,10,11,12]:
        s = switches[r][c]
        v1 = foo(s, -7.065, -3.81)
        v2 = foo(s, -7.065, -7.065)
        Segment(v1, v2).print(net_ids['N-5V-0'], "Front")

    for c in [0,1,2,3,4]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]
        v1 = foo(s1, 7.7, -7.065)
        v2 = foo(s2, -7.7, -7.065)
        Segment(v1, v2).print(net_ids['N-5V-0'], "Front")

    for c in [7,8,9,10,11]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]
        v1 = foo(s1, 7.7, -7.065)
        v2 = foo(s2, -7.7, -7.065)
        Segment(v1, v2).print(net_ids['N-5V-0'], "Front")

    v1 = Vec2(167.935, 102.19)
    v2 = Vec2(167.935, 123.825)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Front")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    s = switches[1][6]
    v1 = foo(s, -1.27, -5.08)
    v2 = Vec2(154.015, 123.825)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Front")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    s = switches[2][6]
    v1 = foo(s, -1.27, -5.08)
    v2 = Vec2(193.406, 123.825)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Front")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    v1 = Vec2(154.015, 123.825)
    v2 = Vec2(193.406, 123.825)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")

    v1 = Vec2(193.406, 123.825)
    v2 = Vec2(203.2, 123.825)
    v3 = Vec2(206.725, 123.026)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    Segment(v2, v3).print(net_ids['N-5V-0'], "Back")

    # RGB 1
    v1 = Vec2(74.657058, 46.496823)
    #v2 = Vec2(84.096, 48.161)
    v2 = Vec2(84.595, 45.329)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 2
    v1 = Vec2(59.432942, 105.195177)
    v2 = Vec2(54.839, 104.385)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 3
    v1 = Vec2(113.567058, 45.233823)
    v2 = Vec2(122.507, 49.73)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 4
    v1 = Vec2(98.342942, 103.933177)
    v2 = Vec2(93.749, 103.123)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 5
    v1 = Vec2(149.178058, 62.683823)
    v2 = bar(v1.x, v1.y, 2.54, 0, 10)
    v3 = Vec2(139.857, 61.04)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    Segment(v2, v3).print(net_ids['N-5V-0'], "Front")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 6
    v1 = Vec2(133.953942, 121.382177)
    v2 = Vec2(129.36, 120.572)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 7
    v1 = Vec2(205.549019, 61.850312)
    v2 = Vec2(214.156, 55.466)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 8
    v1 = Vec2(211.318981, 122.215688)
    v2 = Vec2(206.725, 123.026)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 9
    v1 = Vec2(241.160019, 44.400312)
    v2 = Vec2(250.932, 44.624)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 10
    v1 = Vec2(246.929981, 104.766688)
    v2 = Vec2(242.336, 105.577)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 11
    v1 = Vec2(280.070019, 45.663312)
    v2 = Vec2(289.51, 43.999)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # RGB 12
    v1 = Vec2(285.839981, 106.028688)
    v2 = Vec2(281.246, 106.839)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Back")
    print_via(v2.x, v2.y, net_ids['N-5V-0'])

    # Attach left half to controller VCC pin
    v1 = Vec2(156.991844, 48.890424)
    v2 = Vec2(167.38, 58.6)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Front")

    # Attach right half to controller VCC pin
    v1 = Vec2(194.198734, 51.985707)
    v2 = Vec2(189.23, 48.26)
    v3 = Vec2(180.34, 48.26)
    v4 = Vec2(170.18, 58.42)
    v5 = Vec2(167.38, 58.6)
    Segment(v1, v2).print(net_ids['N-5V-0'], "Front")
    Segment(v2, v3).print(net_ids['N-5V-0'], "Front")
    Segment(v3, v4).print(net_ids['N-5V-0'], "Front")
    Segment(v4, v5).print(net_ids['N-5V-0'], "Front")


def print_gnd_attachment_segments(switches):
    net_id = net_ids['N-GND-0']
    r = 0
    y = -8.335
    for c in [1,2,3,4,5,7,8,9,10,11]:
        s = switches[r][c]
        v1 = foo(s, -7.7, y)
        if c == 1:
            v1 = foo(s, -8.335, y)
        v2 = foo(s, 7.7, y)
        Segment(v1, v2).print(net_id, "Back")

    for c in [1,2,3,4]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]
        v1 = foo(s1, 7.7, y)
        v2 = foo(s2, -7.7, y)
        Segment(v1, v2).print(net_id, "Back")

    for c in [7,8,9,10,11]:
        s1 = switches[r][c]
        s2 = switches[r][c+1]
        v1 = foo(s1, 7.7, y)
        v2 = foo(s2, -7.7, y)
        if c == 11:
            v2 = Vec2(285.444, 28.256)
        Segment(v1, v2).print(net_id, "Back")

    # RGB 1
    c = 1
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = foo(s, -8.335, y)
    v3 = Vec2(63.512, 47.883)
    v4 = Vec2(69.356942, 48.913177)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 2
    c = 2
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(86.09, 29.536)
    v3 = Vec2(73.242, 102.403)
    v4 = Vec2(64.733058, 102.778823)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 3
    c = 3
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(105.659, 28.263)
    v3 = Vec2(102.422, 46.62)
    v4 = Vec2(108.266942, 47.650177)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 4
    c = 4
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(123.71, 35.595)
    v3 = Vec2(111.832, 102.961)
    # Hacky workaround for trace collision
    v4 = Vec2(111.125, 104.14)
    v5 = Vec2(103.643058, 101.516823)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v4.x, v4.y, net_id)

    # RGB 5
    c = 5
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(141.572, 43.999)
    v3 = Vec2(138.033, 64.07)
    v4 = Vec2(143.877942, 65.100177)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 6
    # We have to do some hacky stuff here because we can't run ground down
    # column 6
    v1 = Vec2(138.033, 64.07)
    v2 = Vec2(128.396, 118.727)
    v3 = bar(v2.x, v2.y, 2.54, 0, 10)
    v4 = bar(v3.x, v3.y, 11, 0, 10)
    v5 = Vec2(139.254058, 118.965823)
    #v5 = Vec2(TODO, TODO)
    #v2 = Vec2(136.604, 120.174)
    #v3 = bar(136.604, 120.174, -10, 0, 10)
    Segment(v1, v2).print(net_id, "Front")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Back")
    print_via(v3.x, v3.y, net_id)
    Segment(v3, v4).print(net_id, "Front")
    print_via(v4.x, v4.y, net_id)
    Segment(v4, v5).print(net_id, "Back")
    #Segment(v3, v5).print(net_id, "Front")
    #v6 = bar(v5.x, v5.y, 13.5, 0, 10)
    #Segment(v5, v6).print(net_id, "Front")
    #print_via(v6.x, v6.y, net_id)
    #v7 = Vec2(139.254058, 118.965823)
    #Segment(v6, v7).print(net_id, "Back")

    # RGB 7
    c = 7
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(192.205, 47.995)
    v3 = Vec2(195.55, 66.964)
    v4 = Vec2(201.394981, 65.933688)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 8
    c = 8
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(210.205, 40.375)
    v3 = Vec2(223.084, 113.413)
    v4 = Vec2(215.473019, 118.132312)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 9
    c = 9
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(227.955, 31.329)
    v3 = Vec2(231.367, 50.683)
    v4 = Vec2(237.005981, 48.483688)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 10
    c = 10
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(247.251, 31.06)
    v3 = Vec2(259.273, 99.239)
    v4 = Vec2(258.445, 100.33)
    v5 = Vec2(251.084019, 100.683312)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Front")
    Segment(v4, v5).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v4.x, v4.y, net_id)

    # RGB 11
    c = 11
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(266.593, 31.05)
    v3 = Vec2(270.071, 50.777)
    v4 = Vec2(275.915981, 49.746688)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # RGB 11
    c = 12
    s = switches[r][c]
    v1 = foo(s, -8.335, 0)
    v2 = Vec2(285.444, 28.256)
    v3 = Vec2(298.183, 100.501)
    v4 = Vec2(289.994019, 101.945312)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v1, v3).print(net_id, "Front")
    Segment(v3, v4).print(net_id, "Back")
    print_via(v1.x, v1.y, net_id)
    print_via(v3.x, v3.y, net_id)

    # Attach left half to controller GND pin
    v1 = Vec2(157.212377, 47.639718)
    v2 = Vec2(158.23, 48.26)
    v3 = Vec2(167.38, 53.52)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v2, v3).print(net_id, "Back")

    # Attach right half to controller GND pin
    v1 = Vec2(192.787623, 47.639718)
    v2 = Vec2(191.77, 48.26)
    v3 = Vec2(182.62, 56.06)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v2, v3).print(net_id, "Back")

    # Attach controller GND pins together
    v1 = Vec2(175, 53.52)
    v2 = Vec2(175, 58.6)
    Segment(v1, v2).print(net_id, "Back")
    v1 = Vec2(167.38, 53.52)
    v2 = Vec2(175, 53.52)
    Segment(v1, v2).print(net_id, "Back")
    v1 = Vec2(182.62, 56.06)
    v2 = Vec2(175, 56.06)
    Segment(v1, v2).print(net_id, "Back")
    v1 = Vec2(182.62, 58.6)
    v2 = Vec2(175, 58.6)
    Segment(v1, v2).print(net_id, "Back")


def print_rgb_data_attachment_segments(switches):
    net_id = net_ids['N-RGB-D0']
    v1 = Vec2(182.3914, 78.92)
    v2 = Vec2(191.958672, 78.92)
    v3 = Vec2(191.958672, 64.246657)
    v4 = Vec2(200.821942, 62.683823)
    Segment(v1, v2).print(net_id, "Front")
    Segment(v2, v3).print(net_id, "Front")
    print_via(v3.x, v3.y, net_id)
    Segment(v3, v4).print(net_id, "Back")

    net_id = net_ids['N-RGB-D1']
    v1 = Vec2(206.122058, 65.100177)
    v2 = Vec2(236.432942, 45.233823)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D2']
    v1 = Vec2(241.733058, 47.650177)
    v2 = Vec2(275.342942, 46.496823)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D3']
    v1 = Vec2(69.929981, 45.663312)
    v4 = Vec2(60.005981, 101.945312)
    v2 = bar(v1.x, v1.y, -7.125, 0, 10)
    v3 = bar(v4.x, v4.y, -7.125, 0, 10)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    print_via(v3.x, v3.y, net_id)
    Segment(v3, v4).print(net_id, "Back")

    net_id = net_ids['N-RGB-D4']
    v1 = Vec2(251.657058, 103.933177)
    v2 = Vec2(285.266942, 102.778823)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D5']
    v1 = Vec2(216.046058, 121.382177)
    v2 = Vec2(246.356942, 101.516823)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D6']
    v1 = Vec2(138.681019, 122.215688)
    v2 = bar(v1.x, v1.y, 5, 0, 10)
    v3 = Vec2(151.765, 122.555)
    v4 = Vec2(198.235, 122.555)
    v6 = Vec2(210.745942, 118.965823)
    v5 = bar(v6.x, v6.y, -5, 0, 10)
    Segment(v1, v2).print(net_id, "Back")
    Segment(v2, v3).print(net_id, "Back")
    Segment(v3, v4).print(net_id, "Back")
    Segment(v4, v5).print(net_id, "Back")
    Segment(v5, v6).print(net_id, "Back")

    net_id = net_ids['N-RGB-D7']
    v1 = Vec2(103.070019, 104.766688)
    v2 = Vec2(134.526981, 118.132312)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D8']
    v1 = Vec2(64.160019, 106.028688)
    v2 = Vec2(98.915981, 100.683312)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D9']
    v1 = Vec2(280.643058, 48.913177)
    v4 = Vec2(290.567058, 105.195177)
    v2 = bar(v1.x, v1.y, 7.125, 0, -10)
    v3 = bar(v4.x, v4.y, 7.125, 0, -10)
    Segment(v1, v2).print(net_id, "Back")
    print_via(v2.x, v2.y, net_id)
    Segment(v2, v3).print(net_id, "Front")
    print_via(v3.x, v3.y, net_id)
    Segment(v3, v4).print(net_id, "Back")

    net_id = net_ids['N-RGB-D10']
    v1 = Vec2(74.084019, 49.746688)
    v2 = Vec2(108.839981, 44.400312)
    Segment(v1, v2).print(net_id, "Back")

    net_id = net_ids['N-RGB-D11']
    v1 = Vec2(112.994019, 48.483688)
    v2 = Vec2(144.450981, 61.850312)
    Segment(v1, v2).print(net_id, "Back")


def main():
    raw_switch_data = []
    with open('switches.json', 'r') as f:
        raw_switch_data = json.loads(f.read())

    switches = {}
    for switch in raw_switch_data:
        row = switch['row']
        column = switch['column']

        if not row in switches:
            switches[row] = {}
        switches[row][column] = switch

    column_segments = calculate_column_segments(switches)
    print_column_segments(column_segments)
    print_switch_to_column_segments(switches)

    print_led_to_resistor_segments(switches)

    print_switch_to_diode_segments(switches)

    mosfet_column_segments = calculate_mosfet_column_segments(switches)
    print_mosfet_column_segments(mosfet_column_segments)
    print_switch_to_mosfet_column_segments(switches)
    print_switch_to_mosfet_column_vias(switches)

    vcc_column_segments = calculate_vcc_column_segments(switches)
    print_vcc_column_segments(vcc_column_segments)
    print_switch_to_vcc_column_segments(switches)

    print_diode_to_row_components(switches)
    print_row_attachment_segments(switches)

    print_column_attachment_segments(switches)

    print_mosfet_attachment_segments(switches)
    print_vcc_attachment_segments(switches)
    print_gnd_attachment_segments(switches)

    print_rgb_data_attachment_segments(switches)



if __name__ == '__main__':
    main()

