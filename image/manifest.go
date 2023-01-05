// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2022 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package image

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/snapcore/snapd/snap"
)

// The seed.manifest generated by ubuntu-image contains entries in the following
// format:
// <snap-name> <snap-revision>
// The goal in a future iteration of this will be to move the generation of the
// seed.manifest to this package, out of ubuntu-image.
// TODO: Move generation of seed.manifest from ubuntu-image to here

// ReadSeedManifest reads a seed.manifest generated by ubuntu-image, and returns
// a map containing the snap names and their revisions.
func ReadSeedManifest(manifestFile string) (map[string]snap.Revision, error) {
	f, err := os.Open(manifestFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	revisions := make(map[string]snap.Revision)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, " ") {
			return nil, fmt.Errorf("line cannot start with any spaces: %q", line)
		}

		tokens := strings.Fields(line)
		// Expect exactly two tokens
		if len(tokens) != 2 {
			return nil, fmt.Errorf("line is illegally formatted: %q", line)
		}

		snapName := tokens[0]
		revString := tokens[1]
		if err := snap.ValidateName(snapName); err != nil {
			return nil, err
		}

		rev, err := snap.ParseRevision(revString)
		if err != nil {
			return nil, err
		}

		// Values that are higher than 0 indicate the revision comes from the store, and values
		// lower than 0 indicate the snap was sourced locally. We allow both in the seed.manifest as
		// long as the user can provide us with the correct snaps. The only number we won't accept is
		// 0.
		if rev.Unset() {
			return nil, fmt.Errorf("cannot use revision %d for snap %q: revision must not be 0", rev, snapName)
		}
		revisions[snapName] = rev
	}
	return revisions, nil
}

// WriteSeedManifest generates the seed.manifest contents from the provided map of
// snaps and their revisions, and stores them in the given file path
func WriteSeedManifest(filePath string, revisions map[string]snap.Revision) error {
	if len(revisions) == 0 {
		return nil
	}

	keys := make([]string, 0, len(revisions))
	for k := range revisions {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := bytes.NewBuffer(nil)
	for _, key := range keys {
		rev := revisions[key]
		if rev.Unset() {
			return fmt.Errorf("revision must not be 0 for snap %q", key)
		}
		fmt.Fprintf(buf, "%s %s\n", key, rev)
	}
	return ioutil.WriteFile(filePath, buf.Bytes(), 0755)
}