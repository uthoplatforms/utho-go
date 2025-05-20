package utho

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCloudInstanceService_Create_happyPath(t *testing.T) {
	token := "token"

	var payload CreateCloudInstanceParams
	_ = faker.FakeData(&payload)

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp CreateCloudInstanceResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/deploy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	var want CreateCloudInstanceResponse
	_ = json.Unmarshal(respBytes, &want)

	got, err := client.CloudInstances().Create(payload)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_Create_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().Create(CreateCloudInstanceParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_Read_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeInstance CloudInstance
	_ = faker.FakeData(&fakeInstance)
	serverResp := struct {
		Cloud []CloudInstance `json:"cloud"`
	}{
		Cloud: []CloudInstance{fakeInstance},
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeInstance)

	mux.HandleFunc("/cloud/"+fakeInstance.ID, func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want CloudInstance
	_ = json.Unmarshal(expectedResponse, &want)

	got, err := client.CloudInstances().Read(fakeInstance.ID)
	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_Read_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	apikey, err := client.CloudInstances().Read("someId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if apikey != nil {
		t.Errorf("Was not expecting any apikey to be returned, instead got %v", apikey)
	}
}

func TestCloudInstanceService_List_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeInstances []CloudInstance
	for i := 0; i < 2; i++ {
		var inst CloudInstance
		_ = faker.FakeData(&inst)
		fakeInstances = append(fakeInstances, inst)
	}
	serverResp := struct {
		Cloud []CloudInstance `json:"cloud"`
	}{
		Cloud: fakeInstances,
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeInstances)

	mux.HandleFunc("/cloud", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []CloudInstance
	_ = json.Unmarshal(expectedResponse, &want)

	got, err := client.CloudInstances().List()
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCloudInstanceService_List_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	cloudinstance, err := client.CloudInstances().List()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if cloudinstance != nil {
		t.Errorf("Was not expecting any cloudinstance to be returned, instead got %v", cloudinstance)
	}
}

func TestCloudInstanceService_Delete_happyPath(t *testing.T) {
	token := "token"
	cloudInstanceId := "someCloudInstanceId"
	deleteCloudInstanceParams := DeleteCloudInstanceParams{Confirm: "I am aware this action will delete data and server permanently"}

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+cloudInstanceId+"/destroy", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	var want DeleteResponse
	_ = json.Unmarshal(respBytes, &want)

	got, _ := client.CloudInstances().Delete(cloudInstanceId, deleteCloudInstanceParams)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestCloudInstanceService_Delete_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.CloudInstances().Delete("someCloudInstanceId", DeleteCloudInstanceParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any reponse to be returned, instead got %v", delResponse)
	}
}

func TestCloudInstanceService_ListOsImages_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	var fakeImages []OsImage
	for i := 0; i < 2; i++ {
		var img OsImage
		_ = faker.FakeData(&img)
		fakeImages = append(fakeImages, img)
	}
	serverResp := struct {
		Images []OsImage `json:"images"`
	}{
		Images: fakeImages,
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakeImages)

	mux.HandleFunc("/cloud/images", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []OsImage
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.CloudInstances().ListOsImages()
	if len(got) != len(want) {
		t.Errorf("Was expecting %d cloudinstance to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestCloudInstanceService_ListOsImages_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	cloudinstance, err := client.CloudInstances().ListOsImages()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if cloudinstance != nil {
		t.Errorf("Was not expecting any cloudinstance to be returned, instead got %v", cloudinstance)
	}
}

func TestCloudInstanceService_ListResizePlans_happyPath(t *testing.T) {
	client, mux, _, teardown := setup("token")
	defer teardown()

	instanceId := "someId"
	var fakePlans []Plan
	for i := 0; i < 2; i++ {
		var plan Plan
		_ = faker.FakeData(&plan)
		fakePlans = append(fakePlans, plan)
	}
	serverResp := struct {
		Plans []Plan `json:"plans"`
	}{
		Plans: fakePlans,
	}
	serverResponse, _ := json.Marshal(serverResp)
	expectedResponse, _ := json.Marshal(fakePlans)

	mux.HandleFunc("/cloud/"+instanceId+"/resizeplans", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "GET")
		testHeader(t, req, "Authorization", "Bearer token")
		w.Write(serverResponse)
	})

	var want []Plan
	_ = json.Unmarshal(expectedResponse, &want)

	got, _ := client.CloudInstances().ListResizePlans(instanceId)
	if len(got) != len(want) {
		t.Errorf("Was expecting %d cloudinstance to be returned, instead got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Response = %v, want %v", got, want)
	}
}

func TestCloudInstanceService_ListResizePlans_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	cloudinstance, err := client.CloudInstances().ListResizePlans("id")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if cloudinstance != nil {
		t.Errorf("Was not expecting any cloudinstance to be returned, instead got %v", cloudinstance)
	}
}

