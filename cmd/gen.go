/*
Copyright © 2021 ras0q <asymptote.k.k@gmail.com>

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
	"github.com/pkg/errors"
	"github.com/ras0q/gcg/internal/handler"
	"github.com/ras0q/gcg/internal/util/injector"
	"github.com/spf13/cobra"
)

var genOpts handler.GenOpts

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate constructors",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		h := injector.NewHandlers()
		if err := h.ExecuteGen(args[0], genOpts); err != nil {
			return errors.Wrap(err, "Could not generate constructors")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&genOpts.Output, "output", "o", "", "Output file")
	genCmd.Flags().BoolVarP(&genOpts.IsPrivate, "private", "p", false, "Generate private constructors")
}
