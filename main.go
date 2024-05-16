package main

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/slides/internal/cmd"
	"github.com/maaslalani/slides/internal/model"
	"github.com/maaslalani/slides/internal/navigation"
	"github.com/maaslalani/slides/internal/remote"
	"github.com/muesli/coral"
)

var (
	rootCmd = &coral.Command{
		Use:   "slides <file.md>",
		Short: "Terminal based presentation tool",
		Args:  coral.ArbitraryArgs,
		RunE: func(cmd *coral.Command, args []string) error {
			var err error
			var fileName string

			if len(args) > 0 {
				fileName = args[0]
			}

			isRemoteOnly := isRemoteControlledOnly()

			presentation := model.Model{
				Page:             0,
				Date:             time.Now().Format("2006-01-02"),
				FileName:         fileName,
				Search:           navigation.NewSearch(),
				DisableLocalKeys: isRemoteOnly,
			}
			err = presentation.Load()
			if err != nil {
				return err
			}

			p := tea.NewProgram(presentation, tea.WithAltScreen())
			if listener := initRemoteListener(p); listener != nil {
				defer listener.Close()
			}
			_, err = p.Run()
			return err
		},
	}
)

// isRemoteControlledOnly checks if SLIDES_REMOTE_ONLY environment
// var is set to true values
func isRemoteControlledOnly() bool {
	remoteControlledOnly := os.Getenv("SLIDES_REMOTE_ONLY")
	switch remoteControlledOnly {
	case "y", "Y", "yes", "YES", "1", "on", "ON", "true", "TRUE", "t", "T":
		return true
	}
	return false
}

// initialize and start socket remote listener if required
func initRemoteListener(p *tea.Program) (remoteListener *remote.SocketRemoteListener) {
	// init if env var is given
	// TODO: decide whether to use flags also or not
	remoteSocketPath := os.Getenv("SLIDES_REMOTE_SOCKET")
	if remoteSocketPath != "" {
		var err error
		relay := remote.NewCommandRelay(p)
		remoteListener, err = remote.NewSocketRemoteListener(
			remoteSocketPath, relay)
		if err != nil {
			fmt.Errorf(err.Error())
			os.Exit(1)
		}
		remoteListener.Start()
	}

	return remoteListener
}

func init() {
	rootCmd.AddCommand(
		cmd.ServeCmd,
	)
	rootCmd.AddCommand(
		cmd.RemoteCmd,
	)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
