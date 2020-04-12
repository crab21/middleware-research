package retailcloud

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ServiceInstance is a nested struct in retailcloud response
type ServiceInstance struct {
	AppId        int64                `json:"AppId" xml:"AppId"`
	EnvId        int64                `json:"EnvId" xml:"EnvId"`
	Headless     bool                 `json:"Headless" xml:"Headless"`
	K8sServiceId string               `json:"K8sServiceId" xml:"K8sServiceId"`
	Name         string               `json:"Name" xml:"Name"`
	ServiceId    int64                `json:"ServiceId" xml:"ServiceId"`
	ServiceType  string               `json:"ServiceType" xml:"ServiceType"`
	PortMappings []ServicePortMapping `json:"PortMappings" xml:"PortMappings"`
}
