package alikafka

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

// DescribeAcls invokes the alikafka.DescribeAcls API synchronously
// api document: https://help.aliyun.com/api/alikafka/describeacls.html
func (client *Client) DescribeAcls(request *DescribeAclsRequest) (response *DescribeAclsResponse, err error) {
	response = CreateDescribeAclsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAclsWithChan invokes the alikafka.DescribeAcls API asynchronously
// api document: https://help.aliyun.com/api/alikafka/describeacls.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAclsWithChan(request *DescribeAclsRequest) (<-chan *DescribeAclsResponse, <-chan error) {
	responseChan := make(chan *DescribeAclsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAcls(request)
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

// DescribeAclsWithCallback invokes the alikafka.DescribeAcls API asynchronously
// api document: https://help.aliyun.com/api/alikafka/describeacls.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAclsWithCallback(request *DescribeAclsRequest, callback func(response *DescribeAclsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAclsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAcls(request)
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

// DescribeAclsRequest is the request struct for api DescribeAcls
type DescribeAclsRequest struct {
	*requests.RpcRequest
	AclResourceType string `position:"Query" name:"AclResourceType"`
	AclResourceName string `position:"Query" name:"AclResourceName"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	Username        string `position:"Query" name:"Username"`
}

// DescribeAclsResponse is the response struct for api DescribeAcls
type DescribeAclsResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      bool         `json:"Success" xml:"Success"`
	Code         int          `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	KafkaAclList KafkaAclList `json:"KafkaAclList" xml:"KafkaAclList"`
}

// CreateDescribeAclsRequest creates a request to invoke DescribeAcls API
func CreateDescribeAclsRequest() (request *DescribeAclsRequest) {
	request = &DescribeAclsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "DescribeAcls", "alikafka", "openAPI")
	return
}

// CreateDescribeAclsResponse creates a response to parse from DescribeAcls response
func CreateDescribeAclsResponse() (response *DescribeAclsResponse) {
	response = &DescribeAclsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
