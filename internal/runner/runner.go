package runner

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

func Run(name string, args []string, wasm []byte, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	ctx := context.Background()

	rt := wazero.NewRuntime(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, rt)

	args = append([]string{name}, args...)

	cfg := wazero.NewModuleConfig().
		WithSysNanosleep().
		WithSysNanotime().
		WithSysWalltime().
		WithStderr(stderr).
		WithStdout(stdout).
		WithStdin(stdin).
		WithRandSource(rand.Reader).
		WithArgs(args...)
	for _, env := range os.Environ() {
		k, v, _ := strings.Cut(env, "=")
		cfg = cfg.WithEnv(k, v)
	}

	_, err := rt.InstantiateWithConfig(ctx, wasm, cfg)
	if err != nil {
		if sErr, ok := err.(*sys.ExitError); ok { //nolint:errorlint
			return int(sErr.ExitCode())
		}
		log.Fatal(err)
	}
	return 0
}
