package imm

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

// CreateGrabFrameTask invokes the imm.CreateGrabFrameTask API synchronously
// api document: https://help.aliyun.com/api/imm/creategrabframetask.html
func (client *Client) CreateGrabFrameTask(request *CreateGrabFrameTaskRequest) (response *CreateGrabFrameTaskResponse, err error) {
	response = CreateCreateGrabFrameTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGrabFrameTaskWithChan invokes the imm.CreateGrabFrameTask API asynchronously
// api document: https://help.aliyun.com/api/imm/creategrabframetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGrabFrameTaskWithChan(request *CreateGrabFrameTaskRequest) (<-chan *CreateGrabFrameTaskResponse, <-chan error) {
	responseChan := make(chan *CreateGrabFrameTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGrabFrameTask(request)
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

// CreateGrabFrameTaskWithCallback invokes the imm.CreateGrabFrameTask API asynchronously
// api document: https://help.aliyun.com/api/imm/creategrabframetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGrabFrameTaskWithCallback(request *CreateGrabFrameTaskRequest, callback func(response *CreateGrabFrameTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGrabFrameTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateGrabFrameTask(request)
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

// CreateGrabFrameTaskRequest is the request struct for api CreateGrabFrameTask
type CreateGrabFrameTaskRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	CustomMessage   string `position:"Query" name:"CustomMessage"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	TargetList      string `position:"Query" name:"TargetList"`
	VideoUri        string `position:"Query" name:"VideoUri"`
}

// CreateGrabFrameTaskResponse is the response struct for api CreateGrabFrameTask
type CreateGrabFrameTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateGrabFrameTaskRequest creates a request to invoke CreateGrabFrameTask API
func CreateCreateGrabFrameTaskRequest() (request *CreateGrabFrameTaskRequest) {
	request = &CreateGrabFrameTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateGrabFrameTask", "imm", "openAPI")
	return
}

// CreateCreateGrabFrameTaskResponse creates a response to parse from CreateGrabFrameTask response
func CreateCreateGrabFrameTaskResponse() (response *CreateGrabFrameTaskResponse) {
	response = &CreateGrabFrameTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
