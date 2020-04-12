package outboundbot

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

// CreateScriptWaveform invokes the outboundbot.CreateScriptWaveform API synchronously
// api document: https://help.aliyun.com/api/outboundbot/createscriptwaveform.html
func (client *Client) CreateScriptWaveform(request *CreateScriptWaveformRequest) (response *CreateScriptWaveformResponse, err error) {
	response = CreateCreateScriptWaveformResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScriptWaveformWithChan invokes the outboundbot.CreateScriptWaveform API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/createscriptwaveform.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateScriptWaveformWithChan(request *CreateScriptWaveformRequest) (<-chan *CreateScriptWaveformResponse, <-chan error) {
	responseChan := make(chan *CreateScriptWaveformResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScriptWaveform(request)
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

// CreateScriptWaveformWithCallback invokes the outboundbot.CreateScriptWaveform API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/createscriptwaveform.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateScriptWaveformWithCallback(request *CreateScriptWaveformRequest, callback func(response *CreateScriptWaveformResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScriptWaveformResponse
		var err error
		defer close(result)
		response, err = client.CreateScriptWaveform(request)
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

// CreateScriptWaveformRequest is the request struct for api CreateScriptWaveform
type CreateScriptWaveformRequest struct {
	*requests.RpcRequest
	ScriptId      string `position:"Query" name:"ScriptId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	FileName      string `position:"Query" name:"FileName"`
	ScriptContent string `position:"Query" name:"ScriptContent"`
	FileId        string `position:"Query" name:"FileId"`
}

// CreateScriptWaveformResponse is the response struct for api CreateScriptWaveform
type CreateScriptWaveformResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Success          bool   `json:"Success" xml:"Success"`
	Code             string `json:"Code" xml:"Code"`
	Message          string `json:"Message" xml:"Message"`
	HttpStatusCode   int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ScriptWaveformId string `json:"ScriptWaveformId" xml:"ScriptWaveformId"`
}

// CreateCreateScriptWaveformRequest creates a request to invoke CreateScriptWaveform API
func CreateCreateScriptWaveformRequest() (request *CreateScriptWaveformRequest) {
	request = &CreateScriptWaveformRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateScriptWaveform", "outboundbot", "openAPI")
	return
}

// CreateCreateScriptWaveformResponse creates a response to parse from CreateScriptWaveform response
func CreateCreateScriptWaveformResponse() (response *CreateScriptWaveformResponse) {
	response = &CreateScriptWaveformResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
