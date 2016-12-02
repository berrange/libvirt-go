// +build integration

package libvirt

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func buildTestQEMUConnection() VirConnection {
	conn, err := NewVirConnection("qemu:///system")
	if err != nil {
		panic(err)
	}
	return conn
}

func buildTestQEMUDomain() (Domain, VirConnection) {
	conn := buildTestQEMUConnection()
	dom, err := conn.DomainDefineXML(`<domain type="qemu">
		<name>` + strings.Replace(time.Now().String(), " ", "_", -1) + `</name>
		<memory unit="KiB">128</memory>
		<os>
			<type>hvm</type>
		</os>
	</domain>`)
	if err != nil {
		panic(err)
	}
	return dom, conn
}

func TestMultipleCloseCallback(t *testing.T) {
	nbCall1 := 0
	nbCall2 := 0
	nbCall3 := 0
	conn := buildTestQEMUConnection()
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
		if nbCall1 != 0 || nbCall2 != 0 || nbCall3 != 1 {
			t.Errorf("Wrong number of calls to callback, got %v, expected %v",
				[]int{nbCall1, nbCall2, nbCall3},
				[]int{0, 0, 1})
		}
	}()

	callback := func(conn VirConnection, reason int, opaque func()) {
		if reason != VIR_CONNECT_CLOSE_REASON_KEEPALIVE {
			t.Errorf("Expected close reason to be %d, got %d",
				VIR_CONNECT_CLOSE_REASON_KEEPALIVE, reason)
		}
		opaque()
	}
	err := conn.RegisterCloseCallback(callback, func() {
		nbCall1++
	})
	if err != nil {
		t.Fatalf("Unable to register close callback: %+v", err)
	}
	err = conn.RegisterCloseCallback(callback, func() {
		nbCall2++
	})
	if err != nil {
		t.Fatalf("Unable to register close callback: %+v", err)
	}
	err = conn.RegisterCloseCallback(callback, func() {
		nbCall3++
	})
	if err != nil {
		t.Fatalf("Unable to register close callback: %+v", err)
	}

	// To trigger a disconnect, we use a keepalive
	if err := conn.SetKeepAlive(1, 1); err != nil {
		t.Fatalf("Unable to enable keeplive: %+v", err)
	}
	EventRunDefaultImpl()
	time.Sleep(2 * time.Second)
	EventRunDefaultImpl()
}

func TestUnregisterCloseCallback(t *testing.T) {
	nbCall := 0
	conn := buildTestQEMUConnection()
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
		if nbCall != 0 {
			t.Errorf("Expected no call to close callback, got %d", nbCall)
		}
	}()

	callback := func(conn VirConnection, reason int, opaque func()) {
		nbCall++
	}
	err := conn.RegisterCloseCallback(callback, nil)
	if err != nil {
		t.Fatalf("Unable to register close callback: %+v", err)
	}
	err = conn.UnregisterCloseCallback()
	if err != nil {
		t.Fatalf("Unable to unregister close callback: %+v", err)
	}
}

func TestSetKeepalive(t *testing.T) {
	EventRegisterDefaultImpl()        // We need the event loop for keepalive
	conn := buildTestQEMUConnection() // The test driver doesn't support keepalives
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	if err := conn.SetKeepAlive(1, 1); err != nil {
		t.Error(err)
		return
	}

	// It should block until we have a keepalive message
	done := make(chan struct{})
	timeout := time.After(5 * time.Second)
	go func() {
		EventRunDefaultImpl()
		close(done)
	}()
	select {
	case <-done: // OK!
	case <-timeout:
		t.Fatalf("timeout reached while waiting for keepalive")
	}
}

func TestConnectionWithAuth(t *testing.T) {
	conn, err := NewVirConnectionWithAuth("test+tcp://127.0.0.1/default", "user", "pass")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := conn.CloseConnection()
	if err != nil {
		t.Error(err)
		return
	}
	if res != 0 {
		t.Errorf("CloseConnection() == %d, expected 0", res)
	}
}

func TestConnectionWithWrongCredentials(t *testing.T) {
	conn, err := NewVirConnectionWithAuth("test+tcp://127.0.0.1/default", "user", "wrongpass")
	if err == nil {
		conn.CloseConnection()
		t.Error(err)
		return
	}
}

