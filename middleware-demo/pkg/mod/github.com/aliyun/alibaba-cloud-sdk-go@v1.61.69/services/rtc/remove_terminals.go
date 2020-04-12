package rtc

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

// RemoveTerminals invokes the rtc.RemoveTerminals API synchronously
// api document: https://help.aliyun.com/api/rtc/removeterminals.html
func (client *Client) RemoveTerminals(request *RemoveTerminalsRequest) (response *RemoveTerminalsResponse, err error) {
	response = CreateRemoveTerminalsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveTerminalsWithChan invokes the rtc.RemoveTerminals API asynchronously
// api document: https://help.aliyun.com/api/rtc/removeterminals.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveTerminalsWithChan(request *RemoveTerminalsRequest) (<-chan *RemoveTerminalsResponse, <-chan error) {
	responseChan := make(chan *RemoveTerminalsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveTerminals(request)
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

// RemoveTerminalsWithCallback invokes the rtc.RemoveTerminals API asynchronously
// api document: https://help.aliyun.com/api/rtc/removeterminals.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveTerminalsWithCallback(request *RemoveTerminalsRequest, callback func(response *RemoveTerminalsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveTerminalsResponse
		var err error
		defer close(result)
		response, err = client.RemoveTerminals(request)
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

// RemoveTerminalsRequest is the request struct for api RemoveTerminals
type RemoveTerminalsRequest struct {
	*requests.RpcRequest
	TerminalIds *[]string        `position:"Query" name:"TerminalIds"  type:"Repeated"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	ChannelId   string           `position:"Query" name:"ChannelId"`
}

// RemoveTerminalsResponse is the response struct for api RemoveTerminals
type RemoveTerminalsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Terminals Terminals `json:"Terminals" xml:"Terminals"`
}

// CreateRemoveTerminalsRequest creates a request to invoke RemoveTerminals API
func CreateRemoveTerminalsRequest() (request *RemoveTerminalsRequest) {
	request = &RemoveTerminalsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "RemoveTerminals", "", "")
	return
}

// CreateRemoveTerminalsResponse creates a response to parse from RemoveTerminals response
func CreateRemoveTerminalsResponse() (response *RemoveTerminalsResponse) {
	response = &RemoveTerminalsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
