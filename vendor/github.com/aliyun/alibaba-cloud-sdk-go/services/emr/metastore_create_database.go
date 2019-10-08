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

// MetastoreCreateDatabase invokes the emr.MetastoreCreateDatabase API synchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatedatabase.html
func (client *Client) MetastoreCreateDatabase(request *MetastoreCreateDatabaseRequest) (response *MetastoreCreateDatabaseResponse, err error) {
	response = CreateMetastoreCreateDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreCreateDatabaseWithChan invokes the emr.MetastoreCreateDatabase API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatedatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreCreateDatabaseWithChan(request *MetastoreCreateDatabaseRequest) (<-chan *MetastoreCreateDatabaseResponse, <-chan error) {
	responseChan := make(chan *MetastoreCreateDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreCreateDatabase(request)
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

// MetastoreCreateDatabaseWithCallback invokes the emr.MetastoreCreateDatabase API asynchronously
// api document: https://help.aliyun.com/api/emr/metastorecreatedatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreCreateDatabaseWithCallback(request *MetastoreCreateDatabaseRequest, callback func(response *MetastoreCreateDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreCreateDatabaseResponse
		var err error
		defer close(result)
		response, err = client.MetastoreCreateDatabase(request)
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

// MetastoreCreateDatabaseRequest is the request struct for api MetastoreCreateDatabase
type MetastoreCreateDatabaseRequest struct {
	*requests.RpcRequest
	DbSource        string           `position:"Query" name:"DbSource"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DbName          string           `position:"Query" name:"DbName"`
	DataSourceId    string           `position:"Query" name:"DataSourceId"`
	Description     string           `position:"Query" name:"Description"`
	Comment         string           `position:"Query" name:"Comment"`
	LocationUri     string           `position:"Query" name:"LocationUri"`
	ClusterBizId    string           `position:"Query" name:"ClusterBizId"`
}

// MetastoreCreateDatabaseResponse is the response struct for api MetastoreCreateDatabase
type MetastoreCreateDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMetastoreCreateDatabaseRequest creates a request to invoke MetastoreCreateDatabase API
func CreateMetastoreCreateDatabaseRequest() (request *MetastoreCreateDatabaseRequest) {
	request = &MetastoreCreateDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreCreateDatabase", "emr", "openAPI")
	return
}

// CreateMetastoreCreateDatabaseResponse creates a response to parse from MetastoreCreateDatabase response
func CreateMetastoreCreateDatabaseResponse() (response *MetastoreCreateDatabaseResponse) {
	response = &MetastoreCreateDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}