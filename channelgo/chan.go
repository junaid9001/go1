package channelgo

func Channel1(n chan int){
	for i:=0;i<=9;i++{
		n<-i
	}
	close(n)
}