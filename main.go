package main 

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item){
	
}

func (p *Player) DropItem(itemName string){

}

func (p *Player) UseItem(itemName string){

}

func main(){

}