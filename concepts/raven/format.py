#!/usr/bin/env python3

def fmt(n):
    return '{}'.format(round(n,3))

def main():
    mt = '<AbsoluteMoveTo> <EndPoint X="{}" Y="{}" /> </AbsoluteMoveTo>'
    lt = '<AbsoluteLineTo> <EndPoint X="{}" Y="{}" /> </AbsoluteLineTo>'

    print(mt.format(fmt(0), fmt(-33.83)))
    print(lt.format(fmt(16.102172463211986), fmt(-33.83)))
    print(lt.format(fmt(51.632032223028077), fmt(-51.2392421709312)))
    print(lt.format(fmt(69.664950700161214), fmt(-54.418903476604655)))
    print(lt.format(fmt(128.27477522926552), fmt(-56.49719397203603)))
    print(lt.format(fmt(144.97445682087536), fmt(38.211162080973452)))
    print(lt.format(fmt(107.42906113433914), fmt(44.831326831802869)))
    print(lt.format(fmt(50.733061134339017), fmt(57.8753268318029)))
    print(lt.format(fmt(14.035487342421575), fmt(64.346)))
    print(lt.format(fmt(-14.035487342421575), fmt(64.346)))
    print(lt.format(fmt(-50.733061134339017), fmt(57.8753268318029)))
    print(lt.format(fmt(-107.42906113433914), fmt(44.831326831802869)))
    print(lt.format(fmt(-144.97445682087536), fmt(38.211162080973509)))
    print(lt.format(fmt(-128.27477522926552), fmt(-56.497193972035959)))
    print(lt.format(fmt(-69.664950700161214), fmt(-54.418903476604655)))
    print(lt.format(fmt(-51.632032223028062), fmt(-51.2392421709312)))
    print(lt.format(fmt(-16.102172463211978), fmt(-33.83)))
    print(lt.format(fmt(0), fmt(-33.83)))

if __name__ == '__main__':
    main()

