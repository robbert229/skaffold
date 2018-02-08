/*
Copyright 2018 Google LLC

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

package flags

import (
	"fmt"
	"reflect"
	"text/template"

	"github.com/pkg/errors"
)

type TemplateFlag struct {
	rawTemplate string
	template    *template.Template
	context     interface{}
}

func (t *TemplateFlag) String() string {
	return t.rawTemplate
}

func (t *TemplateFlag) Usage() string {
	defaultUsage := "Format output with go-template."
	if t.context != nil {
		goType := reflect.TypeOf(t.context)
		url := fmt.Sprintf("https://godoc.org/%s#%s", goType.PkgPath(), goType.Name())
		defaultUsage += fmt.Sprintf(" For full struct documentation, see %s", url)
	}
	return defaultUsage
}

func (t *TemplateFlag) Set(value string) error {
	tmpl, err := template.New("flagtemplate").Parse(value)
	if err != nil {
		return errors.Wrap(err, "setting template flag")
	}
	t.rawTemplate = value
	t.template = tmpl
	return nil
}

func (t *TemplateFlag) Type() string {
	return fmt.Sprintf("%T", t)
}

func (t *TemplateFlag) Template() *template.Template {
	return t.template
}

func NewTemplateFlag(value string, context interface{}) *TemplateFlag {
	return &TemplateFlag{
		template:    template.Must(template.New("flagtemplate").Parse(value)),
		rawTemplate: value,
		context:     context,
	}
}