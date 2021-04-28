要解決問題，首先要確定問題來源
查看服務的CPU、I/O、記憶體各項數值是否正常
確認Sql執行計畫，針對結果進行優化
> 
RDBMS的讀取瓶頸可能的原因有: 資料量體過大、Command過於複雜、沒有index
寫入瓶頸大多是因為資料量體大時又有過多index而導致
改善具體作法逐點如下:

* 1. 使用Cache機制減低資料庫壓力
>  
* 2. 將不同業務的資料存放於不同資料庫，缺點是分庫之後就不能使用join
所以domain劃分要清楚
>  
* 3. 建立archive流程，將一定時間以前的檔案封存，減少資料量體。
ex: 一年前的完成訂單資料封存後移至history。
這個做法快速，只需要提供封存資料的查詢介面
>  
* 4. 若資料有即時性，那就要進行分表作業，將不常用到的欄位分開(垂直分表)
資料有sequence的話，每一定數量就新增表去存放(水平分表)
這方式要注意業務邏輯
>  
* 5. 建立讀寫分離資料庫，降低併發鎖表的機率
另一方面index優化可以只針對讀取資料庫，可增進寫入效率