func TestQemuMonitorCommand(t *testing.T) {
	dom, conn := buildTestQEMUDomain()
	defer func() {
		dom.Destroy()
		dom.Undefine()
		dom.Free()
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	if err := dom.Create(); err != nil {
		t.Error(err)
		return
	}

	if _, err := dom.QemuMonitorCommand(VIR_DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT, "{\"execute\" : \"query-cpus\"}"); err != nil {
		t.Error(err)
		return
	}

	if _, err := dom.QemuMonitorCommand(VIR_DOMAIN_QEMU_MONITOR_COMMAND_HMP, "info cpus"); err != nil {
		t.Error(err)
		return
	}
}

func TestDomainCreateWithFlags(t *testing.T) {
	dom, conn := buildTestQEMUDomain()
	defer func() {
		dom.Destroy()
		dom.Undefine()
		dom.Free()
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	if err := dom.CreateWithFlags(VIR_DOMAIN_START_PAUSED); err != nil {
		state, err := dom.GetState()
		if err != nil {
			t.Error(err)
			return
		}

		if state[0] != VIR_DOMAIN_PAUSED {
			t.Fatalf("Domain should be paused")
		}
	}
}

func defineTestLxcDomain(conn VirConnection, title string) (Domain, error) {
	if title == "" {
		title = time.Now().String()
	}
	xml := `<domain type='lxc'>
	  <name>` + title + `</name>
	  <title>` + title + `</title>
	  <memory>102400</memory>
	  <os>
	    <type>exe</type>
	    <init>/bin/sh</init>
	  </os>
	  <devices>
	    <console type='pty'/>
	  </devices>
	</domain>`
	dom, err := conn.DomainDefineXML(xml)
	return dom, err
}

// Integration tests are run against LXC using Libvirt 1.2.x
// on Debian Wheezy (libvirt from wheezy-backports)
//
// To run,
// 		go test -tags integration

func TestIntegrationGetMetadata(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	title := time.Now().String()
	dom, err := defineTestLxcDomain(conn, title)
	if err != nil {
		t.Error(err)
		return
	}
	defer dom.Free()
	if err := dom.Create(); err != nil {
		t.Error(err)
		return
	}
	v, err := dom.GetMetadata(VIR_DOMAIN_METADATA_TITLE, "", 0)
	dom.Destroy()
	if err != nil {
		t.Error(err)
		return
	}
	if v != title {
		t.Fatal("title didnt match: expected %s, got %s", title, v)
		return
	}
	if err := dom.Undefine(); err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationSetMetadata(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()
	const domTitle = "newtitle"
	if err := dom.SetMetadata(VIR_DOMAIN_METADATA_TITLE, domTitle, "", "", 0); err != nil {
		t.Error(err)
		return
	}
	v, err := dom.GetMetadata(VIR_DOMAIN_METADATA_TITLE, "", 0)
	if err != nil {
		t.Error(err)
		return
	}
	if v != domTitle {
		t.Fatalf("VIR_DOMAIN_METADATA_TITLE should have been %s, not %s", domTitle, v)
		return
	}
}

func TestIntegrationGetSysinfo(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	info, err := conn.GetSysinfo(0)
	if err != nil {
		t.Error(err)
		return
	}
	if strings.Index(info, "<sysinfo") != 0 {
		t.Fatalf("Sysinfo not valid: %s", info)
		return
	}
}

func testNWFilterXML(name, chain string) string {
	defName := name
	if defName == "" {
		defName = time.Now().String()
	}
	return `<filter name='` + defName + `' chain='` + chain + `'>
            <rule action='drop' direction='out' priority='500'>
            <ip match='no' srcipaddr='$IP'/>
            </rule>
			</filter>`
}

func TestIntergrationDefineUndefineNWFilterXML(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := filter.Undefine(); err != nil {
			t.Fatal(err)
		}
		filter.Free()
	}()
	_, err = conn.NWFilterDefineXML(testNWFilterXML("", "bad"))
	if err == nil {
		t.Fatal("Should have had an error")
	}
}

func TestIntegrationNWFilterGetName(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetName(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationNWFilterGetUUID(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetUUID(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationNWFilterGetUUIDString(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetUUIDString(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationNWFilterGetXMLDesc(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetXMLDesc(0); err != nil {
		t.Error(err)
	}
}

func TestIntegrationLookupNWFilterByName(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	origName := time.Now().String()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML(origName, "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	filter2, err := conn.LookupNWFilterByName(origName)
	if err != nil {
		t.Error(err)
		return
	}
	defer filter2.Free()
	var newName string
	newName, err = filter2.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	if newName != origName {
		t.Fatalf("expected filter name: %s ,got: %s", origName, newName)
	}
}

func TestIntegrationLookupNWFilterByUUIDString(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	origName := time.Now().String()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML(origName, "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	filter2, err := conn.LookupNWFilterByName(origName)
	if err != nil {
		t.Error(err)
		return
	}
	defer filter2.Free()
	var filterUUID string
	filterUUID, err = filter2.GetUUIDString()
	if err != nil {
		t.Error(err)
		return
	}
	filter3, err := conn.LookupNWFilterByUUIDString(filterUUID)
	defer filter3.Free()
	if err != nil {
		t.Error(err)
		return
	}
	name, err := filter3.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	if name != origName {
		t.Fatalf("fetching by UUID: expected filter name: %s ,got: %s", name, origName)
	}
}

func TestIntegrationDomainAttachDetachDevice(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()
	const nwXml = `<interface type='network'>
		<mac address='52:54:00:37:aa:c7'/>
		<source network='default'/>
		<model type='virtio'/>
		</interface>`
	if err := dom.AttachDeviceFlags(nwXml, VIR_DOMAIN_DEVICE_MODIFY_CONFIG); err != nil {
		t.Error(err)
		return
	}
	if err := dom.DetachDeviceFlags(nwXml, VIR_DOMAIN_DEVICE_MODIFY_CONFIG); err != nil {
		t.Error(err)
		return
	}
}

func TestStorageVolResize(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()
	const newCapacityInBytes = 12582912
	if err := vol.Resize(newCapacityInBytes, VIR_STORAGE_VOL_RESIZE_ALLOCATE); err != nil {
		t.Fatal(err)
	}
}

func TestStorageVolWipe(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()
	if err := vol.Wipe(0); err != nil {
		t.Fatal(err)
	}
}

func TestStorageVolWipePattern(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()
	if err := vol.WipePattern(VIR_STORAGE_VOL_WIPE_ALG_ZERO, 0); err != nil {
		t.Fatal(err)
	}
}

func testSecretTypeCephFromXML(name string) string {
	var setName string
	if name == "" {
		setName = time.Now().String()
	} else {
		setName = name
	}
	return `<secret ephemeral='no' private='no'>
            <usage type='ceph'>
            <name>` + setName + `</name>
            </usage>
            </secret>`
}

func TestIntegrationSecretDefineUndefine(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer sec.Free()

	if err := sec.Undefine(); err != nil {
		t.Fatal(err)
	}
}

func TestIntegrationSecretGetUUID(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	if _, err := sec.GetUUID(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSecretGetUUIDString(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	if _, err := sec.GetUUIDString(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSecretGetXMLDesc(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	if _, err := sec.GetXMLDesc(0); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSecretGetUsageType(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	uType, err := sec.GetUsageType()
	if err != nil {
		t.Error(err)
		return
	}
	if uType != VIR_SECRET_USAGE_TYPE_CEPH {
		t.Fatal("unexpected usage type.Expected usage type is Ceph")
	}
}

func TestIntegrationSecretGetUsageID(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	setUsageID := time.Now().String()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(setUsageID), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	recUsageID, err := sec.GetUsageID()
	if err != nil {
		t.Error(err)
		return
	}
	if recUsageID != setUsageID {
		t.Fatalf("exepected usage ID: %s, got: %s", setUsageID, recUsageID)
	}
}

func TestIntegrationLookupSecretByUsage(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	usageID := time.Now().String()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(usageID), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	sec2, err := conn.LookupSecretByUsage(VIR_SECRET_USAGE_TYPE_CEPH, usageID)
	if err != nil {
		t.Fatal(err)
	}
	sec2.Free()
}

func TestIntegrationGetDomainCPUStats(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()

	if err := dom.Create(); err != nil {
		t.Fatal(err)
	}
	defer dom.Destroy()

	// ... if @params is NULL and @nparams is 0 and @ncpus is 0, the
	// number of cpus available to query is returned. From the host perspective,
	ncpus, err := dom.GetCPUStats(nil, 0, 0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if ncpus != 1 {
		t.Fatal("Number of CPUs should be 1, got", ncpus)
	}

	// ... if @params is NULL and @nparams is 0 and @ncpus is 1,
	// and the return value will be how many statistics are available for the given @start_cpu.
	nparams, err := dom.GetCPUStats(nil, 0, 0, 1, 0)
	if err != nil {
		t.Fatal(err)
	}

	const lxcNumParams = 1
	const lxcParamName = "cpu_time"

	if nparams != lxcNumParams {
		t.Fatal("Number of parameters for this hypervisor should be 2, got ", nparams)
	}
	var params VirTypedParameters
	if _, err = dom.GetCPUStats(&params, nparams, 0, uint32(ncpus), 0); err != nil {
		t.Fatal(err)
	}

	if len(params) != lxcNumParams {
		t.Fatalf("Wanted %d returned parameters, got %d", lxcNumParams, len(params))
	}
	param := params[0]
	if param.Name != lxcParamName {
		t.Fatalf("Wanted param '%s', got '%s'", lxcParamName, param.Name)
	}
	if _, ok := param.Value.(uint64); !ok {
		t.Fatalf("Wanted uint64 param, got %v instead", param.Value)
	}
}

// Not supported on libvirt driver, so no integration test
// func TestGetInterfaceParameters(t *testing.T) {
// 	dom, conn := buildTestDomain()
// 	defer func() {
// 		dom.Undefine()
// 		dom.Free()
// 		if res, _ := conn.CloseConnection(); res != 0 {
// 			t.Errorf("CloseConnection() == %d, expected 0", res)
// 		}
// 	}()
// 	iface := "either mac or path to interface"
// 	nparams := int(0)
// 	if _, err := dom.GetInterfaceParameters(iface, nil, &nparams, 0); err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	var params VirTypedParameters
// 	if _, err := dom.GetInterfaceParameters(iface, &params, &nparams, 0); err != nil {
// 		t.Error(err)
// 		return
// 	}
// }

func TestIntegrationListAllInterfaces(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	ifaces, err := conn.ListAllInterfaces(0)
	if err != nil {
		t.Fatal(err)
	}
	lookingFor := "lo"
	found := false
	for _, iface := range ifaces {
		name, err := iface.GetName()
		if err != nil {
			t.Fatal(err)
		}
		if name == lookingFor {
			found = true
		}
		iface.Free()
	}
	if found == false {
		t.Fatalf("interface %s not found", lookingFor)
	}
}

func TestIntergrationListAllNWFilters(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	testNWFilterName := time.Now().String()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML(testNWFilterName, "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()

	filters, err := conn.ListAllNWFilters(0)
	if len(filters) == 0 {
		t.Fatal("length of []NWFilter shouldn't be 0")
	}

	found := false
	for _, f := range filters {
		name, _ := f.GetName()
		if name == testNWFilterName {
			found = true
		}
		f.Free()
	}
	if found == false {
		t.Fatalf("NWFilter %s not found", testNWFilterName)
	}
}

func TestIntegrationDomainBlockStatsFlags(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()

	if err := dom.Create(); err != nil {
		t.Fatal(err)
	}
	defer dom.Destroy()

	// special case, count number of parameters
	_, err = dom.BlockStatsFlags("", nil, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegrationDomainInterfaceStats(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()
	const nwXml = `<interface type='network'>
		<mac address='52:54:00:37:aa:c7'/>
		<source network='default'/>
		<model type='virtio'/>
		</interface>`
	if err := dom.AttachDeviceFlags(nwXml, VIR_DOMAIN_DEVICE_MODIFY_CONFIG); err != nil {
		t.Fatal(err)
	}

	if err := dom.Create(); err != nil {
		t.Fatal(err)
	}

	if _, err := dom.InterfaceStats("vnet0"); err != nil {
		t.Error(err)
	}

	if err := dom.Destroy(); err != nil {
		t.Fatal(err)
	}

	if err := dom.DetachDeviceFlags(nwXml, VIR_DOMAIN_DEVICE_MODIFY_CONFIG); err != nil {
		t.Fatal(err)
	}
}

func TestStorageVolUploadDownload(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()

	data := []byte{1, 2, 3, 4, 5, 6}

	// write above data to the vol
	// 1. create a stream
	stream, err := NewStream(&conn, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		stream.Free()
	}()

	// 2. set it up to upload from stream
	if err := vol.Upload(stream, 0, uint64(len(data)), 0); err != nil {
		stream.Abort()
		t.Fatal(err)
	}

	// 3. do the actual writing
	if n, err := stream.Write(data); err != nil || n != len(data) {
		t.Fatal(err, n)
	}

	// 4. finish!
	if err := stream.Close(); err != nil {
		t.Fatal(err)
	}

	// read back the data
	// 1. create a stream
	downStream, err := NewStream(&conn, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		downStream.Free()
	}()

	// 2. set it up to download from stream
	if err := vol.Download(downStream, 0, uint64(len(data)), 0); err != nil {
		downStream.Abort()
		t.Fatal(err)
	}

	// 3. do the actual reading
	buf := make([]byte, 1024)
	if n, err := downStream.Read(buf); err != nil || n != len(data) {
		t.Fatal(err, n)
	}

	t.Logf("read back: %#v", buf[:len(data)])

	// 4. finish!
	if err := downStream.Close(); err != nil {
		t.Fatal(err)
	}
}

/*func TestDomainMemoryStats(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()
	if err := dom.Create(); err != nil {
		t.Fatal(err)
	}
	defer dom.Destroy()

	ms, err := dom.MemoryStats(1, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(ms) != 1 {
		t.Fatal("Should have got one result, got", len(ms))
	}
}*/

func TestDomainListAllInterfaceAddresses(t *testing.T) {
	dom, conn := buildTestQEMUDomain()
	defer func() {
		dom.Free()
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()
	if err := dom.Create(); err != nil {
		t.Error(err)
		return
	}
	defer dom.Destroy()
	ifaces, err := dom.ListAllInterfaceAddresses(0)
	if err != nil {
		t.Fatal(err)
	}

	if len(ifaces) != 0 {
		t.Fatal("should have 0 interfaces", len(ifaces))
	}
}
