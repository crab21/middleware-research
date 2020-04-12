package voicenavigator

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

// ModifyKnowledge invokes the voicenavigator.ModifyKnowledge API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyknowledge.html
func (client *Client) ModifyKnowledge(request *ModifyKnowledgeRequest) (response *ModifyKnowledgeResponse, err error) {
	response = CreateModifyKnowledgeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyKnowledgeWithChan invokes the voicenavigator.ModifyKnowledge API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyknowledge.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyKnowledgeWithChan(request *ModifyKnowledgeRequest) (<-chan *ModifyKnowledgeResponse, <-chan error) {
	responseChan := make(chan *ModifyKnowledgeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyKnowledge(request)
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

// ModifyKnowledgeWithCallback invokes the voicenavigator.ModifyKnowledge API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyknowledge.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyKnowledgeWithCallback(request *ModifyKnowledgeRequest, callback func(response *ModifyKnowledgeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyKnowledgeResponse
		var err error
		defer close(result)
		response, err = client.ModifyKnowledge(request)
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

// ModifyKnowledgeRequest is the request struct for api ModifyKnowledge
type ModifyKnowledgeRequest struct {
	*requests.RpcRequest
	UserUtterance      string           `position:"Query" name:"UserUtterance"`
	Interruptible      requests.Boolean `position:"Query" name:"Interruptible"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
	NavigationScriptId string           `position:"Query" name:"NavigationScriptId"`
	Answer             string           `position:"Query" name:"Answer"`
	SimilarUtterances  *[]string        `position:"Query" name:"SimilarUtterances"  type:"Repeated"`
}

// ModifyKnowledgeResponse is the response struct for api ModifyKnowledge
type ModifyKnowledgeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyKnowledgeRequest creates a request to invoke ModifyKnowledge API
func CreateModifyKnowledgeRequest() (request *ModifyKnowledgeRequest) {
	request = &ModifyKnowledgeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "ModifyKnowledge", "voicebot", "openAPI")
	return
}

// CreateModifyKnowledgeResponse creates a response to parse from ModifyKnowledge response
func CreateModifyKnowledgeResponse() (response *ModifyKnowledgeResponse) {
	response = &ModifyKnowledgeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
