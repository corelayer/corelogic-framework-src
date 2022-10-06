#/*
# * Copyright 2022 CoreLayer BV
# *
# *    Licensed under the Apache License, Version 2.0 (the "License");
# *    you may not use this file except in compliance with the License.
# *    You may obtain a copy of the License at
# *
# *        http://www.apache.org/licenses/LICENSE-2.0
# *
# *    Unless required by applicable law or agreed to in writing, software
# *    distributed under the License is distributed on an "AS IS" BASIS,
# *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# *    See the License for the specific language governing permissions and
# *    limitations under the License.
# */

build:
	sh scripts/build.sh

clean:
	sh scripts/clean.sh

#coverage:
#	sh scripts/coverage.sh

generate:
	./output/generate

pre-commit:
	make clean

#test:
#	sh scripts/test.sh

verify:
	./output/verify
