// Code generated DO NOT EDIT

package mesh



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus/v5"
)

var Node1Interface = "org.bluez.mesh.Node1"


// NewNode1 create a new instance of Node1
//
// Args:

func NewNode1(objectPath dbus.ObjectPath) (*Node1, error) {
	a := new(Node1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez.mesh",
			Iface: Node1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Node1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Node1 Mesh Node Hierarchy

*/
type Node1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Node1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Node1Properties contains the exposed properties of an interface
type Node1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Beacon This property indicates whether the periodic beaconing is
		enabled (true) or disabled (false).

	uint8 BeaconFlags [read-only]

		This property may be read at any time to determine the flag
		field setting on sent and received beacons of the primary
		network key.
	*/
	Beacon bool

	/*
	IvIndex This property may be read at any time to determine the IV_Index
		that the current network is on. This information is only useful
		for provisioning.
	*/
	IvIndex uint32

	/*
	Proxy Indicates support for GATT proxy
	*/
	Proxy bool

	/*
	Relay Indicates support for relaying messages

	If a key is absent from the dictionary, the feature is not supported.
	Otherwise, true means that the feature is enabled and false means that
	the feature is disabled.
	*/
	Relay bool

	/*
	LowPower Indicates support for operating in Low Power node mode
	*/
	LowPower bool

	/*
	SecondsSinceLastHeard This property may be read at any time to determine the number of
		seconds since mesh network layer traffic was last detected on
		this node's network.
	*/
	SecondsSinceLastHeard uint32

	/*
	Addresses This property contains unicast addresses of node's elements.
	*/
	Addresses []uint16

	/*
	SequenceNumber This property may be read at any time to determine the
		sequence number.
	*/
	SequenceNumber uint32

	/*
	Features The dictionary that contains information about feature support.
		The following keys are defined:
	*/
	Features map[string]interface{}

	/*
	Friend Indicates the ability to establish a friendship with a
			Low Power node
	*/
	Friend bool

}

//Lock access to properties
func (p *Node1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Node1Properties) Unlock() {
	p.lock.Unlock()
}




// SetBeacon set Beacon value
func (a *Node1) SetBeacon(v bool) error {
	return a.SetProperty("Beacon", v)
}



// GetBeacon get Beacon value
func (a *Node1) GetBeacon() (bool, error) {
	v, err := a.GetProperty("Beacon")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}




// SetIvIndex set IvIndex value
func (a *Node1) SetIvIndex(v uint32) error {
	return a.SetProperty("IvIndex", v)
}



// GetIvIndex get IvIndex value
func (a *Node1) GetIvIndex() (uint32, error) {
	v, err := a.GetProperty("IvIndex")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}




// SetProxy set Proxy value
func (a *Node1) SetProxy(v bool) error {
	return a.SetProperty("Proxy", v)
}



// GetProxy get Proxy value
func (a *Node1) GetProxy() (bool, error) {
	v, err := a.GetProperty("Proxy")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}




// SetRelay set Relay value
func (a *Node1) SetRelay(v bool) error {
	return a.SetProperty("Relay", v)
}



// GetRelay get Relay value
func (a *Node1) GetRelay() (bool, error) {
	v, err := a.GetProperty("Relay")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}




// SetLowPower set LowPower value
func (a *Node1) SetLowPower(v bool) error {
	return a.SetProperty("LowPower", v)
}



// GetLowPower get LowPower value
func (a *Node1) GetLowPower() (bool, error) {
	v, err := a.GetProperty("LowPower")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}




// SetSecondsSinceLastHeard set SecondsSinceLastHeard value
func (a *Node1) SetSecondsSinceLastHeard(v uint32) error {
	return a.SetProperty("SecondsSinceLastHeard", v)
}



// GetSecondsSinceLastHeard get SecondsSinceLastHeard value
func (a *Node1) GetSecondsSinceLastHeard() (uint32, error) {
	v, err := a.GetProperty("SecondsSinceLastHeard")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}




// SetAddresses set Addresses value
func (a *Node1) SetAddresses(v []uint16) error {
	return a.SetProperty("Addresses", v)
}



// GetAddresses get Addresses value
func (a *Node1) GetAddresses() ([]uint16, error) {
	v, err := a.GetProperty("Addresses")
	if err != nil {
		return []uint16{}, err
	}
	return v.Value().([]uint16), nil
}




// SetSequenceNumber set SequenceNumber value
func (a *Node1) SetSequenceNumber(v uint32) error {
	return a.SetProperty("SequenceNumber", v)
}



// GetSequenceNumber get SequenceNumber value
func (a *Node1) GetSequenceNumber() (uint32, error) {
	v, err := a.GetProperty("SequenceNumber")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}




// SetFeatures set Features value
func (a *Node1) SetFeatures(v map[string]interface{}) error {
	return a.SetProperty("Features", v)
}



// GetFeatures get Features value
func (a *Node1) GetFeatures() (map[string]interface{}, error) {
	v, err := a.GetProperty("Features")
	if err != nil {
		return map[string]interface{}{}, err
	}
	return v.Value().(map[string]interface{}), nil
}




// SetFriend set Friend value
func (a *Node1) SetFriend(v bool) error {
	return a.SetProperty("Friend", v)
}



// GetFriend get Friend value
func (a *Node1) GetFriend() (bool, error) {
	v, err := a.GetProperty("Friend")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}



// Close the connection
func (a *Node1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Node1 object path
func (a *Node1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Node1 dbus client
func (a *Node1) Client() *bluez.Client {
	return a.client
}

// Interface return Node1 interface
func (a *Node1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Node1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Node1Properties to map
func (a *Node1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Node1Properties
func (a *Node1Properties) FromMap(props map[string]interface{}) (*Node1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Node1Properties
func (a *Node1Properties) FromDBusMap(props map[string]dbus.Variant) (*Node1Properties, error) {
	s := new(Node1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Node1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Node1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Node1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Node1) GetProperties() (*Node1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Node1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Node1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Node1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Node1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Node1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Node1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Send 
		This method is used to send a message originated by a local
		model.

		The element_path parameter is the object path of an element from
		a collection of the application elements (see Mesh Application
		Hierarchy section).

		The destination parameter contains the destination address. This
		destination must be a uint16 to a unicast address, or a well
		known group address.

		The key_index parameter determines which application key to use
		for encrypting the message. The key_index must be valid for that
		element, i.e., the application key must be bound to a model on
		this element. Otherwise, org.bluez.mesh.Error.NotAuthorized will
		be returned.

		The data parameter is an outgoing message to be encypted by the
		bluetooth-meshd daemon and sent on.

		Possible errors:
			org.bluez.mesh.Error.NotAuthorized
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.NotFound


*/
func (a *Node1) Send(element_path dbus.ObjectPath, destination uint16, key_index uint16, data []byte) error {
	
	return a.client.Call("Send", 0, element_path, destination, key_index, data).Store()
	
}

/*
DevKeySend 
		This method is used to send a message originated by a local
		model encoded with the device key of the remote node.

		The element_path parameter is the object path of an element from
		a collection of the application elements (see Mesh Application
		Hierarchy section).

		The destination parameter contains the destination address. This
		destination must be a uint16 to a unicast address, or a well
		known group address.

		The remote parameter, if true, looks up the device key by the
		destination address in the key database to encrypt the message.
		If remote is true, but requested key does not exist, a NotFound
		error will be returned. If set to false, the local node's
		device key is used.

		The net_index parameter is the subnet index of the network on
		which the message is to be sent.

		The data parameter is an outgoing message to be encypted by the
		meshd daemon and sent on.

		Possible errors:
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.NotFound


*/
func (a *Node1) DevKeySend(element_path dbus.ObjectPath, destination uint16, remote bool, net_index uint16, data []byte) error {
	
	return a.client.Call("DevKeySend", 0, element_path, destination, remote, net_index, data).Store()
	
}

/*
AddNetKey 
		This method is used to send add or update network key originated
		by the local configuration client to a remote configuration
		server.

		The element_path parameter is the object path of an element from
		a collection of the application elements (see Mesh Application
		Hierarchy section).

		The destination parameter contains the destination address. This
		destination must be a uint16 to a nodes primary unicast address.

		The subnet_index parameter refers to the subnet index of the
		network that is being added or updated. This key must exist in
		the local key database.

		The net_index parameter is the subnet index of the network on
		which the message is to be sent.

		The update parameter indicates if this is an addition or an
		update. If true, the subnet key must be in the phase 1 state of
		the key update procedure.

		Possible errors:
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.NotFound


*/
func (a *Node1) AddNetKey(element_path dbus.ObjectPath, destination uint16, subnet_index uint16, net_index uint16, update bool) error {
	
	return a.client.Call("AddNetKey", 0, element_path, destination, subnet_index, net_index, update).Store()
	
}

/*
AddAppKey 
		This method is used to send add or update network key originated
		by the local configuration client to a remote configuration
		server.

		The element_path parameter is the object path of an element from
		a collection of the application elements (see Mesh Application
		Hierarchy section).

		The destination parameter contains the destination address. This
		destination must be a uint16 to a nodes primary unicast address.

		The app_index parameter refers to the application key which is
		being added or updated. This key must exist in the local key
		database.

		The net_index parameter is the subnet index of the network on
		which the message is to be sent.

		The update parameter indicates if this is an addition or an
		update. If true, the subnet key must be in the phase 1 state of
		the key update procedure.

		Possible errors:
			org.bluez.mesh.Error.InvalidArguments
			org.bluez.mesh.Error.NotFound


*/
func (a *Node1) AddAppKey(element_path dbus.ObjectPath, destination uint16, app_index uint16, net_index uint16, update bool) error {
	
	return a.client.Call("AddAppKey", 0, element_path, destination, app_index, net_index, update).Store()
	
}

/*
Publish 
		This method is used to send a publication originated by a local
		model. If the model does not exist, or it has no publication
		record, the method returns org.bluez.mesh.Error.DoesNotExist
		error.

		The element_path parameter is the object path of an element from
		a collection of the application elements (see Mesh Application
		Hierarchy section).

		The model parameter contains a model ID, as defined by the
		Bluetooth SIG.

		Since only one Publish record may exist per element-model, the
		destination and key_index are obtained from the Publication
		record cached by the daemon.

		Possible errors:
			org.bluez.mesh.Error.DoesNotExist
			org.bluez.mesh.Error.InvalidArguments


*/
func (a *Node1) Publish(element_path dbus.ObjectPath, model uint16, data []byte) error {
	
	return a.client.Call("Publish", 0, element_path, model, data).Store()
	
}

/*
VendorPublish 
		This method is used to send a publication originated by a local
		vendor model. If the model does not exist, or it has no
		publication record, the method returns
		org.bluez.mesh.Error.DoesNotExist error.

		The element_path parameter is the object path of an element from
		a collection of the application elements (see Mesh Application
		Hierarchy section).

		The vendor parameter is a 16-bit Bluetooth-assigned Company ID.

		The model_id parameter is a 16-bit vendor-assigned Model
		Identifier.

		Since only one Publish record may exist per element-model, the
		destination and key_index are obtained from the Publication
		record cached by the daemon.

		Possible errors:
			org.bluez.mesh.Error.DoesNotExist
			org.bluez.mesh.Error.InvalidArguments

Properties:
	dict Features [read-only]

		The dictionary that contains information about feature support.
		The following keys are defined:

		boolean Friend

			Indicates the ability to establish a friendship with a
			Low Power node

		boolean LowPower

			Indicates support for operating in Low Power node mode

		boolean Proxy

			Indicates support for GATT proxy

		boolean Relay
			Indicates support for relaying messages

	If a key is absent from the dictionary, the feature is not supported.
	Otherwise, true means that the feature is enabled and false means that
	the feature is disabled.

	boolean Beacon [read-only]

		This property indicates whether the periodic beaconing is
		enabled (true) or disabled (false).

	uint8 BeaconFlags [read-only]

		This property may be read at any time to determine the flag
		field setting on sent and received beacons of the primary
		network key.

	uint32 IvIndex [read-only]

		This property may be read at any time to determine the IV_Index
		that the current network is on. This information is only useful
		for provisioning.

	uint32 SecondsSinceLastHeard [read-only]

		This property may be read at any time to determine the number of
		seconds since mesh network layer traffic was last detected on
		this node's network.

	array{uint16} Addresses [read-only]

		This property contains unicast addresses of node's elements.

	uint32 SequenceNumber [read-only]

		This property may be read at any time to determine the
		sequence number.


*/
func (a *Node1) VendorPublish(element_path dbus.ObjectPath, vendor uint16, model_id uint16, data []byte) error {
	
	return a.client.Call("VendorPublish", 0, element_path, vendor, model_id, data).Store()
	
}

