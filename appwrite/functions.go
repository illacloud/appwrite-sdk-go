package appwrite

import (
	"strings"
)

// Functions service
type Functions struct {
	client Client
}

func NewFunctions(clt Client) *Functions {
	return &Functions{
		client: clt,
	}
}

// List get a list of all the project's functions. You can use the query
// params to filter your results.
func (srv *Functions) List(Queries []interface{}, Search string) (*ClientResponse, error) {
	path := "/functions"
	
	params := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Create create a new function. You can pass a list of
// [permissions](/docs/permissions) to allow different project users or team
// with access to execute the function using the client API.
func (srv *Functions) Create(FunctionId string, Name string, Execute []interface{}, Runtime string, Events []interface{}, Schedule string, Timeout int, Enabled bool) (*ClientResponse, error) {
	path := "/functions"
	
	params := map[string]interface{}{
		"functionId": FunctionId,
		"name": Name,
		"execute": Execute,
		"runtime": Runtime,
		"events": Events,
		"schedule": Schedule,
		"timeout": Timeout,
		"enabled": Enabled,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// ListRuntimes get a list of all runtimes that are currently active on your
// instance.
func (srv *Functions) ListRuntimes() (*ClientResponse, error) {
	path := "/functions/runtimes"
	
	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Get get a function by its unique ID.
func (srv *Functions) Get(FunctionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Update update function by its unique ID.
func (srv *Functions) Update(FunctionId string, Name string, Execute []interface{}, Events []interface{}, Schedule string, Timeout int, Enabled bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	params := map[string]interface{}{
		"name": Name,
		"execute": Execute,
		"events": Events,
		"schedule": Schedule,
		"timeout": Timeout,
		"enabled": Enabled,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// Delete delete a function by its unique ID.
func (srv *Functions) Delete(FunctionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListDeployments get a list of all the project's code deployments. You can
// use the query params to filter your results.
func (srv *Functions) ListDeployments(FunctionId string, Queries []interface{}, Search string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")

	params := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateDeployment create a new function code deployment. Use this endpoint
// to upload a new version of your code function. To execute your newly
// uploaded code, you'll need to update the function's deployment to use your
// new deployment UID.
// 
// This endpoint accepts a tar.gz file compressed with your code. Make sure to
// include any dependencies your code has within the compressed file. You can
// learn more about code packaging in the [Appwrite Cloud Functions
// tutorial](/docs/functions).
// 
// Use the "command" param to set the entry point used to execute your code.
func (srv *Functions) CreateDeployment(FunctionId string, Entrypoint string, Code string, Activate bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")

	params := map[string]interface{}{
		"entrypoint": Entrypoint,
		"code": Code,
		"activate": Activate,
	}

	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetDeployment get a code deployment by its unique ID.
func (srv *Functions) GetDeployment(FunctionId string, DeploymentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateDeployment update the function code deployment ID using the unique
// function ID. Use this endpoint to switch the code deployment that should be
// executed by the execution endpoint.
func (srv *Functions) UpdateDeployment(FunctionId string, DeploymentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// DeleteDeployment delete a code deployment by its unique ID.
func (srv *Functions) DeleteDeployment(FunctionId string, DeploymentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// CreateBuild
func (srv *Functions) CreateBuild(FunctionId string, DeploymentId string, BuildId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId, "{buildId}", BuildId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}/builds/{buildId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// ListExecutions get a list of all the current user function execution logs.
// You can use the query params to filter your results.
func (srv *Functions) ListExecutions(FunctionId string, Queries []interface{}, Search string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")

	params := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateExecution trigger a function execution. The returned object will
// return you the current execution status. You can ping the `Get Execution`
// endpoint to get updates on the current execution status. Once this endpoint
// is called, your function execution process will start asynchronously.
func (srv *Functions) CreateExecution(FunctionId string, Data string, Async bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")

	params := map[string]interface{}{
		"data": Data,
		"async": Async,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetExecution get a function execution log by its unique ID.
func (srv *Functions) GetExecution(FunctionId string, ExecutionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/functions/{functionId}/executions/{executionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListVariables get a list of all variables of a specific function.
func (srv *Functions) ListVariables(FunctionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/variables")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateVariable create a new function variable. These variables can be
// accessed within function in the `env` object under the request variable.
func (srv *Functions) CreateVariable(FunctionId string, Key string, Value string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/variables")

	params := map[string]interface{}{
		"key": Key,
		"value": Value,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetVariable get a variable by its unique ID.
func (srv *Functions) GetVariable(FunctionId string, VariableId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateVariable update variable by its unique ID.
func (srv *Functions) UpdateVariable(FunctionId string, VariableId string, Key string, Value string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")

	params := map[string]interface{}{
		"key": Key,
		"value": Value,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteVariable delete a variable by its unique ID.
func (srv *Functions) DeleteVariable(FunctionId string, VariableId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}
