package gif_test

import (
	"fmt"
	"os"

	"github.com/artyom/image/gif"
)

func ExampleFrameDecoder() {
	f, err := os.Open("../testdata/video-001.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	dec := gif.NewFrameDecoder(f)
	for dec.Next() {
		img, delay, disposal := dec.Frame()
		fmt.Printf("%T, %d, %v\n", img, delay, disposal)
	}
	if err := dec.Err(); err != nil {
		fmt.Println(err)
		return
	}
	// Output:
	// *image.Paletted, 0, 0
}
