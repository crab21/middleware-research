package alimt

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

// TranslateECommerce invokes the alimt.TranslateECommerce API synchronously
// api document: https://help.aliyun.com/api/alimt/translateecommerce.html
func (client *Client) TranslateECommerce(request *TranslateECommerceRequest) (response *TranslateECommerceResponse, err error) {
	response = CreateTranslateECommerceResponse()
	err = client.DoAction(request, response)
	return
}

// TranslateECommerceWithChan invokes the alimt.TranslateECommerce API asynchronously
// api document: https://help.aliyun.com/api/alimt/translateecommerce.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TranslateECommerceWithChan(request *TranslateECommerceRequest) (<-chan *TranslateECommerceResponse, <-chan error) {
	responseChan := make(chan *TranslateECommerceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TranslateECommerce(request)
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

// TranslateECommerceWithCallback invokes the alimt.TranslateECommerce API asynchronously
// api document: https://help.aliyun.com/api/alimt/translateecommerce.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TranslateECommerceWithCallback(request *TranslateECommerceRequest, callback func(response *TranslateECommerceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TranslateECommerceResponse
		var err error
		defer close(result)
		response, err = client.TranslateECommerce(request)
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

// TranslateECommerceRequest is the request struct for api TranslateECommerce
type TranslateECommerceRequest struct {
	*requests.RpcRequest
	SourceLanguage string `position:"Body" name:"SourceLanguage"`
	SourceText     string `position:"Body" name:"SourceText"`
	FormatType     string `position:"Body" name:"FormatType"`
	Scene          string `position:"Body" name:"Scene"`
	TargetLanguage string `position:"Body" name:"TargetLanguage"`
}

// TranslateECommerceResponse is the response struct for api TranslateECommerce
type TranslateECommerceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTranslateECommerceRequest creates a request to invoke TranslateECommerce API
func CreateTranslateECommerceRequest() (request *TranslateECommerceRequest) {
	request = &TranslateECommerceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "TranslateECommerce", "alimtct", "openAPI")
	return
}

// CreateTranslateECommerceResponse creates a response to parse from TranslateECommerce response
func CreateTranslateECommerceResponse() (response *TranslateECommerceResponse) {
	response = &TranslateECommerceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}