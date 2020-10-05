#!/usr/bin/env python3

# This script should be deleted once it has outlived its usefulness. It's
# purpose was to translate the coordinates for the points around the keywells
# and case edges.

def foo():

    min_x = 999999
    max_x = -999999
    min_y = 999999
    max_y = -999999


    xs = [ 32.80, 71.95, 91, 109, 128.05, 167.2, 148.15, 129.1, 110.05, 89.95, 70.9 ]
    max_x = max(xs)
    min_x = min(xs)

    ys = [ 140.175, 137.175, 132.175, 138.175, 143.175, 65.925, 27.825, 22.825, 16.825, 21.825, 24.825 ]
    max_y = max(ys)
    min_y = min(ys)

    w = max_x - min_x
    h = max_y - min_y

    mid_x = min_x + (w / 2)
    mid_y = min_y + (h / 2)

    for x in xs:
        print('<Constant Name="XTODO"  Value="{}" />'.format(round(x - min_x - (w/2), 3)))
        
    for y in ys:
        print('<Constant Name="YTODO"  Value="{}" />'.format(round(y - min_y - (h/2), 3)))

    print('min_x = {}, max_x = {}, min_y = {}, max_y = {}'.format(min_x, max_x, min_y, max_y))
    print('mid_x = {}, mid_y = {}'.format(round(mid_x, 3), round(mid_y, 3)))

def bar():
    dx = 0
    print('<Constant Name="X01"  Value="{}" />'.format(32.80 + dx))
    print('<Constant Name="X02"  Value="{}" />'.format(71.95 + dx))
    print('<Constant Name="X03"  Value="{}" />'.format(91 + dx))
    print('<Constant Name="X04"  Value="{}" />'.format(109 + dx))
    print('<Constant Name="X05"  Value="{}" />'.format(128.05 + dx))
    print('<Constant Name="X06"  Value="{}" />'.format(167.2 + dx))
    print('<Constant Name="X07"  Value="{}" />'.format(148.15 + dx))
    print('<Constant Name="X08"  Value="{}" />'.format(129.1 + dx))
    print('<Constant Name="X09"  Value="{}" />'.format(110.05 + dx))
    print('<Constant Name="X10"  Value="{}" />'.format(89.95 + dx))
    print('<Constant Name="X11"  Value="{}" />'.format(70.9 + dx))

    dy = -80
    print('<Constant Name="Y1"  Value="{}" />'.format(140.175 + dy))
    print('<Constant Name="Y2"  Value="{}" />'.format(137.175 + dy))
    print('<Constant Name="Y3"  Value="{}" />'.format(132.175 + dy))
    print('<Constant Name="Y4"  Value="{}" />'.format(138.175 + dy))
    print('<Constant Name="Y5"  Value="{}" />'.format(143.175 + dy))
    print('<Constant Name="Y6"  Value="{}" />'.format(65.925 + dy))
    print('<Constant Name="Y7"  Value="{}" />'.format(27.825 + dy))
    print('<Constant Name="Y8"  Value="{}" />'.format(22.825 + dy))
    print('<Constant Name="Y9"  Value="{}" />'.format(16.825 + dy))
    print('<Constant Name="Y10"  Value="{}" />'.format(21.825 + dy))
    print('<Constant Name="Y11"  Value="{}" />'.format(24.825 + dy))

def main():
    #bar()
    foo()

if __name__ == '__main__':
    main()

