#!/usr/bin/env python3

COLS=13
ROWS=5

def print_keysw_component(uid, index, x, y):
    print('$Comp')
    print('U 1 1 {}'.format(uid))
    print('L keyboard_parts:KEYSW K{}'.format(index))
    print('P {} {}'.format(x, y))
    print('F 0 "K{}" H {} {} 60  0000 C CNN'.format(index, x, y+233))
    print('F 1 "KEYSW" H {} {} 60  0001 C CNN'.format(x, y-100))
    print('F 2 "" H {} {} 60  0000 C CNN'.format(x,y))
    print('F 3 "" H {} {} 60  0000 C CNN'.format(x,y))
    print('\t1    {} {}'.format(x,y))
    print('\t1    0    0    -1  ')
    print('$EndComp')

def print_diode_component(uid, index, x, y):
    print('$Comp')
    print('L Device:D D{}'.format(index))
    print('U 1 1 {}'.format(uid))
    print('P {} {}'.format(x, y))
    print('F 0 "D{}" V {} {} 50  0000 R CNN'.format(index,x+46,y-79))
    print('F 1 "D" V {} {} 50  0000 R CNN'.format(x-45,y-79))
    print('F 2 "" H {} {} 50  0001 C CNN'.format(x,y))
    print('F 3 "~" H {} {} 50  0001 C CNN'.format(x,y))
    print('\t1    {} {}'.format(x,y))
    print('\t0    -1   -1   0   ')
    print('$EndComp')

def print_wires(x, y, last_row, last_col):
    print('Wire Wire Line')
    print('\t{} {} {} {}'.format(x-350,y+50,x-350,y))
    print('Wire Wire Line')
    print('\t{} {} {} {}'.format(x-350,y,x-300,y))

    if not last_col:
        print('Wire Wire Line')
        print('\t{} {} {} {}'.format(x-350,y+350,x+450,y+350))

    if not last_row:
        print('Wire Wire Line')
        print('\t{} {} {} {}'.format(x+300,y,x+300,y+650))

def print_row_label(x,y,row):
    print('Text Label {} {} 0    50   ~ 0'.format(x-550,y+350))
    print('row{}'.format(row))
    print('Wire Wire Line')
    print('\t{} {} {} {}'.format(x-550,y+350,x-350,y+350))

def print_col_label(x,y,col):
    print('Text Label {} {} 0    50   ~ 0'.format(x+300,y-200))
    print('col{}'.format(col))
    print('Wire Wire Line')
    print('\t{} {} {} {}'.format(x+300,y-200,x+300,y))

def main():
    index=1
    for c in range(0,COLS):
        x = 6550 + (c * 800)
        for r in range(0,ROWS):
            last_col = (c == COLS-1) or (c == 5 and r <2)
            last_row = r == ROWS-1
            if c == 6 and r < 2:
                continue
            y = 1900 + (r * 650)
            keysw_uid = '5D514A{0:02d}'.format(index*2-1)
            diode_uid = '5D514A{0:02d}'.format(index*2)
            print_keysw_component(keysw_uid, index, x, y)
            print_diode_component(diode_uid, index, x-350, y+200)
            print_wires(x,y,last_row,last_col)
            if c == 0:
                print_row_label(x,y,r)
            if r == 0 or (c == 6 and r == 2):
                print_col_label(x,y,c)
            index += 1

if __name__ == '__main__':
    main()

