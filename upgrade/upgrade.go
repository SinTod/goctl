package upgrade

import (
	"fmt"
	"runtime"

	"github.com/SinTod/goctl/rpc/execx"
	"github.com/spf13/cobra"
)

// upgrade gets the latest goctl by
// go install github.com/SinTod/goctl@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/SinTod/goctl@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install github.com/SinTod/goctl@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
