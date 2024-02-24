package biz

type all struct {
	createItemBiz
	getItemBiz
	updateItemBiz
	deleteItemBiz
	listItemBiz
}

func NewItemUseCase() *all {
	return &all{}
}
