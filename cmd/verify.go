/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiaolin2004/signer/tools"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify [httpMethod] [path] [clientId] [rspTimeStr] [rspBody] [signature] [alipayPublicKey]",
	Short: "Verify the response signature",
	Long: `Verify if the signature from Alipay response is valid.
Required inputs:
- httpMethod: The HTTP request method used (e.g., POST).
- path: The API request path (e.g., /api/v1/payments/subscription).
- clientId: Your client ID for authentication.
- rspTimeStr: Response timestamp from header.
- rspBody: The raw JSON response body.
- signature: The signature string received from Alipay (base64 encoded).
- alipayPublicKey: The Alipay public key PEM (base64 or file content).
Example:
verify POST /api/v1/payments/subscription 2021000118635012 1710922210239123000 '{"result":"ok"}' XXXXXXXXXXXX ALIPAY_PUBKEY`,
	Args: cobra.ExactArgs(7),
	Run: func(cmd *cobra.Command, args []string) {
		httpMethod := args[0]
		path := args[1]
		clientId := args[2]
		rspTimeStr := args[3]
		rspBody := args[4]
		signature := args[5]
		alipayPublicKey := args[6]

		isValid, err := tools.Verify(httpMethod, path, clientId, rspTimeStr, rspBody, signature, alipayPublicKey)
		if err != nil {
			fmt.Printf("❌ 验签失败: %v\n", err)
			return
		}
		if isValid {
			fmt.Println("✅ 签名验证成功")
		} else {
			fmt.Println("❌ 签名不合法")
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
