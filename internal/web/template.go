package web

import (
	"html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/yargevad/filepathx"
)

type Template = template.Template

const (
	templateFileSuffix   = ".html"
	templateIgnorePrefix = "__"
)

type templateParser struct{}

func (templateParser) isIgnoreFile(filePath string) bool {
	return strings.HasPrefix(filepath.Base(filePath), templateIgnorePrefix)
}

func (templateParser) isIgnoreDirectory(filePath string) bool {
	for _, dir := range strings.Split(filepath.ToSlash(filepath.Dir(filePath)), "/") {
		if strings.HasPrefix(dir, templateIgnorePrefix) {
			return true
		}
	}
	return false
}

func (p templateParser) templateName(basePath, filePath string) string {
	return strings.TrimPrefix(strings.TrimSuffix(filePath, templateFileSuffix), basePath+"/")
}

func (p templateParser) files(basePath string) ([]string, error) {
	globPath := path.Join(basePath, "**", "*"+templateFileSuffix)
	files, err := filepathx.Glob(globPath)
	result := make([]string, 0, len(files))
	for _, file := range files {
		if p.isIgnoreDirectory(file) || p.isIgnoreFile(file) {
			continue
		}
		result = append(result, filepath.ToSlash(file))
	}
	return result, err
}

func (p templateParser) parse(rootTemplate *Template, includeTargetPath bool, rootPath, targetPath string) (*Template, error) {
	basePath := path.Join(rootPath, targetPath)
	files, err := p.files(basePath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		t, err := template.ParseFiles(file)
		if err != nil {
			return nil, err
		}
		templateName := p.templateName(basePath, file)
		if includeTargetPath {
			templateName = path.Join(targetPath, templateName)
		}
		if rootTemplate == nil {
			rootTemplate = template.New(templateName)
		}
		_, err = rootTemplate.AddParseTree(templateName, t.Tree)
		if err != nil {
			return nil, err
		}
	}
	return rootTemplate, nil
}

func ParseTemplate(rootPath string, includeTargetPath bool, targetPath ...string) (template *Template, err error) {
	parser := templateParser{}
	for _, target := range targetPath {
		template, err = parser.parse(template, includeTargetPath, rootPath, target)
		if err != nil {
			return nil, err
		}
	}
	return template, nil
}
