package bybit_ws

/*
bybit_ws.send('{"op": "subscribe", "args":["topic", "topic.filter"]}')。
// 同じタイプのフィルターが複数ある場合は、'|'で分割されます。
// 例：BTCUSDの1分足と3分足のクラインを購読する
bybit_ws.send('{"op": "subscribe", "args":["kline.BTCUSD.1m|3m"]}')です。
// 同じ種類のフィルターのすべてのデータを購読する場合は「*」を使う
// すべての製品のすべてのインターバル・クラインを購読する場合
bybit_ws.send('{"op": "subscribe", "args":["kline.*. *"]}')
*/
type Cmd struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}
