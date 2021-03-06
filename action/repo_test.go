package action

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/helm/helm/log"
)

func TestListRepos(t *testing.T) {
	var b bytes.Buffer

	log.IsDebugging = true
	log.Stdout = &b
	defer func() { log.Stdout = os.Stdout }()

	homedir := createTmpHome()
	fakeUpdate(homedir)
	ListRepos(homedir)

	out := b.String()
	if !strings.Contains(out, "charts*\thttps://github.com/helm/charts") {
		t.Errorf("Unexpectedly got %s", out)
	}
}
