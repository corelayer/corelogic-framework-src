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

package main

import (
	"github.com/corelayer/corelogic-framework-src/internal/pkg/framework/packages/contentswitching"
	"github.com/corelayer/corelogic-framework-src/internal/pkg/framework/packages/core"
	"github.com/corelayer/corelogic-framework-src/internal/pkg/framework/packages/loadbalancing"
)

func main() {
	core.GenerateClientIpExpressions("ipv4")
	core.GenerateClientIpExpressions("ipv6")

	core.GenerateServiceGroups("HTTP")
	core.GenerateServiceGroupBindings("HTTP")
	core.GenerateServiceGroups("TCP")
	core.GenerateServiceGroupBindings("TCP")
	core.GenerateServiceGroups("UDP")
	core.GenerateServiceGroupBindings("UDP")
	core.GenerateStringmapIpCheck("CSV_IPFILTER", "CS_VSERVER")
	core.GenerateStringmapIpCheck("CSV_IPZONE", "CS_VSERVER")
	core.GenerateStringmapIpCheck("LBV_IPFILTER", "LB_VSERVER")

	loadbalancing.GenerateVserverIpCheck("CSV_IPFILTER", "http")
	loadbalancing.GenerateVserverIpCheck("CSV_IPFILTER", "tcp")
	loadbalancing.GenerateVserverIpCheck("CSV_IPFILTER", "udp")

	loadbalancing.GenerateVserverIpCheck("LBV_IPFILTER", "http")
	loadbalancing.GenerateVserverIpCheck("LBV_IPFILTER", "tcp")
	loadbalancing.GenerateVserverIpCheck("LBV_IPFILTER", "udp")

	contentswitching.GenerateContentSwitchingActionsIpCheck("CSV_IPFILTER", "http")
	contentswitching.GenerateContentSwitchingActionsIpCheck("CSV_IPFILTER", "tcp")
	contentswitching.GenerateContentSwitchingActionsIpCheck("CSV_IPFILTER", "udp")

	contentswitching.GenerateContentSwitchingActionsIpCheck("LBV_IPFILTER", "http")
	contentswitching.GenerateContentSwitchingActionsIpCheck("LBV_IPFILTER", "tcp")
	contentswitching.GenerateContentSwitchingActionsIpCheck("LBV_IPFILTER", "udp")

	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPFILTER", "http", "tcp", "allow")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPFILTER", "tcp", "tcp", "allow")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPFILTER", "udp", "udp", "allow")

	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPFILTER", "http", "tcp", "block")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPFILTER", "tcp", "tcp", "block")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPFILTER", "udp", "udp", "block")

	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPZONE", "http", "tcp", "lan")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPZONE", "tcp", "tcp", "lan")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("CSV_IPZONE", "udp", "udp", "lan")

	contentswitching.GenerateContentSwitchingPoliciesIpCheck("LBV_IPFILTER", "http", "tcp", "allow")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("LBV_IPFILTER", "tcp", "tcp", "allow")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("LBV_IPFILTER", "udp", "udp", "allow")

	contentswitching.GenerateContentSwitchingPoliciesIpCheck("LBV_IPFILTER", "http", "tcp", "block")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("LBV_IPFILTER", "tcp", "tcp", "block")
	contentswitching.GenerateContentSwitchingPoliciesIpCheck("LBV_IPFILTER", "udp", "udp", "block")

	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("CSV_IPFILTER", "http")
	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("CSV_IPFILTER", "tcp")
	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("CSV_IPFILTER", "udp")

	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("CSV_IPZONE", "http")
	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("CSV_IPZONE", "tcp")
	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("CSV_IPZONE", "udp")

	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("LBV_IPFILTER", "http")
	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("LBV_IPFILTER", "tcp")
	contentswitching.GenerateContentSwitchingPolicyLabelsIpCheck("LBV_IPFILTER", "udp")

	contentswitching.GenerateContentSwitchingPolicyLabelBindingsIpCheck("CSV_IPFILTER", "http", "tcp", "allow", "101")
	contentswitching.GenerateContentSwitchingPolicyLabelBindingsIpCheck("CSV_IPFILTER", "http", "tcp", "block", "101")

	//responder.GenerateResponderIpCheck("CSV_IPFILTER", "http", "CS_VSERVER")
	//responder.GenerateResponderIpCheck("CSV_IPFILTER", "tcp", "CS_VSERVER")
	//responder.GenerateResponderIpCheck("CSV_IPFILTER", "udp", "CS_VSERVER")
	//
	//responder.GenerateResponderIpCheck("LBV_IPFILTER", "http", "LB_VSERVER")
	//responder.GenerateResponderIpCheck("LBV_IPFILTER", "tcp", "LB_VSERVER")
	//responder.GenerateResponderIpCheck("LBV_IPFILTER", "udp", "LB_VSERVER")

}
