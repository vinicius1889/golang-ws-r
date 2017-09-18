package main
import (
	"fmt"
	"time"
	"os"
	"os/exec"
)
func main(){

	fmt.Println("running now "+time.Now().Format("15:04:05"))

	home := (os.Getenv("HOME")+"/rscripts")

	out,err:= exec.Command("bash","-c","Rscript "+home+"/hello.r ").Output()
	if err != nil {
		fmt.Printf(" error %s",err)
		panic(err)
	}
	fmt.Printf("%s",out)
	fmt.Println("running now "+time.Now().Format("15:04:05"))

}
