package gopherjs

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/events"
	"github.com/gu-io/gu/notifications/mque"
	"github.com/gu-io/gu/shell"
)

// GetEvent returns the appropriate event from the provided structures.
func GetEvent(ev *js.Object, handle mque.End) *events.BaseEvent {
	if ev == nil || ev == js.Undefined {
		return nil
	}

	c := ev.Get("constructor")

	switch c {
	case js.Global.Get("AnimationEvent"):
		return events.NewBaseEvent(&events.AnimationEvent{
			Core:          ev,
			AnimationName: ev.Get("animationName").String(),
			ElapsedTime:   ev.Get("elapsedTime").Float(),
		}, handle)
	case js.Global.Get("AudioProcessingEvent"):
		return events.NewBaseEvent(&events.AudioProcessingEvent{
			Core:         ev,
			PlaybackTime: ev.Get("playbackTime").Float(),
		}, handle)
	case js.Global.Get("BeforeInputEvent"):
		return events.NewBaseEvent(&events.BeforeInputEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("BeforeUnloadEvent"):
		return events.NewBaseEvent(&events.BeforeUnloadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("BlobEvent"):
		return events.NewBaseEvent(&events.BlobEvent{
			Core: ev,
			Data: fromBlob(ev.Get("data")),
		}, handle)
	case js.Global.Get("ChangeEvent"):
		return events.NewBaseEvent(&events.ChangeEvent{
			Core:  ev,
			Value: ev.Get("target").Get("value").String(),
		}, handle)
	case js.Global.Get("ClipboardEvent"):
		return events.NewBaseEvent(&events.ClipboardEvent{
			Core: ev,
			Data: fromDataTransfer(ev.Get("clipboardData")),
		}, handle)
	case js.Global.Get("CloseEvent"):
		return events.NewBaseEvent(&events.CloseEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("CompositionEvent"):
		return events.NewBaseEvent(&events.CompositionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("CSSFontFaceLoadEvent"):
		return events.NewBaseEvent(&events.CSSFontFaceLoadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("CustomEvent"):
		return events.NewBaseEvent(&events.CustomEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DeviceLightEvent"):
		return events.NewBaseEvent(&events.DeviceLightEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DeviceMotionEvent"):
		return events.NewBaseEvent(&events.DeviceMotionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DeviceOrientationEvent"):
		return events.NewBaseEvent(&events.DeviceOrientationEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DeviceProximityEvent"):
		return events.NewBaseEvent(&events.DeviceProximityEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DOMTransactionEvent"):
		return events.NewBaseEvent(&events.DOMTransactionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragStartEvent"):
		return events.NewBaseEvent(&events.DragStartEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragEndEvent"):
		return events.NewBaseEvent(&events.DragEndEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragEnterEvent"):
		return events.NewBaseEvent(&events.DragEnterEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragLeaveEvent"):
		return events.NewBaseEvent(&events.DragLeaveEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragOverEvent"):
		return events.NewBaseEvent(&events.DragOverEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DropEvent"):
		return events.NewBaseEvent(&events.DropEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("DragEvent"):
		return events.NewBaseEvent(&events.DragEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("EditingBeforeInputEvent"):
		return events.NewBaseEvent(&events.EditingBeforeInputEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("ErrorEvent"):
		return events.NewBaseEvent(&events.ErrorEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("FocusEvent"):
		return events.NewBaseEvent(&events.FocusEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("GamepadEvent"):
		return events.NewBaseEvent(&events.GamepadEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("HashChangeEvent"):
		return events.NewBaseEvent(&events.HashChangeEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("IDBVersionChangeEvent"):
		return events.NewBaseEvent(&events.IDBVersionChangeEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("KeyboardEvent"):
		return events.NewBaseEvent(&events.KeyboardEvent{
			Core:          ev,
			CharCode:      ev.Get("charCode").Int(),
			Key:           ev.Get("key").String(),
			Locale:        ev.Get("locale").String(),
			AltKey:        ev.Get("altKey").Bool(),
			CtrlKey:       ev.Get("ctrlKey").Bool(),
			MetaKey:       ev.Get("metaKey").Bool(),
			ShiftKey:      ev.Get("shiftKey").Bool(),
			Repeat:        ev.Get("repeat").Bool(),
			Location:      ev.Get("location").Int(),
			ModifiedState: ev.Get("getModifierState").Bool(),
			KeyIdentifier: ev.Get("keyIdentifier").String(),
			KeyLocation:   events.KeyLocation(ev.Get("KeyLocation").Uint64()),
			KeyCode:       events.KeyCode(ev.Get("keyCode").Uint64()),
		}, handle)
	case js.Global.Get("MediaStreamEvent"):
		return events.NewBaseEvent(&events.MediaStreamEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("MessageEvent"):
		return events.NewBaseEvent(&events.MessageEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("MouseEvent"):
		return events.NewBaseEvent(&events.MouseEvent{
			UIEvent: &events.UIEvent{
				Core: ev,
			},
			ClientX:  ev.Get("clientX").Float(),
			ClientY:  ev.Get("clientY").Float(),
			OffsetX:  ev.Get("offsetX").Float(),
			OffsetY:  ev.Get("offsetY").Float(),
			PageX:    ev.Get("pageX").Float(),
			PageY:    ev.Get("pageY").Float(),
			ScreenX:  ev.Get("screenX").Float(),
			ScreenY:  ev.Get("screenU").Float(),
			MovemenX: ev.Get("movementX").Float(),
			MovemenY: ev.Get("movementY").Float(),
			Button:   ev.Get("button").Int(),
			Detail:   ev.Get("clientX").Int(),
			AltKey:   ev.Get("altKey").Bool(),
			CtrlKey:  ev.Get("ctrlKey").Bool(),
			MetaKey:  ev.Get("metaKey").Bool(),
			ShiftKey: ev.Get("shiftKey").Bool(),
		}, handle)
	case js.Global.Get("MutationEvent"):
		return events.NewBaseEvent(&events.MutationEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("OfflineAudioCompletionEvent"):
		return events.NewBaseEvent(&events.OfflineAudioCompletionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("PageTransitionEvent"):
		return events.NewBaseEvent(&events.PageTransitionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("PointerEvent"):
		return events.NewBaseEvent(&events.PointerEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("PopStateEvent"):
		return events.NewBaseEvent(&events.PopStateEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("ProgressEvent"):
		return events.NewBaseEvent(&events.ProgressEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("RelatedEvent"):
		return events.NewBaseEvent(&events.RelatedEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("RTCPeerConnectionIceEvent"):
		return events.NewBaseEvent(&events.RTCPeerConnectionIceEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("SensorEvent"):
		return events.NewBaseEvent(&events.SensorEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("StorageEvent"):
		return events.NewBaseEvent(&events.StorageEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("SVGEvent"):
		return events.NewBaseEvent(&events.SVGEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("SVGZoomEvent"):
		return events.NewBaseEvent(&events.SVGZoomEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TimeEvent"):
		return events.NewBaseEvent(&events.TimeEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TouchEvent"):
		return events.NewBaseEvent(&events.TouchEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TrackEvent"):
		return events.NewBaseEvent(&events.TrackEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("TransitionEvent"):
		return events.NewBaseEvent(&events.TransitionEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("UIEvent"):
		return events.NewBaseEvent(&events.UIEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("UserProximityEvent"):
		return events.NewBaseEvent(&events.UserProximityEvent{
			Core: ev,
		}, handle)
	case js.Global.Get("WheelEvent"):
		return events.NewBaseEvent(&events.WheelEvent{
			Core:      ev,
			DeltaX:    ev.Get("deltaX").Float(),
			DeltaY:    ev.Get("deltaX").Float(),
			DeltaZ:    ev.Get("deltaX").Float(),
			DeltaMode: events.DeltaMode(ev.Get("deltaMode").Uint64()),
		}, handle)
	}

	return events.NewBaseEvent(ev, handle)
}

// fromBlob transform the providded js.Object blob into a byte slice.
func fromBlob(o *js.Object) []byte {
	var buf []byte
	var wg sync.WaitGroup

	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onloadend", js.MakeFunc(func(this *js.Object, _ []*js.Object) interface{} {
		defer wg.Done()
		buf = js.Global.Get("Uint8Array").New(this.Get("result")).Interface().([]uint8)
		return nil
	}))

	fileReader.Call("readAsArrayBuffer", o)

	wg.Wait()
	return buf
}

// fromFile transform the providded js.Object blob into a byte slice.
func fromFile(o *js.Object) []byte {
	var buf []byte
	var wg sync.WaitGroup

	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onloadend", js.MakeFunc(func(this *js.Object, _ []*js.Object) interface{} {
		defer wg.Done()
		buf = js.Global.Get("Uint8Array").New(this.Get("result")).Interface().([]uint8)
		return nil
	}))

	fileReader.Call("readAsArrayBuffer", o)
	// fileReader.Call("readAsBinaryString", o)

	wg.Wait()
	return buf
}

// fromDataTransfer returns a transfer object from the js.Object.
func fromDataTransfer(o *js.Object) events.DataTransfer {
	var dt events.DataTransfer
	dt.DropEffect = o.Get("dropEffect").String()
	dt.EffectAllowed = o.Get("effectAllowed").String()

	types := o.Get("types")
	if types != nil && types != js.Undefined {
		dt.Types = shell.ObjectToStringList(types)
	}

	var dItems []events.DataTransferItem

	items := o.Get("items")
	for i := 0; i < items.Length(); i++ {
		item := items.Index(i)
		dItems = append(dItems, events.DataTransferItem{
			Name: item.Get("name").String(),
			Size: item.Get("size").Int(),
			Data: fromFile(item),
		})
	}

	var dFiles []events.DataTransferItem

	files := o.Get("files")
	for i := 0; i < files.Length(); i++ {
		item := files.Index(i)
		dFiles = append(dFiles, events.DataTransferItem{
			Name: item.Get("name").String(),
			Size: item.Get("size").Int(),
			Data: fromFile(item),
		})
	}

	dt.Items = events.DataTransferItemList{Items: dItems}
	dt.Files = dFiles
	return dt
}
