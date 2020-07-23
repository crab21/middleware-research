package slb

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

// DescribeMainSubordinateServerGroupAttribute invokes the slb.DescribeMainSubordinateServerGroupAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describemasterslaveservergroupattribute.html
func (client *Client) DescribeMainSubordinateServerGroupAttribute(request *DescribeMainSubordinateServerGroupAttributeRequest) (response *DescribeMainSubordinateServerGroupAttributeResponse, err error) {
	response = CreateDescribeMainSubordinateServerGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMainSubordinateServerGroupAttributeWithChan invokes the slb.DescribeMainSubordinateServerGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describemasterslaveservergroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMainSubordinateServerGroupAttributeWithChan(request *DescribeMainSubordinateServerGroupAttributeRequest) (<-chan *DescribeMainSubordinateServerGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeMainSubordinateServerGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMainSubordinateServerGroupAttribute(request)
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

// DescribeMainSubordinateServerGroupAttributeWithCallback invokes the slb.DescribeMainSubordinateServerGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describemasterslaveservergroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMainSubordinateServerGroupAttributeWithCallback(request *DescribeMainSubordinateServerGroupAttributeRequest, callback func(response *DescribeMainSubordinateServerGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMainSubordinateServerGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeMainSubordinateServerGroupAttribute(request)
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

// DescribeMainSubordinateServerGroupAttributeRequest is the request struct for api DescribeMainSubordinateServerGroupAttribute
type DescribeMainSubordinateServerGroupAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId              string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	Tags                     string           `position:"Query" name:"Tags"`
	MainSubordinateServerGroupId string           `position:"Query" name:"MainSubordinateServerGroupId"`
}

// DescribeMainSubordinateServerGroupAttributeResponse is the response struct for api DescribeMainSubordinateServerGroupAttribute
type DescribeMainSubordinateServerGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId                  string                                                             `json:"RequestId" xml:"RequestId"`
	LoadBalancerId             string                                                             `json:"LoadBalancerId" xml:"LoadBalancerId"`
	MainSubordinateServerGroupId   string                                                             `json:"MainSubordinateServerGroupId" xml:"MainSubordinateServerGroupId"`
	MainSubordinateServerGroupName string                                                             `json:"MainSubordinateServerGroupName" xml:"MainSubordinateServerGroupName"`
	MainSubordinateBackendServers  MainSubordinateBackendServersInDescribeMainSubordinateServerGroupAttribute `json:"MainSubordinateBackendServers" xml:"MainSubordinateBackendServers"`
}

// CreateDescribeMainSubordinateServerGroupAttributeRequest creates a request to invoke DescribeMainSubordinateServerGroupAttribute API
func CreateDescribeMainSubordinateServerGroupAttributeRequest() (request *DescribeMainSubordinateServerGroupAttributeRequest) {
	request = &DescribeMainSubordinateServerGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeMainSubordinateServerGroupAttribute", "slb", "openAPI")
	return
}

// CreateDescribeMainSubordinateServerGroupAttributeResponse creates a response to parse from DescribeMainSubordinateServerGroupAttribute response
func CreateDescribeMainSubordinateServerGroupAttributeResponse() (response *DescribeMainSubordinateServerGroupAttributeResponse) {
	response = &DescribeMainSubordinateServerGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}