/*
Copyright © 2025 xiaolin2004 <1553367438@qq.com>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	tools2 "github.com/xiaolin2004/signer/tools"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use: "check [httpMethod] [path] [clientId] [respTime] [respBody] [respSign] [publicKey]",
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.ExactArgs(7)(cmd, args)
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		httpMethod, path, clientId, respTime, respBody, respSign, publicKey := args[0], args[1], args[2], args[3], args[4], args[5], args[6]
		valid, err := tools2.CheckRspSign(httpMethod, path, clientId, respTime, respBody, respSign, publicKey)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(valid)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
