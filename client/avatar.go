package client

import "github.com/hypebeast/go-osc/osc"

func (v *VRCOSC) SendAvatarParam(param string, value interface{}) error {
	msg := osc.NewMessage("/avatar/parameters/" + param)

	// For some reason sending "true" isn't working, so we have to send an int value of 1 for bool
	switch value.(type) {
	case bool:
		if value == true {
			value = int32(1)
		} else {
			value = int32(0)
		}
	}

	msg.Append(value)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}
