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

// SilenceTimeout invokes the voicenavigator.SilenceTimeout API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/silencetimeout.html
func (client *Client) SilenceTimeout(request *SilenceTimeoutRequest) (response *SilenceTimeoutResponse, err error) {
	response = CreateSilenceTimeoutResponse()
	err = client.DoAction(request, response)
	return
}

// SilenceTimeoutWithChan invokes the voicenavigator.SilenceTimeout API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/silencetimeout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SilenceTimeoutWithChan(request *SilenceTimeoutRequest) (<-chan *SilenceTimeoutResponse, <-chan error) {
	responseChan := make(chan *SilenceTimeoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SilenceTimeout(request)
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

// SilenceTimeoutWithCallback invokes the voicenavigator.SilenceTimeout API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/silencetimeout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SilenceTimeoutWithCallback(request *SilenceTimeoutRequest, callback func(response *SilenceTimeoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SilenceTimeoutResponse
		var err error
		defer close(result)
		response, err = client.SilenceTimeout(request)
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

// SilenceTimeoutRequest is the request struct for api SilenceTimeout
type SilenceTimeoutRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// SilenceTimeoutResponse is the response struct for api SilenceTimeout
type SilenceTimeoutResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateSilenceTimeoutRequest creates a request to invoke SilenceTimeout API
func CreateSilenceTimeoutRequest() (request *SilenceTimeoutRequest) {
	request = &SilenceTimeoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "SilenceTimeout", "voicebot", "openAPI")
	return
}

// CreateSilenceTimeoutResponse creates a response to parse from SilenceTimeout response
func CreateSilenceTimeoutResponse() (response *SilenceTimeoutResponse) {
	response = &SilenceTimeoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
