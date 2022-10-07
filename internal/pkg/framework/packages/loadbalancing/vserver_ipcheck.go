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

package loadbalancing

import (
	shared "github.com/corelayer/corelogic-framework-src/internal/pkg"
	"github.com/corelayer/go-corelogic-framework-models/pkg/models"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

func GenerateVserverIpCheck(elementName string, protocol string) {
	m := models.Module{
		Name: "loadbalancing",
	}

	s := models.Section{
		Name: "trafficmanagement.loadbalancing.virtualservers",
	}

	e := models.Element{
		Name: elementName + "_" + strings.ToUpper(protocol),
		Expressions: models.Expression{
			Install:   "add lb vserver <<name>> <<protocol>> 0.0.0.0 0",
			Uninstall: "rm lb vserver <<name>>",
		},
	}

	e.Fields = append(e.Fields, models.Field{
		Id:   "name",
		Data: "<<prefix>>_" + elementName + "_" + strings.ToUpper(protocol),
	})

	e.Fields = append(e.Fields, models.Field{
		Id:   "protocol",
		Data: strings.ToUpper(protocol),
	})

	s.Elements = append(s.Elements, e)
	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/loadbalancing"
	filename := "vserver_" + elementName + "_" + protocol
	shared.WriteToFile(path, filename, d)
}
