import (
	"fmt"
	"math"
	"sort"

	"github.com/gonum/stat"
)


func main() {
	x := []float64{2, 4, 6, 8}
	y := []float64{81, 93, 91, 97}

	mx := stat.Mean(x, nil)
	my := stat.Mean(y, nil)

	fmt.Printf("x mean  =     %v\n", mx)
	fmt.Printf("y mean  =     %v\n", my)

	var divisor float64
	for i := range x {
		divisor += math.Pow(float64(mx-x[i]), 2)
		//divisor += math.Pow(float64(x[i]-mx), 2)
	}

	var dividend float64
	for i := 0; i < len(x); i++ {
		dividend += (x[i] - mx) * (y[i] - my)
	}

	fmt.Printf("분모    =     %v\n", divisor)
	fmt.Printf("분자    =     %v\n", dividend)

	a := dividend / divisor
	b := my - (mx * a)

	fmt.Printf("기울기 a=     %v\n", a)
	fmt.Printf("y 절편 b=     %v\n", b)
}
