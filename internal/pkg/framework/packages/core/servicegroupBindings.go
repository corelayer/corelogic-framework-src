/*
 * Copyright 2022 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package core

import (
	shared "github.com/corelayer/corelogic-framework-src/internal/pkg"
	"github.com/corelayer/go-corelogic-framework-models/pkg/models"
	"gopkg.in/yaml.v3"
	"log"
)

func GenerateServiceGroupBindings(protocol string) {
	m := models.Module{
		Name:     "servicegroups",
		Tags:     nil,
		Sections: nil,
	}

	s := models.Section{
		Name:     "trafficmanagement.loadbalancing.servicegroupbindings",
		Elements: nil,
	}

	e := models.Element{
		Name:   "DUMMY_" + protocol,
		Tags:   nil,
		Fields: nil,
		Expressions: models.Expression{
			Install:   "bind servicegroup <<servicegroup>> <<server>> *",
			Uninstall: "unbind servicegroup <<servicegroup>> <server>> *",
		},
	}
	e.Fields = append(e.Fields, generateServiceGroupBindingFields(protocol)...)
	s.Elements = append(s.Elements, e)

	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/core"
	filename := "servicegroupbindings_" + protocol
	shared.WriteToFile(path, filename, d)
	//shared.AddFileToGit(path, filename)
}

func generateServiceGroupBindingFields(protocol string) []models.Field {
	output := make([]models.Field, 0)

	output = append(output, models.Field{
		Id:   "name",
		Data: "<<prefix>>_DUMMY_" + protocol,
	})

	output = append(output, models.Field{
		Id:   "servicegroup",
		Data: "<<core.servicegroups.trafficmanagement.loadbalancing.servicegroups.DUMMY_" + protocol + "/name>",
	})

	output = append(output, models.Field{
		Id:   "server",
		Data: "<<core.servers.trafficmanagement.loadbalancing.servers.DUMMY/name>>",
	})

	return output
}
