func maxProfit(prices []int) int {
    max:=0;
    
    buy:=0
    sell:=0
    
    for i:=0; i<len(prices); i++{
        if(prices[i]<prices[buy]){
            buy = i;
            sell = i;
        }
        if(prices[i]>prices[sell]){
            sell = i;
        }
        if((prices[sell]-prices[buy])>max){
            //fmt.Println(prices[sell],prices[buy])
            max = prices[sell]-prices[buy];
        }
    }
    return max;
}
