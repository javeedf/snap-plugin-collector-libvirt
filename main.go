/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (

	// Import the snap plugin library
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	// Import our collector plugin implementation
	"github.com/intelsdi-x/snap-plugin-collector-libvirt/libvirtcollector"
)

// plugin bootstrap
func main() {
	plugin.StartCollector(libvirtcollector.LibvirtCollector{}, libvirtcollector.Name, libvirtcollector.Version)

}
