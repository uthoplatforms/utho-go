package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestObjectStorageService_CreateBucket_happyPath(t *testing.T) {
	token := "token"
	payload := CreateBucketParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for CreateBucketParams: %v", err)
	}
	payload.Billing = "monthly"
	payload.Dcslug = "innoida"
	payload.Name = faker.Word()

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/bucket/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"Bucket created successfully"}`)
	})

	got, err := client.ObjectStorage().CreateBucket(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Bucket created successfully"}`), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestObjectStorageService_CreateBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ObjectStorage().CreateBucket(CreateBucketParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestObjectStorageService_ReadBucket_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	dcslug := faker.Word()
	bucketName := faker.Word()

	var want Bucket
	err := faker.FakeData(&want)
	if err != nil {
		t.Fatalf("Failed to fake data for Bucket: %v", err)
	}
	want.Dcslug = dcslug
	want.Name = bucketName

	serverResponseMap := map[string][]Bucket{"buckets": {want}}
	serverResponse, err := json.Marshal(serverResponseMap)
	if err != nil {
		t.Fatalf("Failed to marshal faked bucket: %v", err)
	}

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket", dcslug), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ObjectStorage().ReadBucket(dcslug, bucketName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_ReadBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ObjectStorage().ReadBucket(faker.Word(), faker.Word())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestObjectStorageService_ListBucket_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	dcslug := faker.Word()
	var dummyBuckets []Bucket
	for i := 0; i < 3; i++ {
		var bucket Bucket
		_ = faker.FakeData(&bucket)
		dummyBuckets = append(dummyBuckets, bucket)
	}

	serverResponse, _ := json.Marshal(Buckets{
		Buckets: dummyBuckets,
		Status:  "success",
	})

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket", dcslug), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ObjectStorage().ListBuckets(dcslug)
	assert.Equal(t, dummyBuckets, got)
}

func TestObjectStorageService_ListBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	objectstorage, err := client.ObjectStorage().ListBuckets(faker.Word())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if objectstorage != nil {
		t.Errorf("Was not expecting any objectstorage to be returned, instead got %v", objectstorage)
	}
}

func TestObjectStorageService_DeleteBucket_happyPath(t *testing.T) {
	token := "token"
	bucketName := faker.Word()
	dcslug := faker.Word()

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/delete", dcslug, bucketName), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"success"}`)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.ObjectStorage().DeleteBucket(dcslug, bucketName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_DeleteBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ObjectStorage().DeleteBucket(faker.Word(), faker.Word())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

func TestObjectStorageService_CreateAccessKey_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAccessKeyParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for CreateAccessKeyParams: %v", err)
	}
	payload.Dcslug = faker.Word()
	payload.AccesskeyName = faker.Word() + "-key"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/accesskey/create", payload.Dcslug), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"Access Key created successfully","access_key":"faketoken","secret_key":"fakesecret"}`)
	})

	got, err := client.ObjectStorage().CreateAccessKey(payload)

	var want CreateAccessKeyResponse
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Access Key created successfully","access_key":"faketoken","secret_key":"fakesecret"}`), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestObjectStorageService_CreateAccessKey_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ObjectStorage().CreateAccessKey(CreateAccessKeyParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestObjectStorageService_ReadAccessKey_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ObjectStorage().ReadAccessKey(faker.Word(), faker.Word())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestObjectStorageService_ListAccessKey_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	dcslug := faker.Word()
	var dummyAccessKeys []AccessKey
	for i := 0; i < 3; i++ {
		var accessKey AccessKey
		_ = faker.FakeData(&accessKey)
		dummyAccessKeys = append(dummyAccessKeys, accessKey)
	}

	serverResponse, _ := json.Marshal(AccessKeys{
		AccessKeys: dummyAccessKeys,
		Status:     "success",
	})

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/accesskeys", dcslug), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ObjectStorage().ListAccessKeys(dcslug)
	assert.Equal(t, dummyAccessKeys, got)
}

func TestObjectStorageService_ListAccessKey_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	objectstorage, err := client.ObjectStorage().ListAccessKeys(faker.Word())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if objectstorage != nil {
		t.Errorf("Was not expecting any objectstorage to be returned, instead got %v", objectstorage)
	}
}

