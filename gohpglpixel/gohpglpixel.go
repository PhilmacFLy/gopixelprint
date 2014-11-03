package gohpglpixel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
import "strconv"

type Pixelart struct {
	height  int
	width   int
	canvas  [][]int
	scaling int
	filling int
	border  bool
	Lala    string
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
	p.width = width
	p.height = height
	p.canvas = make([][]int, height)
	for i := range p.canvas {
		p.canvas[i] = make([]int, width)
	}
	fmt.Println("%v\n", p.canvas)
}

func (p *Pixelart) SetPixel(x int, y int, color int) {
	if x < p.width {
		if y < p.height {
			p.canvas[y][x] = color
		}
	}
}

func (p *Pixelart) SetScaling(scaling int) {
	p.scaling = scaling
}

func (p *Pixelart) SetFilling(filling int) {
	p.filling = filling
}

func (p *Pixelart) SetBorder(border bool) {
	p.border = border
}

func (p *Pixelart) Print() {
	fmt.Println(p.canvas)
}

func (p *Pixelart) SaveHPGL(filename string) {
	hpgl := make([]string, p.width*p.height*6+1)
	color := make([][]string, 6)
	for i := range color {
		color[i] = make([]string, p.width*p.height)
	}
	count := make([]int, 6)
	count[0] = 0
	count[1] = 0
	count[2] = 0
	count[3] = 0
	count[4] = 0
	count[5] = 0
	hpgl[0] = "SC0,210,0,210;"
	for j := 0; j < p.width*p.height; j++ {
		q := j / p.width
		r := j % p.width
		if p.canvas[r][q] != 0 {
			color[p.canvas[r][q]-1][count[p.canvas[r][q]-1]] = p.genratesquare(r, q)
			count[p.canvas[r][q]-1] = count[p.canvas[r][q]-1] + 1
		}
	}
	i := 1
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
	writeLines(hpgl, filename)
}

func (p *Pixelart) genratesquare(x int, y int) string {
	Result := ""
	if p.border {
		Result = Result + "PU" + strconv.Itoa(x*p.scaling) + "," + strconv.Itoa(y*p.scaling) + ";"
		Result = Result + "PD" + strconv.Itoa((x+1)*p.scaling) + "," + strconv.Itoa(y*p.scaling) + ";"
		Result = Result + "PD" + strconv.Itoa((x+1)*p.scaling) + "," + strconv.Itoa((1+y)*p.scaling) + ";"
		Result = Result + "PD" + strconv.Itoa(x*p.scaling) + "," + strconv.Itoa((1+y)*p.scaling) + ";"
		Result = Result + "PD" + strconv.Itoa(x*p.scaling) + "," + strconv.Itoa(y*p.scaling) + ";"
	}
	i := 0
	if p.filling > 0 {
		for i < 10 {
			if (p.filling == 1) || (p.filling == 2) {
				Result = Result + "PU" + strconv.Itoa(x*p.scaling+i) + "," + strconv.Itoa(y*p.scaling) + ";"
				Result = Result + "PD" + strconv.Itoa(x*p.scaling) + "," + strconv.Itoa(y*p.scaling+i) + ";"
			}
			if (p.filling == 1) || (p.filling == 3) {
				Result = Result + "PU" + strconv.Itoa(x*p.scaling+i) + "," + strconv.Itoa((y+1)*p.scaling) + ";"
				Result = Result + "PD" + strconv.Itoa((x+1)*p.scaling) + "," + strconv.Itoa(y*p.scaling+i) + ";"
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
