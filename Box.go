package main
import(
	"ethos/syscall"
	"ethos/efmt"
	"math"
)

func main(){
me:=syscall.GetUser()

box_toRead := BoxType{lower_x:0, lower_y:0, upper_x:10, upper_y:10}
efmt.Printf("lower left x xoordinate: %v lower left y coordinate: %v\n", box_toRead.lower_x, box_toRead.lower_y)
efmt.Printf("upper left x xoordinate: %v upper left y coordinate: %v\n", box_toRead.upper_x, box_toRead.upper_y)
efmt.Printf(me)
x_square := math.Pow(box_toRead.lower_x - box_toRead.upper_x,2 ) 
y_square :=  math.Pow(box_toRead.lower_y - box_toRead.upper_y,2) 
distance :=  math.Sqrt(x_square + y_square ) 
efmt.Println("The distance between the points is: ", distance)

path := "/user/"+me+"/myDir/"
data := BoxType{lower_x:5, lower_y:5, upper_x:50, upper_y:50}
data.WriteVar(path+"BoxFile")

}
