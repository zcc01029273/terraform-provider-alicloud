package ddoscoo

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

// DescribeWebRules invokes the ddoscoo.DescribeWebRules API synchronously
func (client *Client) DescribeWebRules(request *DescribeWebRulesRequest) (response *DescribeWebRulesResponse, err error) {
	response = CreateDescribeWebRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebRulesWithChan invokes the ddoscoo.DescribeWebRules API asynchronously
func (client *Client) DescribeWebRulesWithChan(request *DescribeWebRulesRequest) (<-chan *DescribeWebRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeWebRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebRules(request)
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

// DescribeWebRulesWithCallback invokes the ddoscoo.DescribeWebRules API asynchronously
func (client *Client) DescribeWebRulesWithCallback(request *DescribeWebRulesRequest, callback func(response *DescribeWebRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebRules(request)
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

// DescribeWebRulesRequest is the request struct for api DescribeWebRules
type DescribeWebRulesRequest struct {
	*requests.RpcRequest
	PageNumber         requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId    string           `position:"Query" name:"ResourceGroupId"`
	SourceIp           string           `position:"Query" name:"SourceIp"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	InstanceIds        *[]string        `position:"Query" name:"InstanceIds"  type:"Repeated"`
	QueryDomainPattern string           `position:"Query" name:"QueryDomainPattern"`
	Domain             string           `position:"Query" name:"Domain"`
}

// DescribeWebRulesResponse is the response struct for api DescribeWebRules
type DescribeWebRulesResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int64     `json:"TotalCount" xml:"TotalCount"`
	WebRules   []WebRule `json:"WebRules" xml:"WebRules"`
}

// CreateDescribeWebRulesRequest creates a request to invoke DescribeWebRules API
func CreateDescribeWebRulesRequest() (request *DescribeWebRulesRequest) {
	request = &DescribeWebRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeWebRulesResponse creates a response to parse from DescribeWebRules response
func CreateDescribeWebRulesResponse() (response *DescribeWebRulesResponse) {
	response = &DescribeWebRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
