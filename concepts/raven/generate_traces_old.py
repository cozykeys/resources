#!/usr/bin/env python3

# A lot of this is ridiculously hacky and makes major assumptions but it's
# still better than drawing traces by hand. :)

import json
import math

# Lookup table for net ids
net_ids = {
    "N-row-0": 1, "N-row-1": 2, "N-row-2": 3, "N-row-3": 4, "N-row-4": 5,

    "N-col-0":   6, "N-col-1":   7, "N-col-2":   8, "N-col-3":  9, "N-col-4": 10,
    "N-col-5":  11, "N-col-6":  12, "N-col-7":  13, "N-col-8": 14, "N-col-9": 15,
    "N-col-10": 16, "N-col-11": 17, "N-col-12": 18,

    "N-diode-0-0":  19, "N-diode-0-1":  20, "N-diode-0-2":  21, "N-diode-0-3":  22,
    "N-diode-0-4":  23, "N-diode-0-5":  24, "N-diode-0-6":  25, "N-diode-0-7":  26,
    "N-diode-0-8":  27, "N-diode-0-9":  28, "N-diode-0-10": 29, "N-diode-0-11": 30,
    "N-diode-0-12": 31, "N-diode-1-0":  32, "N-diode-1-1":  33, "N-diode-1-2":  34,
    "N-diode-1-3":  35, "N-diode-1-4":  36, "N-diode-1-5":  37, "N-diode-1-6":  38,
    "N-diode-1-7":  39, "N-diode-1-8":  40, "N-diode-1-9":  41, "N-diode-1-10": 42,
    "N-diode-1-11": 43, "N-diode-1-12": 44, "N-diode-2-0":  45, "N-diode-2-1":  46,
    "N-diode-2-2":  47, "N-diode-2-3":  48, "N-diode-2-4":  49, "N-diode-2-5":  50,
    "N-diode-2-6":  51, "N-diode-2-7":  52, "N-diode-2-8":  53, "N-diode-2-9":  54,
    "N-diode-2-10": 55, "N-diode-2-11": 56, "N-diode-2-12": 57, "N-diode-3-0":  58,
    "N-diode-3-1":  59, "N-diode-3-2":  60, "N-diode-3-3":  61, "N-diode-3-4":  62,
    "N-diode-3-5":  63, "N-diode-3-7":  64, "N-diode-3-8":  65, "N-diode-3-9":  66,
    "N-diode-3-10": 67, "N-diode-3-11": 68, "N-diode-3-12": 69, "N-diode-4-0":  70,
    "N-diode-4-1":  71, "N-diode-4-2":  72, "N-diode-4-3":  73, "N-diode-4-4":  74,
    "N-diode-4-5":  75, "N-diode-4-7":  76, "N-diode-4-8":  77, "N-diode-4-9":  78,
    "N-diode-4-10": 79, "N-diode-4-11": 80, "N-diode-4-12": 81,
}

column_segments = {}

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

class vec2:
    def __init__(self, x = 0, y = 0):
        self.x = x
        self.y = y
    
    def __str__(self):
        return "({0}, {1})".format(self.x, self.y)

class segment:
    def __init__(self, start = vec2(), end = vec2()):
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


def print_switch_to_diode_segment(key):
    SWITCH_PAD_DX    = 2.54
    SWITCH_PAD_DY    = 5.08
    SWITCH_PAD_MAG   = math.sqrt(math.pow(SWITCH_PAD_DX, 2) + math.pow(SWITCH_PAD_DY, 2))
    SWITCH_PAD_THETA = math.asin(SWITCH_PAD_DY / SWITCH_PAD_MAG)

    position = vec2(key['x'], key['y'])
    rotation = math.radians(-key['rotation'])
    net_id = net_ids['N-diode-{0}-{1}'.format(key['row'], key['column'])]

    start_x = position.x + (SWITCH_PAD_MAG * math.cos(SWITCH_PAD_THETA + rotation))
    start_y = position.y - (SWITCH_PAD_MAG * math.sin(SWITCH_PAD_THETA + rotation))

    mid_dx    = 7.0
    mid_dy    = 5.08
    mid_mag   = math.sqrt(math.pow(mid_dx, 2) + math.pow(mid_dy, 2))
    mid_theta = 0.0
    if key['diode_position'] == "left":
        mid_theta = math.pi - math.asin(mid_dy / mid_mag)
    elif key['diode_position'] == "right":
        mid_theta = math.asin(mid_dy / mid_mag)
    mid_x = position.x + (mid_mag * math.cos(mid_theta + rotation))
    mid_y = position.y - (mid_mag * math.sin(mid_theta + rotation))

    end_x = start_x
    end_y = start_y
    if key['diode_position'] == "left":
        theta = math.radians(100)
        end_x = key['x'] + (8.0 * math.cos(math.radians(190))) + (3.81 * math.cos(theta))
        end_y = key['y'] - (8.0 * math.sin(math.radians(190))) - (3.81 * math.sin(theta))
    elif key['diode_position'] == "right":
        if rotation == 0.0:
            theta = math.radians(90)
            end_x = key['x'] + 9 + (3.81 * math.cos(theta))
            end_y = key['y'] - (3.81 * math.sin(theta))
        else:
            theta = math.radians(80)
            end_x = key['x'] + (8.0 * math.cos(math.radians(350))) + (3.81 * math.cos(theta))
            end_y = key['y'] - (8.0 * math.sin(math.radians(350))) - (3.81 * math.sin(theta))
    else:
        return

    print_segment(start_x, start_y, mid_x, mid_y, net_id, "Back")
    print_segment(mid_x, mid_y, end_x, end_y, net_id, "Back")

