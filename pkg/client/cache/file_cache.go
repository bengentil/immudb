/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/rogpeppe/go-internal/lockedfile"
	"os"
	"path/filepath"
	"strings"

	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/golang/protobuf/proto"
)

// STATE_FN ...
const STATE_FN = ".state-"

type fileCache struct {
	Dir string
	stateFile *lockedfile.File
}

// NewFileCache returns a new file cache
func NewFileCache(dir string) Cache {
	return &fileCache{Dir: dir}
}

func (w *fileCache) Get(serverUUID string, db string) (*schema.ImmutableState, error) {
	if w.stateFile == nil {
		return nil, ErrCacheNotLocked
	}
	scanner := bufio.NewScanner(w.stateFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, db+":") {
			r := strings.Split(line, ":")
			if r[1] == "" {
				return nil, ErrPrevStateNotFound
			}
			oldState, err := base64.StdEncoding.DecodeString(r[1])
			if err != nil {
				return nil, ErrLocalStateCorrupted
			}
			state := &schema.ImmutableState{}
			if err = proto.Unmarshal(oldState, state); err != nil {
				return nil, ErrLocalStateCorrupted
			}
			return state, nil
		}
	}
	return nil, ErrPrevStateNotFound
}

func (w *fileCache) Set(serverUUID string, db string, state *schema.ImmutableState) error {
	if w.stateFile == nil {
		return ErrCacheNotLocked
	}
	raw, err := proto.Marshal(state)
	if err != nil {
		return err
	}

	newState := db + ":" + base64.StdEncoding.EncodeToString(raw)
	var exists bool

	scanner := bufio.NewScanner(w.stateFile)
	scanner.Split(bufio.ScanLines)
	var lines [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, db+":") {
			exists = true
			lines = append(lines, []byte(newState))
		}
	}
	if !exists {
		lines = append(lines, []byte(newState))
	}
	output := bytes.Join(lines, []byte("\n"))

	_, err = w.stateFile.WriteAt( output, 0)
	if err != nil {
		return err
	}
	return nil
}


func (w *fileCache) Lock(serverUUID string) (err error) {
	if w.stateFile != nil {
		return ErrCacheAlreadyLocked
	}
	w.stateFile, err = lockedfile.OpenFile(w.getStateFilePath(serverUUID), os.O_RDWR | os.O_CREATE, 0755)
	return err
}

func (w *fileCache) Unlock() (err error) {
	return w.stateFile.Close()
}

func (w *fileCache) getStateFilePath(UUID string) string {
	return filepath.Join(w.Dir, STATE_FN+UUID)
}
