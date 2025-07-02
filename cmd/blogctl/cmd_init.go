package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const defaultEnvExample = `# credentails
BLOGCTL_HATENA_ID=
BLOGCTL_BLOG_ID=
BLOGCTL_APIKEY=

# configurations
BLOGCTL_ENTRY_SLUG_PREFIX=entry
BLOGCTL_ENTRY_SLUG_TAILSLASH_MAPPING=_index.md
`

func doInitEnv() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get the current working directory: %v", err)
	}

	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		log.Fatalf("this command required to `EDITOR` environment variable.")
	}

	fn := filepath.Join(cwd, ".env")
	if _, err := os.Stat(fn); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(fn, []byte(defaultEnvExample), 0600)
		if err != nil {
			log.Fatalf("failed to write default `.env` template: %v", err)
		}
	}

	cmd := exec.Command(editor, fn)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to edit `.env` file: %v", err)
	}

	os.Exit(0)
}
