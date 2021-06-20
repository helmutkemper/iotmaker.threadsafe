package iotmaker_threadsafe

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"sync"
)

type platFunc func(iotmaker_platform_IDraw.IDraw)
type platformType iotmaker_platform_IDraw.IDraw

var closeTheDoorAndMakeTheWorldWait sync.Mutex

func ScratchPad(platform platformType, prepareGradientFilter, drawInvisible, getImageData, clearRectangle platFunc) {

	if platform == nil {
		return
	}

	closeTheDoorAndMakeTheWorldWait.Lock()

	if prepareGradientFilter != nil {
		prepareGradientFilter(platform)
	}

	if drawInvisible != nil {
		drawInvisible(platform)
	}

	platform.Stroke()

	if getImageData != nil {
		getImageData(platform)
	}

	if clearRectangle != nil {
		clearRectangle(platform)
	}

	closeTheDoorAndMakeTheWorldWait.Unlock()
}
