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
	"github.com/corelayer/corelogic-framework-src/internal/pkg/framework/packages/core"
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
	core.GenerateStringmapIpFilter("CSV_IPFILTER", "CS_VSERVER")
	core.GenerateStringmapIpFilter("LBV_IPFILTER", "LB_VSERVER")
}
