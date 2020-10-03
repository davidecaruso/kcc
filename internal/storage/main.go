/*
Copyright © 2020 Davide Caruso <davide.caruso93@gmail.com>

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
package storage

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"kcc/internal/service"
	"kcc/tools"
	"os"
	"os/exec"
	"strconv"
)

type Storage struct {
	path string
	data map[string]service.Service
}

var (
	cwd, _ = os.Getwd()
	S      = Storage{path: cwd + "/assets/.hosts"}
)

func (s *Storage) lock() error {
	cmd := exec.Command("/bin/sh", "-c", "sudo chown root:root "+s.path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (s *Storage) unlock() error {
	uid := strconv.Itoa(os.Getuid())
	gid := strconv.Itoa(os.Getgid())
	cmd := exec.Command("/bin/sh", "-c", "sudo chown "+uid+":"+gid+" "+s.path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (s *Storage) create() error {
	if _, err := os.Stat(s.path); err != nil {
		if os.IsNotExist(err) {
			if err = ioutil.WriteFile(s.path, []byte("{}"), 0600); err != nil {
				return err
			}
		}
	}
	return s.lock()
}

func (s *Storage) update() (bool, error) {
	if err := s.create(); err != nil {
		return false, err
	}

	data, err := json.Marshal(&s.data)
	if err != nil {
		return false, err
	}

	if err := s.unlock(); err != nil {
		return false, err
	}
	err = ioutil.WriteFile(s.path, data, 0600)
	if err != nil {
		return false, err
	}
	return true, s.lock()
}

func (s *Storage) load() error {
	if err := s.create(); err != nil {
		return err
	}
	if err := s.unlock(); err != nil {
		return err
	}
	file, err := ioutil.ReadFile(s.path)
	if err != nil {
		return err
	}
	if err := s.lock(); err != nil {
		return err
	}
	return json.Unmarshal(file, &s.data)
}

func (s *Storage) Add(se service.Service) (bool, error) {
	if err := s.load(); err != nil {
		return false, err
	}

	if _, exists := s.data[se.Key()]; exists {
		if c := tools.Confirm("Credentials already exist, would you replace them?"); !c {
			return true, nil
		}
	}

	s.data[se.Key()] = se
	return s.update()
}

func (s *Storage) Get(se service.Service) error {
	if err := s.load(); err != nil {
		return err
	}

	if _, exists := s.data[se.Key()]; !exists {
		fmt.Println("Not found")
		return nil
	}

	return clipboard.WriteAll(s.data[se.Key()].Password)
}

func (s *Storage) Delete(se service.Service) (bool, error) {
	if err := s.load(); err != nil {
		return false, err
	}

	if _, exists := s.data[se.Key()]; exists {
		if c := tools.Confirm("Do you really want to delete service credentials?"); c {
			delete(s.data, se.Key())
		}
	}

	return s.update()
}

func init() {
	if err := S.load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
