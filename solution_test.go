package square

import (
	"fmt"
	"testing"
)

func TestTrinagle(t *testing.T) {
	fmt.Println(CalcSquare(1, SidesTriangle))
	fmt.Println(CalcSquare(4.4, SidesTriangle))
	fmt.Println(CalcSquare(15.67, SidesTriangle))
}
