package client

import "github.com/hypebeast/go-osc/osc"

func (v *VRCOSC) Vertical(input float32) error {
	msg := osc.NewMessage("/input/Vertical")
	msg.Append(input)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) Horizontal(input float32) error {
	msg := osc.NewMessage("/input/Horizontal")
	msg.Append(input)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) LookVertical(input float32) error {
	msg := osc.NewMessage("/input/LookVertical")
	msg.Append(input)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) LookHorizontal(input float32) error {
	msg := osc.NewMessage("/input/LookHorizontal")
	msg.Append(input)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}
