package game

import (
	"fmt"
	"strings"
)

func (b Board) String() string {
	f := boardStr
	guide := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "V", "w", "x", "y", "z", "~",
	}
	for i, hex := range b {
		fill := []string{" ", " ", " ", " ", " ", " ", " "}
		if hex.piece != nil {
			fill = hex.piece.Fill()
		}
		for _, k := range fill {
			f = strings.Replace(f, guide[i], k, 1)
		}
	}
	return f
}

func (s *State) String() string {
	return s.board.String() + fmt.Sprintf("\nTurn: %v", s.turn)
}

func (m *Move) String() string {
	if m.action == TRAVEL {
		return fmt.Sprintf("[Travel From:%v Direction:%v]\n", m.hex, m.target)
	}
	return fmt.Sprintf("[Rotate Hex:%v Amount:%v]\n", m.hex, m.target)
}

func (p *Piece) Fill() []string {
	out := []string{" ", " ", " ", " ", " ", " ", " "}

	if p.team == RED {
		out[3] = "•"
	}

	if p.arrows[0] {
		out[1] = "^"
	}
	if p.arrows[1] {
		out[2] = ">"
	}
	if p.arrows[2] {
		out[6] = ">"
	}
	if p.arrows[3] {
		out[5] = "v"
	}
	if p.arrows[4] {
		out[4] = "<"
	}
	if p.arrows[5] {
		out[0] = "<"
	}

	return out
}

const boardStr = `
                          -------
                         / x x x \
                  -------    x    -------
                 / w w w \ x x x / y y y \
          -------    w    -------    y    -------
         / V V V \ w w w / q q q \ y y y / z z z \
  -------    V    -------    q    -------    z    -------  
 / u u u \ V V V / p p p \ q q q / r r r \ z z z / ~ ~ ~ \
/    u    -------    p    -------    r    -------    ~    \
 \ u u u / o o o \ p p p / j j j \ r r r / s s s \ ~ ~ ~ /
  -------    o    -------    j    -------    s    -------
 / n n n \ o o o / i i i \ j j j / k k k \ s s s / t t t \
/    n    -------    i    -------    k    -------    t    \
 \ n n n / h h h \ i i i / c c c \ k k k / l l l \ t t t /
  -------    h    -------    c    -------    l    -------
 / g g g \ h h h / b b b \ c c c / d d d \ l l l / m m m \
\    g    -------    b    -------    d    -------    m    /
 \ g g g / a a a \ b b b / 6 6 6 \ d d d / e e e \ m m m /
  -------    a    -------    6    -------    e    -------
 / 9 9 9 \ a a a / 5 5 5 \ 6 6 6 / 7 7 7 \ e e e / f f f \
\    9    -------    5    -------    7    -------    f    /
 \ 9 9 9 / 4 4 4 \ 5 5 5 / 2 2 2 \ 7 7 7 / 8 8 8 \ f f f /
  -------    4    -------    2    -------    8    ------- 
         \ 4 4 4 / 1 1 1 \ 2 2 2 / 3 3 3 \ 8 8 8 /
          -------    1    -------    3    -------
                 \ 1 1 1 / 0 0 0 \ 3 3 3 /
                  -------    0    -------
                         \ 0 0 0 /   
                          -------
`
