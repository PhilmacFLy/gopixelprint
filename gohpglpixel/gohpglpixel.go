package gohpglpixel

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
import "strconv"

type Pixelart struct {
	Height  int
	Width   int
	Canvas  [][]int
	Scaling int
	Filling int
	Border  bool
	Title   string
}

func (p *Pixelart) ReadFile(filename string) error {
	p.SetDim(1024, 1024)
	linecount := 0
	offset := 0
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	lines, err := lineCounter(file)
	if err != nil {
		return err
	}
	file.Close()
	file, _ = os.Open(filename)
	llen := maxLineLength(file)
	p.SetDim(llen-1, lines)
	file.Close()
	file, _ = os.Open(filename)
	r := bufio.NewReader(file)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		offset = 0
		for _, r := range s {
			p.SetPixel(offset, linecount, int(r-'0'))
			offset++
		}
		linecount++
	}
	return nil
}
func maxLineLength(r io.Reader) int {
	llen := 0
	br := bufio.NewReader(r)

	for {
		s, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if len(s) > llen {
			llen = len(s)
		}
	}
	return llen
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 8196)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}
		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func (p *Pixelart) SetDim(width int, height int) {
	p.Width = width
	p.Height = height
	p.Canvas = make([][]int, height)
	for i := range p.Canvas {
		p.Canvas[i] = make([]int, width)
	}
}

func (p *Pixelart) SetPixel(x int, y int, color int) {
	if x < p.Width {
		if y < p.Height {
			p.Canvas[y][x] = color
		}
	}
}

func (p *Pixelart) SetScaling(scaling int) {
	p.Scaling = scaling
}

func (p *Pixelart) SetFilling(filling int) {
	p.Filling = filling
}

func (p *Pixelart) SetBorder(border bool) {
	p.Border = border
}

func (p *Pixelart) SetTitle(title string) {
	p.Title = title
}

func (p *Pixelart) Print() {
	fmt.Println(p.Canvas)
}

func (p *Pixelart) SaveHPGL(filename string) {
	hpgl := make([]string, p.Width*p.Height*6+1)
	color := make([][]string, 6)
	for i := range color {
		color[i] = make([]string, p.Width*p.Height)
	}
	count := make([]int, 6)
	count[0] = 0
	count[1] = 0
	count[2] = 0
	count[3] = 0
	count[4] = 0
	count[5] = 0
	i := 0
	hpgl[i] = "IN;IP0,0,4000,4000;SC0,100,0,100;"
	i++
	switch p.Filling {
	default:
		hpgl[i] = ""
	case 1:
		hpgl[i] = "FT1"
	case 2:
		hpgl[i] = "FT2"
	case 3:
		hpgl[i] = "FT3,2,45"
	case 4:
		hpgl[i] = "FT4,2,45"
	}
	i++
	for j := 0; j < p.Width*p.Height; j++ {
		q := j / p.Width
		r := j % p.Width
		if p.Canvas[r][q] != 0 {
			color[p.Canvas[r][q]-1][count[p.Canvas[r][q]-1]] = p.generatesquare(r, q)
			count[p.Canvas[r][q]-1] = count[p.Canvas[r][q]-1] + 1
		}
	}
	for j := 0; j < 6; j++ {
		if count[j] > 0 {
			hpgl[i] = "SP" + strconv.Itoa(j+1)
			i++
			for k := 0; k < count[j]; k++ {
				hpgl[i] = color[j][k]
				i++
			}
		}
	}
	if p.Title != "" {
		titley := (p.Height + 2) * p.Scaling
		titlex := 0
		hpgl[i] = "SP1;SI0.8,1.0;PA" + strconv.Itoa(titley) + "," + strconv.Itoa(titlex) + ";LO6;DI0,1;LB" + p.Title
	} else {
		hpgl[i] = ""
	}
	writeLines(hpgl, filename)
}

func (p *Pixelart) generatesquare(x int, y int) string {
	Result := ""
	command := ""
	if p.Border {
		command = "ER"
	} else {
		command = "RR"
	}

	Result = "PA" + strconv.Itoa(x*p.Scaling) + "," + strconv.Itoa(y*p.Scaling) + ";"
	Result = Result + command + strconv.Itoa(p.Scaling) + "," + strconv.Itoa(p.Scaling) + ";"

	return Result
}

func (p *Pixelart) generatemanualsquare(x int, y int) string {
	Result := ""
	if p.Border {
		Result = Result + "PU" + strconv.Itoa(x*p.Scaling) + "," + strconv.Itoa(y*p.Scaling) + ";"
		Result = Result + "PD" + strconv.Itoa((x+1)*p.Scaling) + "," + strconv.Itoa(y*p.Scaling) + ";"
		Result = Result + "PD" + strconv.Itoa((x+1)*p.Scaling) + "," + strconv.Itoa((1+y)*p.Scaling) + ";"
		Result = Result + "PD" + strconv.Itoa(x*p.Scaling) + "," + strconv.Itoa((1+y)*p.Scaling) + ";"
		Result = Result + "PD" + strconv.Itoa(x*p.Scaling) + "," + strconv.Itoa(y*p.Scaling) + ";"
	}
	i := 0
	if p.Filling > 0 {
		for i < 10 {
			if (p.Filling == 1) || (p.Filling == 2) {
				Result = Result + "PU" + strconv.Itoa(x*p.Scaling+i) + "," + strconv.Itoa(y*p.Scaling) + ";"
				Result = Result + "PD" + strconv.Itoa(x*p.Scaling) + "," + strconv.Itoa(y*p.Scaling+i) + ";"
			}
			if (p.Filling == 1) || (p.Filling == 3) {
				Result = Result + "PU" + strconv.Itoa(x*p.Scaling+i) + "," + strconv.Itoa((y+1)*p.Scaling) + ";"
				Result = Result + "PD" + strconv.Itoa((x+1)*p.Scaling) + "," + strconv.Itoa(y*p.Scaling+i) + ";"
			}
			i = i + 2
		}
	}
	Result = Result + "PU"
	return Result
}

func (p *Pixelart) WritePixelart(filename string) error {
	b, err := json.MarshalIndent(&p, "", "    ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(filename+".json", b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pixelart) LoadPixelart(filename string) error {
	filename = filename + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return err
	}
	return nil
}
