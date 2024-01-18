package cmd

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/adrg/xdg"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client/config"
	"strings"
)

var (
	cfgFile   string
	verbosity *int
)

// RootCmd is clusterctl root CLI command.
var RootCmd = &cobra.Command{
	Use:          "clusterctl",
	SilenceUsage: true,
	Short:        "clusterctl controls the lifecycle of a Cluster API management cluster",
	Long: LongDesc(`
		Get started with Cluster API using clusterctl to create a management cluster,
		install providers, and create templates for your workload cluster.`),
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		// Check if clusterctl needs an upgrade "AFTER" running each command
		// and sub-command.
		configClient, err := config.New(ctx, cfgFile)
		if err != nil {
			return err
		}
		disable, err := configClient.Variables().Get("CLUSTERCTL_DISABLE_VERSIONCHECK")
		if err == nil && disable == "true" {
			// version check is disabled. Return early.
			return nil
		}
		checker, err := newVersionChecker(ctx, configClient.Variables())
		if err != nil {
			return err
		}
		output, err := checker.Check(ctx)
		if err != nil {
			return errors.Wrap(err, "unable to verify clusterctl version")
		}
		if output != "" {
			// Print the output in yellow so it is more visible.
			fmt.Fprintf(os.Stderr, "\033[33m%s\033[0m", output)
		}

		configDirectory, err := xdg.ConfigFile(config.ConfigFolderXDG)
		if err != nil {
			return err
		}

		// clean the downloaded config if was fetched from remote
		downloadConfigFile := filepath.Join(configDirectory, config.DownloadConfigFile)
		if _, err := os.Stat(downloadConfigFile); err == nil {
			if verbosity != nil && *verbosity >= 5 {
				fmt.Fprintf(os.Stdout, "Removing downloaded clusterctl config file: %s\n", config.DownloadConfigFile)
			}
			_ = os.Remove(downloadConfigFile)
		}

		return nil
	},
}

// Execute executes the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if verbosity != nil && *verbosity >= 5 {
			if err, ok := err.(stackTracer); ok {
				for _, f := range err.StackTrace() {
					fmt.Fprintf(os.Stderr, "%+s:%d\n", f, f)
				}
			}
		}
		// TODO: print cmd help if validation error
		os.Exit(1)
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// LongDesc normalizes a command's long description to follow the conventions.
func LongDesc(s string) string {
	if s == "" {
		return s
	}
	return normalizer{s}.heredoc().trim().string
}

// TODO: document this, what does it do? Why is it here?
type normalizer struct {
	string
}

func (s normalizer) heredoc() normalizer {
	s.string = heredoc.Doc(s.string)
	return s
}

func (s normalizer) trim() normalizer {
	s.string = strings.TrimSpace(s.string)
	return s
}
