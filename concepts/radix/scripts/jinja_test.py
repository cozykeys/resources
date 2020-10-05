#!/usr/bin/env python3

from jinja2 import Environment, FileSystemLoader

def main():
    env = Environment(loader=FileSystemLoader('./templates'))

    switches = [
        { 'x': 10, 'y': 10, 'rotation': 10 },
        { 'x': 20, 'y': 20, 'rotation': 20 },
        { 'x': 30, 'y': 30, 'rotation': 30 }
    ]

    print(env.get_template('switch_footprint.j2').render(switches=switches))

if __name__ == '__main__':
    main()
