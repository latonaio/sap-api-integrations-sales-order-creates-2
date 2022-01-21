# sap-api-integrations-sales-order-creates  
sap-api-integrations-sales-order-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で受注データを登録するマイクロサービスです。  
sap-api-integrations-sales-order-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-sales-order-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_SALES_ORDER_SRV_0001/overview  

## 動作環境  
sap-api-integrations-sales-order-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-sales-order-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-order-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_ORDER_SRV_0001/overview  
* APIサービス名(=baseURL): API_PRODUCT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-order-creates には、次の API をコールするためのリソースが含まれています。  

* A_SalesOrder（受注 - ヘッダ）※受注の詳細データを取得するために、ToHeaderPartner、ToItem、ToItemPricingElement、ToItemScheduleLine、と合わせて利用されます。
* A_SalesOrderItem（受注 - 明細）※受注明細の詳細データを取得するために、ToItemPricingElement、ToItemScheduleLine、と合わせて利用されます。
* ToHeaderPartner（受注 - ヘッダ取引先）
* ToItem（受注 - 明細）
* ToItemPricingElement（受注 - 明細価格条件）
* ToItemScheduleLine（受注 - 明細納入日程行）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header", "Item" が指定されています。

```
"api_schema": "SAPSalesOrderCreate",
"accepter": ["Header", "Item"],
"sales_order": "1",
"deleted": false
```
  
* 全データを登録する際のsample.jsonの記載例(2)  

全データを登録する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPSalesOrderCreate",
"accepter": ["All"],
"sales_order": "1",
"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostSalesOrder(
	header             *requests.Header,
	item               *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
				func() {
					c.Item(item)
					wg.Done()
				}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 受注 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"Material" ～ "ProductStandardID" は、/SAP_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-order-creates/SAP_API_Caller/caller.go#L93",
	"function": "sap-api-integrations-sales-order-creates/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"Material": "A001",
			"Product_desc": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description",
			"BaseUnit": "AU",
			"ValidityStartDate": "",
			"ProductGroup": "A001",
			"Division": "00",
			"GrossWeight": "0.000",
			"WeightUnit": "KG",
			"SizeOrDimensionText": "",
			"ProductStandardID": ""
		}
	],
	"time": "2021-12-22T10:19:54.310555+09:00"
}
```
