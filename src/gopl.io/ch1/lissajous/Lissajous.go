package lissajous

import (
	"image/color"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"fmt"
)

var palette = []color.Color{color.White, color.Black}
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

//func main() {
//	lissajous(os.Stdout)
//}

func Lissajous(out io.Writer, r *http.Request) {
	//关于 Go 语言中 const 的问题:https://segmentfault.com/q/1010000008620205
	//int和float不能直接相乘
	var cycles  int = 5     // number of complete x oscillator revolutions
	const (
		//cycles = 5
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	//fmt.Println(reflect.TypeOf(cycles))

	r.ParseForm()
	reqCycles := r.Form.Get("cycles")
	if reqCycles != "" {
		cycles,_ = strconv.Atoi(reqCycles)
	}
	fmt.Println("cycles=" + strconv.Itoa(cycles))

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
