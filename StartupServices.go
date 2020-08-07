package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func showLaunchAgent() {

	app := "ls"
	arg0 := "-la"
	arg1 := "/Library/LaunchAgents/"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))

	/* ALTERNATIVE (old)
	out, err := exec.Command(app, arg0).Output()
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println("Command successfully executed")
		output := string(out[:])
		fmt.Println(output)
	}
	*/

}

func showLaunchDaemons() {

	app := "ls"
	arg0 := "-la"
	arg1 := "/Library/LaunchDaemons/"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))
}

func showStartupItems() {

	app := "ls"
	arg0 := "-la"
	arg1 := "/Library/StartupItems/"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))
}

func main() {

	if runtime.GOOS == "windows" {
		fmt.Println("cant execute this on a windows machine")
	} else {
		fmt.Println("1. Show Launch Agents")
		fmt.Println("2. Show Launch Daemons")
		fmt.Println("3. Show Startup Items")

		var i int
		_, err := fmt.Scanf("%d", &i)

		if err != nil {
			fmt.Println("Sorry that was nto a valid option")
			fmt.Println(err)
			return
		}

		if i == 1 {
			showLaunchAgent()
		} else if i == 2 {
			showLaunchDaemons()
		} else if i == 3 {
			showStartupItems()
		} else {
			fmt.Println("Sorry selection not in scope")
			return
		}

	}

}