func TestObjectStorageService_UpdateBucketAccessControl_happyPath(t *testing.T) {
	token := "token"
	payload := UpdateBucketAccessControlParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for UpdateBucketAccessControlParams: %v", err)
	}
	payload.Dcslug = faker.Word()
	payload.BucketName = faker.Word()
	payload.Policy = "public"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/policy/%s", payload.Dcslug, payload.BucketName, payload.Policy), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"Bucket access control updated successfully"}`)
	})

	got, err := client.ObjectStorage().UpdateBucketAccessControl(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Bucket access control updated successfully"}`), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestObjectStorageService_UpdateBucketAccessControl_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ObjectStorage().UpdateBucketAccessControl(UpdateBucketAccessControlParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestObjectStorageService_CreateDirectroy_happyPath(t *testing.T) {
	token := "token"
	payload := CreateDirectroyParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for CreateDirectroyParams: %v", err)
	}
	payload.Dcslug = faker.Word()
	payload.BucketName = faker.Word()
	payload.Path = "/" + faker.Word()

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/createdirectory", payload.Dcslug, payload.BucketName), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"Directory created successfully"}`)
	})

	got, err := client.ObjectStorage().CreateDirectroy(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Directory created successfully"}`), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestObjectStorageService_CreateDirectroy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ObjectStorage().CreateDirectroy(CreateDirectroyParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestObjectStorageService_ListBucketObjectsAndDirectories_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	dcslug := faker.Word()
	bucketName := faker.Word()
	path := fmt.Sprintf("/%s/%s", faker.Word(), faker.Word())

	var dummyObjects []Object
	for i := 0; i < 3; i++ {
		var object Object
		_ = faker.FakeData(&object)
		dummyObjects = append(dummyObjects, object)
	}

	serverResponse, _ := json.Marshal(Objects{
		Objects: dummyObjects,
		Status:  "success",
	})

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/objects", dcslug, bucketName), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		if req.URL.Query().Get("path") != path {
			t.Errorf("Expected path query parameter '%s', got '%s'", path, req.URL.Query().Get("path"))
		}
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ObjectStorage().ListBucketObjectsAndDirectories(dcslug, bucketName, path)
	assert.Equal(t, dummyObjects, got)
}

func TestObjectStorageService_ListBucketObjectsAndDirectories_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	objectstorage, err := client.ObjectStorage().ListBucketObjectsAndDirectories(faker.Word(), faker.Word(), faker.URL())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if objectstorage != nil {
		t.Errorf("Was not expecting any objectstorage to be returned, instead got %v", objectstorage)
	}
}

func TestObjectStorageService_DeleteDirectroy_happyPath(t *testing.T) {
	token := "token"
	bucketName := faker.Word()
	dcslug := faker.Word()
	directoryName := faker.Word()

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/delete/object", dcslug, bucketName), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"success"}`)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.ObjectStorage().DeleteDirectroy(dcslug, bucketName, directoryName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_DeleteDirectroy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ObjectStorage().DeleteDirectroy(faker.Word(), faker.Word(), faker.Word())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

func TestObjectStorageService_GetSharableUrlOfObject_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	dcslug := faker.Word()
	bucketName := faker.Word()
	path := fmt.Sprintf("/%s/%s.txt", faker.Word(), faker.Word())

	var want GetSharableUrlOfObject
	want.URL = faker.URL()

	serverResponse, err := json.Marshal(map[string]string{"url": want.URL})
	if err != nil {
		t.Fatalf("Failed to marshal faked sharable URL: %v", err)
	}

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/download", dcslug, bucketName), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		if req.URL.Query().Get("path") != path {
			t.Errorf("Expected path query parameter '%s', got '%s'", path, req.URL.Query().Get("path"))
		}
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ObjectStorage().GetSharableUrlOfObject(dcslug, bucketName, path)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_GetSharableUrlOfObject_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ObjectStorage().GetSharableUrlOfObject(faker.Word(), faker.Word(), faker.URL())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestObjectStorageService_ListSubscriptionPlanPricing_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var dummyPricing []Pricing
	for i := 0; i < 3; i++ {
		var pricing Pricing
		_ = faker.FakeData(&pricing)
		dummyPricing = append(dummyPricing, pricing)
	}

	serverResponse, _ := json.Marshal(PlanList{
		Pricing: dummyPricing,
		Status:  "success",
	})

	mux.HandleFunc("/pricing/objectstorage", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, string(serverResponse))
	})

	got, _ := client.ObjectStorage().ListSubscriptionPlanPricing()
	assert.Equal(t, dummyPricing, got)
}

func TestObjectStorageService_UpdateBucketAccessKeyPermission_happyPath(t *testing.T) {
	token := "token"
	payload := UpdateBucketAccessKeyPermissionParams{}
	err := faker.FakeData(&payload)
	if err != nil {
		t.Fatalf("Failed to fake data for UpdateBucketAccessKeyPermissionParams: %v", err)
	}
	payload.Dcslug = faker.Word()
	payload.BucketName = faker.Word()
	payload.PermissionName = "read"
	payload.AccessKeyId = faker.UUIDDigit()

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/objectstorage/%s/bucket/%s/permission/%s/accesskey/%s", payload.Dcslug, payload.BucketName, payload.PermissionName, payload.AccessKeyId), func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, `{"status":"success","message":"Bucket access key permission updated successfully"}`)
	})

	got, err := client.ObjectStorage().UpdateBucketAccessKeyPermission(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(`{"status":"success","message":"Bucket access key permission updated successfully"}`), &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestObjectStorageService_UpdateBucketAccessKeyPermission_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.ObjectStorage().UpdateBucketAccessKeyPermission(UpdateBucketAccessKeyPermissionParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
