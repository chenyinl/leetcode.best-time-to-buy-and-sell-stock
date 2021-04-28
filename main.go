func maxProfit(prices []int) int {
    max:=0;
    
    buy:=0
    sell:=0
    
    for i:=0; i+1<len(prices); i++{
        if(prices[i]<prices[buy]){
            buy = i;
            sell = i+1;
        }
        if(prices[i+1]>prices[sell]){
            sell = i+1;
        }
        
        if((prices[sell]-prices[buy])>max){
            //fmt.Println(prices[sell],prices[buy])
            max = prices[sell]-prices[buy];
        }
    }
    return max;
}
