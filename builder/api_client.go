package builder

import (
	"time"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
)

// APIClient represents SAKURA CLOUD api client
type APIClient interface {
	ServerNew() *sacloud.Server
	ServerRead(serverID int64) (*sacloud.Server, error)
	ServerCreate(value *sacloud.Server) (*sacloud.Server, error)
	ServerSleepUntilUp(serverID int64, timeout time.Duration) error
	ServerInsertCDROM(serverID int64, cdromID int64) (bool, error)
	ServerBoot(serverID int64) (bool, error)

	SSHKeyNew() *sacloud.SSHKey
	SSHKeyCreate(value *sacloud.SSHKey) (*sacloud.SSHKey, error)
	SSHKeyDelete(sshKeyID int64) (*sacloud.SSHKey, error)
	SSHKeyGenerate(name string, passPhrase string, desc string) (*sacloud.SSHKeyGenerated, error)

	NoteNew() *sacloud.Note
	NoteCreate(value *sacloud.Note) (*sacloud.Note, error)
	NoteDelete(noteID int64) (*sacloud.Note, error)

	DiskNew() *sacloud.Disk
	DiskNewCondig() *sacloud.DiskEditValue
	DiskCreate(value *sacloud.Disk) (*sacloud.Disk, error)
	DiskCreateWithConfig(value *sacloud.Disk, config *sacloud.DiskEditValue, bootAtAvailable bool) (*sacloud.Disk, error)
	DiskSleepWhileCopying(id int64, timeout time.Duration) error
	DiskConnectToServer(diskID int64, serverID int64) (bool, error)

	InterfaceConnectToPacketFilter(interfaceID int64, packetFilterID int64) (bool, error)
	InterfaceSetDisplayIPAddress(interfaceID int64, ip string) (bool, error) // Interface

	ServerPlanGetBySpec(core int, memGB int, gen sacloud.PlanGenerations) (*sacloud.ProductServer, error)

	ArchiveFindByOSType(os ostype.ArchiveOSTypes) (*sacloud.Archive, error)

	GetTimeoutDuration() time.Duration
}

// NewAPIClient create new APICLient from *api.Client
func NewAPIClient(client *api.Client) APIClient {
	return &apiClient{client: client}
}

type apiClient struct {
	client *api.Client
}

func (a *apiClient) ServerNew() *sacloud.Server {
	return a.client.Server.New()
}

func (a *apiClient) ServerRead(serverID int64) (*sacloud.Server, error) {
	return a.client.Server.Read(serverID)
}

func (a *apiClient) ServerCreate(value *sacloud.Server) (*sacloud.Server, error) {
	return a.client.Server.Create(value)
}

func (a *apiClient) ServerSleepUntilUp(serverID int64, timeout time.Duration) error {
	return a.client.Server.SleepUntilUp(serverID, timeout)
}

func (a *apiClient) ServerInsertCDROM(serverID int64, cdromID int64) (bool, error) {
	return a.client.Server.InsertCDROM(serverID, cdromID)
}

func (a *apiClient) ServerBoot(serverID int64) (bool, error) {
	return a.client.Server.Boot(serverID)
}

func (a *apiClient) SSHKeyNew() *sacloud.SSHKey {
	return a.client.SSHKey.New()
}

func (a *apiClient) SSHKeyCreate(value *sacloud.SSHKey) (*sacloud.SSHKey, error) {
	return a.client.SSHKey.Create(value)
}

func (a *apiClient) SSHKeyDelete(sshKeyID int64) (*sacloud.SSHKey, error) {
	return a.client.SSHKey.Delete(sshKeyID)
}

func (a *apiClient) SSHKeyGenerate(name string, passPhrase string, desc string) (*sacloud.SSHKeyGenerated, error) {
	return a.client.SSHKey.Generate(name, passPhrase, desc)
}

func (a *apiClient) NoteNew() *sacloud.Note {
	return a.client.Note.New()
}

func (a *apiClient) NoteCreate(value *sacloud.Note) (*sacloud.Note, error) {
	return a.client.Note.Create(value)
}

func (a *apiClient) NoteDelete(noteID int64) (*sacloud.Note, error) {
	return a.client.Note.Delete(noteID)
}

func (a *apiClient) DiskNew() *sacloud.Disk {
	return a.client.Disk.New()
}

func (a *apiClient) DiskNewCondig() *sacloud.DiskEditValue {
	return a.client.Disk.NewCondig()
}

func (a *apiClient) DiskCreate(value *sacloud.Disk) (*sacloud.Disk, error) {
	return a.client.Disk.Create(value)
}

func (a *apiClient) DiskCreateWithConfig(
	value *sacloud.Disk,
	config *sacloud.DiskEditValue,
	bootAtAvailable bool) (*sacloud.Disk, error) {
	return a.client.Disk.CreateWithConfig(value, config, bootAtAvailable)
}

func (a *apiClient) DiskSleepWhileCopying(id int64, timeout time.Duration) error {
	return a.client.Disk.SleepWhileCopying(id, timeout)
}

func (a *apiClient) DiskConnectToServer(diskID int64, serverID int64) (bool, error) {
	return a.client.Disk.ConnectToServer(diskID, serverID)
}

func (a *apiClient) InterfaceConnectToPacketFilter(interfaceID int64, packetFilterID int64) (bool, error) {
	return a.client.Interface.ConnectToPacketFilter(interfaceID, packetFilterID)
}

func (a *apiClient) InterfaceSetDisplayIPAddress(interfaceID int64, ip string) (bool, error) {
	return a.client.Interface.SetDisplayIPAddress(interfaceID, ip)
}

func (a *apiClient) ServerPlanGetBySpec(core int, memGB int, gen sacloud.PlanGenerations) (*sacloud.ProductServer, error) {
	return a.client.Product.Server.GetBySpec(core, memGB, gen)
}

func (a *apiClient) ArchiveFindByOSType(os ostype.ArchiveOSTypes) (*sacloud.Archive, error) {
	return a.client.Archive.FindByOSType(os)
}

func (a *apiClient) GetTimeoutDuration() time.Duration {
	return a.client.DefaultTimeoutDuration
}
