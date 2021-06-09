package cmd

import (
	"fmt"
	"log"
	"strings"
	//"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Serve files from filesystem",
	Long: `
Location can be relative or absolute.

If no location is specified then the current directory is used.
If no port is set then port 8080 is used.`,
Example: `helper http `+ color.HiRedString("-d ")+`../ `+ color.HiRedString("-p ")+`8080`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.HasAvailablePersistentFlags()
		log.Default()
		fmt.Printf("%#v\n", args)
		fmt.Println(strings.Contains(strings.Join(args, ","),"-d"))
		fmt.Printf("args: %v\n", args)

		/* if (args[0] == "") {
			
		} */
		/* fsys := http.Dir(".")
		
		// Simple static webserver:
		http.ListenAndServe(":8080", http.FileServer(fsys)) */
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	httpCmd.PersistentFlags().StringP("directory", "d","./", "location (relative OR absolute")
	httpCmd.PersistentFlags().StringP("port", "p","8080","")
}
