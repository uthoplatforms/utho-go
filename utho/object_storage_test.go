package utho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectStorageService_CreateBucket_happyPath(t *testing.T) {
	token := "token"
	payload := CreateBucketParams{
		Billing: "monthly",
		Dcslug:  "innoida",
		Name:    "mybucketg5cdt60unocvqavv",
		Price:   "125",
		Size:    "10120",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/bucket/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.ObjectStorage().CreateBucket(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

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

	dcslug := "innoida"
	bucketName := "examplename"

	expectedResponse := dummyReadBucketRes
	serverResponse := dummyReadBucketServerRes

	mux.HandleFunc("/objectstorage/"+dcslug+"/bucket", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want Bucket
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().ReadBucket(dcslug, bucketName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_ReadBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ObjectStorage().ReadBucket("innoida", "name")
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

	dcslug := "innoida"
	expectedResponse := dummyReadBucketRes
	serverResponse := "[" + dummyReadBucketServerRes + "]"

	mux.HandleFunc("/objectstorage/"+dcslug+"/bucket", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Bucket
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().ListBuckets(dcslug)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d objectstorage to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestObjectStorageService_ListBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	objectstorage, err := client.ObjectStorage().ListBuckets("dcslug")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if objectstorage != nil {
		t.Errorf("Was not expecting any objectstorage to be returned, instead got %v", objectstorage)
	}
}

func TestObjectStorageService_DeleteBucket_happyPath(t *testing.T) {
	token := "token"
	bucketName := "examplename"
	dcslug := "innoida"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/"+dcslug+"/bucket/"+bucketName+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.ObjectStorage().DeleteBucket(dcslug, bucketName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_DeleteBucket_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ObjectStorage().DeleteBucket("innoida", "name")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

// Access key
func TestObjectStorageService_CreateAccessKey_happyPath(t *testing.T) {
	token := "token"
	payload := CreateAccessKeyParams{
		Dcslug:        "innoida",
		AccesskeyName: "access-key1",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/"+payload.Dcslug+"/accesskey/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateAccessKeyResponseJson)
	})

	got, err := client.ObjectStorage().CreateAccessKey(payload)

	var want CreateAccessKeyResponse
	_ = json.Unmarshal([]byte(dummyCreateAccessKeyResponseJson), &want)

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

func TestObjectStorageService_ReadAccessKey_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	dcslug := "innoida"
	accesskeyName := "access-key1"

	expectedResponse := dummyReadAccessKeyRes
	serverResponse := dummyReadAccessKeyServerRes

	mux.HandleFunc("/objectstorage/"+dcslug+"/accesskeys", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want AccessKey
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().ReadAccessKey(dcslug, accesskeyName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_ReadAccessKey_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ObjectStorage().ReadAccessKey("innoida", "name")
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

	dcslug := "innoida"
	expectedResponse := dummyReadAccessKeyRes
	serverResponse := "[" + dummyReadAccessKeyServerRes + "]"

	mux.HandleFunc("/objectstorage/"+dcslug+"/accesskeys", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []AccessKey
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().ListAccessKeys(dcslug)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d objectstorage to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestObjectStorageService_ListAccessKey_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	objectstorage, err := client.ObjectStorage().ListAccessKeys("dcslug")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if objectstorage != nil {
		t.Errorf("Was not expecting any objectstorage to be returned, instead got %v", objectstorage)
	}
}

// func TestObjectStorageService_DeleteAccessKey_happyPath(t *testing.T) {
// 	token := "token"
// 	accesskeyName := "examplename"
// 	dcslug := "innoida"

// 	client, mux, _, teardown := setup(token)
// 	defer teardown()

// 	mux.HandleFunc("/objectstorage/"+dcslug+"/accesskey/"+accesskeyName+"/delete", func(w http.ResponseWriter, req *http.Request) {
// 		testHttpMethod(t, req, "DELETE")
// 		testHeader(t, req, "Authorization", "Bearer "+token)
// 		fmt.Fprint(w, dummyDeleteResponseJson)
// 	})

// 	want := DeleteResponse{Status: "success", Message: "success"}

// 	got, _ := client.ObjectStorage().DeleteAccessKey(dcslug, accesskeyName)
// 	if !reflect.DeepEqual(*got, want) {
// 		t.Errorf("Response = %v, want %v", *got, want)
// 	}
// }

// func TestObjectStorageService_DeleteAccessKey_invalidServer(t *testing.T) {
// 	client, _ := NewClient("token")

// 	delResponse, err := client.ObjectStorage().DeleteAccessKey("innoida", "name")
// 	if err == nil {
// 		t.Errorf("Expected error to be returned")
// 	}
// 	if delResponse != nil {
// 		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
// 	}
// }

func TestObjectStorageService_UpdateBucketAccessControl_happyPath(t *testing.T) {
	token := "token"
	payload := UpdateBucketAccessControlParams{
		Dcslug:     "innoida",
		Policy:     "public",
		BucketName: "examplename",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/"+payload.Dcslug+"/bucket/"+payload.BucketName+"/policy/"+payload.Policy, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.ObjectStorage().UpdateBucketAccessControl(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

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
	payload := CreateDirectroyParams{
		Dcslug:     "innoida",
		Path:       "/example",
		BucketName: "examplename",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/"+payload.Dcslug+"/bucket/"+payload.BucketName+"/createdirectory", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.ObjectStorage().CreateDirectroy(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

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

	dcslug := "innoida"
	bucketName := "examplename"
	path := "/pathname"
	expectedResponse := dummyListBucketObjectsAndDirectoriesRes
	serverResponse := "[" + dummyListBucketObjectsAndDirectoriesServerRes + "]"

	mux.HandleFunc("/objectstorage/"+dcslug+"/bucket/"+bucketName+"/objects?path="+path, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Object
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().ListBucketObjectsAndDirectories(dcslug, bucketName, path)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d objectstorage to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestObjectStorageService_ListBucketObjectsAndDirectories_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	objectstorage, err := client.ObjectStorage().ListBucketObjectsAndDirectories("dcslug", "name", "/path")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if objectstorage != nil {
		t.Errorf("Was not expecting any objectstorage to be returned, instead got %v", objectstorage)
	}
}

func TestObjectStorageService_DeleteDirectroy_happyPath(t *testing.T) {
	token := "token"
	bucketName := "examplename"
	dcslug := "innoida"
	directoryName := "pathname"

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/"+dcslug+"/bucket/"+bucketName+"/delete/object", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyDeleteResponseJson)
	})

	want := DeleteResponse{Status: "success", Message: "success"}

	got, _ := client.ObjectStorage().DeleteDirectroy(dcslug, bucketName, directoryName)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_DeleteDirectroy_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.ObjectStorage().DeleteDirectroy("innoida", "name", "directoryName")
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

	dcslug := "innoida"
	bucketName := "examplename"
	path := "/pathname"

	expectedResponse := dummyGetSharableUrlOfObjectRes
	serverResponse := dummyGetSharableUrlOfObjectServerRes

	mux.HandleFunc("/objectstorage/"+dcslug+"/bucket/"+bucketName+"/download", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want GetSharableUrlOfObject
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().GetSharableUrlOfObject(dcslug, bucketName, path)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestObjectStorageService_GetSharableUrlOfObject_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.ObjectStorage().GetSharableUrlOfObject("innoida", "name", "path")
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

	expectedResponse := dummyListSubscriptionPlanPricingRes
	serverResponse := dummyListSubscriptionPlanPricingServerRes

	mux.HandleFunc("/pricing/objectstorage", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		fmt.Fprint(w, serverResponse)
	})

	var want []Pricing
	_ = json.Unmarshal([]byte(expectedResponse), &want)

	got, _ := client.ObjectStorage().ListSubscriptionPlanPricing()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d objectstorage to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestObjectStorageService_UpdateBucketAccessKeyPermission_happyPath(t *testing.T) {
	token := "token"
	payload := UpdateBucketAccessKeyPermissionParams{
		Dcslug:         "innoida",
		BucketName:     "examplename",
		PermissionName: "read",
		AccessKeyId:    "xzM9235pSqYW7RdkhXJa4sftViE1vCIT6AOZ",
	}

	client, mux, _, teardown := setup(token)
	defer teardown()

	mux.HandleFunc("/objectstorage/"+payload.Dcslug+"/bucket/"+payload.BucketName+"/permission/"+payload.PermissionName+"/accesskey/"+payload.AccessKeyId, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		fmt.Fprint(w, dummyCreateResponseJson)
	})

	got, err := client.ObjectStorage().UpdateBucketAccessKeyPermission(payload)

	var want CreateResponse
	_ = json.Unmarshal([]byte(dummyCreateResponseJson), &want)

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