def print_switch_to_column_segment(key):
    SWITCH_PAD_DX    = -3.81
    SWITCH_PAD_DY    = -2.54
    SWITCH_PAD_MAG   = math.sqrt(math.pow(SWITCH_PAD_DX, 2) + math.pow(SWITCH_PAD_DY, 2))
    SWITCH_PAD_THETA = math.asin(SWITCH_PAD_DY / SWITCH_PAD_MAG) + math.pi

    position = vec2(key['x'], key['y'])
    rotation = math.radians(-key['rotation'])
    net_id = net_ids['N-col-{0}'.format(key['column'])]

    start_x = position.x + (SWITCH_PAD_MAG * math.cos(SWITCH_PAD_THETA + rotation))
    start_y = position.y - (SWITCH_PAD_MAG * math.sin(SWITCH_PAD_THETA + rotation))

    SEGMENT_END_DX    = -3.81 + 0.6985
    SEGMENT_END_DY    = -2.54
    SEGMENT_END_MAG   = math.sqrt(math.pow(SEGMENT_END_DX, 2) + math.pow(SEGMENT_END_DY, 2))
    SEGMENT_END_THETA = math.asin(SEGMENT_END_DY / SEGMENT_END_MAG) + math.pi

    end_x = position.x + (SEGMENT_END_MAG * math.cos(SEGMENT_END_THETA + rotation))
    end_y = position.y - (SEGMENT_END_MAG * math.sin(SEGMENT_END_THETA + rotation))

    global column_segments
    if key['column'] != 6:
        if key['column'] not in column_segments:
            column_segments[key['column']] = segment()
        if key['row'] == 0:
            column_segments[key['column']].start = vec2(end_x, end_y)
        if key['row'] == 4:
            column_segments[key['column']].end = vec2(end_x, end_y)

    print_segment(start_x, start_y, end_x, end_y, net_id)


def get_bottom_diode_pad(key):
    if key['diode_position'] == "left":
        theta = math.radians(280)
        return vec2(
            key['x'] + (8.0 * math.cos(math.radians(190))) + (3.81 * math.cos(theta)),
            key['y'] - (8.0 * math.sin(math.radians(190))) - (3.81 * math.sin(theta)))
            #key['x'] - 9 + (3.81 * math.cos(theta)),
            #key['y'] - (3.81 * math.sin(theta)))

    if key['diode_position'] == "right":
        theta = math.radians(260)
        return vec2(
            key['x'] + (8.0 * math.cos(math.radians(350))) + (3.81 * math.cos(theta)),
            key['y'] - (8.0 * math.sin(math.radians(350))) - (3.81 * math.sin(theta)))
            #key['x'] + 9 + (3.81 * math.cos(theta)),
            #key['y'] - (3.81 * math.sin(theta)))

    raise 'Not supported exception'


def get_attachment_point(attachment, key, rotation):
    dx = -6.35 if attachment == 'left' else 6.35
    dy = -4
    m = math.sqrt(math.pow(dx, 2) + math.pow(dy, 2))
    switch_pos = vec2(key['x'], key['y'])
    attach_pos = vec2(key['x'] + dx, key['y'] + dy)

    theta = segment(switch_pos, attach_pos).theta() + math.radians(rotation)

    return vec2(
        switch_pos.x + (m * math.cos(theta)),
        switch_pos.y - (m * math.sin(theta)))

def print_switch_pair_row_segments_left(lhs, rhs):
    ldp = get_bottom_diode_pad(lhs)
    rdp = get_bottom_diode_pad(rhs)
    lap = get_attachment_point('left', rhs, -10)
    rap = get_attachment_point('right', rhs, -10)

    net_id = net_ids["N-row-{}".format(lhs['row'])]
    print_segment(ldp.x, ldp.y, lap.x, lap.y, net_id, "Back")
    print_segment(lap.x, lap.y, rap.x, rap.y, net_id, "Back")
    print_segment(rap.x, rap.y, rdp.x, rdp.y, net_id, "Back")

def print_switch_pair_row_segments_right(lhs, rhs):
    ldp = get_bottom_diode_pad(lhs)
    rdp = get_bottom_diode_pad(rhs)
    lap = get_attachment_point('left', rhs, 10)
    rap = get_attachment_point('right', rhs, 10)

    net_id = net_ids["N-row-{}".format(lhs['row'])]
    print_segment(ldp.x, ldp.y, rap.x, rap.y, net_id, "Back")
    print_segment(rap.x, rap.y, lap.x, lap.y, net_id, "Back")
    print_segment(lap.x, lap.y, rdp.x, rdp.y, net_id, "Back")

def print_column_segments():
    print('/* Column Segments */')

    global column_segments
    for col in column_segments:
        s = column_segments[col]
        net_id = net_ids['N-col-{0}'.format(col)]
        print_segment(s.start.x, s.start.y, s.end.x, s.end.y, net_id)

def theta_test(actual, expected):
    if actual - expected > 0.01:
        print('FAILURE')
    else:
        print('SUCCESS')


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

        print_switch_to_diode_segment(switch)
        print_switch_to_column_segment(switch)

    print_column_segments()

    for r in range(0,5):
        for c in range(0,5):
            print_switch_pair_row_segments_left(switches[r][c], switches[r][c+1])
    #print_switch_pair_row_segments_left(switches[4][5], switches[0][6])

    for r in range(0,5):
        for c in range(12,7,-1):
            print_switch_pair_row_segments_right(switches[r][c], switches[r][c-1])
    #print_switch_pair_row_segments_right(switches[4][7], switches[1][6])


if __name__ == '__main__':
    main()

