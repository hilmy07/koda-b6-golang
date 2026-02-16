package service

var historyOrders [][]Menu

func AddHistory(order []Menu) {
	historyOrders = append(historyOrders, order)
}

func GetHistory() [][]Menu {
	return historyOrders
}

func HistoryIsEmpty() bool {
	return len(historyOrders) == 0
}
