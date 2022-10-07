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

func GenerateServiceGroups(protocol string) {
	m := models.Module{
		Name: "servicegroups",
	}

	s := models.Section{
		Name: "trafficmanagement.loadbalancing.servicegroups",
	}

	e := models.Element{
		Name: "DUMMY_" + protocol,
		Expressions: models.Expression{
			Install:   "add servicegroup <<name>> <<type>> -healthMonitor NO",
			Uninstall: "rm servicegroup <<name>>",
		},
	}
	e.Fields = append(e.Fields, generateServiceGroupFields(protocol)...)
	s.Elements = append(s.Elements, e)

	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/core"
	filename := "servicegroup_" + protocol
	shared.WriteToFile(path, filename, d)
}

func generateServiceGroupFields(protocol string) []models.Field {
	output := make([]models.Field, 0)

	output = append(output, models.Field{
		Id:   "name",
		Data: "<<prefix>>_DUMMY_" + protocol,
	})

	output = append(output, models.Field{
		Id:   "type",
		Data: protocol,
	})

	return output
}
