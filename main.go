package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
	termbox "github.com/nsf/termbox-go"
)
func imgtype(imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	imgText := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	termbox.SetOutputMode(termbox.Output256)
	const mojiretsu = "0123456789 ░#" //{' ', '░', '▒', '▓', '█'}
	nihongoRune := []rune(mojiretsu)
	termbox.SetCell(0, 0, nihongoRune[1], termbox.ColorRed, termbox.ColorBlue)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			r, g, b, _ := imgText.At(j, i).RGBA()
			//h, l, s := rgb2hls(int32(r), int32(g), int32(b))
			r = uint32(r * 3 / 65535)
			g = uint32(g * 3 / 65535)
			b = uint32(b * 3 / 65535)

			chi := 10
			cidx := 0
			cbidx := 0
			if r == g && g == b && r > 0 {
				// white
				cidx = 8
				if r > 2 {
					cidx = cidx + 512
				}
			} else {
				if r > 0 {
					// r > 0
					if g == 0 && b == 0 {
						//red
						cidx = 2
						if r > 1 {
							// light red
							cidx = cidx + 512
						}
					}
					if g > 0 {
						// r>0,g>0
						if b == 0 {
							// r>0,g>0,b=0
							if r >= g && g > 1 {
								// yellow
								cidx = 4
								if g > 1 {
									cidx = cidx + 512
								}
							} else {
								// g=1
								if r > 2 {
									// red
									cidx = 2
									if r > 2 {
										cidx = cidx + 512
									}
								}
							}
						} else {
							// r>0,b>0
							if b > g && r >= b {
								// magenta
								cidx = 6
							} else {
								if r >= g && g > 2 {
									// yellow
									cidx = 4
								} else {
									cidx = 8
								}
							}
						}

					} else {
						// g=0
						if b > 0 {
							if r >= b && b > 1 {
								// magenta
								cidx = 6
							}
						}
					}
				} else {
					// r = 0
					if g > 0 {
						if b == 0 {
							//green
							cidx = 3
							if g > 1 {
								cidx = cidx + 512
							}
						} else {
							// r=0,g>0,b>0
							if g < b {
								//blue
								cidx = 5
								if b > 2 {
									cidx = cidx + 512
								}

							} else {
								// cyan
								cidx = 7
								if b > 1 {
									cidx = cidx + 512
								}
							}
						}
					} else {
						// r =0,g = 0
						if b > 0 {
							// blue
							cidx = 5
							if b > 1 {
								cidx = cidx + 512
							}
						} else {
							// r=0,g=0,b=0
							// black
							cidx = 1
						}
					}
				}
			}
			if r >= 2 && g >= 2 && b < 2 {
				// yellow
				cidx = 4
			}
			if r == 3 && g == 3 && b == 2 {
				// white
				cidx = 8 + 512
			}
			if r == 2 && g == 2 && b == 1 {
				// white
				cidx = 8
			}
			if r < 2 && g >= 2 && b >= 2 {
				// cyan
				cidx = 7
			}
			if r >= 2 && g < 2 && b >= 2 {
				// magenta
				cidx = 6
			}
			if r == 3 && g < 2 && b < 2 {
				// red
				cidx = 2
			}
			if r == 3 && g == 1 && b == 1 {
				// white
				cbidx = 2
				cidx = 8
				chi = 11
			}
			if r < 2 && g == 3 && b < 2 {
				// green
				cidx = 3
			}
			if r < 2 && g < 2 && b == 3 {
				// blue
				cidx = 5
			}
			if r == 2 && g == 1 && b == 0 {
				chi = 11
				// yellow
				cidx = 4
				cbidx = 2
			}
			if r == 1 && g == 1 && b == 0 {
				chi = 11
				// yellow
				cidx = 1
				cbidx = 4
			}
			if r == 3 && g == 3 && b == 2 {
				chi = 11
				// yellow
				cidx = 8
				cbidx = 4
			}
			if r == 1 && g == 0 && b == 1 {
				chi = 11
				// magenta
				cidx = 1
				cbidx = 6
			}
			if r == 2 && g == 0 && b == 1 {
				chi = 11
				// magenta
				cidx = 6
				cbidx = 2
			}
			if r == 1 && g == 0 && b == 2 {
				chi = 11
				// magenta
				cidx = 6
				cbidx = 5
			}
			if r == 3 && g == 2 && b == 3 {
				chi = 11
				// magenta
				cidx = 8
				cbidx = 6
			}
			if r == 0 && g == 1 && b == 1 {
				chi = 11
				// cyan
				cidx = 1
				cbidx = 7
			}
			if r == 0 && g == 2 && b == 1 {
				chi = 11
				// cyan
				cidx = 7
				cbidx = 3
			}
			if r == 0 && g == 1 && b == 2 {
				chi = 11
				// cyan
				cidx = 7
				cbidx = 5
			}
			if r == 2 && g == 3 && b == 3 {
				chi = 11
				// cyan
				cidx = 8
				cbidx = 7
			}
			if r == 1 && g == 1 && b == 1 {
				chi = 11
				// white
				cidx = 1
				cbidx = 8
			}
			if r == 2 && g == 2 && b == 2 {
				chi = 11
				// white
				cidx = 8
				cbidx = 8 + 512
			}
			if r == 1 && g == 0 && b == 0 {
				chi = 11
				// red
				cidx = 1
				cbidx = 2
			}
			if r == 0 && g == 1 && b == 0 {
				chi = 11
				// green
				cidx = 1
				cbidx = 3
			}
			if r == 0 && g == 0 && b == 1 {
				chi = 11
				// blue
				cidx = 1
				cbidx = 5
			}
			termbox.SetCell(j, i, nihongoRune[chi], termbox.Attribute(cbidx), termbox.Attribute(cidx))
		}
	}
	termbox.Flush()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				//Escキー押した時の処理
				return
			case termbox.KeyArrowUp:
				//上キー押した時の処理
			case termbox.KeyArrowDown:
				//下キー押した時の処理
			default:
			}
		default:
		}
	}

}
func main() {
	imagePath := "test.jpg"
	if len(os.Args) == 2 {
		imagePath = os.Args[1]
	}
	imgtype(imagePath)
}
