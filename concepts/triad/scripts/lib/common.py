import enum


class TriadPiece(enum.Enum):
    LEFT = 0
    CENTER = 1
    RIGHT = 2


class Corner(enum.Enum):
    TOP_LEFT = 0
    TOP_RIGHT = 1
    BOTTOM_LEFT = 2
    BOTTOM_RIGHT = 3
