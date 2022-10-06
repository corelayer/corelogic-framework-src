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
	"strconv"
	"strings"
)

func GenerateClientIpExpressions(ipVersion string) {
	m := models.Module{
		Name: "core",
	}

	s := models.Section{
		Name: "appexpert.expressions.advanced",
	}

	eSrc := models.Element{
		Name: "CLIENT_SRC_" + strings.ToUpper(ipVersion),
	}
	eSrc.Fields = append(eSrc.Fields, models.Field{
		Id:   "name",
		Data: "<<prefix>>_CLIENT_SRC_" + strings.ToUpper(ipVersion),
	})
	eSrc.Fields = append(eSrc.Fields, models.Field{
		Id:   "data",
		Data: strings.ToUpper(getNsObjectForClientIP(ipVersion)),
	})
	s.Elements = append(s.Elements, eSrc)

	// Generate element to check whether an IP address is either IPv4 or IPv6
	eSrcIsIpVersion := models.Element{
		Name: "CLIENT_SRC_IS_" + strings.ToUpper(ipVersion),
	}
	eSrcIsIpVersion.Fields = append(eSrcIsIpVersion.Fields, models.Field{
		Id:   "name",
		Data: "<<prefix>>_CLIENT_SRC_IS_" + strings.ToUpper(ipVersion),
	})
	eSrcIsIpVersion.Fields = append(eSrcIsIpVersion.Fields, models.Field{
		Id:   "data",
		Data: "<<core.placeholders.appexpert.expressions.advanced.CLIENT_SRC_" + strings.ToUpper(ipVersion) + "/data>>.IS_" + strings.ToUpper(getOtherIpVersion(ipVersion)) + ".NOT",
	})
	s.Elements = append(s.Elements, eSrcIsIpVersion)

	// Generate element to check if IP address belongs to a subnet
	eSubnet := models.Element{
		Name: "CLIENT_SRC_" + strings.ToUpper(ipVersion) + "_SUBNET",
	}
	if ipVersion == "ipv4" {
		eSubnet.Fields = append(eSubnet.Fields, generateClientIpFields("ipv4", 1, 32)...)
	} else if ipVersion == "ipv6" {
		eSubnet.Fields = append(eSubnet.Fields, generateClientIpFields("ipv6", 1, 128)...)
	}
	s.Elements = append(s.Elements, eSubnet)

	// Add section to module
	m.Sections = append(m.Sections, s)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}

	path := "framework/packages/core"
	filename := "client_" + ipVersion
	shared.WriteToFile(path, filename, d)
	shared.AddFileToGit(path, filename)
}

func getNsObjectForClientIP(ipVersion string) string {
	output := ""
	switch ipVersion {
	case "ipv4":
		output = "CLIENT.IP.SRC"
	case "ipv6":
		output = "CLIENT.IPV6.SRC"
	}
	return output
}

func getOtherIpVersion(ipVersion string) string {
	output := ""
	switch ipVersion {
	case "ipv4":
		output = "ipv6"
	case "ipv6":
		output = "ipv4"
	}

	return output
}

func generateClientIpFields(ipVersion string, subnetLow int, subnetHigh int) []models.Field {
	output := make([]models.Field, 0, subnetHigh)
	for i := subnetHigh; i >= subnetLow; i-- {
		subnet := fmt.Sprintf("%03d", i)

		// Generate this field when subnet cidr is either /32 (IPv4) or /128 (IPv6)
		if i == subnetHigh {
			output = append(output, models.Field{
				Id:   subnet,
				Data: "<<core.placeholders.appexpert.expressions.advanced.CLIENT_SRC_" + strings.ToUpper(ipVersion) + "/data>>.TYPECAST_TEXT_T + \"/" + strconv.Itoa(i) + "\"",
			})
		} else {
			output = append(output, models.Field{
				Id:   subnet,
				Data: "<<core.placeholders.appexpert.expressions.advanced.CLIENT_SRC_" + strings.ToUpper(ipVersion) + "/data>>.SUBNET(" + strconv.Itoa(i) + ").TYPECAST_TEXT_T + \"/" + strconv.Itoa(i) + "\"",
			})
		}
	}

	return output
}
