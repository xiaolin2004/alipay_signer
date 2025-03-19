/*
Copyright © 2025 xiaolin2004 <1553367438@qq.com>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiaolin2004/signer/tools"
	"strconv"
	"time"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use: "sign [clientId] [path] [httpMethod] [privateKey] [jsonBody]",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(5)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		clientId, path, httpMethod, privateKey, jsonBody := args[0], args[1], args[2], args[3], args[4]
		reqTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		sign, err := tools.GenSign(fmt.Sprintf("%s", httpMethod), path, clientId, reqTime, jsonBody, privateKey)
		if err != nil {
			fmt.Printf("%v", err)
		}
		header := tools.BuildBaseHeader(reqTime, clientId, "1", sign)
		for key, value := range header {
			fmt.Printf("%s:%s\n", key, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
