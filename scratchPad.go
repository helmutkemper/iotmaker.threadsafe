package iotmaker_threadsafe

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"sync"
)

type platFunc func(iotmaker_platform_IDraw.IDraw)
type platformType iotmaker_platform_IDraw.IDraw

var closeTheDoorAndMakeTheWorldWait sync.Mutex

func ScratchPad(platform platformType, drawInvisible, getImageData, clearRectangle platFunc) {
	closeTheDoorAndMakeTheWorldWait.Lock()

	drawInvisible(platform)
	platform.Stroke()
	getImageData(platform)
	clearRectangle(platform)

	closeTheDoorAndMakeTheWorldWait.Unlock()
}
