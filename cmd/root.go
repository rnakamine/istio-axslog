/*
Copyright © 2022 Ryo Nakamine <rnakamine8080@gmail.com>

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
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/Songmu/go-ltsv"
	"github.com/goccy/go-json"
	"github.com/rnakamine/istio-axslog/parser"
	"github.com/rnakamine/istio-axslog/version"
	"github.com/spf13/cobra"
)

var output string

var rootCmd = &cobra.Command{
	Use:          "istio-axslog",
	Short:        "istio-axslog is parsed istio-proxy(envoy) access log and output in any format",
	Long:         `istio-axslog is parsed istio-proxy(envoy) access log and output in any format.`,
	Version:      version.Version,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if output != "json" && output != "ltsv" {
			return errors.New("unsupported output format. supported formats are json, ltsv")
		}
		p := parser.New()
		scanner := bufio.NewScanner(cmd.InOrStdin())
		for scanner.Scan() {
			accessLog, err := p.Parse(scanner.Text())
			if err != nil {
				return err
			}
			if *accessLog != (parser.EnvoyAccessLog{}) {
				switch output {
				case "json":
					out, err := json.Marshal(accessLog)
					if err != nil {
						return err
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(out))
				case "ltsv":
					out, err := ltsv.Marshal(accessLog)
					if err != nil {
						return err
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(out))
				}
			}
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "json", "output format. supported formats are json, ltsv")
}
