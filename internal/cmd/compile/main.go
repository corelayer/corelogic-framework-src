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
	"fmt"
	"github.com/corelayer/go-corelogic-framework-models/pkg/controllers"
	"log"
	"sort"
)

func main() {
	c := controllers.FrameworkLoader{}
	f, err := c.LoadFromDisk("framework")
	if err != nil {
		log.Fatal(err)
	}

	elements, err := f.GetElements()
	if err != nil {
		log.Fatal(err)
	}
	elementKeys := make([]string, len(elements))
	for key := range elements {
		elementKeys = append(elementKeys, key)
	}
	sort.Strings(elementKeys)
	for _, key := range elementKeys {
		fmt.Println(key, elements[key])
	}
	fmt.Println("-----------------------------")
	//fields, err := f.GetFields()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fieldKeys := make([]string, len(fields))
	//for key := range fields {
	//	fieldKeys = append(fieldKeys, key)
	//}
	//sort.Strings(fieldKeys)
	//for _, key := range fieldKeys {
	//	fmt.Println(key, fields[key])
	//}
}
