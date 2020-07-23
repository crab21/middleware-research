package ehpc

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

// AddLocalNodes invokes the ehpc.AddLocalNodes API synchronously
// api document: https://help.aliyun.com/api/ehpc/addlocalnodes.html
func (client *Client) AddLocalNodes(request *AddLocalNodesRequest) (response *AddLocalNodesResponse, err error) {
	response = CreateAddLocalNodesResponse()
	err = client.DoAction(request, response)
	return
}

// AddLocalNodesWithChan invokes the ehpc.AddLocalNodes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/addlocalnodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLocalNodesWithChan(request *AddLocalNodesRequest) (<-chan *AddLocalNodesResponse, <-chan error) {
	responseChan := make(chan *AddLocalNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLocalNodes(request)
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

// AddLocalNodesWithCallback invokes the ehpc.AddLocalNodes API asynchronously
// api document: https://help.aliyun.com/api/ehpc/addlocalnodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLocalNodesWithCallback(request *AddLocalNodesRequest, callback func(response *AddLocalNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLocalNodesResponse
		var err error
		defer close(result)
		response, err = client.AddLocalNodes(request)
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

// AddLocalNodesRequest is the request struct for api AddLocalNodes
type AddLocalNodesRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
	Nodes     string `position:"Query" name:"Nodes"`
}

// AddLocalNodesResponse is the response struct for api AddLocalNodes
type AddLocalNodesResponse struct {
	*responses.BaseResponse
	RequestId   string                     `json:"RequestId" xml:"RequestId"`
	InstanceIds InstanceIdsInAddLocalNodes `json:"InstanceIds" xml:"InstanceIds"`
}

// CreateAddLocalNodesRequest creates a request to invoke AddLocalNodes API
func CreateAddLocalNodesRequest() (request *AddLocalNodesRequest) {
	request = &AddLocalNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "AddLocalNodes", "ehs", "openAPI")
	return
}

// CreateAddLocalNodesResponse creates a response to parse from AddLocalNodes response
func CreateAddLocalNodesResponse() (response *AddLocalNodesResponse) {
	response = &AddLocalNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}