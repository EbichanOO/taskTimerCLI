package cmd

import (
	"fmt"
	"time"
	"strconv"

	"github.com/spf13/cobra"
)

func getNowTime()time.Time{
	var jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return time.Now().In(jst)
}

func time2str(t time.Time)string{
	return t.Format("01-02 15:04:05")
}

func timer(){
	printer([]string{"stop","show"})
}

func printer(commands []string){
	var printStr = ""
	for _,v := range commands {
		printStr += v + "\n"
	}
	printStr += "\033[" + strconv.Itoa(len(commands)+1) + "A"
	for ;;{
		fmt.Printf(time2str(getNowTime()) + "\n" + printStr)
		time.Sleep(1)
	}
}

var timerCmd = &cobra.Command{
	Use: "timer",
	Short: "hoge",
	Long: "hoge",
	Run: func(cmd *cobra.Command, args []string){
		timer()
	},
}

func init(){
	rootCmd.AddCommand(timerCmd)
}