func TestCloudInstanceService_CreateSnapshot_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var payload CreateSnapshotParams
	_ = faker.FakeData(&payload)

	var fakeResp CreateBasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/snapshot/create", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().CreateSnapshot(instanceId, payload)

	var want CreateBasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_CreateSnapshot_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().CreateSnapshot("instanceId", CreateSnapshotParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_DeleteSnapshot_happyPath(t *testing.T) {
	token := "token"
	cloudInstanceId := "someCloudInstanceId"
	snapshotId := "somesnapshotId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp DeleteResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Message = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+cloudInstanceId+"/snapshot/"+snapshotId+"/delete", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, "DELETE")
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	var want DeleteResponse
	_ = json.Unmarshal(respBytes, &want)

	got, _ := client.CloudInstances().DeleteSnapshot(cloudInstanceId, snapshotId)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Response = %v, want %v", *got, want)
	}
}

func TestCloudInstanceService_DeleteSnapshot_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	delResponse, err := client.CloudInstances().DeleteSnapshot("someCloudInstanceId", "somesnapshotId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if delResponse != nil {
		t.Errorf("Was not expecting any response to be returned, instead got %v", delResponse)
	}
}

func TestCloudInstanceService_EnableBackup_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/backups/enable", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().EnableBackup(instanceId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_EnableBackup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().EnableBackup("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_DisableBackup_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/backups/disable", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().DisableBackup(instanceId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_DisableBackup_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().DisableBackup("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_HardReboot_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/hardreboot", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().HardReboot(instanceId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_HardReboot_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().HardReboot("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_PowerCycle_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/powercycle", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().PowerCycle(instanceId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_PowerCycle_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().PowerCycle("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_PowerOff_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/poweroff", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().PowerOff(instanceId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_PowerOff_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().PowerOff("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_PowerOn_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/poweron", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().PowerOn(instanceId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_PowerOn_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().PowerOn("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_Rebuild_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/rebuild", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	payload := RebuildCloudInstanceParams{
		Image:   "almalinux-9.2-x86_64",
		Confirm: "I am aware this action will delete data permanently and build a fresh server",
	}
	got, err := client.CloudInstances().Rebuild(instanceId, payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_Rebuild_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().Rebuild("instanceId", RebuildCloudInstanceParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_ResetPassword_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp ResetPasswordResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	fakeResp.Password = "newpassword123"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/resetpassword", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().ResetPassword(instanceId)

	var want ResetPasswordResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_ResetPassword_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().ResetPassword("instanceId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_Resize_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/resize", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	payload := ResizeCloudInstanceParams{
		Type: "ramcpu",
		Plan: "11111",
	}
	got, err := client.CloudInstances().Resize(instanceId, payload)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_Resize_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().Resize("instanceId", ResizeCloudInstanceParams{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestCloudInstanceService_RestoreSnapshot_happyPath(t *testing.T) {
	token := "token"
	instanceId := "someId"
	snapshotId := "snapshotId"

	client, mux, _, teardown := setup(token)
	defer teardown()

	var fakeResp BasicResponse
	_ = faker.FakeData(&fakeResp)
	fakeResp.Status = "success"
	respBytes, _ := json.Marshal(fakeResp)

	mux.HandleFunc("/cloud/"+instanceId+"/snapshot/"+snapshotId+"/restore", func(w http.ResponseWriter, req *http.Request) {
		testHttpMethod(t, req, http.MethodPost)
		testHeader(t, req, "Authorization", "Bearer "+token)
		w.Write(respBytes)
	})

	got, err := client.CloudInstances().RestoreSnapshot(instanceId, snapshotId)

	var want BasicResponse
	_ = json.Unmarshal(respBytes, &want)

	assert.Nil(t, err)
	assert.Equal(t, want, *got)
}

func TestCloudInstanceService_RestoreSnapshot_invalidServer(t *testing.T) {
	client, _ := NewClient("token")

	_, err := client.CloudInstances().RestoreSnapshot("instanceId", "snapshotId")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
