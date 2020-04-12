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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateCluster invokes the retailcloud.CreateCluster API synchronously
// api document: https://help.aliyun.com/api/retailcloud/createcluster.html
func (client *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	response = CreateCreateClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateClusterWithChan invokes the retailcloud.CreateCluster API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/createcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateClusterWithChan(request *CreateClusterRequest) (<-chan *CreateClusterResponse, <-chan error) {
	responseChan := make(chan *CreateClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCluster(request)
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

// CreateClusterWithCallback invokes the retailcloud.CreateCluster API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/createcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateClusterWithCallback(request *CreateClusterRequest, callback func(response *CreateClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateCluster(request)
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

// CreateClusterRequest is the request struct for api CreateCluster
type CreateClusterRequest struct {
	*requests.RpcRequest
	BusinessCode              string           `position:"Query" name:"BusinessCode"`
	CreateWithLogIntegration  requests.Boolean `position:"Query" name:"CreateWithLogIntegration"`
	Vswitchids                *[]string        `position:"Query" name:"Vswitchids"  type:"Repeated"`
	CloudMonitorFlags         requests.Integer `position:"Query" name:"CloudMonitorFlags"`
	ClusterEnvType            string           `position:"Query" name:"ClusterEnvType"`
	CreateWithArmsIntegration requests.Boolean `position:"Query" name:"CreateWithArmsIntegration"`
	KeyPair                   string           `position:"Query" name:"KeyPair"`
	ClusterTitle              string           `position:"Query" name:"ClusterTitle"`
	PodCIDR                   string           `position:"Query" name:"PodCIDR"`
	ClusterId                 requests.Integer `position:"Query" name:"ClusterId"`
	ClusterType               string           `position:"Query" name:"ClusterType"`
	Password                  string           `position:"Query" name:"Password"`
	SnatEntry                 requests.Integer `position:"Query" name:"SnatEntry"`
	NetPlug                   string           `position:"Query" name:"NetPlug"`
	VpcId                     string           `position:"Query" name:"VpcId"`
	RegionName                string           `position:"Query" name:"RegionName"`
	PrivateZone               requests.Boolean `position:"Query" name:"PrivateZone"`
	ServiceCIDR               string           `position:"Query" name:"ServiceCIDR"`
	PublicSlb                 requests.Integer `position:"Query" name:"PublicSlb"`
}

// CreateClusterResponse is the response struct for api CreateCluster
type CreateClusterResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateClusterRequest creates a request to invoke CreateCluster API
func CreateCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "CreateCluster", "retailcloud", "openAPI")
	return
}

// CreateCreateClusterResponse creates a response to parse from CreateCluster response
func CreateCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
