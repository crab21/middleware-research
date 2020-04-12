package cr_ee

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateRepoBuildRule invokes the cr.UpdateRepoBuildRule API synchronously
// api document: https://help.aliyun.com/api/cr/updaterepobuildrule.html
func (client *Client) UpdateRepoBuildRule(request *UpdateRepoBuildRuleRequest) (response *UpdateRepoBuildRuleResponse, err error) {
	response = CreateUpdateRepoBuildRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRepoBuildRuleWithChan invokes the cr.UpdateRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/updaterepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRepoBuildRuleWithChan(request *UpdateRepoBuildRuleRequest) (<-chan *UpdateRepoBuildRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateRepoBuildRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRepoBuildRule(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateRepoBuildRuleWithCallback invokes the cr.UpdateRepoBuildRule API asynchronously
// api document: https://help.aliyun.com/api/cr/updaterepobuildrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRepoBuildRuleWithCallback(request *UpdateRepoBuildRuleRequest, callback func(response *UpdateRepoBuildRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRepoBuildRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateRepoBuildRule(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateRepoBuildRuleRequest is the request struct for api UpdateRepoBuildRule
type UpdateRepoBuildRuleRequest struct {
	*requests.RpcRequest
	RepoId             string `position:"Query" name:"RepoId"`
	PushName           string `position:"Query" name:"PushName"`
	DockerfileName     string `position:"Query" name:"DockerfileName"`
	DockerfileLocation string `position:"Query" name:"DockerfileLocation"`
	BuildRuleId        string `position:"Query" name:"BuildRuleId"`
	InstanceId         string `position:"Query" name:"InstanceId"`
	ImageTag           string `position:"Query" name:"ImageTag"`
	PushType           string `position:"Query" name:"PushType"`
}

// UpdateRepoBuildRuleResponse is the response struct for api UpdateRepoBuildRule
type UpdateRepoBuildRuleResponse struct {
	*responses.BaseResponse
	UpdateRepoBuildRuleIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                         string `json:"Code" xml:"Code"`
	RequestId                    string `json:"RequestId" xml:"RequestId"`
	BuildRuleId                  string `json:"BuildRuleId" xml:"BuildRuleId"`
}

// CreateUpdateRepoBuildRuleRequest creates a request to invoke UpdateRepoBuildRule API
func CreateUpdateRepoBuildRuleRequest() (request *UpdateRepoBuildRuleRequest) {
	request = &UpdateRepoBuildRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "UpdateRepoBuildRule", "cr", "openAPI")
	return
}

// CreateUpdateRepoBuildRuleResponse creates a response to parse from UpdateRepoBuildRule response
func CreateUpdateRepoBuildRuleResponse() (response *UpdateRepoBuildRuleResponse) {
	response = &UpdateRepoBuildRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
