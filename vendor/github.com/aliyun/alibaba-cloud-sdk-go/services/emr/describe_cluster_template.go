package emr

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

// DescribeClusterTemplate invokes the emr.DescribeClusterTemplate API synchronously
func (client *Client) DescribeClusterTemplate(request *DescribeClusterTemplateRequest) (response *DescribeClusterTemplateResponse, err error) {
	response = CreateDescribeClusterTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterTemplateWithChan invokes the emr.DescribeClusterTemplate API asynchronously
func (client *Client) DescribeClusterTemplateWithChan(request *DescribeClusterTemplateRequest) (<-chan *DescribeClusterTemplateResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterTemplate(request)
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

// DescribeClusterTemplateWithCallback invokes the emr.DescribeClusterTemplate API asynchronously
func (client *Client) DescribeClusterTemplateWithCallback(request *DescribeClusterTemplateRequest, callback func(response *DescribeClusterTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterTemplateResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterTemplate(request)
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

// DescribeClusterTemplateRequest is the request struct for api DescribeClusterTemplate
type DescribeClusterTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	BizId           string           `position:"Query" name:"BizId"`
}

// DescribeClusterTemplateResponse is the response struct for api DescribeClusterTemplate
type DescribeClusterTemplateResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	TemplateInfo TemplateInfo `json:"TemplateInfo" xml:"TemplateInfo"`
}

// CreateDescribeClusterTemplateRequest creates a request to invoke DescribeClusterTemplate API
func CreateDescribeClusterTemplateRequest() (request *DescribeClusterTemplateRequest) {
	request = &DescribeClusterTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterTemplate", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterTemplateResponse creates a response to parse from DescribeClusterTemplate response
func CreateDescribeClusterTemplateResponse() (response *DescribeClusterTemplateResponse) {
	response = &DescribeClusterTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
