package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gitchander/go-bluetooth/api/service"
	"github.com/gitchander/go-bluetooth/bluez/profile/agent"
	"github.com/gitchander/go-bluetooth/bluez/profile/gatt"
	"github.com/gitchander/go-bluetooth/utils/random"
)

func main() {

	const adapterID = "hci0"

	log.SetLevel(log.DebugLevel)

	err := serve(context.Background(), adapterID)
	if err != nil {
		log.Error(err)
		return
	}

	// fmt.Println("stop")
	// time.Sleep(60 * time.Second)
	// fmt.Println("start")
	time.Sleep(2 * time.Second)

	err = serve(context.Background(), adapterID)
	if err != nil {
		log.Error(err)
		return
	}
}

func serve(ctx context.Context, adapterID string) error {

	var (
		//agentCap = agent.CapDisplayOnly
		agentCap = agent.CapDisplayYesNo
		//agentCap = agent.CapKeyboardOnly
		//agentCap = agent.CapNoInputNoOutput
		//agentCap = agent.CapKeyboardDisplay
	)

	options := service.AppOptions{
		AdapterID: adapterID,

		AgentCaps:         agentCap,
		AgentSetAsDefault: false,

		UUID:       "1234",
		UUIDSuffix: "-0000-1000-8000-00805F9B34FB",
	}

	app, err := service.NewApp(options)
	if err != nil {
		return err
	}
	defer app.Close()

	// Set passkey
	{
		agent1 := app.Agent()
		as, ok := agent1.(*agent.SimpleAgent)
		if ok {
			as.SetPassKey(237710)
			//as.SetPassCode("54321")
		}
	}

	app.SetName("go_bluetooth")

	// Init adapter
	{
		adapter := app.Adapter()

		log.Infof("HW address %s", adapter.Properties.Address)

		if !(adapter.Properties.Powered) {
			err = adapter.SetPowered(true)
			if err != nil {
				log.Fatalf("Failed to power the adapter: %s", err)
			}
		}

		err = adapter.SetPairable(true)
		if err != nil {
			return err
		}
		// err = adapter.SetPairableTimeout(1000)
		// if err != nil {
		// 	return err
		// }
	}

	service1, err := app.NewService("2233")
	if err != nil {
		return err
	}

	err = app.AddService(service1)
	if err != nil {
		return err
	}

	char1, err := service1.NewChar("3344")
	if err != nil {
		return err
	}

	// Set char flags
	{
		//char1.Properties.Flags = []string{
		// gatt.FlagCharacteristicRead,
		// gatt.FlagCharacteristicWrite,
		// gatt.FlagCharacteristicNotify,
		//}

		char1.Properties.Flags = []string{
			gatt.FlagCharacteristicEncryptRead,
			gatt.FlagCharacteristicEncryptWrite,
			gatt.FlagCharacteristicNotify,
		}

		// char1.Properties.Flags = []string{
		// 	"encrypt-read",
		// 	"encrypt-write",
		// 	"notify",
		// 	//			"encrypt-notify", // It not works!!!
		// }
	}

	char1.OnRead(
		func(c *service.Char, options service.Options) ([]byte, error) {

			log.Debug("CharReadCallback:")
			printJSON("options: ", options)

			value := []byte("Hello, World!")
			log.Debugf("value string: %s\n", value)
			log.Debugf("value hex: [% x]\n", value)

			return value, nil
		})

	char1.OnWrite(
		func(c *service.Char, value []byte, options service.Options) ([]byte, error) {

			log.Debug("CharWriteCallback:")
			log.Debugf("value: [% x]\n", value)
			printJSON("options: ", options)

			return value, nil
		})

	char1.OnNotify(
		func(c *service.Char, notify bool) error {
			var message string
			if notify {
				message = "Notify start"
			} else {
				message = "Notify stop"
			}
			log.Debug("OnNotify: ", message)
			return nil
		})

	err = service1.AddChar(char1)
	if err != nil {
		return err
	}

	err = app.Run()
	if err != nil {
		return err
	}

	log.Infof("Exposed service %s", service1.Properties.UUID)

	// Advertise
	{
		adv := app.GetAdvertisement()
		printJSON("adv:", adv)
	}

	timeout := uint32(6 * 3600) // 6h
	log.Infof("Advertising for %ds...", timeout)
	cancel, err := app.Advertise(timeout)
	if err != nil {
		return err
	}
	defer cancel()

	//data := make([]byte, 4096)
	data := make([]byte, 100)

	r := random.NewRandomer(random.NextRand())

	dur := 2 * time.Second
	cctx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(dur))
	defer cancelFunc()

	for {
		select {
		case <-cctx.Done():
			return nil
		case <-time.After(10 * time.Second):
			r.FillBytes(data)
			char1.WriteValue(data, nil)
		}
	}
}

func printJSON(prefix string, v any) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(prefix + string(data))
}
