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
	"fmt"
	shared "github.com/corelayer/corelogic-framework-src/internal/pkg"
	"github.com/corelayer/go-corelogic-framework-models/pkg/models"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

func GenerateStringmapIpFilter(elementName string, nsObject string) {
	m := models.Module{
		Name: "stringmaps",
	}

	s := models.Section{
		Name: "appexpert.stringmaps",
	}

	e := models.Element{
		Name: elementName,
		Expressions: models.Expression{
			Install:   "add policy stringmap <<name>>",
			Uninstall: "rm policy stringmap <<name>>",
		},
	}

	e.Fields = append(e.Fields, models.Field{
		Id:   "name",
		Data: "<<prefix>>_" + elementName,
	})

	e.Fields = append(e.Fields, generateStringmapIpFilterFields("ipv4", "tcp", 1, 32, nsObject)...)
	e.Fields = append(e.Fields, generateStringmapIpFilterFields("ipv6", "tcp", 1, 128, nsObject)...)
	e.Fields = append(e.Fields, generateStringmapIpFilterFields("ipv4", "udp", 1, 32, nsObject)...)
	e.Fields = append(e.Fields, generateStringmapIpFilterFields("ipv6", "udp", 1, 128, nsObject)...)

	s.Elements = append(s.Elements, e)
	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/core"
	filename := "stringmap_" + elementName
	shared.WriteToFile(path, filename, d)
	shared.AddFileToGit(path, filename)
}

func generateStringmapIpFilterFields(ipVersion string, protocol string, subnetLow int, subnetHigh int, nsObject string) []models.Field {
	output := make([]models.Field, 0, subnetHigh)
	for i := subnetHigh; i >= subnetLow; i-- {
		subnet := fmt.Sprintf("%03d", i)

		output = append(output, models.Field{
			Id:   strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_key_" + subnet,
			Data: "(\"csv=\" + CLIENT." + strings.ToUpper(protocol) + "." + nsObject + ".NAME + \";address=\" + <<core.placeholders.appexpert.expressions.advanced.CLIENT_SRC_" + strings.ToUpper(ipVersion) + "_SUBNET/" + subnet + ">> + \";\").SET_TEXT_MODE(NOIGNORECASE)",
		})

		output = append(output, models.Field{
			Id:   strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_key_exists_" + subnet,
			Data: "<<core.stringmaps.appexpert.stringmaps.CSV_IPFILTER/" + strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_key_" + subnet + ">>.IS_STRINGMAP_KEY(\"<<core.stringmaps.appexpert.stringmaps.CSV_IPFILTER/name>>\")",
		})

		output = append(output, models.Field{
			Id:   strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_value_" + subnet,
			Data: "<<core.stringmaps.appexpert.stringmaps.CSV_IPFILTER/" + strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_key_" + subnet + ">>.MAP_STRING(\"<<core.stringmaps.appexpert.stringmaps.CSV_IPFILTER/name>>\").TYPECAST_NVLIST_T(';','=')",
		})

		output = append(output, models.Field{
			Id:   strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_value_description_" + subnet,
			Data: "<<core.stringmaps.appexpert.stringmaps.CSV_IPFILTER/" + strings.ToLower(ipVersion) + "_" + strings.ToLower(protocol) + "_value_" + subnet + ">>.VALUE(\"description\")",
		})
	}

	return output
}
