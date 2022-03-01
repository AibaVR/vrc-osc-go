package client

import (
	"github.com/hypebeast/go-osc/osc"
	"time"
)

type VRCOSC struct {
	Client *osc.Client
}

const SEND_PORT = 9000

func NewVRCOSC() *VRCOSC {
	client := osc.NewClient("127.0.0.1", SEND_PORT)
	return &VRCOSC{Client: client}
}

func NewOSC(addr string, port int) *VRCOSC {
	client := osc.NewClient(addr, port)
	return &VRCOSC{Client: client}
}

func (v *VRCOSC) sendStartInput(input string) error {
	msg := osc.NewMessage("/input/" + input)
	msg.Append(int32(1))
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) sendStopInput(input string) error {
	msg := osc.NewMessage("/input/" + input)
	msg.Append(false)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) JumpAuto() error {
	// Send the jump message
	if err := v.Jump(); err != nil {
		return err
	}

	// We need a delay to stop the jump or else the jump won't execute
	//time.Sleep(time.Nanosecond * 10)
	time.Sleep(time.Nanosecond)
	// Send the jump message
	if err := v.StopJump(); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) ToggleProne() error {
	msg := osc.NewMessage("/input/ToggleSitStand")
	msg.Append(int32(-1))
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) ToggleSit() error {
	msg := osc.NewMessage("/input/ToggleSitStand")
	msg.Append(int32(1))
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) ToggleStand() error {
	msg := osc.NewMessage("/input/ToggleSitStand")
	msg.Append(false)
	if err := v.Client.Send(msg); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) MoveForward() error {
	if err := v.sendStartInput("MoveForward"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopMoveForward() error {
	if err := v.sendStopInput("MoveForward"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) MoveBackward() error {
	if err := v.sendStartInput("MoveBackward"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopMoveBackward() error {
	if err := v.sendStopInput("MoveBackward"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) MoveLeft() error {
	if err := v.sendStartInput("MoveLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopMoveLeft() error {
	if err := v.sendStopInput("MoveLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) MoveRight() error {
	if err := v.sendStartInput("MoveRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopMoveRight() error {
	if err := v.sendStopInput("MoveRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) LookLeft() error {
	if err := v.sendStartInput("LookLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopLookLeft() error {
	if err := v.sendStopInput("LookLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) LookRight() error {
	if err := v.sendStartInput("LookRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopLookRight() error {
	if err := v.sendStopInput("LookRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) LookDown() error {
	if err := v.sendStartInput("LookDown"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopLookDown() error {
	if err := v.sendStopInput("LookDown"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) LookUp() error {
	if err := v.sendStartInput("LookUp"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopLookUp() error {
	if err := v.sendStopInput("LookUp"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) Jump() error {
	if err := v.sendStartInput("Jump"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopJump() error {
	if err := v.sendStopInput("Jump"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) Run() error {
	if err := v.sendStartInput("Run"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopRun() error {
	if err := v.sendStopInput("Run"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) Back() error {
	if err := v.sendStartInput("Back"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopBack() error {
	if err := v.sendStopInput("Back"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) Menu() error {
	if err := v.sendStartInput("Menu"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopMenu() error {
	if err := v.sendStopInput("Menu"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) ComfortLeft() error {
	if err := v.sendStartInput("ComfortLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopComfortLeft() error {
	if err := v.sendStopInput("ComfortLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) ComfortRight() error {
	if err := v.sendStartInput("ComfortRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopComfortRight() error {
	if err := v.sendStopInput("ComfortRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) DropRight() error {
	if err := v.sendStartInput("DropRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopDropRight() error {
	if err := v.sendStopInput("DropRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) UseRight() error {
	if err := v.sendStartInput("UseRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopUseRight() error {
	if err := v.sendStopInput("UseRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) GrabRight() error {
	if err := v.sendStartInput("GrabRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopGrabRight() error {
	if err := v.sendStopInput("GrabRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) DropLeft() error {
	if err := v.sendStartInput("DropLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopDropLeft() error {
	if err := v.sendStopInput("DropLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) UseLeft() error {
	if err := v.sendStartInput("UseLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopUseLeft() error {
	if err := v.sendStopInput("UseLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) GrabLeft() error {
	if err := v.sendStartInput("GrabLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopGrabLeft() error {
	if err := v.sendStopInput("GrabLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) PanicButton() error {
	if err := v.sendStartInput("PanicButton"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopPanicButton() error {
	if err := v.sendStopInput("PanicButton"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) QuickMenuToggleLeft() error {
	if err := v.sendStartInput("QuickMenuToggleLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopQuickMenuToggleLeft() error {
	if err := v.sendStopInput("QuickMenuToggleLeft"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) QuickMenuToggleRight() error {
	if err := v.sendStartInput("QuickMenuToggleRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopQuickMenuToggleRight() error {
	if err := v.sendStopInput("QuickMenuToggleRight"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) AFKToggle() error {
	if err := v.sendStartInput("AFKToggle"); err != nil {
		return err
	}

	return nil
}

func (v *VRCOSC) StopAFKToggle() error {
	if err := v.sendStopInput("AFKToggle"); err != nil {
		return err
	}

	return nil
}
