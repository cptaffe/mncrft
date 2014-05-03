package main

import(
	"os/exec"
	"os/user"
	"fmt"
)

const(
	// colors (linux)
	BLUE="\033[35m"
	//BLUE="\033[38;5;220m" -- 256 colors
	RED="\033[31m"
	ENDC="\033[0m"
	BEGIN=BLUE+"# "+ENDC
	CMDOUT=RED+"> "+ENDC
)

func main(){
	fmt.Printf(BEGIN+"starting minecraft...\n")
	usr, err := user.Current()
	if err != nil{
		fmt.Printf(BEGIN+"error getting user.")
	}
	cmd := exec.Command("java", "-jar", fmt.Sprintf("%s/.minecraft-jar/minecraft.jar", usr.HomeDir))
	s, err := cmd.Output()
	if err != nil{
		sub := ""
		for i := 0; i < len(s); i++ {
			if s[i] != '\n'{
				sub += string(s[i])
			} else {
				fmt.Printf(CMDOUT+"%s\n", sub)
				sub = ""
			}
		}
	} else {
		fmt.Printf(BEGIN+"no error.\n")
	}
}