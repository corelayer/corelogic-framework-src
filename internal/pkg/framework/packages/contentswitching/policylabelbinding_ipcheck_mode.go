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

func GenerateContentSwitchingPolicyLabelBindingsIpCheck(elementName string, protocol string, baseProtocol string, checkMode string, priorityPrefix string) {
	m := models.Module{
		Name: "contentswitching",
	}

	s := models.Section{
		Name: "trafficmanagement.contentswitching.policylabelbindings",
	}

	s.Elements = append(s.Elements, generateContentSwitchingPolicyLabelBindingsIpCheckElementsPerIpVersion(elementName, "ipv4", protocol, baseProtocol, 1, 32, checkMode, priorityPrefix)...)
	s.Elements = append(s.Elements, generateContentSwitchingPolicyLabelBindingsIpCheckElementsPerIpVersion(elementName, "ipv6", protocol, baseProtocol, 1, 128, checkMode, priorityPrefix)...)

	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/contentswitching"
	filename := "policylabelbinding_" + elementName + "_" + protocol + "_" + checkMode
	shared.WriteToFile(path, filename, d)
}

func generateContentSwitchingPolicyLabelBindingsIpCheckFullName(elementName string, ipVersion string, protocol string, checkMode string, subnet string) string {
	return elementName + "_" + strings.ToUpper(ipVersion) + "_" + strings.ToUpper(protocol) + "_" + strings.ToUpper(checkMode) + "_" + subnet
}

func generateReverseSubnetNumberForCidr(i int, ipVersion string) string {
	var output string
	switch ipVersion {
	case "ipv4":
		output = fmt.Sprintf("%03d", 33-i)
	case "ipv6":
		output = fmt.Sprintf("%03d", 129-i)
	}
	return output
}

func generateContentSwitchingPolicyLabelBindingsIpCheckElementsPerIpVersion(elementName string, ipVersion string, protocol string, baseProtocol string, subnetLow int, subnetHigh int, checkMode string, priorityPrefix string) []models.Element {
	output := make([]models.Element, 0, subnetHigh)

	for i := subnetHigh; i >= subnetLow; i-- {
		subnet := fmt.Sprintf("%03d", i)
		priority := fmt.Sprintf("%s%s%s", priorityPrefix, strings.TrimPrefix(ipVersion, "ipv"), generateReverseSubnetNumberForCidr(i, ipVersion))
		e := models.Element{
			Name: generateContentSwitchingPolicyLabelBindingsIpCheckFullName(elementName, ipVersion, protocol, checkMode, subnet),
		}

		switch checkMode {
		case "block":
			e.Expressions = models.Expression{
				Install:   "bind cs policylabel <<labelname>> <<policyname>> <<priority>>",
				Uninstall: "unbind cs policylabel <<labelname>>",
			}

			e.Fields = append(e.Fields, generateContentSwitchingPolicyLabelBindingsIpCheckFields(elementName, "ipv4", protocol, baseProtocol, checkMode, subnet, priority)...)
		case "allow":
			e.Expressions = models.Expression{
				Install:   "bind cs policylabel <<labelname>> <<policyname>> <<priority>> -gotoPriorityExpression <<gotoPriority>>",
				Uninstall: "unbind cs policylabel <<labelname>>",
			}

			e.Fields = append(e.Fields, generateContentSwitchingPolicyLabelBindingsIpCheckFields(elementName, "ipv4", protocol, baseProtocol, checkMode, subnet, priority)...)
		case "lan":
			e.Expressions = models.Expression{
				Install:   "bind cs policylabel <<labelname>> <<policyname>> <<priority>>",
				Uninstall: "unbind cs policylabel <<labelname>>",
			}

			e.Fields = append(e.Fields, generateContentSwitchingPolicyLabelBindingsIpCheckFields(elementName, "ipv4", protocol, baseProtocol, checkMode, subnet, priority)...)
		}

		output = append(output, e)
	}
	return output
}

func generateContentSwitchingPolicyLabelBindingsIpCheckFields(elementName string, ipVersion string, protocol string, baseProtocol string, checkMode string, subnet string, priority string) []models.Field {
	output := make([]models.Field, 0)

	output = append(output, models.Field{
		Id:   "labelname",
		Data: "<<contentswitching.trafficmanagement.contentswitching.policylabels." + getContentSwitchingPolicyLabelIpCheckFullName(elementName, protocol) + "/name>>",
	})

	output = append(output, models.Field{
		Id:   "policyname",
		Data: "<<contentswitching.trafficmanagement.contentswitching.policies." + getContentSwitchingPolicyIpCheckFullName(elementName, ipVersion, protocol, checkMode, subnet) + "/name>>",
	})

	output = append(output, models.Field{
		Id:   "priority",
		Data: priority,
	})

	if checkMode == "allow" {
		output = append(output, models.Field{
			Id:   "gotoPriority",
			Data: priority,
		})
	}

	return output
}
