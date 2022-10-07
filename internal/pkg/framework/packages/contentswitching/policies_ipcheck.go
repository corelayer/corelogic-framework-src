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
	"fmt"
	shared "github.com/corelayer/corelogic-framework-src/internal/pkg"
	"github.com/corelayer/go-corelogic-framework-models/pkg/models"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

func GenerateContentSwitchingPoliciesIpCheck(elementName string, protocol string, baseProtocol string, filterMode string) {
	m := models.Module{
		Name: "contentswitching",
	}

	s := models.Section{
		Name: "trafficmanagement.contentswitching.policies",
	}

	s.Elements = append(s.Elements, generateContentSwitchingPolicyIpCheckElements(elementName, "ipv4", protocol, baseProtocol, 1, 32, filterMode)...)
	s.Elements = append(s.Elements, generateContentSwitchingPolicyIpCheckElements(elementName, "ipv6", protocol, baseProtocol, 1, 128, filterMode)...)

	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/contentswitching"
	filename := "policies_" + elementName + "_" + protocol + "_" + filterMode
	shared.WriteToFile(path, filename, d)
	//shared.AddFileToGit(path, filename)
}

func getContentSwitchingPolicyIpCheckFullName(elementName string, ipVersion string, protocol string, checkMode string, subnet string) string {
	return elementName + "_" + strings.ToUpper(ipVersion) + "_" + strings.ToUpper(protocol) + "_" + strings.ToUpper(checkMode) + "_" + subnet
}

func generateContentSwitchingPolicyIpCheckElements(elementName string, ipVersion string, protocol string, baseProtocol string, subnetLow int, subnetHigh int, checkMode string) []models.Element {
	output := make([]models.Element, 0, subnetHigh)

	for i := subnetHigh; i >= subnetLow; i-- {
		subnet := fmt.Sprintf("%03d", i)

		e := models.Element{
			Name: getContentSwitchingPolicyIpCheckFullName(elementName, ipVersion, protocol, checkMode, subnet),
		}

		switch checkMode {
		case "block":
			e.Expressions = models.Expression{
				Install:   "add cs policy <<name>> -rule q{<<expression>>} -action <<action>>",
				Uninstall: "rm cs policy <<name>>",
			}

			e.Fields = append(e.Fields, generateContentSwitchingPolicyIpCheckFields(elementName, "ipv4", protocol, baseProtocol, checkMode, subnet)...)
		case "allow":
			e.Expressions = models.Expression{
				Install:   "add cs policy <<name>> -rule q{<<expression>>}",
				Uninstall: "rm cs policy <<name>>",
			}

			e.Fields = append(e.Fields, generateContentSwitchingPolicyIpCheckFields(elementName, "ipv4", protocol, baseProtocol, checkMode, subnet)...)
		case "lan":
			e.Expressions = models.Expression{
				Install:   "add cs policy <<name>> -rule q{<<expression>>}",
				Uninstall: "rm cs policy <<name>>",
			}

			e.Fields = append(e.Fields, generateContentSwitchingPolicyIpCheckFields(elementName, "ipv4", protocol, baseProtocol, checkMode, subnet)...)
		}

		output = append(output, e)
	}
	return output
}

func generateContentSwitchingPolicyIpCheckFields(elementName string, ipVersion string, protocol string, baseProtocol string, checkMode string, subnet string) []models.Field {
	output := make([]models.Field, 0)

	output = append(output, models.Field{
		Id:   "name",
		Data: "<<prefix>>_" + getContentSwitchingPolicyIpCheckFullName(elementName, ipVersion, protocol, checkMode, subnet),
	})

	output = append(output, models.Field{
		Id:   "expression",
		Data: "<<core.stringmaps.appexpert.stringmaps." + elementName + "/" + strings.ToLower(ipVersion) + "_" + strings.ToLower(baseProtocol) + "_key_exists_" + subnet + ">>",
	})

	if checkMode == "block" {
		output = append(output, models.Field{
			Id:   "action",
			Data: "<<contentswitching.trafficmanagement.contentswitching.actions." + elementName + "_" + strings.ToUpper(protocol) + "/name>>",
		})
	}

	return output
}
