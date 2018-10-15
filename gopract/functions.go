package main

import(
  "os"
  "fmt"
  "reflect"
  // "time"
  // "math"
)

func add( x,y int) int{
  return x + y
}
func printStr(x string) string{
    return x
}


func swap(x, y string) (string, string) {
	return y, x
}

func backwards(x,y,z int)(string){
  backString str :=  fmt.Sprintf("The inputs backwards are %d %d %d", y,z,x)

  return "backwards is %v \n",x
}

func main() {
  inputString := os.Args[len(os.Args)-1:]
  // argsWithoutProg := os.Args[1:]
  // arg := os.Args[3]
  fmt.Printf("the type of this input is %s \n",reflect.TypeOf(inputString))
  fmt.Println(inputString)
  fmt.Printf("The addition of 22 and 33 is %v \n",add(22,33))
  fmt.Printf("1,2,3 backwards is %v \n", backwards(1,2,3))
  // fmt.Println(argsWithoutProg)
  // fmt.Println(arg)
}
