package bybit_ws

import "github.com/chuckpreslar/emission"

//特定のイベントにリスナーを追加します。
func (b *ByBitWS) On(event interface{}, listener interface{}) *emission.Emitter {
	return b.emitter.On(event, listener)
}

//イベントを発信する
func (b *ByBitWS) Emit(event interface{}, arguments ...interface{}) *emission.Emitter {
	return b.emitter.Emit(event, arguments...)
}

//イベントのリスナーを削除します。
func (b *ByBitWS) Off(event interface{}, listener interface{}) *emission.Emitter {
	return b.emitter.Off(event, listener)
}
