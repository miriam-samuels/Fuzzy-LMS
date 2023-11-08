package types

type Env struct{
	ConnectionPort string `barfenv:"key=CONNECTION_PORT;required=true"`
	AppName			string `barfenv:"key=APP_NAME;required=true"`
	JWTSecret 		string	`barfenv:"key=JWT_SECRET;required=true"`
	LoanDbDatasourceUri string `barfenv:"key=LOAN_DB_DATASOURCE_URI;required=true"`
	StorageBucket string `barfenv:"key=STORAGE_BUCKET;required=true"`
}


