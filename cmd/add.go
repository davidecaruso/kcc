/*
Copyright © 2020 Davide Caruso <davide.caruso93@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kcc/internal/service"
	"kcc/internal/storage"
)

var s service.Service

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Store service credentials",
	Long: `Examples:

kcc add -s facebook.com -u john@doe.com -p secret
kcc add -s 176.69.100.144 -u johndoe -p secret`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := storage.S.Add(s); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Ok")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&s.User, "user", "u", "", "The user name")
	addCmd.Flags().StringVarP(&s.Host, "service", "s", "", "The service name")
	addCmd.Flags().StringVarP(&s.Password, "password", "p", "", "The user password")
}
