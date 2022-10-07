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

package shared

import (
	"log"
	"os"
	"strings"
)

func getFileNameWithExtension(fileName string) string {
	return fileName + ".generated.yaml"
}

func printCopyright() string {
	output := "#/*\n# * Copyright 2022 CoreLayer BV\n# *\n# *    Licensed under the Apache License, Version 2.0 (the \"License\");\n# *    you may not use this file except in compliance with the License.\n# *    You may obtain a copy of the License at\n# *\n# *        http://www.apache.org/licenses/LICENSE-2.0\n# *\n# *    Unless required by applicable law or agreed to in writing, software\n# *    distributed under the License is distributed on an \"AS IS\" BASIS,\n# *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n# *    See the License for the specific language governing permissions and\n# *    limitations under the License.\n# */\n"
	return output
}

func WriteToFile(path string, fileName string, data []byte) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
	}

	f, err := os.Create(path + "/" + strings.ToLower(getFileNameWithExtension(fileName)))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(printCopyright())
	_, err = f.Write(data)
	if err != nil {
		log.Fatal()
	}

}
