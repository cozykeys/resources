#!/usr/bin/env python3

import math

def pos(x,y): return (x + (18.6 / 2), y + (18.6 / 2))

def foo(x_direction, y_direction, x, y):
    #delta = math.sin(math.pi/4)*10 # 45*
    delta = 10
    w = 18.6
    h = 18.6
    dx = 0
    dy = 0
    if x_direction == 'right':
        dx = x + (18.1 / 2) + 1 + delta
    else:
        dx = x - (18.1 / 2) - 1 - delta
    if y_direction == 'up':
        dy = y + (18.1 / 2) + 1 + delta
    else:
        dy = y - (18.1 / 2) - 1 - delta
    return (dx,dy)


def main():
    delta = math.sin(math.pi/4)*10 # 45*

    (x,y) = pos(284.154, 157.193)
    print((x,y))
    print((x-20.05,y+20.05))
    print('{}, adjusted = {}'.format(
        foo('left', 'up', x, y),
        foo('left', 'up', x + 0.5, y + 0.5)))

if __name__ == '__main__':
    main()

