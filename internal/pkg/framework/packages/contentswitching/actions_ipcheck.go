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

package contentswitching

import (
	shared "github.com/corelayer/corelogic-framework-src/internal/pkg"
	"github.com/corelayer/go-corelogic-framework-models/pkg/models"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

func GenerateContentSwitchingActionsIpCheck(elementName string, protocol string) {
	m := models.Module{
		Name: "contentswitching",
	}

	s := models.Section{
		Name: "trafficmanagement.contentswitching.actions",
	}

	e := models.Element{
		Name: getContentSwitchingActionIpCheckFullName(elementName, protocol),
		Expressions: models.Expression{
			Install:   "add cs action <<name>> -targetVserverExpr <<expression>>",
			Uninstall: "rm cs action <<name",
		},
	}

	e.Fields = append(e.Fields, models.Field{
		Id:   "name",
		Data: "<<prefix>>_" + getContentSwitchingActionIpCheckFullName(elementName, protocol),
	})

	e.Fields = append(e.Fields, models.Field{
		Id:   "expression",
		Data: "<<loadbalancing.trafficmanagement.loadbalancing.virtualservers." + getContentSwitchingActionIpCheckFullName(elementName, protocol) + "/name>>",
	})

	s.Elements = append(s.Elements, e)
	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/contentswitching"
	filename := "actions_" + elementName + "_" + protocol
	shared.WriteToFile(path, filename, d)
	//shared.AddFileToGit(path, filename)
}

func getContentSwitchingActionIpCheckFullName(elementName string, protocol string) string {
	return elementName + "_" + strings.ToUpper(protocol)
}
