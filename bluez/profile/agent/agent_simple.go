package agent

import (
	"fmt"
	"sync/atomic"

	"github.com/godbus/dbus/v5"
	log "github.com/sirupsen/logrus"

	"github.com/gitchander/go-bluetooth/bluez/profile/adapter"
)

const (
	AgentBasePath             = "/agent/simple%d"
	SimpleAgentPinCode        = "0000"
	SimpleAgentPassKey uint32 = 1024
)

var agentInstances atomic.Uint32

func nextAgentInstances() int {
	return int(agentInstances.Add(1) - 1)
}

func NextAgentPath() dbus.ObjectPath {
	ai := nextAgentInstances()
	p := dbus.ObjectPath(fmt.Sprintf(AgentBasePath, ai))
	return p
}

// NewDefaultSimpleAgent return a SimpleAgent instance with default pincode and passcode
func NewDefaultSimpleAgent() *SimpleAgent {
	ag := &SimpleAgent{
		path:    NextAgentPath(),
		passKey: SimpleAgentPassKey,
		pinCode: SimpleAgentPinCode,
	}

	return ag
}

// NewSimpleAgent return a SimpleAgent instance
func NewSimpleAgent() *SimpleAgent {
	ag := &SimpleAgent{
		path: NextAgentPath(),
	}
	return ag
}

// SimpleAgent implement interface Agent1Client
type SimpleAgent struct {
	path    dbus.ObjectPath
	pinCode string
	passKey uint32
}

var _ Agent1Client = &SimpleAgent{}

func (self *SimpleAgent) SetPassKey(passkey uint32) {
	self.passKey = passkey
}

func (self *SimpleAgent) SetPassCode(pinCode string) {
	self.pinCode = pinCode
}

func (self *SimpleAgent) PassKey() uint32 {
	return self.passKey
}

func (self *SimpleAgent) PassCode() string {
	return self.pinCode
}

func (self *SimpleAgent) Path() dbus.ObjectPath {
	return self.path
}

func (self *SimpleAgent) Interface() string {
	return Agent1Interface
}

func (self *SimpleAgent) Release() error {
	return nil
}

func (self *SimpleAgent) RequestPinCode(path dbus.ObjectPath) (pincode string, err error) {

	log.Debugf("SimpleAgent: RequestPinCode: %s", path)

	adapterID, err := adapter.ParseAdapterID(path)
	if err != nil {
		log.Warnf("SimpleAgent.RequestPinCode: Failed to load adapter %s", err)
		return "", dbus.MakeFailedError(err)
	}

	err = SetTrusted(adapterID, path)
	if err != nil {
		log.Errorf("SimpleAgent.RequestPinCode SetTrusted failed: %s", err)
		return "", dbus.MakeFailedError(err)
	}

	pincode = self.pinCode
	log.Debugf("SimpleAgent: Returning pin code: %s", pincode)
	return pincode, nil
}

func (self *SimpleAgent) DisplayPinCode(device dbus.ObjectPath, pincode string) error {
	log.Info(fmt.Sprintf("SimpleAgent: DisplayPinCode (%s, %s)", device, pincode))
	return nil
}

func (self *SimpleAgent) RequestPasskey(path dbus.ObjectPath) (passkey uint32, err error) {

	adapterID, err := adapter.ParseAdapterID(path)
	if err != nil {
		log.Warnf("SimpleAgent.RequestPassKey: Failed to load adapter %s", err)
		return 0, dbus.MakeFailedError(err)
	}

	err = SetTrusted(adapterID, path)
	if err != nil {
		log.Errorf("SimpleAgent.RequestPassKey: SetTrusted %s", err)
		return 0, dbus.MakeFailedError(err)
	}

	log.Debugf("RequestPasskey: returning %d", self.passKey)

	passkey = self.passKey

	return passkey, nil
}

func (self *SimpleAgent) DisplayPasskey(device dbus.ObjectPath, passkey uint32, entered uint16) error {
	log.Debugf("SimpleAgent: DisplayPasskey %s, %06d entered %d", device, passkey, entered)
	return nil
}

func (self *SimpleAgent) RequestConfirmation(path dbus.ObjectPath, passkey uint32) error {

	log.Debugf("SimpleAgent: RequestConfirmation (%s, %06d)", path, passkey)

	adapterID, err := adapter.ParseAdapterID(path)
	if err != nil {
		log.Warnf("SimpleAgent: Failed to load adapter %s", err)
		return dbus.MakeFailedError(err)
	}

	err = SetTrusted(adapterID, path)
	if err != nil {
		log.Warnf("Failed to set trust for %s: %s", path, err)
		return dbus.MakeFailedError(err)
	}

	log.Debug("SimpleAgent: RequestConfirmation OK")
	return nil
}

func (self *SimpleAgent) RequestAuthorization(device dbus.ObjectPath) error {
	log.Debugf("SimpleAgent: RequestAuthorization (%s)", device)
	return nil
}

func (self *SimpleAgent) AuthorizeService(device dbus.ObjectPath, uuid string) error {
	log.Debugf("SimpleAgent: AuthorizeService (%s, %s)", device, uuid) // directly authorized
	return nil
}

func (self *SimpleAgent) Cancel() error {
	log.Debugf("SimpleAgent: Cancel")
	return nil
}
