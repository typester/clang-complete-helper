package main

import (
	"path/filepath"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

type Project struct {
	file string
	data []byte
	root string
}

func FindProject(file string) (*Project, error) {
	f, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	for dir := filepath.Dir(f); dir != string(filepath.Separator); dir = filepath.Dir(dir) {
		var p string
		p, err = find_xcodeproj(dir)

		if p != "" {
			pbxproj := filepath.Join(p, "project.pbxproj")
			data, err := ioutil.ReadFile(pbxproj)
			if err != nil {
				return nil, err
			}

			return &Project{p, data, filepath.Dir(p)}, nil
		}
	}

	return nil, err
}

func find_xcodeproj(dir string) (string, error) {
	m, err := filepath.Glob( filepath.Join(dir, "*.xcodeproj") )
	if err != nil {
		return "", err
	}

	if len(m) >= 1 {
		f := m[0]

		if !filepath.IsAbs(f) {
			f, err = filepath.Abs(f)
			if err != nil {
				return "", err
			}
		}

		return f, nil
	} else {
		return "", errors.New("xcodeproj not found")
	}
}

func (p *Project) Type() string {
	re := regexp.MustCompile(`SDKROOT\s*=\s*"?(.*?)"?;`)
	m := re.FindSubmatch(p.data)

	if len(m) >= 2 {
		return string(m[1])
	}

	return ""
}

func (p *Project) Cflags() []string {
	re := regexp.MustCompile(`(?s)HEADER_SEARCH_PATHS\s*=\s*\((.*?)\);`)
	m := re.FindSubmatch(p.data)

	var flags []string
	if len(m) >= 2 {
		re := regexp.MustCompile(`\s*"?(.*?)"?,`)

		m := re.FindAllStringSubmatch(string(m[1]), -1)

		for _, l := range(m) {
			if len(l) >= 2 {
				if l[1] == "$(inherited)" { // ignore
					continue
				}

				l[1] = strings.Replace(l[1], "$(SRCROOT)", p.root, -1)
				flags = append(flags, l[1])
			}
		}
	}

	return flags
}
