package appwrite

import (
	"strings"
)

// Databases service
type Databases struct {
	client Client
}

func NewDatabases(clt Client) *Databases {
	return &Databases{
		client: clt,
	}
}

// List get a list of all databases from the current Appwrite project. You can
// use the search parameter to filter your results.
func (srv *Databases) List(Queries []interface{}, Search string) (*ClientResponse, error) {
	path := "/databases"
	
	params := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Create create a new Database.
// 
func (srv *Databases) Create(DatabaseId string, Name string) (*ClientResponse, error) {
	path := "/databases"
	
	params := map[string]interface{}{
		"databaseId": DatabaseId,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Get get a database by its unique ID. This endpoint response returns a JSON
// object with the database metadata.
func (srv *Databases) Get(DatabaseId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Update update a database by its unique ID.
func (srv *Databases) Update(DatabaseId string, Name string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")

	params := map[string]interface{}{
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// Delete delete a database by its unique ID. Only API keys with with
// databases.write scope can delete a database.
func (srv *Databases) Delete(DatabaseId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListCollections get a list of all collections that belong to the provided
// databaseId. You can use the search parameter to filter your results.
func (srv *Databases) ListCollections(DatabaseId string, Queries []interface{}, Search string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")

	params := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateCollection create a new Collection. Before using this route, you
// should create a new database resource using either a [server
// integration](/docs/server/databases#databasesCreateCollection) API or
// directly from your database console.
func (srv *Databases) CreateCollection(DatabaseId string, CollectionId string, Name string, Permissions []interface{}, DocumentSecurity bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")

	params := map[string]interface{}{
		"collectionId": CollectionId,
		"name": Name,
		"permissions": Permissions,
		"documentSecurity": DocumentSecurity,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetCollection get a collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
func (srv *Databases) GetCollection(DatabaseId string, CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateCollection update a collection by its unique ID.
func (srv *Databases) UpdateCollection(DatabaseId string, CollectionId string, Name string, Permissions []interface{}, DocumentSecurity bool, Enabled bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")

	params := map[string]interface{}{
		"name": Name,
		"permissions": Permissions,
		"documentSecurity": DocumentSecurity,
		"enabled": Enabled,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
func (srv *Databases) DeleteCollection(DatabaseId string, CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListAttributes
func (srv *Databases) ListAttributes(DatabaseId string, CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateBooleanAttribute create a boolean attribute.
// 
func (srv *Databases) CreateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default bool, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/boolean")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateDatetimeAttribute
func (srv *Databases) CreateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/datetime")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateEmailAttribute create an email attribute.
// 
func (srv *Databases) CreateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/email")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateEnumAttribute
func (srv *Databases) CreateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []interface{}, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/enum")

	params := map[string]interface{}{
		"key": Key,
		"elements": Elements,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateFloatAttribute create a float attribute. Optionally, minimum and
// maximum values can be provided.
// 
func (srv *Databases) CreateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Min float64, Max float64, Default float64, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/float")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"min": Min,
		"max": Max,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateIntegerAttribute create an integer attribute. Optionally, minimum and
// maximum values can be provided.
// 
func (srv *Databases) CreateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Min int, Max int, Default int, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/integer")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"min": Min,
		"max": Max,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateIpAttribute create IP address attribute.
// 
func (srv *Databases) CreateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/ip")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateStringAttribute create a string attribute.
// 
func (srv *Databases) CreateStringAttribute(DatabaseId string, CollectionId string, Key string, Size int, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/string")

	params := map[string]interface{}{
		"key": Key,
		"size": Size,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateUrlAttribute create a URL attribute.
// 
func (srv *Databases) CreateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/url")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetAttribute
func (srv *Databases) GetAttribute(DatabaseId string, CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteAttribute
func (srv *Databases) DeleteAttribute(DatabaseId string, CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListDocuments get a list of all the user's documents in a given collection.
// You can use the query params to filter your results.
func (srv *Databases) ListDocuments(DatabaseId string, CollectionId string, Queries []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"queries": Queries,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](/docs/server/databases#databasesCreateCollection) API or
// directly from your database console.
func (srv *Databases) CreateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, Permissions []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"documentId": DocumentId,
		"data": Data,
		"permissions": Permissions,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
func (srv *Databases) GetDocument(DatabaseId string, CollectionId string, DocumentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
func (srv *Databases) UpdateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, Permissions []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
		"data": Data,
		"permissions": Permissions,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// DeleteDocument delete a document by its unique ID.
func (srv *Databases) DeleteDocument(DatabaseId string, CollectionId string, DocumentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListIndexes
func (srv *Databases) ListIndexes(DatabaseId string, CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateIndex
func (srv *Databases) CreateIndex(DatabaseId string, CollectionId string, Key string, Type string, Attributes []interface{}, Orders []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")

	params := map[string]interface{}{
		"key": Key,
		"type": Type,
		"attributes": Attributes,
		"orders": Orders,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetIndex
func (srv *Databases) GetIndex(DatabaseId string, CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteIndex
func (srv *Databases) DeleteIndex(DatabaseId string, CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}
